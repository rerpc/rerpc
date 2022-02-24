// Code generated by protoc-gen-go-connect. DO NOT EDIT.
// versions:
// - protoc-gen-go-connect v0.0.1
// - protoc              v3.17.3
// source: connect/ping/v1test/ping.proto

package pingv1test

import (
	context "context"
	connect "github.com/bufbuild/connect"
	clientstream "github.com/bufbuild/connect/clientstream"
	protobuf "github.com/bufbuild/connect/codec/protobuf"
	protojson "github.com/bufbuild/connect/codec/protojson"
	gzip "github.com/bufbuild/connect/compress/gzip"
	handlerstream "github.com/bufbuild/connect/handlerstream"
	v1test "github.com/bufbuild/connect/internal/gen/proto/go/connect/ping/v1test"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// connect package are compatible. If you get a compiler error that this
// constant isn't defined, this code was generated with a version of connect
// newer than the one compiled into your binary. You can fix the problem by
// either regenerating this code with an older version of connect or updating
// the connect version compiled into your binary.
const _ = connect.IsAtLeastVersion0_0_1

// PingServiceClient is a client for the connect.ping.v1test.PingService
// service.
type PingServiceClient interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse]
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect.Request[v1test.CountUpRequest]) (*clientstream.Server[v1test.CountUpResponse], error)
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]
}

// NewPingServiceClient constructs a client for the
// connect.ping.v1test.PingService service. By default, it uses the binary
// protobuf codec.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(baseURL string, doer connect.Doer, opts ...connect.ClientOption) (PingServiceClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	opts = append([]connect.ClientOption{
		connect.WithGRPC(true),
		connect.WithCodec(protobuf.Name, protobuf.New()),
		connect.WithCompressor(gzip.Name, gzip.New()),
	}, opts...)
	var (
		client pingServiceClient
		err    error
	)
	client.ping, err = connect.NewUnaryClientImplementation[v1test.PingRequest, v1test.PingResponse](
		doer,
		baseURL,
		"connect.ping.v1test.PingService/Ping",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.fail, err = connect.NewUnaryClientImplementation[v1test.FailRequest, v1test.FailResponse](
		doer,
		baseURL,
		"connect.ping.v1test.PingService/Fail",
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.sum, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"connect.ping.v1test.PingService/Sum",
		connect.StreamTypeClient,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.countUp, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"connect.ping.v1test.PingService/CountUp",
		connect.StreamTypeServer,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client.cumSum, err = connect.NewStreamClientImplementation(
		doer,
		baseURL,
		"connect.ping.v1test.PingService/CumSum",
		connect.StreamTypeBidirectional,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping    func(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	fail    func(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	sum     func(context.Context) (connect.Sender, connect.Receiver)
	countUp func(context.Context) (connect.Sender, connect.Receiver)
	cumSum  func(context.Context) (connect.Sender, connect.Receiver)
}

var _ PingServiceClient = (*pingServiceClient)(nil) // verify interface implementation

// Ping calls connect.ping.v1test.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error) {
	return c.ping(ctx, req)
}

// Fail calls connect.ping.v1test.PingService.Fail.
func (c *pingServiceClient) Fail(ctx context.Context, req *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error) {
	return c.fail(ctx, req)
}

// Sum calls connect.ping.v1test.PingService.Sum.
func (c *pingServiceClient) Sum(ctx context.Context) *clientstream.Client[v1test.SumRequest, v1test.SumResponse] {
	sender, receiver := c.sum(ctx)
	return clientstream.NewClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
}

// CountUp calls connect.ping.v1test.PingService.CountUp.
func (c *pingServiceClient) CountUp(ctx context.Context, req *connect.Request[v1test.CountUpRequest]) (*clientstream.Server[v1test.CountUpResponse], error) {
	sender, receiver := c.countUp(ctx)
	for key, values := range req.Header() {
		sender.Header()[key] = append(sender.Header()[key], values...)
	}
	for key, values := range req.Trailer() {
		sender.Trailer()[key] = append(sender.Trailer()[key], values...)
	}
	if err := sender.Send(req.Msg); err != nil {
		_ = sender.Close(err)
		_ = receiver.Close()
		return nil, err
	}
	if err := sender.Close(nil); err != nil {
		_ = receiver.Close()
		return nil, err
	}
	return clientstream.NewServer[v1test.CountUpResponse](receiver), nil
}

// CumSum calls connect.ping.v1test.PingService.CumSum.
func (c *pingServiceClient) CumSum(ctx context.Context) *clientstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse] {
	sender, receiver := c.cumSum(ctx)
	return clientstream.NewBidirectional[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
}

// PingServiceHandler is an implementation of the
// connect.ping.v1test.PingService service.
type PingServiceHandler interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

// WithPingServiceHandler wraps the service implementation in a
// connect.MuxOption, which can then be passed to connect.NewServeMux.
//
// By default, services support the gRPC and gRPC-Web protocols with the binary
// protobuf and JSON codecs.
func WithPingServiceHandler(svc PingServiceHandler, opts ...connect.HandlerOption) connect.MuxOption {
	handlers := make([]connect.Handler, 0, 5)
	opts = append([]connect.HandlerOption{
		connect.WithGRPC(true),
		connect.WithGRPCWeb(true),
		connect.WithCodec(protobuf.Name, protobuf.New()),
		connect.WithCodec(protojson.Name, protojson.New()),
		connect.WithCompressor(gzip.Name, gzip.New()),
	}, opts...)

	ping, err := connect.NewUnaryHandler(
		"connect.ping.v1test.PingService/Ping", // procedure name
		"connect.ping.v1test.PingService",      // reflection name
		svc.Ping,
		opts...,
	)
	if err != nil {
		return connect.WithHandlers(nil, err)
	}
	handlers = append(handlers, *ping)

	fail, err := connect.NewUnaryHandler(
		"connect.ping.v1test.PingService/Fail", // procedure name
		"connect.ping.v1test.PingService",      // reflection name
		svc.Fail,
		opts...,
	)
	if err != nil {
		return connect.WithHandlers(nil, err)
	}
	handlers = append(handlers, *fail)

	sum, err := connect.NewStreamingHandler(
		connect.StreamTypeClient,
		"connect.ping.v1test.PingService/Sum", // procedure name
		"connect.ping.v1test.PingService",     // reflection name
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := handlerstream.NewClient[v1test.SumRequest, v1test.SumResponse](sender, receiver)
			err := svc.Sum(ctx, typed)
			_ = receiver.Close()
			_ = sender.Close(err)
		},
		opts...,
	)
	if err != nil {
		return connect.WithHandlers(nil, err)
	}
	handlers = append(handlers, *sum)

	countUp, err := connect.NewStreamingHandler(
		connect.StreamTypeServer,
		"connect.ping.v1test.PingService/CountUp", // procedure name
		"connect.ping.v1test.PingService",         // reflection name
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := handlerstream.NewServer[v1test.CountUpResponse](sender)
			req, err := connect.ReceiveRequest[v1test.CountUpRequest](receiver)
			if err != nil {
				_ = receiver.Close()
				_ = sender.Close(err)
				return
			}
			if err = receiver.Close(); err != nil {
				_ = sender.Close(err)
				return
			}
			err = svc.CountUp(ctx, req, typed)
			_ = sender.Close(err)
		},
		opts...,
	)
	if err != nil {
		return connect.WithHandlers(nil, err)
	}
	handlers = append(handlers, *countUp)

	cumSum, err := connect.NewStreamingHandler(
		connect.StreamTypeBidirectional,
		"connect.ping.v1test.PingService/CumSum", // procedure name
		"connect.ping.v1test.PingService",        // reflection name
		func(ctx context.Context, sender connect.Sender, receiver connect.Receiver) {
			typed := handlerstream.NewBidirectional[v1test.CumSumRequest, v1test.CumSumResponse](sender, receiver)
			err := svc.CumSum(ctx, typed)
			_ = receiver.Close()
			_ = sender.Close(err)
		},
		opts...,
	)
	if err != nil {
		return connect.WithHandlers(nil, err)
	}
	handlers = append(handlers, *cumSum)

	return connect.WithHandlers(handlers, nil)
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

var _ PingServiceHandler = (*UnimplementedPingServiceHandler)(nil) // verify interface implementation

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect.Request[v1test.PingRequest]) (*connect.Response[v1test.PingResponse], error) {
	return nil, connect.Errorf(connect.CodeUnimplemented, "connect.ping.v1test.PingService.Ping isn't implemented")
}

func (UnimplementedPingServiceHandler) Fail(context.Context, *connect.Request[v1test.FailRequest]) (*connect.Response[v1test.FailResponse], error) {
	return nil, connect.Errorf(connect.CodeUnimplemented, "connect.ping.v1test.PingService.Fail isn't implemented")
}

func (UnimplementedPingServiceHandler) Sum(context.Context, *handlerstream.Client[v1test.SumRequest, v1test.SumResponse]) error {
	return connect.Errorf(connect.CodeUnimplemented, "connect.ping.v1test.PingService.Sum isn't implemented")
}

func (UnimplementedPingServiceHandler) CountUp(context.Context, *connect.Request[v1test.CountUpRequest], *handlerstream.Server[v1test.CountUpResponse]) error {
	return connect.Errorf(connect.CodeUnimplemented, "connect.ping.v1test.PingService.CountUp isn't implemented")
}

func (UnimplementedPingServiceHandler) CumSum(context.Context, *handlerstream.Bidirectional[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return connect.Errorf(connect.CodeUnimplemented, "connect.ping.v1test.PingService.CumSum isn't implemented")
}
