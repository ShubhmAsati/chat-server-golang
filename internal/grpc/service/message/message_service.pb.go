// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: message/message_service.proto

package message_grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserChatMessageNotificationType int32

const (
	UserChatMessage_MESSAGE     UserChatMessageNotificationType = 0
	UserChatMessage_IMAGE       UserChatMessageNotificationType = 1
	UserChatMessage_FILE        UserChatMessageNotificationType = 2
	UserChatMessage_WEBRTCEVENT UserChatMessageNotificationType = 3
)

// Enum value maps for UserChatMessageNotificationType.
var (
	UserChatMessageNotificationType_name = map[int32]string{
		0: "MESSAGE",
		1: "IMAGE",
		2: "FILE",
		3: "WEBRTCEVENT",
	}
	UserChatMessageNotificationType_value = map[string]int32{
		"MESSAGE":     0,
		"IMAGE":       1,
		"FILE":        2,
		"WEBRTCEVENT": 3,
	}
)

func (x UserChatMessageNotificationType) Enum() *UserChatMessageNotificationType {
	p := new(UserChatMessageNotificationType)
	*p = x
	return p
}

func (x UserChatMessageNotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserChatMessageNotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_message_service_proto_enumTypes[0].Descriptor()
}

func (UserChatMessageNotificationType) Type() protoreflect.EnumType {
	return &file_message_message_service_proto_enumTypes[0]
}

func (x UserChatMessageNotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserChatMessageNotificationType.Descriptor instead.
func (UserChatMessageNotificationType) EnumDescriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{2, 0}
}

type UserById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserById) Reset() {
	*x = UserById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserById) ProtoMessage() {}

func (x *UserById) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserById.ProtoReflect.Descriptor instead.
func (*UserById) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserById) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId        string `protobuf:"bytes,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UtcTimestamp string `protobuf:"bytes,3,opt,name=UtcTimestamp,proto3" json:"UtcTimestamp,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *Msg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Msg) GetUtcTimestamp() string {
	if x != nil {
		return x.UtcTimestamp
	}
	return ""
}

type UserChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nt                  UserChatMessageNotificationType `protobuf:"varint,1,opt,name=nt,proto3,enum=message_grpc.UserChatMessageNotificationType" json:"nt,omitempty"`
	IsGroup             bool                            `protobuf:"varint,2,opt,name=isGroup,proto3" json:"isGroup,omitempty"`
	FromUserId          string                          `protobuf:"bytes,3,opt,name=fromUserId,proto3" json:"fromUserId,omitempty"`
	ToUserId            string                          `protobuf:"bytes,4,opt,name=toUserId,proto3" json:"toUserId,omitempty"`
	Msg                 *Msg                            `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`
	ChatId              string                          `protobuf:"bytes,7,opt,name=chatId,proto3" json:"chatId,omitempty"`
	SentUtcTimeStamp    string                          `protobuf:"bytes,8,opt,name=sentUtcTimeStamp,proto3" json:"sentUtcTimeStamp,omitempty"`
	RecivedUtdTimeStamp string                          `protobuf:"bytes,9,opt,name=recivedUtdTimeStamp,proto3" json:"recivedUtdTimeStamp,omitempty"`
	IceCandidate        string                          `protobuf:"bytes,10,opt,name=iceCandidate,proto3" json:"iceCandidate,omitempty"`
	DynamicJSONString   string                          `protobuf:"bytes,11,opt,name=dynamicJSONString,proto3" json:"dynamicJSONString,omitempty"`
	WebRTCEventType     string                          `protobuf:"bytes,12,opt,name=webRTCEventType,proto3" json:"webRTCEventType,omitempty"`
}

func (x *UserChatMessage) Reset() {
	*x = UserChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChatMessage) ProtoMessage() {}

func (x *UserChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChatMessage.ProtoReflect.Descriptor instead.
func (*UserChatMessage) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserChatMessage) GetNt() UserChatMessageNotificationType {
	if x != nil {
		return x.Nt
	}
	return UserChatMessage_MESSAGE
}

func (x *UserChatMessage) GetIsGroup() bool {
	if x != nil {
		return x.IsGroup
	}
	return false
}

func (x *UserChatMessage) GetFromUserId() string {
	if x != nil {
		return x.FromUserId
	}
	return ""
}

func (x *UserChatMessage) GetToUserId() string {
	if x != nil {
		return x.ToUserId
	}
	return ""
}

func (x *UserChatMessage) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *UserChatMessage) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *UserChatMessage) GetSentUtcTimeStamp() string {
	if x != nil {
		return x.SentUtcTimeStamp
	}
	return ""
}

func (x *UserChatMessage) GetRecivedUtdTimeStamp() string {
	if x != nil {
		return x.RecivedUtdTimeStamp
	}
	return ""
}

func (x *UserChatMessage) GetIceCandidate() string {
	if x != nil {
		return x.IceCandidate
	}
	return ""
}

func (x *UserChatMessage) GetDynamicJSONString() string {
	if x != nil {
		return x.DynamicJSONString
	}
	return ""
}

func (x *UserChatMessage) GetWebRTCEventType() string {
	if x != nil {
		return x.WebRTCEventType
	}
	return ""
}

type MsgSent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSent           bool   `protobuf:"varint,1,opt,name=isSent,proto3" json:"isSent,omitempty"`
	SentUtcTimeStamp string `protobuf:"bytes,2,opt,name=sentUtcTimeStamp,proto3" json:"sentUtcTimeStamp,omitempty"`
}

func (x *MsgSent) Reset() {
	*x = MsgSent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSent) ProtoMessage() {}

func (x *MsgSent) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSent.ProtoReflect.Descriptor instead.
func (*MsgSent) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{3}
}

func (x *MsgSent) GetIsSent() bool {
	if x != nil {
		return x.IsSent
	}
	return false
}

func (x *MsgSent) GetSentUtcTimeStamp() string {
	if x != nil {
		return x.SentUtcTimeStamp
	}
	return ""
}

type StatusOkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	SuccessCode int32  `protobuf:"varint,2,opt,name=successCode,proto3" json:"successCode,omitempty"`
	Err         *Error `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *StatusOkResponse) Reset() {
	*x = StatusOkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusOkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusOkResponse) ProtoMessage() {}

func (x *StatusOkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusOkResponse.ProtoReflect.Descriptor instead.
func (*StatusOkResponse) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{4}
}

func (x *StatusOkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StatusOkResponse) GetSuccessCode() int32 {
	if x != nil {
		return x.SuccessCode
	}
	return 0
}

func (x *StatusOkResponse) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg     string `protobuf:"bytes,1,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	ErrorCode    int32  `protobuf:"varint,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorDetails string `protobuf:"bytes,3,opt,name=errorDetails,proto3" json:"errorDetails,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *Error) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *Error) GetErrorDetails() string {
	if x != nil {
		return x.ErrorDetails
	}
	return ""
}

type EmitWebRTCeventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUserId     string `protobuf:"bytes,1,opt,name=fromUserId,proto3" json:"fromUserId,omitempty"`
	ToUserId       string `protobuf:"bytes,2,opt,name=toUserId,proto3" json:"toUserId,omitempty"`
	ICECandidate   string `protobuf:"bytes,3,opt,name=ICECandidate,proto3" json:"ICECandidate,omitempty"`
	EventType      string `protobuf:"bytes,4,opt,name=eventType,proto3" json:"eventType,omitempty"`
	DynamicPayload string `protobuf:"bytes,5,opt,name=dynamicPayload,proto3" json:"dynamicPayload,omitempty"`
}

func (x *EmitWebRTCeventsRequest) Reset() {
	*x = EmitWebRTCeventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmitWebRTCeventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitWebRTCeventsRequest) ProtoMessage() {}

func (x *EmitWebRTCeventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitWebRTCeventsRequest.ProtoReflect.Descriptor instead.
func (*EmitWebRTCeventsRequest) Descriptor() ([]byte, []int) {
	return file_message_message_service_proto_rawDescGZIP(), []int{6}
}

func (x *EmitWebRTCeventsRequest) GetFromUserId() string {
	if x != nil {
		return x.FromUserId
	}
	return ""
}

func (x *EmitWebRTCeventsRequest) GetToUserId() string {
	if x != nil {
		return x.ToUserId
	}
	return ""
}

func (x *EmitWebRTCeventsRequest) GetICECandidate() string {
	if x != nil {
		return x.ICECandidate
	}
	return ""
}

func (x *EmitWebRTCeventsRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *EmitWebRTCeventsRequest) GetDynamicPayload() string {
	if x != nil {
		return x.DynamicPayload
	}
	return ""
}

var File_message_message_service_proto protoreflect.FileDescriptor

var file_message_message_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x22, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x59, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x74, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x55, 0x74, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x87, 0x04, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3f, 0x0a, 0x02, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x02, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x74, 0x55, 0x74, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x65, 0x6e, 0x74, 0x55, 0x74, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x69, 0x76, 0x65, 0x64, 0x55, 0x74, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x63, 0x69, 0x76, 0x65, 0x64, 0x55, 0x74, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x77, 0x65, 0x62, 0x52, 0x54, 0x43, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77,
	0x65, 0x62, 0x52, 0x54, 0x43, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46,
	0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x49, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x45, 0x42, 0x52, 0x54, 0x43, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x22, 0x4d, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x6e,
	0x74, 0x55, 0x74, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x6e, 0x74, 0x55, 0x74, 0x63, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x75, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x65, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x17, 0x65, 0x6d, 0x69, 0x74, 0x57, 0x65, 0x62, 0x52,
	0x54, 0x43, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x49,
	0x43, 0x45, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x49, 0x43, 0x45, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x83, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x42, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x74, 0x12,
	0x43, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x73, 0x67,
	0x53, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x59, 0x0a, 0x10, 0x65, 0x6d, 0x69, 0x74, 0x57, 0x65, 0x62, 0x52, 0x54, 0x43, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x65, 0x6d, 0x69, 0x74, 0x57, 0x65, 0x62, 0x52, 0x54, 0x43, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_message_message_service_proto_rawDescOnce sync.Once
	file_message_message_service_proto_rawDescData = file_message_message_service_proto_rawDesc
)

func file_message_message_service_proto_rawDescGZIP() []byte {
	file_message_message_service_proto_rawDescOnce.Do(func() {
		file_message_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_message_service_proto_rawDescData)
	})
	return file_message_message_service_proto_rawDescData
}

var file_message_message_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_message_message_service_proto_goTypes = []interface{}{
	(UserChatMessageNotificationType)(0), // 0: message_grpc.UserChatMessage.notification_type
	(*UserById)(nil),                     // 1: message_grpc.UserById
	(*Msg)(nil),                          // 2: message_grpc.Msg
	(*UserChatMessage)(nil),              // 3: message_grpc.UserChatMessage
	(*MsgSent)(nil),                      // 4: message_grpc.MsgSent
	(*StatusOkResponse)(nil),             // 5: message_grpc.StatusOkResponse
	(*Error)(nil),                        // 6: message_grpc.Error
	(*EmitWebRTCeventsRequest)(nil),      // 7: message_grpc.emitWebRTCeventsRequest
}
var file_message_message_service_proto_depIdxs = []int32{
	0, // 0: message_grpc.UserChatMessage.nt:type_name -> message_grpc.UserChatMessage.notification_type
	2, // 1: message_grpc.UserChatMessage.msg:type_name -> message_grpc.Msg
	6, // 2: message_grpc.StatusOkResponse.err:type_name -> message_grpc.Error
	1, // 3: message_grpc.Message.Connect:input_type -> message_grpc.UserById
	3, // 4: message_grpc.Message.BroadCastMessage:input_type -> message_grpc.UserChatMessage
	3, // 5: message_grpc.Message.SendMessage:input_type -> message_grpc.UserChatMessage
	1, // 6: message_grpc.Message.RemoveConnection:input_type -> message_grpc.UserById
	7, // 7: message_grpc.Message.emitWebRTCevents:input_type -> message_grpc.emitWebRTCeventsRequest
	3, // 8: message_grpc.Message.Connect:output_type -> message_grpc.UserChatMessage
	4, // 9: message_grpc.Message.BroadCastMessage:output_type -> message_grpc.MsgSent
	4, // 10: message_grpc.Message.SendMessage:output_type -> message_grpc.MsgSent
	5, // 11: message_grpc.Message.RemoveConnection:output_type -> message_grpc.StatusOkResponse
	5, // 12: message_grpc.Message.emitWebRTCevents:output_type -> message_grpc.StatusOkResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_message_service_proto_init() }
func file_message_message_service_proto_init() {
	if File_message_message_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserById); i {
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
		file_message_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_message_message_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChatMessage); i {
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
		file_message_message_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSent); i {
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
		file_message_message_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusOkResponse); i {
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
		file_message_message_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_message_message_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmitWebRTCeventsRequest); i {
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
			RawDescriptor: file_message_message_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_message_service_proto_goTypes,
		DependencyIndexes: file_message_message_service_proto_depIdxs,
		EnumInfos:         file_message_message_service_proto_enumTypes,
		MessageInfos:      file_message_message_service_proto_msgTypes,
	}.Build()
	File_message_message_service_proto = out.File
	file_message_message_service_proto_rawDesc = nil
	file_message_message_service_proto_goTypes = nil
	file_message_message_service_proto_depIdxs = nil
}
