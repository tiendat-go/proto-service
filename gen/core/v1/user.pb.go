// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: core/v1/user.proto

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

type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_EC_UNSPECIFIED  ErrorCode = 0
	ErrorCode_ERROR_CODE_SUCCESS         ErrorCode = 1
	ErrorCode_ERROR_CODE_INVALID_REQUEST ErrorCode = 2
	ErrorCode_ERROR_CODE_DUPLICATE_USER  ErrorCode = 3
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_CODE_EC_UNSPECIFIED",
		1: "ERROR_CODE_SUCCESS",
		2: "ERROR_CODE_INVALID_REQUEST",
		3: "ERROR_CODE_DUPLICATE_USER",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_EC_UNSPECIFIED":  0,
		"ERROR_CODE_SUCCESS":         1,
		"ERROR_CODE_INVALID_REQUEST": 2,
		"ERROR_CODE_DUPLICATE_USER":  3,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_user_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_core_v1_user_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_user_proto_rawDescGZIP(), []int{0}
}

type SignUpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	mi := &file_core_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorCode     ErrorCode              `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=core.v1.ErrorCode" json:"error_code,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	DisplayName   string                 `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	mi := &file_core_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SignUpResponse) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_ERROR_CODE_EC_UNSPECIFIED
}

func (x *SignUpResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *SignUpResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_core_v1_user_proto protoreflect.FileDescriptor

const file_core_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x12core/v1/user.proto\x12\acore.v1\"H\n" +
	"\rSignUpRequest\x12\x1b\n" +
	"\tuser_name\x18\x01 \x01(\tR\buserName\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\xa3\x01\n" +
	"\x0eSignUpResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x121\n" +
	"\n" +
	"error_code\x18\x02 \x01(\x0e2\x12.core.v1.ErrorCodeR\terrorCode\x12#\n" +
	"\rerror_message\x18\x03 \x01(\tR\ferrorMessage\x12!\n" +
	"\fdisplay_name\x18\x04 \x01(\tR\vdisplayName*\x81\x01\n" +
	"\tErrorCode\x12\x1d\n" +
	"\x19ERROR_CODE_EC_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12ERROR_CODE_SUCCESS\x10\x01\x12\x1e\n" +
	"\x1aERROR_CODE_INVALID_REQUEST\x10\x02\x12\x1d\n" +
	"\x19ERROR_CODE_DUPLICATE_USER\x10\x032H\n" +
	"\vUserService\x129\n" +
	"\x06SignUp\x12\x16.core.v1.SignUpRequest\x1a\x17.core.v1.SignUpResponseBf\n" +
	"\vcom.core.v1B\tUserProtoP\x01Z\x0f/gen/go/core/v1\xa2\x02\x03CXX\xaa\x02\aCore.V1\xca\x02\aCore\\V1\xe2\x02\x13Core\\V1\\GPBMetadata\xea\x02\bCore::V1b\x06proto3"

var (
	file_core_v1_user_proto_rawDescOnce sync.Once
	file_core_v1_user_proto_rawDescData []byte
)

func file_core_v1_user_proto_rawDescGZIP() []byte {
	file_core_v1_user_proto_rawDescOnce.Do(func() {
		file_core_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_core_v1_user_proto_rawDesc), len(file_core_v1_user_proto_rawDesc)))
	})
	return file_core_v1_user_proto_rawDescData
}

var file_core_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_v1_user_proto_goTypes = []any{
	(ErrorCode)(0),         // 0: core.v1.ErrorCode
	(*SignUpRequest)(nil),  // 1: core.v1.SignUpRequest
	(*SignUpResponse)(nil), // 2: core.v1.SignUpResponse
}
var file_core_v1_user_proto_depIdxs = []int32{
	0, // 0: core.v1.SignUpResponse.error_code:type_name -> core.v1.ErrorCode
	1, // 1: core.v1.UserService.SignUp:input_type -> core.v1.SignUpRequest
	2, // 2: core.v1.UserService.SignUp:output_type -> core.v1.SignUpResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_v1_user_proto_init() }
func file_core_v1_user_proto_init() {
	if File_core_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_core_v1_user_proto_rawDesc), len(file_core_v1_user_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_v1_user_proto_goTypes,
		DependencyIndexes: file_core_v1_user_proto_depIdxs,
		EnumInfos:         file_core_v1_user_proto_enumTypes,
		MessageInfos:      file_core_v1_user_proto_msgTypes,
	}.Build()
	File_core_v1_user_proto = out.File
	file_core_v1_user_proto_goTypes = nil
	file_core_v1_user_proto_depIdxs = nil
}
