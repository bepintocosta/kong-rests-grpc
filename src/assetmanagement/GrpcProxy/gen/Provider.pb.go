// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.6.1
// source: Provider.proto

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

// Create Provider
type CreateProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode string `protobuf:"bytes,1,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	QueueName  string `protobuf:"bytes,3,opt,name=queueName,proto3" json:"queueName,omitempty"`
}

func (x *CreateProviderRequest) Reset() {
	*x = CreateProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProviderRequest) ProtoMessage() {}

func (x *CreateProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProviderRequest.ProtoReflect.Descriptor instead.
func (*CreateProviderRequest) Descriptor() ([]byte, []int) {
	return file_Provider_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProviderRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *CreateProviderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProviderRequest) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

type CreateProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateProviderResponse) Reset() {
	*x = CreateProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProviderResponse) ProtoMessage() {}

func (x *CreateProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProviderResponse.ProtoReflect.Descriptor instead.
func (*CreateProviderResponse) Descriptor() ([]byte, []int) {
	return file_Provider_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProviderResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// List Providers
type ListProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode string `protobuf:"bytes,1,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
	Index      int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Size       int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Sort       string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Order      string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ListProvidersRequest) Reset() {
	*x = ListProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvidersRequest) ProtoMessage() {}

func (x *ListProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListProvidersRequest) Descriptor() ([]byte, []int) {
	return file_Provider_proto_rawDescGZIP(), []int{2}
}

func (x *ListProvidersRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *ListProvidersRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ListProvidersRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListProvidersRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListProvidersRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type ListProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*ProviderResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalCount int32               `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListProvidersResponse) Reset() {
	*x = ListProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvidersResponse) ProtoMessage() {}

func (x *ListProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvidersResponse.ProtoReflect.Descriptor instead.
func (*ListProvidersResponse) Descriptor() ([]byte, []int) {
	return file_Provider_proto_rawDescGZIP(), []int{3}
}

func (x *ListProvidersResponse) GetData() []*ProviderResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListProvidersResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProviderResponse) Reset() {
	*x = ProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderResponse) ProtoMessage() {}

func (x *ProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderResponse.ProtoReflect.Descriptor instead.
func (*ProviderResponse) Descriptor() ([]byte, []int) {
	return file_Provider_proto_rawDescGZIP(), []int{4}
}

func (x *ProviderResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProviderResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_Provider_proto protoreflect.FileDescriptor

var file_Provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x69, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xcd, 0x01,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a,
	0x06, 0x2e, 0x2f, 0x3b, 0x67, 0x65, 0x6e, 0xaa, 0x02, 0x1f, 0x6e, 0x65, 0x74, 0x70, 0x6f, 0x6e,
	0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Provider_proto_rawDescOnce sync.Once
	file_Provider_proto_rawDescData = file_Provider_proto_rawDesc
)

func file_Provider_proto_rawDescGZIP() []byte {
	file_Provider_proto_rawDescOnce.Do(func() {
		file_Provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_Provider_proto_rawDescData)
	})
	return file_Provider_proto_rawDescData
}

var file_Provider_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Provider_proto_goTypes = []interface{}{
	(*CreateProviderRequest)(nil),  // 0: AssetManagement.CreateProviderRequest
	(*CreateProviderResponse)(nil), // 1: AssetManagement.CreateProviderResponse
	(*ListProvidersRequest)(nil),   // 2: AssetManagement.ListProvidersRequest
	(*ListProvidersResponse)(nil),  // 3: AssetManagement.ListProvidersResponse
	(*ProviderResponse)(nil),       // 4: AssetManagement.ProviderResponse
}
var file_Provider_proto_depIdxs = []int32{
	4, // 0: AssetManagement.ListProvidersResponse.data:type_name -> AssetManagement.ProviderResponse
	0, // 1: AssetManagement.Provider.CreateProvider:input_type -> AssetManagement.CreateProviderRequest
	2, // 2: AssetManagement.Provider.ListProviders:input_type -> AssetManagement.ListProvidersRequest
	1, // 3: AssetManagement.Provider.CreateProvider:output_type -> AssetManagement.CreateProviderResponse
	3, // 4: AssetManagement.Provider.ListProviders:output_type -> AssetManagement.ListProvidersResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Provider_proto_init() }
func file_Provider_proto_init() {
	if File_Provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProviderRequest); i {
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
		file_Provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProviderResponse); i {
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
		file_Provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProvidersRequest); i {
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
		file_Provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProvidersResponse); i {
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
		file_Provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderResponse); i {
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
			RawDescriptor: file_Provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Provider_proto_goTypes,
		DependencyIndexes: file_Provider_proto_depIdxs,
		MessageInfos:      file_Provider_proto_msgTypes,
	}.Build()
	File_Provider_proto = out.File
	file_Provider_proto_rawDesc = nil
	file_Provider_proto_goTypes = nil
	file_Provider_proto_depIdxs = nil
}
