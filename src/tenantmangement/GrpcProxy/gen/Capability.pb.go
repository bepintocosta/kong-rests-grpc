// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.6.1
// source: Capability.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Create Capability
type CreateCapabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode string `protobuf:"bytes,1,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCapabilityRequest) Reset() {
	*x = CreateCapabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Capability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCapabilityRequest) ProtoMessage() {}

func (x *CreateCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Capability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCapabilityRequest.ProtoReflect.Descriptor instead.
func (*CreateCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_Capability_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCapabilityRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *CreateCapabilityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCapabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateCapabilityResponse) Reset() {
	*x = CreateCapabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Capability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCapabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCapabilityResponse) ProtoMessage() {}

func (x *CreateCapabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Capability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCapabilityResponse.ProtoReflect.Descriptor instead.
func (*CreateCapabilityResponse) Descriptor() ([]byte, []int) {
	return file_Capability_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCapabilityResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// List Capabilities By Tenant
type ListCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode string `protobuf:"bytes,1,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
	Index      int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Size       int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Sort       string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Order      string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ListCapabilitiesRequest) Reset() {
	*x = ListCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Capability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapabilitiesRequest) ProtoMessage() {}

func (x *ListCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Capability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*ListCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_Capability_proto_rawDescGZIP(), []int{2}
}

func (x *ListCapabilitiesRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ListCapabilitiesRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListCapabilitiesRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type ListCapabilitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*CapabilitieResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalCount int32                  `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListCapabilitiesResponse) Reset() {
	*x = ListCapabilitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Capability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapabilitiesResponse) ProtoMessage() {}

func (x *ListCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Capability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapabilitiesResponse.ProtoReflect.Descriptor instead.
func (*ListCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return file_Capability_proto_rawDescGZIP(), []int{3}
}

func (x *ListCapabilitiesResponse) GetData() []*CapabilitieResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListCapabilitiesResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CapabilitieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CapabilitieResponse) Reset() {
	*x = CapabilitieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Capability_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilitieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilitieResponse) ProtoMessage() {}

func (x *CapabilitieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Capability_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilitieResponse.ProtoReflect.Descriptor instead.
func (*CapabilitieResponse) Descriptor() ([]byte, []int) {
	return file_Capability_proto_rawDescGZIP(), []int{4}
}

func (x *CapabilitieResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CapabilitieResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_Capability_proto protoreflect.FileDescriptor

var file_Capability_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x8d, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x74, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xde, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x67, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x28, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2a, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x67, 0x65, 0x6e, 0xaa, 0x02, 0x1f, 0x6e,
	0x65, 0x74, 0x70, 0x6f, 0x6e, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Capability_proto_rawDescOnce sync.Once
	file_Capability_proto_rawDescData = file_Capability_proto_rawDesc
)

func file_Capability_proto_rawDescGZIP() []byte {
	file_Capability_proto_rawDescOnce.Do(func() {
		file_Capability_proto_rawDescData = protoimpl.X.CompressGZIP(file_Capability_proto_rawDescData)
	})
	return file_Capability_proto_rawDescData
}

var file_Capability_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Capability_proto_goTypes = []interface{}{
	(*CreateCapabilityRequest)(nil),  // 0: AssetManagement.CreateCapabilityRequest
	(*CreateCapabilityResponse)(nil), // 1: AssetManagement.CreateCapabilityResponse
	(*ListCapabilitiesRequest)(nil),  // 2: AssetManagement.ListCapabilitiesRequest
	(*ListCapabilitiesResponse)(nil), // 3: AssetManagement.ListCapabilitiesResponse
	(*CapabilitieResponse)(nil),      // 4: AssetManagement.CapabilitieResponse
}
var file_Capability_proto_depIdxs = []int32{
	4, // 0: AssetManagement.ListCapabilitiesResponse.data:type_name -> AssetManagement.CapabilitieResponse
	0, // 1: AssetManagement.Capability.CreateCapability:input_type -> AssetManagement.CreateCapabilityRequest
	2, // 2: AssetManagement.Capability.ListCapabilities:input_type -> AssetManagement.ListCapabilitiesRequest
	1, // 3: AssetManagement.Capability.CreateCapability:output_type -> AssetManagement.CreateCapabilityResponse
	3, // 4: AssetManagement.Capability.ListCapabilities:output_type -> AssetManagement.ListCapabilitiesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Capability_proto_init() }
func file_Capability_proto_init() {
	if File_Capability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Capability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCapabilityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Capability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCapabilityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Capability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCapabilitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Capability_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCapabilitiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Capability_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilitieResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Capability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Capability_proto_goTypes,
		DependencyIndexes: file_Capability_proto_depIdxs,
		MessageInfos:      file_Capability_proto_msgTypes,
	}.Build()
	File_Capability_proto = out.File
	file_Capability_proto_rawDesc = nil
	file_Capability_proto_goTypes = nil
	file_Capability_proto_depIdxs = nil
}