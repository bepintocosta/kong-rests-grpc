// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: Provider.proto

package gen

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

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderResponse, error)
	UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*UpdateProviderResponse, error)
	ProviderDetail(ctx context.Context, in *ProviderDetailRequest, opts ...grpc.CallOption) (*ProviderDetailResponse, error)
	ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error)
	ListProvidersbyTenantAndCapability(ctx context.Context, in *ListProvidersbyTenantAndCapabilityRequest, opts ...grpc.CallOption) (*ListProvidersbyTenantAndCapabilityResponse, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderResponse, error) {
	out := new(CreateProviderResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Provider/CreateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*UpdateProviderResponse, error) {
	out := new(UpdateProviderResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Provider/UpdateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ProviderDetail(ctx context.Context, in *ProviderDetailRequest, opts ...grpc.CallOption) (*ProviderDetailResponse, error) {
	out := new(ProviderDetailResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Provider/ProviderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error) {
	out := new(ListProvidersResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Provider/ListProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ListProvidersbyTenantAndCapability(ctx context.Context, in *ListProvidersbyTenantAndCapabilityRequest, opts ...grpc.CallOption) (*ListProvidersbyTenantAndCapabilityResponse, error) {
	out := new(ListProvidersbyTenantAndCapabilityResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Provider/ListProvidersbyTenantAndCapability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderResponse, error)
	UpdateProvider(context.Context, *UpdateProviderRequest) (*UpdateProviderResponse, error)
	ProviderDetail(context.Context, *ProviderDetailRequest) (*ProviderDetailResponse, error)
	ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error)
	ListProvidersbyTenantAndCapability(context.Context, *ListProvidersbyTenantAndCapabilityRequest) (*ListProvidersbyTenantAndCapabilityResponse, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedProviderServer) UpdateProvider(context.Context, *UpdateProviderRequest) (*UpdateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvider not implemented")
}
func (UnimplementedProviderServer) ProviderDetail(context.Context, *ProviderDetailRequest) (*ProviderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProviderDetail not implemented")
}
func (UnimplementedProviderServer) ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedProviderServer) ListProvidersbyTenantAndCapability(context.Context, *ListProvidersbyTenantAndCapabilityRequest) (*ListProvidersbyTenantAndCapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProvidersbyTenantAndCapability not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Provider/CreateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).CreateProvider(ctx, req.(*CreateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_UpdateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).UpdateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Provider/UpdateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).UpdateProvider(ctx, req.(*UpdateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ProviderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ProviderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Provider/ProviderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ProviderDetail(ctx, req.(*ProviderDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Provider/ListProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ListProviders(ctx, req.(*ListProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ListProvidersbyTenantAndCapability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvidersbyTenantAndCapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ListProvidersbyTenantAndCapability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Provider/ListProvidersbyTenantAndCapability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ListProvidersbyTenantAndCapability(ctx, req.(*ListProvidersbyTenantAndCapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AssetManagement.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProvider",
			Handler:    _Provider_CreateProvider_Handler,
		},
		{
			MethodName: "UpdateProvider",
			Handler:    _Provider_UpdateProvider_Handler,
		},
		{
			MethodName: "ProviderDetail",
			Handler:    _Provider_ProviderDetail_Handler,
		},
		{
			MethodName: "ListProviders",
			Handler:    _Provider_ListProviders_Handler,
		},
		{
			MethodName: "ListProvidersbyTenantAndCapability",
			Handler:    _Provider_ListProvidersbyTenantAndCapability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Provider.proto",
}
