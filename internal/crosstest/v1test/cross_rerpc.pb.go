// Code generated by protoc-gen-go-rerpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-rerpc v0.0.1
// - protoc             v3.17.3
// source: internal/crosstest/v1test/cross.proto

package crosspb

import (
	context "context"
	errors "errors"
	rerpc "github.com/rerpc/rerpc"
	callstream "github.com/rerpc/rerpc/callstream"
	handlerstream "github.com/rerpc/rerpc/handlerstream"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the
// rerpc package are compatible. If you get a compiler error that this constant
// isn't defined, this code was generated with a version of rerpc newer than the
// one compiled into your binary. You can fix the problem by either regenerating
// this code with an older version of rerpc or updating the rerpc version
// compiled into your binary.
const _ = rerpc.SupportsCodeGenV0 // requires reRPC v0.0.1 or later

// CrossServiceClientReRPC is a client for the
// internal.crosstest.v1test.CrossService service.
type CrossServiceClientReRPC interface {
	Ping(ctx context.Context, req *PingRequest) (*PingResponse, error)
	Fail(ctx context.Context, req *FailRequest) (*FailResponse, error)
	Sum(ctx context.Context) *callstream.Client[SumRequest, SumResponse]
	CountUp(ctx context.Context, req *CountUpRequest) (*callstream.Server[CountUpResponse], error)
	CumSum(ctx context.Context) *callstream.Bidirectional[CumSumRequest, CumSumResponse]
}

type crossServiceClientReRPC struct {
	doer    rerpc.Doer
	baseURL string
	options []rerpc.CallOption
}

// NewCrossServiceClientReRPC constructs a client for the
// internal.crosstest.v1test.CrossService service. Call options passed here
// apply to all calls made with this client.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewCrossServiceClientReRPC(baseURL string, doer rerpc.Doer, opts ...rerpc.CallOption) CrossServiceClientReRPC {
	return &crossServiceClientReRPC{
		baseURL: strings.TrimRight(baseURL, "/"),
		doer:    doer,
		options: opts,
	}
}

// Ping calls internal.crosstest.v1test.CrossService.Ping. Call options passed
// here apply only to this call.
func (c *crossServiceClientReRPC) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	call := rerpc.NewClientFunc[PingRequest, PingResponse](
		c.doer,
		c.baseURL,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"Ping",                      // protobuf method
		c.options...,
	)
	return call(ctx, req)
}

// Fail calls internal.crosstest.v1test.CrossService.Fail. Call options passed
// here apply only to this call.
func (c *crossServiceClientReRPC) Fail(ctx context.Context, req *FailRequest) (*FailResponse, error) {
	call := rerpc.NewClientFunc[FailRequest, FailResponse](
		c.doer,
		c.baseURL,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"Fail",                      // protobuf method
		c.options...,
	)
	return call(ctx, req)
}

// Sum calls internal.crosstest.v1test.CrossService.Sum. Call options passed
// here apply only to this call.
func (c *crossServiceClientReRPC) Sum(ctx context.Context) *callstream.Client[SumRequest, SumResponse] {
	ctx, call := rerpc.NewClientStream(
		ctx,
		c.doer,
		rerpc.StreamTypeClient,
		c.baseURL,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"Sum",                       // protobuf method
		c.options...,
	)
	stream := call(ctx)
	return callstream.NewClient[SumRequest, SumResponse](stream)
}

// CountUp calls internal.crosstest.v1test.CrossService.CountUp. Call options
// passed here apply only to this call.
func (c *crossServiceClientReRPC) CountUp(ctx context.Context, req *CountUpRequest) (*callstream.Server[CountUpResponse], error) {
	ctx, call := rerpc.NewClientStream(
		ctx,
		c.doer,
		rerpc.StreamTypeServer,
		c.baseURL,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"CountUp",                   // protobuf method
		c.options...,
	)
	stream := call(ctx)
	if err := stream.Send(req); err != nil {
		_ = stream.CloseSend(err)
		_ = stream.CloseReceive()
		return nil, err
	}
	if err := stream.CloseSend(nil); err != nil {
		_ = stream.CloseReceive()
		return nil, err
	}
	return callstream.NewServer[CountUpResponse](stream), nil
}

// CumSum calls internal.crosstest.v1test.CrossService.CumSum. Call options
// passed here apply only to this call.
func (c *crossServiceClientReRPC) CumSum(ctx context.Context) *callstream.Bidirectional[CumSumRequest, CumSumResponse] {
	ctx, call := rerpc.NewClientStream(
		ctx,
		c.doer,
		rerpc.StreamTypeBidirectional,
		c.baseURL,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"CumSum",                    // protobuf method
		c.options...,
	)
	stream := call(ctx)
	return callstream.NewBidirectional[CumSumRequest, CumSumResponse](stream)
}

// CrossServiceReRPC is a server for the internal.crosstest.v1test.CrossService
// service. To make sure that adding methods to this protobuf service doesn't
// break all implementations of this interface, all implementations must embed
// UnimplementedCrossServiceReRPC.
//
// By default, recent versions of grpc-go have a similar forward compatibility
// requirement. See https://github.com/grpc/grpc-go/issues/3794 for a longer
// discussion.
type CrossServiceReRPC interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Fail(context.Context, *FailRequest) (*FailResponse, error)
	Sum(context.Context, *handlerstream.Client[SumRequest, SumResponse]) error
	CountUp(context.Context, *CountUpRequest, *handlerstream.Server[CountUpResponse]) error
	CumSum(context.Context, *handlerstream.Bidirectional[CumSumRequest, CumSumResponse]) error
	mustEmbedUnimplementedCrossServiceReRPC()
}

// NewCrossServiceHandlerReRPC wraps each method on the service implementation
// in a *rerpc.Handler. The returned slice can be passed to rerpc.NewServeMux.
func NewCrossServiceHandlerReRPC(svc CrossServiceReRPC, opts ...rerpc.HandlerOption) []*rerpc.Handler {
	handlers := make([]*rerpc.Handler, 0, 5)

	ping := rerpc.NewUnaryHandler(
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"Ping",                      // protobuf method
		svc.Ping,
		opts...,
	)
	handlers = append(handlers, ping)

	fail := rerpc.NewUnaryHandler(
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"Fail",                      // protobuf method
		svc.Fail,
		opts...,
	)
	handlers = append(handlers, fail)

	sum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeClient,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"Sum",                       // protobuf method
		func(ctx context.Context, sf rerpc.StreamFunc) {
			stream := sf(ctx)
			typed := handlerstream.NewClient[SumRequest, SumResponse](stream)
			err := svc.Sum(stream.Context(), typed)
			_ = stream.CloseReceive()
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = stream.CloseSend(err)
		},
		opts...,
	)
	handlers = append(handlers, sum)

	countUp := rerpc.NewStreamingHandler(
		rerpc.StreamTypeServer,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"CountUp",                   // protobuf method
		func(ctx context.Context, sf rerpc.StreamFunc) {
			stream := sf(ctx)
			typed := handlerstream.NewServer[CountUpResponse](stream)
			var req CountUpRequest
			if err := stream.Receive(&req); err != nil {
				_ = stream.CloseReceive()
				_ = stream.CloseSend(err)
				return
			}
			if err := stream.CloseReceive(); err != nil {
				_ = stream.CloseSend(err)
				return
			}
			err := svc.CountUp(stream.Context(), &req, typed)
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = stream.CloseSend(err)
		},
		opts...,
	)
	handlers = append(handlers, countUp)

	cumSum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeBidirectional,
		"internal.crosstest.v1test", // protobuf package
		"CrossService",              // protobuf service
		"CumSum",                    // protobuf method
		func(ctx context.Context, sf rerpc.StreamFunc) {
			stream := sf(ctx)
			typed := handlerstream.NewBidirectional[CumSumRequest, CumSumResponse](stream)
			err := svc.CumSum(stream.Context(), typed)
			_ = stream.CloseReceive()
			if err != nil {
				if _, ok := rerpc.AsError(err); !ok {
					if errors.Is(err, context.Canceled) {
						err = rerpc.Wrap(rerpc.CodeCanceled, err)
					}
					if errors.Is(err, context.DeadlineExceeded) {
						err = rerpc.Wrap(rerpc.CodeDeadlineExceeded, err)
					}
				}
			}
			_ = stream.CloseSend(err)
		},
		opts...,
	)
	handlers = append(handlers, cumSum)

	return handlers
}

var _ CrossServiceReRPC = (*UnimplementedCrossServiceReRPC)(nil) // verify interface implementation

// UnimplementedCrossServiceReRPC returns CodeUnimplemented from all methods. To
// maintain forward compatibility, all implementations of CrossServiceReRPC must
// embed UnimplementedCrossServiceReRPC.
type UnimplementedCrossServiceReRPC struct{}

func (UnimplementedCrossServiceReRPC) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "internal.crosstest.v1test.CrossService.Ping isn't implemented")
}

func (UnimplementedCrossServiceReRPC) Fail(context.Context, *FailRequest) (*FailResponse, error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "internal.crosstest.v1test.CrossService.Fail isn't implemented")
}

func (UnimplementedCrossServiceReRPC) Sum(context.Context, *handlerstream.Client[SumRequest, SumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.crosstest.v1test.CrossService.Sum isn't implemented")
}

func (UnimplementedCrossServiceReRPC) CountUp(context.Context, *CountUpRequest, *handlerstream.Server[CountUpResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.crosstest.v1test.CrossService.CountUp isn't implemented")
}

func (UnimplementedCrossServiceReRPC) CumSum(context.Context, *handlerstream.Bidirectional[CumSumRequest, CumSumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.crosstest.v1test.CrossService.CumSum isn't implemented")
}

func (UnimplementedCrossServiceReRPC) mustEmbedUnimplementedCrossServiceReRPC() {}
