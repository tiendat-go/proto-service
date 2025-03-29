// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: registry/v1/registry.proto

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
	DiscoveryService_RegisterService_FullMethodName   = "/registry.v1.DiscoveryService/RegisterService"
	DiscoveryService_DeregisterService_FullMethodName = "/registry.v1.DiscoveryService/DeregisterService"
	DiscoveryService_Heartbeat_FullMethodName         = "/registry.v1.DiscoveryService/Heartbeat"
	DiscoveryService_GetServices_FullMethodName       = "/registry.v1.DiscoveryService/GetServices"
	DiscoveryService_GetRandService_FullMethodName    = "/registry.v1.DiscoveryService/GetRandService"
)

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error)
	DeregisterService(ctx context.Context, in *DeregisterServiceRequest, opts ...grpc.CallOption) (*DeregisterServiceResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetRandService(ctx context.Context, in *GetRandServiceRequest, opts ...grpc.CallOption) (*GetRandServiceResponse, error)
}

type discoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryServiceClient(cc grpc.ClientConnInterface) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterServiceResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_RegisterService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) DeregisterService(ctx context.Context, in *DeregisterServiceRequest, opts ...grpc.CallOption) (*DeregisterServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeregisterServiceResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_DeregisterService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_Heartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_GetServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) GetRandService(ctx context.Context, in *GetRandServiceRequest, opts ...grpc.CallOption) (*GetRandServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRandServiceResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_GetRandService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
// All implementations must embed UnimplementedDiscoveryServiceServer
// for forward compatibility.
type DiscoveryServiceServer interface {
	RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error)
	DeregisterService(context.Context, *DeregisterServiceRequest) (*DeregisterServiceResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	GetRandService(context.Context, *GetRandServiceRequest) (*GetRandServiceResponse, error)
	mustEmbedUnimplementedDiscoveryServiceServer()
}

// UnimplementedDiscoveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDiscoveryServiceServer struct{}

func (UnimplementedDiscoveryServiceServer) RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedDiscoveryServiceServer) DeregisterService(context.Context, *DeregisterServiceRequest) (*DeregisterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterService not implemented")
}
func (UnimplementedDiscoveryServiceServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedDiscoveryServiceServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedDiscoveryServiceServer) GetRandService(context.Context, *GetRandServiceRequest) (*GetRandServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandService not implemented")
}
func (UnimplementedDiscoveryServiceServer) mustEmbedUnimplementedDiscoveryServiceServer() {}
func (UnimplementedDiscoveryServiceServer) testEmbeddedByValue()                          {}

// UnsafeDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryServiceServer will
// result in compilation errors.
type UnsafeDiscoveryServiceServer interface {
	mustEmbedUnimplementedDiscoveryServiceServer()
}

func RegisterDiscoveryServiceServer(s grpc.ServiceRegistrar, srv DiscoveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedDiscoveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DiscoveryService_ServiceDesc, srv)
}

func _DiscoveryService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_RegisterService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).RegisterService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_DeregisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).DeregisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_DeregisterService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).DeregisterService(ctx, req.(*DeregisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_Heartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_GetServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_GetRandService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).GetRandService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_GetRandService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).GetRandService(ctx, req.(*GetRandServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscoveryService_ServiceDesc is the grpc.ServiceDesc for DiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.v1.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _DiscoveryService_RegisterService_Handler,
		},
		{
			MethodName: "DeregisterService",
			Handler:    _DiscoveryService_DeregisterService_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _DiscoveryService_Heartbeat_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _DiscoveryService_GetServices_Handler,
		},
		{
			MethodName: "GetRandService",
			Handler:    _DiscoveryService_GetRandService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1/registry.proto",
}
