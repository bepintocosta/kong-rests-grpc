// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: Capability.proto

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

// CapabilityClient is the client API for Capability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CapabilityClient interface {
	CreateCapability(ctx context.Context, in *CreateCapabilityRequest, opts ...grpc.CallOption) (*CreateCapabilityResponse, error)
	UpdateCapability(ctx context.Context, in *UpdateCapabilityRequest, opts ...grpc.CallOption) (*UpdateCapabilityResponse, error)
	CapabilityDetail(ctx context.Context, in *CapabilityDetailRequest, opts ...grpc.CallOption) (*CapabilityDetailResponse, error)
	ListCapabilities(ctx context.Context, in *ListCapabilitiesRequest, opts ...grpc.CallOption) (*ListCapabilitiesResponse, error)
	ListCapabilitiesByTenant(ctx context.Context, in *ListCapabilitiesByTenantRequest, opts ...grpc.CallOption) (*ListCapabilitiesByTenantResponse, error)
}

type capabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewCapabilityClient(cc grpc.ClientConnInterface) CapabilityClient {
	return &capabilityClient{cc}
}

func (c *capabilityClient) CreateCapability(ctx context.Context, in *CreateCapabilityRequest, opts ...grpc.CallOption) (*CreateCapabilityResponse, error) {
	out := new(CreateCapabilityResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Capability/CreateCapability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilityClient) UpdateCapability(ctx context.Context, in *UpdateCapabilityRequest, opts ...grpc.CallOption) (*UpdateCapabilityResponse, error) {
	out := new(UpdateCapabilityResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Capability/UpdateCapability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilityClient) CapabilityDetail(ctx context.Context, in *CapabilityDetailRequest, opts ...grpc.CallOption) (*CapabilityDetailResponse, error) {
	out := new(CapabilityDetailResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Capability/CapabilityDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilityClient) ListCapabilities(ctx context.Context, in *ListCapabilitiesRequest, opts ...grpc.CallOption) (*ListCapabilitiesResponse, error) {
	out := new(ListCapabilitiesResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Capability/ListCapabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilityClient) ListCapabilitiesByTenant(ctx context.Context, in *ListCapabilitiesByTenantRequest, opts ...grpc.CallOption) (*ListCapabilitiesByTenantResponse, error) {
	out := new(ListCapabilitiesByTenantResponse)
	err := c.cc.Invoke(ctx, "/AssetManagement.Capability/ListCapabilitiesByTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CapabilityServer is the server API for Capability service.
// All implementations must embed UnimplementedCapabilityServer
// for forward compatibility
type CapabilityServer interface {
	CreateCapability(context.Context, *CreateCapabilityRequest) (*CreateCapabilityResponse, error)
	UpdateCapability(context.Context, *UpdateCapabilityRequest) (*UpdateCapabilityResponse, error)
	CapabilityDetail(context.Context, *CapabilityDetailRequest) (*CapabilityDetailResponse, error)
	ListCapabilities(context.Context, *ListCapabilitiesRequest) (*ListCapabilitiesResponse, error)
	ListCapabilitiesByTenant(context.Context, *ListCapabilitiesByTenantRequest) (*ListCapabilitiesByTenantResponse, error)
	mustEmbedUnimplementedCapabilityServer()
}

// UnimplementedCapabilityServer must be embedded to have forward compatible implementations.
type UnimplementedCapabilityServer struct {
}

func (UnimplementedCapabilityServer) CreateCapability(context.Context, *CreateCapabilityRequest) (*CreateCapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCapability not implemented")
}
func (UnimplementedCapabilityServer) UpdateCapability(context.Context, *UpdateCapabilityRequest) (*UpdateCapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCapability not implemented")
}
func (UnimplementedCapabilityServer) CapabilityDetail(context.Context, *CapabilityDetailRequest) (*CapabilityDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CapabilityDetail not implemented")
}
func (UnimplementedCapabilityServer) ListCapabilities(context.Context, *ListCapabilitiesRequest) (*ListCapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCapabilities not implemented")
}
func (UnimplementedCapabilityServer) ListCapabilitiesByTenant(context.Context, *ListCapabilitiesByTenantRequest) (*ListCapabilitiesByTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCapabilitiesByTenant not implemented")
}
func (UnimplementedCapabilityServer) mustEmbedUnimplementedCapabilityServer() {}

// UnsafeCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CapabilityServer will
// result in compilation errors.
type UnsafeCapabilityServer interface {
	mustEmbedUnimplementedCapabilityServer()
}

func RegisterCapabilityServer(s grpc.ServiceRegistrar, srv CapabilityServer) {
	s.RegisterService(&Capability_ServiceDesc, srv)
}

func _Capability_CreateCapability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilityServer).CreateCapability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Capability/CreateCapability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilityServer).CreateCapability(ctx, req.(*CreateCapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capability_UpdateCapability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilityServer).UpdateCapability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Capability/UpdateCapability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilityServer).UpdateCapability(ctx, req.(*UpdateCapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capability_CapabilityDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilityDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilityServer).CapabilityDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Capability/CapabilityDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilityServer).CapabilityDetail(ctx, req.(*CapabilityDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capability_ListCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilityServer).ListCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Capability/ListCapabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilityServer).ListCapabilities(ctx, req.(*ListCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capability_ListCapabilitiesByTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCapabilitiesByTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilityServer).ListCapabilitiesByTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetManagement.Capability/ListCapabilitiesByTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilityServer).ListCapabilitiesByTenant(ctx, req.(*ListCapabilitiesByTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Capability_ServiceDesc is the grpc.ServiceDesc for Capability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Capability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AssetManagement.Capability",
	HandlerType: (*CapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCapability",
			Handler:    _Capability_CreateCapability_Handler,
		},
		{
			MethodName: "UpdateCapability",
			Handler:    _Capability_UpdateCapability_Handler,
		},
		{
			MethodName: "CapabilityDetail",
			Handler:    _Capability_CapabilityDetail_Handler,
		},
		{
			MethodName: "ListCapabilities",
			Handler:    _Capability_ListCapabilities_Handler,
		},
		{
			MethodName: "ListCapabilitiesByTenant",
			Handler:    _Capability_ListCapabilitiesByTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Capability.proto",
}
