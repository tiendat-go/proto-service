// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: registry/v1/registry.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceName   string                 `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterServiceRequest) Reset() {
	*x = RegisterServiceRequest{}
	mi := &file_registry_v1_registry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceRequest) ProtoMessage() {}

func (x *RegisterServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_registry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceRequest.ProtoReflect.Descriptor instead.
func (*RegisterServiceRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1_registry_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterServiceRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RegisterServiceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterServiceResponse) Reset() {
	*x = RegisterServiceResponse{}
	mi := &file_registry_v1_registry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceResponse) ProtoMessage() {}

func (x *RegisterServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_registry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceResponse.ProtoReflect.Descriptor instead.
func (*RegisterServiceResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterServiceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetServicesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceName   string                 `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServicesRequest) Reset() {
	*x = GetServicesRequest{}
	mi := &file_registry_v1_registry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesRequest) ProtoMessage() {}

func (x *GetServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_registry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesRequest.ProtoReflect.Descriptor instead.
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1_registry_proto_rawDescGZIP(), []int{2}
}

func (x *GetServicesRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GetServicesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []string               `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServicesResponse) Reset() {
	*x = GetServicesResponse{}
	mi := &file_registry_v1_registry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesResponse) ProtoMessage() {}

func (x *GetServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_registry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesResponse.ProtoReflect.Descriptor instead.
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1_registry_proto_rawDescGZIP(), []int{3}
}

func (x *GetServicesResponse) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_registry_v1_registry_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_registry_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1_registry_proto_rawDescGZIP(), []int{4}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Healthy       bool                   `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_registry_v1_registry_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_registry_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1_registry_proto_rawDescGZIP(), []int{5}
}

func (x *HealthCheckResponse) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

var File_registry_v1_registry_proto protoreflect.FileDescriptor

const file_registry_v1_registry_proto_rawDesc = "" +
	"\n" +
	"\x1aregistry/v1/registry.proto\x12\vregistry.v1\"U\n" +
	"\x16RegisterServiceRequest\x12!\n" +
	"\fservice_name\x18\x01 \x01(\tR\vserviceName\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\"3\n" +
	"\x17RegisterServiceResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"7\n" +
	"\x12GetServicesRequest\x12!\n" +
	"\fservice_name\x18\x01 \x01(\tR\vserviceName\"3\n" +
	"\x13GetServicesResponse\x12\x1c\n" +
	"\taddresses\x18\x01 \x03(\tR\taddresses\"\x14\n" +
	"\x12HealthCheckRequest\"/\n" +
	"\x13HealthCheckResponse\x12\x18\n" +
	"\ahealthy\x18\x01 \x01(\bR\ahealthy2\x94\x02\n" +
	"\x10DiscoveryService\x12\\\n" +
	"\x0fRegisterService\x12#.registry.v1.RegisterServiceRequest\x1a$.registry.v1.RegisterServiceResponse\x12P\n" +
	"\vGetServices\x12\x1f.registry.v1.GetServicesRequest\x1a .registry.v1.GetServicesResponse\x12P\n" +
	"\vHealthCheck\x12\x1f.registry.v1.HealthCheckRequest\x1a .registry.v1.HealthCheckResponseB\x82\x01\n" +
	"\x0fcom.registry.v1B\rRegistryProtoP\x01Z\x13/gen/go/registry/v1\xa2\x02\x03RXX\xaa\x02\vRegistry.V1\xca\x02\vRegistry\\V1\xe2\x02\x17Registry\\V1\\GPBMetadata\xea\x02\fRegistry::V1b\x06proto3"

var (
	file_registry_v1_registry_proto_rawDescOnce sync.Once
	file_registry_v1_registry_proto_rawDescData []byte
)

func file_registry_v1_registry_proto_rawDescGZIP() []byte {
	file_registry_v1_registry_proto_rawDescOnce.Do(func() {
		file_registry_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_registry_v1_registry_proto_rawDesc), len(file_registry_v1_registry_proto_rawDesc)))
	})
	return file_registry_v1_registry_proto_rawDescData
}

var file_registry_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_registry_v1_registry_proto_goTypes = []any{
	(*RegisterServiceRequest)(nil),  // 0: registry.v1.RegisterServiceRequest
	(*RegisterServiceResponse)(nil), // 1: registry.v1.RegisterServiceResponse
	(*GetServicesRequest)(nil),      // 2: registry.v1.GetServicesRequest
	(*GetServicesResponse)(nil),     // 3: registry.v1.GetServicesResponse
	(*HealthCheckRequest)(nil),      // 4: registry.v1.HealthCheckRequest
	(*HealthCheckResponse)(nil),     // 5: registry.v1.HealthCheckResponse
}
var file_registry_v1_registry_proto_depIdxs = []int32{
	0, // 0: registry.v1.DiscoveryService.RegisterService:input_type -> registry.v1.RegisterServiceRequest
	2, // 1: registry.v1.DiscoveryService.GetServices:input_type -> registry.v1.GetServicesRequest
	4, // 2: registry.v1.DiscoveryService.HealthCheck:input_type -> registry.v1.HealthCheckRequest
	1, // 3: registry.v1.DiscoveryService.RegisterService:output_type -> registry.v1.RegisterServiceResponse
	3, // 4: registry.v1.DiscoveryService.GetServices:output_type -> registry.v1.GetServicesResponse
	5, // 5: registry.v1.DiscoveryService.HealthCheck:output_type -> registry.v1.HealthCheckResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_registry_v1_registry_proto_init() }
func file_registry_v1_registry_proto_init() {
	if File_registry_v1_registry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_registry_v1_registry_proto_rawDesc), len(file_registry_v1_registry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_v1_registry_proto_goTypes,
		DependencyIndexes: file_registry_v1_registry_proto_depIdxs,
		MessageInfos:      file_registry_v1_registry_proto_msgTypes,
	}.Build()
	File_registry_v1_registry_proto = out.File
	file_registry_v1_registry_proto_goTypes = nil
	file_registry_v1_registry_proto_depIdxs = nil
}
