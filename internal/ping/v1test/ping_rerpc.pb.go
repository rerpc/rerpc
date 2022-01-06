// Code generated by protoc-gen-go-rerpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-rerpc v0.0.1
// - protoc             v3.17.3
// source: internal/ping/v1test/ping.proto

package pingpb

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

// PingServiceClientReRPC is a client for the internal.ping.v1test.PingService
// service.
type PingServiceClientReRPC interface {
	Ping(ctx context.Context, req *rerpc.Request[PingRequest]) (*rerpc.Response[PingResponse], error)
	Fail(ctx context.Context, req *rerpc.Request[FailRequest]) (*rerpc.Response[FailResponse], error)
	Sum(ctx context.Context) *callstream.Client[SumRequest, SumResponse]
	CountUp(ctx context.Context, req *rerpc.Request[CountUpRequest]) (*callstream.Server[CountUpResponse], error)
	CumSum(ctx context.Context) *callstream.Bidirectional[CumSumRequest, CumSumResponse]
}

type pingServiceClientReRPC struct {
	doer    rerpc.Doer
	baseURL string
	options []rerpc.ClientOption
}

// NewPingServiceClientReRPC constructs a client for the
// internal.ping.v1test.PingService service.
//
// The URL supplied here should be the base URL for the gRPC server (e.g.,
// https://api.acme.com or https://acme.com/grpc).
func NewPingServiceClientReRPC(baseURL string, doer rerpc.Doer, opts ...rerpc.ClientOption) PingServiceClientReRPC {
	return &pingServiceClientReRPC{
		baseURL: strings.TrimRight(baseURL, "/"),
		doer:    doer,
		options: opts,
	}
}

// Ping calls internal.ping.v1test.PingService.Ping.
func (c *pingServiceClientReRPC) Ping(ctx context.Context, req *rerpc.Request[PingRequest]) (*rerpc.Response[PingResponse], error) {
	call := rerpc.NewClientFunc[PingRequest, PingResponse](
		c.doer,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Ping",                 // protobuf method
		c.options...,
	)
	return call(ctx, req)
}

// Fail calls internal.ping.v1test.PingService.Fail.
func (c *pingServiceClientReRPC) Fail(ctx context.Context, req *rerpc.Request[FailRequest]) (*rerpc.Response[FailResponse], error) {
	call := rerpc.NewClientFunc[FailRequest, FailResponse](
		c.doer,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Fail",                 // protobuf method
		c.options...,
	)
	return call(ctx, req)
}

// Sum calls internal.ping.v1test.PingService.Sum.
func (c *pingServiceClientReRPC) Sum(ctx context.Context) *callstream.Client[SumRequest, SumResponse] {
	call := rerpc.NewClientStream(
		c.doer,
		rerpc.StreamTypeClient,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Sum",                  // protobuf method
		c.options...,
	)
	_, stream := call(ctx)
	return callstream.NewClient[SumRequest, SumResponse](stream)
}

// CountUp calls internal.ping.v1test.PingService.CountUp.
func (c *pingServiceClientReRPC) CountUp(ctx context.Context, req *rerpc.Request[CountUpRequest]) (*callstream.Server[CountUpResponse], error) {
	call := rerpc.NewClientStream(
		c.doer,
		rerpc.StreamTypeServer,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CountUp",              // protobuf method
		c.options...,
	)
	_, stream := call(ctx)
	if err := stream.Send(req.Any()); err != nil {
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

// CumSum calls internal.ping.v1test.PingService.CumSum.
func (c *pingServiceClientReRPC) CumSum(ctx context.Context) *callstream.Bidirectional[CumSumRequest, CumSumResponse] {
	call := rerpc.NewClientStream(
		c.doer,
		rerpc.StreamTypeBidirectional,
		c.baseURL,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CumSum",               // protobuf method
		c.options...,
	)
	_, stream := call(ctx)
	return callstream.NewBidirectional[CumSumRequest, CumSumResponse](stream)
}

// PingServiceReRPC is a server for the internal.ping.v1test.PingService
// service.
type PingServiceReRPC interface {
	Ping(context.Context, *rerpc.Request[PingRequest]) (*rerpc.Response[PingResponse], error)
	Fail(context.Context, *rerpc.Request[FailRequest]) (*rerpc.Response[FailResponse], error)
	Sum(context.Context, *handlerstream.Client[SumRequest, SumResponse]) error
	CountUp(context.Context, *rerpc.Request[CountUpRequest], *handlerstream.Server[CountUpResponse]) error
	CumSum(context.Context, *handlerstream.Bidirectional[CumSumRequest, CumSumResponse]) error
}

// NewPingServiceHandlerReRPC wraps each method on the service implementation in
// a *rerpc.Handler. The returned slice can be passed to rerpc.NewServeMux.
func NewPingServiceHandlerReRPC(svc PingServiceReRPC, opts ...rerpc.HandlerOption) []*rerpc.Handler {
	handlers := make([]*rerpc.Handler, 0, 5)

	ping := rerpc.NewUnaryHandler(
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Ping",                 // protobuf method
		svc.Ping,
		opts...,
	)
	handlers = append(handlers, ping)

	fail := rerpc.NewUnaryHandler(
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Fail",                 // protobuf method
		svc.Fail,
		opts...,
	)
	handlers = append(handlers, fail)

	sum := rerpc.NewStreamingHandler(
		rerpc.StreamTypeClient,
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"Sum",                  // protobuf method
		func(ctx context.Context, stream rerpc.Stream) {
			typed := handlerstream.NewClient[SumRequest, SumResponse](stream)
			err := svc.Sum(ctx, typed)
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
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CountUp",              // protobuf method
		func(ctx context.Context, stream rerpc.Stream) {
			typed := handlerstream.NewServer[CountUpResponse](stream)
			req, err := rerpc.NewReceivedRequest[CountUpRequest](stream)
			if err != nil {
				_ = stream.CloseReceive()
				_ = stream.CloseSend(err)
				return
			}
			if err = stream.CloseReceive(); err != nil {
				_ = stream.CloseSend(err)
				return
			}
			err = svc.CountUp(ctx, req, typed)
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
		"internal.ping.v1test", // protobuf package
		"PingService",          // protobuf service
		"CumSum",               // protobuf method
		func(ctx context.Context, stream rerpc.Stream) {
			typed := handlerstream.NewBidirectional[CumSumRequest, CumSumResponse](stream)
			err := svc.CumSum(ctx, typed)
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

var _ PingServiceReRPC = (*UnimplementedPingServiceReRPC)(nil) // verify interface implementation

// UnimplementedPingServiceReRPC returns CodeUnimplemented from all methods.
type UnimplementedPingServiceReRPC struct{}

func (UnimplementedPingServiceReRPC) Ping(context.Context, *rerpc.Request[PingRequest]) (*rerpc.Response[PingResponse], error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.Ping isn't implemented")
}

func (UnimplementedPingServiceReRPC) Fail(context.Context, *rerpc.Request[FailRequest]) (*rerpc.Response[FailResponse], error) {
	return nil, rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.Fail isn't implemented")
}

func (UnimplementedPingServiceReRPC) Sum(context.Context, *handlerstream.Client[SumRequest, SumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.Sum isn't implemented")
}

func (UnimplementedPingServiceReRPC) CountUp(context.Context, *rerpc.Request[CountUpRequest], *handlerstream.Server[CountUpResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.CountUp isn't implemented")
}

func (UnimplementedPingServiceReRPC) CumSum(context.Context, *handlerstream.Bidirectional[CumSumRequest, CumSumResponse]) error {
	return rerpc.Errorf(rerpc.CodeUnimplemented, "internal.ping.v1test.PingService.CumSum isn't implemented")
}
