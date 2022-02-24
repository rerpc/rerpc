package connect

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect/codec"
	"github.com/bufbuild/connect/compress"
)

type handlerConfiguration struct {
	Compressors      map[string]compress.Compressor
	Codecs           map[string]codec.Codec
	MaxRequestBytes  int64
	Registrar        *Registrar
	Interceptor      Interceptor
	Procedure        string
	RegistrationName string
	HandleGRPC       bool
	HandleGRPCWeb    bool
}

func newHandlerConfiguration(procedure, registrationName string, options []HandlerOption) (*handlerConfiguration, *Error) {
	config := handlerConfiguration{
		Procedure:        procedure,
		RegistrationName: registrationName,
		Compressors:      make(map[string]compress.Compressor),
		Codecs:           make(map[string]codec.Codec),
	}
	for _, opt := range options {
		opt.applyToHandler(&config)
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}
	if reg := config.Registrar; reg != nil && config.RegistrationName != "" {
		reg.register(config.RegistrationName)
	}
	return &config, nil
}

func (c *handlerConfiguration) Validate() *Error {
	if _, ok := c.Codecs[""]; ok {
		return NewError(
			CodeUnknown,
			errors.New("can't register codec with an empty name"),
		)
	}
	if _, ok := c.Compressors[""]; ok {
		return NewError(
			CodeUnknown,
			errors.New("can't register compressor with an empty name"),
		)
	}
	if !(c.HandleGRPC || c.HandleGRPCWeb) {
		return NewError(
			CodeUnknown,
			errors.New("handlers must support at least one protocol"),
		)
	}
	return nil
}

func (c *handlerConfiguration) newSpecification(streamType StreamType) Specification {
	return Specification{
		Procedure: c.Procedure,
		Type:      streamType,
	}
}

func (c *handlerConfiguration) newProtocolHandlers(streamType StreamType) ([]protocolHandler, *Error) {
	var protocols []protocol
	if c.HandleGRPC {
		protocols = append(protocols, &protocolGRPC{web: false})
	}
	if c.HandleGRPCWeb {
		protocols = append(protocols, &protocolGRPC{web: true})
	}
	handlers := make([]protocolHandler, 0, len(protocols))
	codecs := newReadOnlyCodecs(c.Codecs)
	compressors := newReadOnlyCompressors(c.Compressors)
	for _, protocol := range protocols {
		protocolHandler, err := protocol.NewHandler(&protocolHandlerParams{
			Spec:            c.newSpecification(streamType),
			Codecs:          codecs,
			Compressors:     compressors,
			MaxRequestBytes: c.MaxRequestBytes,
		})
		if err != nil {
			return nil, NewError(CodeUnknown, err)
		}
		handlers = append(handlers, protocolHandler)
	}
	return handlers, nil
}

// A HandlerOption configures a Handler.
//
// In addition to any options grouped in the documentation below, remember that
// Registrars and Options are also valid HandlerOptions.
type HandlerOption interface {
	applyToHandler(*handlerConfiguration)
}

// A Handler is the server-side implementation of a single RPC defined by a
// protocol buffer service. It's the interface between the connect library and
// the code generated by the connect protoc plugin; most users won't ever need
// to deal with it directly.
//
// To see an example of how Handler is used in the generated code, see the
// internal/gen/proto/go-connect/connect/ping/v1test package.
type Handler struct {
	spec             Specification
	interceptor      Interceptor
	implementation   func(context.Context, Sender, Receiver, error /* client-visible */)
	protocolHandlers []protocolHandler
}

// NewUnaryHandler constructs a Handler for a request-response procedure. It's
// used in generated code, so most users won't need to call it directly.
func NewUnaryHandler[Req, Res any](
	procedure, registrationName string,
	unary func(context.Context, *Request[Req]) (*Response[Res], error),
	options ...HandlerOption,
) (*Handler, error) {
	config, err := newHandlerConfiguration(procedure, registrationName, options)
	if err != nil {
		return nil, err
	}
	// Given a (possibly failed) stream, how should we call the unary function?
	implementation := func(ctx context.Context, sender Sender, receiver Receiver, clientVisibleError error) {
		defer receiver.Close()

		var request *Request[Req]
		if clientVisibleError != nil {
			// The protocol implementation failed to establish a stream. To make the
			// resulting error visible to the interceptor stack, we still want to
			// call the wrapped unary Func. To do that safely, we need a useful
			// Request struct. (Note that we do *not* actually calling the handler's
			// implementation.)
			request = receiveRequestMetadata[Req](receiver)
		} else {
			var err error
			request, err = ReceiveRequest[Req](receiver)
			if err != nil {
				// Interceptors should see this error too. Just as above, they need a
				// useful Request.
				clientVisibleError = err
				request = receiveRequestMetadata[Req](receiver)
			}
		}

		untyped := Func(func(ctx context.Context, request AnyRequest) (AnyResponse, error) {
			if clientVisibleError != nil {
				// We've already encountered an error, short-circuit before calling the
				// handler's implementation.
				return nil, clientVisibleError
			}
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			typed, ok := request.(*Request[Req])
			if !ok {
				return nil, errorf(CodeInternal, "unexpected handler request type %T", request)
			}
			return unary(ctx, typed)
		})
		if ic := config.Interceptor; ic != nil {
			untyped = ic.Wrap(untyped)
		}

		response, err := untyped(ctx, request)
		if err != nil {
			_ = sender.Close(err)
			return
		}
		mergeHeaders(sender.Header(), response.Header())
		mergeHeaders(sender.Trailer(), response.Trailer())
		_ = sender.Close(sender.Send(response.Any()))
	}

	protocolHandlers, err := config.newProtocolHandlers(StreamTypeUnary)
	if err != nil {
		return nil, err
	}
	return &Handler{
		spec:             config.newSpecification(StreamTypeUnary),
		interceptor:      nil, // already applied
		implementation:   implementation,
		protocolHandlers: protocolHandlers,
	}, nil
}

// NewStreamHandler constructs a Handler for a streaming procedure. It's
// used in generated code, so most users won't need to call it directly.
func NewStreamHandler(
	procedure, registrationName string,
	streamType StreamType,
	implementation func(context.Context, Sender, Receiver),
	options ...HandlerOption,
) (*Handler, error) {
	config, err := newHandlerConfiguration(procedure, registrationName, options)
	if err != nil {
		return nil, err
	}
	protocolHandlers, err := config.newProtocolHandlers(streamType)
	if err != nil {
		return nil, err
	}
	return &Handler{
		spec:        config.newSpecification(streamType),
		interceptor: config.Interceptor,
		implementation: func(ctx context.Context, sender Sender, receiver Receiver, clientVisibleErr error) {
			if clientVisibleErr != nil {
				_ = receiver.Close()
				_ = sender.Close(clientVisibleErr)
				return
			}
			implementation(ctx, sender, receiver)
		},
		protocolHandlers: protocolHandlers,
	}, nil
}

// ServeHTTP implements http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// We don't need to defer functions  to close the request body or read to
	// EOF: the stream we construct later on already does that, and we only
	// return early when dealing with misbehaving clients. In those cases, it's
	// okay if we can't re-use the connection.
	isBidi := (h.spec.Type & StreamTypeBidirectional) == StreamTypeBidirectional
	if isBidi && r.ProtoMajor < 2 {
		h.failNegotiation(w, http.StatusHTTPVersionNotSupported)
		return
	}

	methodHandlers := make([]protocolHandler, 0, len(h.protocolHandlers))
	for _, protocolHandler := range h.protocolHandlers {
		if protocolHandler.ShouldHandleMethod(r.Method) {
			methodHandlers = append(methodHandlers, protocolHandler)
		}
	}
	if len(methodHandlers) == 0 {
		// grpc-go returns a 500 here, but interoperability with non-gRPC HTTP
		// clients is better if we return a 405.
		h.failNegotiation(w, http.StatusMethodNotAllowed)
		return
	}

	// TODO: for GETs, we should parse the Accept header and offer each handler
	// each content-type.
	contentType := r.Header.Get("Content-Type")
	for _, protocolHandler := range methodHandlers {
		if !protocolHandler.ShouldHandleContentType(contentType) {
			continue
		}
		ctx := r.Context()
		if ic := h.interceptor; ic != nil {
			ctx = ic.WrapContext(ctx)
		}
		// Most errors returned from protocolHandler.NewStream are caused by
		// invalid requests. For example, the client may have specified an invalid
		// timeout or an unavailable codec. We'd like those errors to be visible to
		// the interceptor chain, so we're going to capture them here and pass them
		// to the implementation.
		sender, receiver, clientVisibleError := protocolHandler.NewStream(w, r.WithContext(ctx))
		// If NewStream errored and the protocol doesn't want the error sent to
		// the client, sender and/or receiver may be nil. We still want the
		// error to be seen by interceptors, so we provide no-op Sender and
		// Receiver implementations.
		if clientVisibleError != nil && sender == nil {
			sender = newNopSender(h.spec, w.Header(), make(http.Header))
		}
		if clientVisibleError != nil && receiver == nil {
			receiver = newNopReceiver(h.spec, r.Header, r.Trailer)
		}
		if ic := h.interceptor; ic != nil {
			// Unary interceptors were handled in NewUnaryHandler.
			sender = ic.WrapSender(ctx, sender)
			receiver = ic.WrapReceiver(ctx, receiver)
		}
		h.implementation(ctx, sender, receiver, clientVisibleError)
		return
	}
	h.failNegotiation(w, http.StatusUnsupportedMediaType)
}

// Path returns the URL pattern to use when registering this handler.
func (h *Handler) path() string {
	return fmt.Sprintf("/" + h.spec.Procedure)
}

func (h *Handler) failNegotiation(w http.ResponseWriter, code int) {
	// None of the registered protocols is able to serve the request.
	for _, ph := range h.protocolHandlers {
		ph.WriteAccept(w.Header())
	}
	w.WriteHeader(code)
}
