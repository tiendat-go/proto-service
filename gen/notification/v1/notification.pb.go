// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: notification/v1/notification.proto

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

type SendNotificationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendNotificationRequest) Reset() {
	*x = SendNotificationRequest{}
	mi := &file_notification_v1_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationRequest) ProtoMessage() {}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationRequest.ProtoReflect.Descriptor instead.
func (*SendNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (x *SendNotificationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendNotificationRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendNotificationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendNotificationResponse) Reset() {
	*x = SendNotificationResponse{}
	mi := &file_notification_v1_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationResponse) ProtoMessage() {}

func (x *SendNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationResponse.ProtoReflect.Descriptor instead.
func (*SendNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{1}
}

func (x *SendNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetNotificationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotificationsRequest) Reset() {
	*x = GetNotificationsRequest{}
	mi := &file_notification_v1_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsRequest) ProtoMessage() {}

func (x *GetNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetNotificationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Notifications []*Notification        `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotificationsResponse) Reset() {
	*x = GetNotificationsResponse{}
	mi := &file_notification_v1_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsResponse) ProtoMessage() {}

func (x *GetNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{3}
}

func (x *GetNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp     int64                  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_notification_v1_notification_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{4}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_notification_v1_notification_proto protoreflect.FileDescriptor

const file_notification_v1_notification_proto_rawDesc = "" +
	"\n" +
	"\"notification/v1/notification.proto\x12\x0fnotification.v1\"L\n" +
	"\x17SendNotificationRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"4\n" +
	"\x18SendNotificationResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"2\n" +
	"\x17GetNotificationsRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"_\n" +
	"\x18GetNotificationsResponse\x12C\n" +
	"\rnotifications\x18\x01 \x03(\v2\x1d.notification.v1.NotificationR\rnotifications\"o\n" +
	"\fNotification\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x12\x1c\n" +
	"\ttimestamp\x18\x04 \x01(\x03R\ttimestamp2\xe7\x01\n" +
	"\x13NotificationService\x12g\n" +
	"\x10SendNotification\x12(.notification.v1.SendNotificationRequest\x1a).notification.v1.SendNotificationResponse\x12g\n" +
	"\x10GetNotifications\x12(.notification.v1.GetNotificationsRequest\x1a).notification.v1.GetNotificationsResponseB\x9e\x01\n" +
	"\x13com.notification.v1B\x11NotificationProtoP\x01Z\x17/gen/go/notification/v1\xa2\x02\x03NXX\xaa\x02\x0fNotification.V1\xca\x02\x0fNotification\\V1\xe2\x02\x1bNotification\\V1\\GPBMetadata\xea\x02\x10Notification::V1b\x06proto3"

var (
	file_notification_v1_notification_proto_rawDescOnce sync.Once
	file_notification_v1_notification_proto_rawDescData []byte
)

func file_notification_v1_notification_proto_rawDescGZIP() []byte {
	file_notification_v1_notification_proto_rawDescOnce.Do(func() {
		file_notification_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notification_v1_notification_proto_rawDesc), len(file_notification_v1_notification_proto_rawDesc)))
	})
	return file_notification_v1_notification_proto_rawDescData
}

var file_notification_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notification_v1_notification_proto_goTypes = []any{
	(*SendNotificationRequest)(nil),  // 0: notification.v1.SendNotificationRequest
	(*SendNotificationResponse)(nil), // 1: notification.v1.SendNotificationResponse
	(*GetNotificationsRequest)(nil),  // 2: notification.v1.GetNotificationsRequest
	(*GetNotificationsResponse)(nil), // 3: notification.v1.GetNotificationsResponse
	(*Notification)(nil),             // 4: notification.v1.Notification
}
var file_notification_v1_notification_proto_depIdxs = []int32{
	4, // 0: notification.v1.GetNotificationsResponse.notifications:type_name -> notification.v1.Notification
	0, // 1: notification.v1.NotificationService.SendNotification:input_type -> notification.v1.SendNotificationRequest
	2, // 2: notification.v1.NotificationService.GetNotifications:input_type -> notification.v1.GetNotificationsRequest
	1, // 3: notification.v1.NotificationService.SendNotification:output_type -> notification.v1.SendNotificationResponse
	3, // 4: notification.v1.NotificationService.GetNotifications:output_type -> notification.v1.GetNotificationsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notification_v1_notification_proto_init() }
func file_notification_v1_notification_proto_init() {
	if File_notification_v1_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notification_v1_notification_proto_rawDesc), len(file_notification_v1_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_v1_notification_proto_goTypes,
		DependencyIndexes: file_notification_v1_notification_proto_depIdxs,
		MessageInfos:      file_notification_v1_notification_proto_msgTypes,
	}.Build()
	File_notification_v1_notification_proto = out.File
	file_notification_v1_notification_proto_goTypes = nil
	file_notification_v1_notification_proto_depIdxs = nil
}
