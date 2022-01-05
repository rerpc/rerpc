package rerpc_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rerpc/rerpc"
	pingpb "github.com/rerpc/rerpc/internal/ping/v1test"
)

func ExampleCallMetadata() {
	logger := rerpc.UnaryInterceptorFunc(func(next rerpc.Func) rerpc.Func {
		return rerpc.Func(func(ctx context.Context, req rerpc.AnyRequest) (rerpc.AnyResponse, error) {
			fmt.Println("calling", req.Spec().Procedure)
			return next(ctx, req)
		})
	})
	// This interceptor prevents the client from making network requests in
	// examples. Leave it out in real code!
	short := ShortCircuit(rerpc.Errorf(rerpc.CodeUnimplemented, "no networking in examples"))
	client := pingpb.NewPingServiceClientReRPC(
		"https://invalid-test-url",
		http.DefaultClient,
		rerpc.Intercept(rerpc.NewChain(logger, short)),
	)
	client.Ping(context.Background(), rerpc.NewRequest(&pingpb.PingRequest{}))

	// Output:
	// calling internal.ping.v1test.PingService/Ping
}

func ExampleChain() {
	outer := rerpc.UnaryInterceptorFunc(func(next rerpc.Func) rerpc.Func {
		return rerpc.Func(func(ctx context.Context, req rerpc.AnyRequest) (rerpc.AnyResponse, error) {
			fmt.Println("outer interceptor: before call")
			res, err := next(ctx, req)
			fmt.Println("outer interceptor: after call")
			return res, err
		})
	})
	inner := rerpc.UnaryInterceptorFunc(func(next rerpc.Func) rerpc.Func {
		return rerpc.Func(func(ctx context.Context, req rerpc.AnyRequest) (rerpc.AnyResponse, error) {
			fmt.Println("inner interceptor: before call")
			res, err := next(ctx, req)
			fmt.Println("inner interceptor: after call")
			return res, err
		})
	})
	// This interceptor prevents the client from making network requests in
	// examples. Leave it out in real code!
	short := ShortCircuit(rerpc.Errorf(rerpc.CodeUnimplemented, "no networking in examples"))
	client := pingpb.NewPingServiceClientReRPC(
		"https://invalid-test-url",
		http.DefaultClient,
		rerpc.Intercept(rerpc.NewChain(outer, inner, short)),
	)
	client.Ping(context.Background(), rerpc.NewRequest(&pingpb.PingRequest{}))

	// Output:
	// outer interceptor: before call
	// inner interceptor: before call
	// inner interceptor: after call
	// outer interceptor: after call
}
