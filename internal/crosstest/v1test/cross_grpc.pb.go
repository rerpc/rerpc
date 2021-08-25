// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package crosspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CrossServiceClient is the client API for CrossService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrossServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Fail(ctx context.Context, in *FailRequest, opts ...grpc.CallOption) (*FailResponse, error)
	Sum(ctx context.Context, opts ...grpc.CallOption) (CrossService_SumClient, error)
	CountUp(ctx context.Context, in *CountUpRequest, opts ...grpc.CallOption) (CrossService_CountUpClient, error)
	CumSum(ctx context.Context, opts ...grpc.CallOption) (CrossService_CumSumClient, error)
}

type crossServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrossServiceClient(cc grpc.ClientConnInterface) CrossServiceClient {
	return &crossServiceClient{cc}
}

func (c *crossServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/internal.crosstest.v1test.CrossService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crossServiceClient) Fail(ctx context.Context, in *FailRequest, opts ...grpc.CallOption) (*FailResponse, error) {
	out := new(FailResponse)
	err := c.cc.Invoke(ctx, "/internal.crosstest.v1test.CrossService/Fail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crossServiceClient) Sum(ctx context.Context, opts ...grpc.CallOption) (CrossService_SumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrossService_ServiceDesc.Streams[0], "/internal.crosstest.v1test.CrossService/Sum", opts...)
	if err != nil {
		return nil, err
	}
	x := &crossServiceSumClient{stream}
	return x, nil
}

type CrossService_SumClient interface {
	Send(*SumRequest) error
	CloseAndRecv() (*SumResponse, error)
	grpc.ClientStream
}

type crossServiceSumClient struct {
	grpc.ClientStream
}

func (x *crossServiceSumClient) Send(m *SumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *crossServiceSumClient) CloseAndRecv() (*SumResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crossServiceClient) CountUp(ctx context.Context, in *CountUpRequest, opts ...grpc.CallOption) (CrossService_CountUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrossService_ServiceDesc.Streams[1], "/internal.crosstest.v1test.CrossService/CountUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &crossServiceCountUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrossService_CountUpClient interface {
	Recv() (*CountUpResponse, error)
	grpc.ClientStream
}

type crossServiceCountUpClient struct {
	grpc.ClientStream
}

func (x *crossServiceCountUpClient) Recv() (*CountUpResponse, error) {
	m := new(CountUpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crossServiceClient) CumSum(ctx context.Context, opts ...grpc.CallOption) (CrossService_CumSumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrossService_ServiceDesc.Streams[2], "/internal.crosstest.v1test.CrossService/CumSum", opts...)
	if err != nil {
		return nil, err
	}
	x := &crossServiceCumSumClient{stream}
	return x, nil
}

type CrossService_CumSumClient interface {
	Send(*CumSumRequest) error
	Recv() (*CumSumResponse, error)
	grpc.ClientStream
}

type crossServiceCumSumClient struct {
	grpc.ClientStream
}

func (x *crossServiceCumSumClient) Send(m *CumSumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *crossServiceCumSumClient) Recv() (*CumSumResponse, error) {
	m := new(CumSumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CrossServiceServer is the server API for CrossService service.
// All implementations must embed UnimplementedCrossServiceServer
// for forward compatibility
type CrossServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Fail(context.Context, *FailRequest) (*FailResponse, error)
	Sum(CrossService_SumServer) error
	CountUp(*CountUpRequest, CrossService_CountUpServer) error
	CumSum(CrossService_CumSumServer) error
	mustEmbedUnimplementedCrossServiceServer()
}

// UnimplementedCrossServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrossServiceServer struct {
}

func (UnimplementedCrossServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCrossServiceServer) Fail(context.Context, *FailRequest) (*FailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fail not implemented")
}
func (UnimplementedCrossServiceServer) Sum(CrossService_SumServer) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCrossServiceServer) CountUp(*CountUpRequest, CrossService_CountUpServer) error {
	return status.Errorf(codes.Unimplemented, "method CountUp not implemented")
}
func (UnimplementedCrossServiceServer) CumSum(CrossService_CumSumServer) error {
	return status.Errorf(codes.Unimplemented, "method CumSum not implemented")
}
func (UnimplementedCrossServiceServer) mustEmbedUnimplementedCrossServiceServer() {}

// UnsafeCrossServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrossServiceServer will
// result in compilation errors.
type UnsafeCrossServiceServer interface {
	mustEmbedUnimplementedCrossServiceServer()
}

func RegisterCrossServiceServer(s grpc.ServiceRegistrar, srv CrossServiceServer) {
	s.RegisterService(&CrossService_ServiceDesc, srv)
}

func _CrossService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrossServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.crosstest.v1test.CrossService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrossServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrossService_Fail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrossServiceServer).Fail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.crosstest.v1test.CrossService/Fail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrossServiceServer).Fail(ctx, req.(*FailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrossService_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CrossServiceServer).Sum(&crossServiceSumServer{stream})
}

type CrossService_SumServer interface {
	SendAndClose(*SumResponse) error
	Recv() (*SumRequest, error)
	grpc.ServerStream
}

type crossServiceSumServer struct {
	grpc.ServerStream
}

func (x *crossServiceSumServer) SendAndClose(m *SumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *crossServiceSumServer) Recv() (*SumRequest, error) {
	m := new(SumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CrossService_CountUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CountUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrossServiceServer).CountUp(m, &crossServiceCountUpServer{stream})
}

type CrossService_CountUpServer interface {
	Send(*CountUpResponse) error
	grpc.ServerStream
}

type crossServiceCountUpServer struct {
	grpc.ServerStream
}

func (x *crossServiceCountUpServer) Send(m *CountUpResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrossService_CumSum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CrossServiceServer).CumSum(&crossServiceCumSumServer{stream})
}

type CrossService_CumSumServer interface {
	Send(*CumSumResponse) error
	Recv() (*CumSumRequest, error)
	grpc.ServerStream
}

type crossServiceCumSumServer struct {
	grpc.ServerStream
}

func (x *crossServiceCumSumServer) Send(m *CumSumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *crossServiceCumSumServer) Recv() (*CumSumRequest, error) {
	m := new(CumSumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CrossService_ServiceDesc is the grpc.ServiceDesc for CrossService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrossService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.crosstest.v1test.CrossService",
	HandlerType: (*CrossServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CrossService_Ping_Handler,
		},
		{
			MethodName: "Fail",
			Handler:    _CrossService_Fail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sum",
			Handler:       _CrossService_Sum_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CountUp",
			Handler:       _CrossService_CountUp_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CumSum",
			Handler:       _CrossService_CumSum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/crosstest/v1test/cross.proto",
}