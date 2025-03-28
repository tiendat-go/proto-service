// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: crypto/v1/crypto.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CryptoService_GetServerTime_FullMethodName     = "/crypto.v1.CryptoService/GetServerTime"
	CryptoService_GetDepth_FullMethodName          = "/crypto.v1.CryptoService/GetDepth"
	CryptoService_GetTrades_FullMethodName         = "/crypto.v1.CryptoService/GetTrades"
	CryptoService_GetKlinesBySymbol_FullMethodName = "/crypto.v1.CryptoService/GetKlinesBySymbol"
)

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	GetServerTime(ctx context.Context, in *GetServerTimeRequest, opts ...grpc.CallOption) (*GetServerTimeResponse, error)
	GetDepth(ctx context.Context, in *GetDepthRequest, opts ...grpc.CallOption) (*GetDepthResponse, error)
	GetTrades(ctx context.Context, in *GetTradesRequest, opts ...grpc.CallOption) (*GetTradesResponse, error)
	GetKlinesBySymbol(ctx context.Context, in *GetKlinesBySymbolRequest, opts ...grpc.CallOption) (*GetKlinesBySymbolResponse, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) GetServerTime(ctx context.Context, in *GetServerTimeRequest, opts ...grpc.CallOption) (*GetServerTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServerTimeResponse)
	err := c.cc.Invoke(ctx, CryptoService_GetServerTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetDepth(ctx context.Context, in *GetDepthRequest, opts ...grpc.CallOption) (*GetDepthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDepthResponse)
	err := c.cc.Invoke(ctx, CryptoService_GetDepth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTrades(ctx context.Context, in *GetTradesRequest, opts ...grpc.CallOption) (*GetTradesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTradesResponse)
	err := c.cc.Invoke(ctx, CryptoService_GetTrades_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetKlinesBySymbol(ctx context.Context, in *GetKlinesBySymbolRequest, opts ...grpc.CallOption) (*GetKlinesBySymbolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKlinesBySymbolResponse)
	err := c.cc.Invoke(ctx, CryptoService_GetKlinesBySymbol_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility.
type CryptoServiceServer interface {
	GetServerTime(context.Context, *GetServerTimeRequest) (*GetServerTimeResponse, error)
	GetDepth(context.Context, *GetDepthRequest) (*GetDepthResponse, error)
	GetTrades(context.Context, *GetTradesRequest) (*GetTradesResponse, error)
	GetKlinesBySymbol(context.Context, *GetKlinesBySymbolRequest) (*GetKlinesBySymbolResponse, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCryptoServiceServer struct{}

func (UnimplementedCryptoServiceServer) GetServerTime(context.Context, *GetServerTimeRequest) (*GetServerTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTime not implemented")
}
func (UnimplementedCryptoServiceServer) GetDepth(context.Context, *GetDepthRequest) (*GetDepthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepth not implemented")
}
func (UnimplementedCryptoServiceServer) GetTrades(context.Context, *GetTradesRequest) (*GetTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrades not implemented")
}
func (UnimplementedCryptoServiceServer) GetKlinesBySymbol(context.Context, *GetKlinesBySymbolRequest) (*GetKlinesBySymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKlinesBySymbol not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}
func (UnimplementedCryptoServiceServer) testEmbeddedByValue()                       {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	// If the following call pancis, it indicates UnimplementedCryptoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_GetServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetServerTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetServerTime(ctx, req.(*GetServerTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetDepth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetDepth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetDepth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetDepth(ctx, req.(*GetDepthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetTrades_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTrades(ctx, req.(*GetTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetKlinesBySymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKlinesBySymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetKlinesBySymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetKlinesBySymbol_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetKlinesBySymbol(ctx, req.(*GetKlinesBySymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.v1.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerTime",
			Handler:    _CryptoService_GetServerTime_Handler,
		},
		{
			MethodName: "GetDepth",
			Handler:    _CryptoService_GetDepth_Handler,
		},
		{
			MethodName: "GetTrades",
			Handler:    _CryptoService_GetTrades_Handler,
		},
		{
			MethodName: "GetKlinesBySymbol",
			Handler:    _CryptoService_GetKlinesBySymbol_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto/v1/crypto.proto",
}
