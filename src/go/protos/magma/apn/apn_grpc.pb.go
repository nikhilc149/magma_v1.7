// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apn

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

// ApnDBCloudClient is the client API for ApnDBCloud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApnDBCloudClient interface {
	// ListApnConfigs lists network wide apn configuration.
	ListApnConfigs(ctx context.Context, in *ListApnConfigRequest, opts ...grpc.CallOption) (*ListApnConfigResponse, error)
	// ListGatewayApnConfigs lists gateway specific apn configuration
	ListGatewayApnConfigs(ctx context.Context, in *ListGatewayApnConfigRequest, opts ...grpc.CallOption) (*ListGatewayApnConfigResponse, error)
}

type apnDBCloudClient struct {
	cc grpc.ClientConnInterface
}

func NewApnDBCloudClient(cc grpc.ClientConnInterface) ApnDBCloudClient {
	return &apnDBCloudClient{cc}
}

func (c *apnDBCloudClient) ListApnConfigs(ctx context.Context, in *ListApnConfigRequest, opts ...grpc.CallOption) (*ListApnConfigResponse, error) {
	out := new(ListApnConfigResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.ApnDBCloud/ListApnConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apnDBCloudClient) ListGatewayApnConfigs(ctx context.Context, in *ListGatewayApnConfigRequest, opts ...grpc.CallOption) (*ListGatewayApnConfigResponse, error) {
	out := new(ListGatewayApnConfigResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.ApnDBCloud/ListGatewayApnConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApnDBCloudServer is the server API for ApnDBCloud service.
// All implementations must embed UnimplementedApnDBCloudServer
// for forward compatibility
type ApnDBCloudServer interface {
	// ListApnConfigs lists network wide apn configuration.
	ListApnConfigs(context.Context, *ListApnConfigRequest) (*ListApnConfigResponse, error)
	// ListGatewayApnConfigs lists gateway specific apn configuration
	ListGatewayApnConfigs(context.Context, *ListGatewayApnConfigRequest) (*ListGatewayApnConfigResponse, error)
	mustEmbedUnimplementedApnDBCloudServer()
}

// UnimplementedApnDBCloudServer must be embedded to have forward compatible implementations.
type UnimplementedApnDBCloudServer struct {
}

func (UnimplementedApnDBCloudServer) ListApnConfigs(context.Context, *ListApnConfigRequest) (*ListApnConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApnConfigs not implemented")
}
func (UnimplementedApnDBCloudServer) ListGatewayApnConfigs(context.Context, *ListGatewayApnConfigRequest) (*ListGatewayApnConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGatewayApnConfigs not implemented")
}
func (UnimplementedApnDBCloudServer) mustEmbedUnimplementedApnDBCloudServer() {}

// UnsafeApnDBCloudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApnDBCloudServer will
// result in compilation errors.
type UnsafeApnDBCloudServer interface {
	mustEmbedUnimplementedApnDBCloudServer()
}

func RegisterApnDBCloudServer(s grpc.ServiceRegistrar, srv ApnDBCloudServer) {
	s.RegisterService(&ApnDBCloud_ServiceDesc, srv)
}

func _ApnDBCloud_ListApnConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApnConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApnDBCloudServer).ListApnConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.ApnDBCloud/ListApnConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApnDBCloudServer).ListApnConfigs(ctx, req.(*ListApnConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApnDBCloud_ListGatewayApnConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayApnConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApnDBCloudServer).ListGatewayApnConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.ApnDBCloud/ListGatewayApnConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApnDBCloudServer).ListGatewayApnConfigs(ctx, req.(*ListGatewayApnConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApnDBCloud_ServiceDesc is the grpc.ServiceDesc for ApnDBCloud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApnDBCloud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.ApnDBCloud",
	HandlerType: (*ApnDBCloudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApnConfigs",
			Handler:    _ApnDBCloud_ListApnConfigs_Handler,
		},
		{
			MethodName: "ListGatewayApnConfigs",
			Handler:    _ApnDBCloud_ListGatewayApnConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apn.proto",
}
