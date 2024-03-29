// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: user/user_service.proto

package user_grpc

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ProfilePic string `protobuf:"bytes,3,opt,name=profilePic,proto3" json:"profilePic,omitempty"`
	MobileNo   string `protobuf:"bytes,4,opt,name=mobileNo,proto3" json:"mobileNo,omitempty"`
	Status     string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	UserName   string `protobuf:"bytes,6,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetProfilePic() string {
	if x != nil {
		return x.ProfilePic
	}
	return ""
}

func (x *User) GetMobileNo() string {
	if x != nil {
		return x.MobileNo
	}
	return ""
}

func (x *User) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type AddUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Err    *Error `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AddUserResponse) Reset() {
	*x = AddUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResponse) ProtoMessage() {}

func (x *AddUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResponse.ProtoReflect.Descriptor instead.
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddUserResponse) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

type UploadProfilePicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePicPath string `protobuf:"bytes,1,opt,name=profilePicPath,proto3" json:"profilePicPath,omitempty"`
	User           *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UploadProfilePicRequest) Reset() {
	*x = UploadProfilePicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadProfilePicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadProfilePicRequest) ProtoMessage() {}

func (x *UploadProfilePicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadProfilePicRequest.ProtoReflect.Descriptor instead.
func (*UploadProfilePicRequest) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadProfilePicRequest) GetProfilePicPath() string {
	if x != nil {
		return x.ProfilePicPath
	}
	return ""
}

func (x *UploadProfilePicRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type MyMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMsg []*UserMsg `protobuf:"bytes,1,rep,name=userMsg,proto3" json:"userMsg,omitempty"`
	Error   *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MyMessages) Reset() {
	*x = MyMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyMessages) ProtoMessage() {}

func (x *MyMessages) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyMessages.ProtoReflect.Descriptor instead.
func (*MyMessages) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *MyMessages) GetUserMsg() []*UserMsg {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *MyMessages) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UserMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUserId string `protobuf:"bytes,1,opt,name=fromUserId,proto3" json:"fromUserId,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UserMsg) Reset() {
	*x = UserMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMsg) ProtoMessage() {}

func (x *UserMsg) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMsg.ProtoReflect.Descriptor instead.
func (*UserMsg) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserMsg) GetFromUserId() string {
	if x != nil {
		return x.FromUserId
	}
	return ""
}

func (x *UserMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetUserByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserByUserIdRequest) Reset() {
	*x = GetUserByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUserIdRequest) ProtoMessage() {}

func (x *GetUserByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserByUserNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *GetUserByUserNameRequest) Reset() {
	*x = GetUserByUserNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUserNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUserNameRequest) ProtoMessage() {}

func (x *GetUserByUserNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUserNameRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUserNameRequest) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserByUserNameRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type SearchUsersByUserNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *SearchUsersByUserNameResponse) Reset() {
	*x = SearchUsersByUserNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersByUserNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersByUserNameResponse) ProtoMessage() {}

func (x *SearchUsersByUserNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersByUserNameResponse.ProtoReflect.Descriptor instead.
func (*SearchUsersByUserNameResponse) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *SearchUsersByUserNameResponse) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *UserName) Reset() {
	*x = UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserName) ProtoMessage() {}

func (x *UserName) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserName.ProtoReflect.Descriptor instead.
func (*UserName) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserName) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type GetExistingUserNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sample string `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
}

func (x *GetExistingUserNamesRequest) Reset() {
	*x = GetExistingUserNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExistingUserNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExistingUserNamesRequest) ProtoMessage() {}

func (x *GetExistingUserNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExistingUserNamesRequest.ProtoReflect.Descriptor instead.
func (*GetExistingUserNamesRequest) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetExistingUserNamesRequest) GetSample() string {
	if x != nil {
		return x.Sample
	}
	return ""
}

type GetExistingUserNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName []*UserName `protobuf:"bytes,1,rep,name=userName,proto3" json:"userName,omitempty"`
}

func (x *GetExistingUserNamesResponse) Reset() {
	*x = GetExistingUserNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExistingUserNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExistingUserNamesResponse) ProtoMessage() {}

func (x *GetExistingUserNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExistingUserNamesResponse.ProtoReflect.Descriptor instead.
func (*GetExistingUserNamesResponse) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetExistingUserNamesResponse) GetUserName() []*UserName {
	if x != nil {
		return x.UserName
	}
	return nil
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
		mi := &file_user_user_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusOkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusOkResponse) ProtoMessage() {}

func (x *StatusOkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[11]
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
	return file_user_user_service_proto_rawDescGZIP(), []int{11}
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
		mi := &file_user_user_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[12]
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
	return file_user_user_service_proto_rawDescGZIP(), []int{12}
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

var File_user_user_service_proto protoreflect.FileDescriptor

var file_user_user_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65,
	0x72, 0x72, 0x22, 0x66, 0x0a, 0x17, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69,
	0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x0a, 0x4d, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3b,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x16, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a,
	0x18, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x1d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x1c, 0x67, 0x65,
	0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x10, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x65, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0xa0, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x0f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12, 0x22,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0f, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x36, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x15, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x67, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_user_service_proto_rawDescOnce sync.Once
	file_user_user_service_proto_rawDescData = file_user_user_service_proto_rawDesc
)

func file_user_user_service_proto_rawDescGZIP() []byte {
	file_user_user_service_proto_rawDescOnce.Do(func() {
		file_user_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_service_proto_rawDescData)
	})
	return file_user_user_service_proto_rawDescData
}

var file_user_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_user_user_service_proto_goTypes = []interface{}{
	(*User)(nil),                          // 0: user_grpc.User
	(*AddUserResponse)(nil),               // 1: user_grpc.AddUserResponse
	(*UploadProfilePicRequest)(nil),       // 2: user_grpc.UploadProfilePicRequest
	(*MyMessages)(nil),                    // 3: user_grpc.MyMessages
	(*UserMsg)(nil),                       // 4: user_grpc.UserMsg
	(*GetUserByUserIdRequest)(nil),        // 5: user_grpc.getUserByUserIdRequest
	(*GetUserByUserNameRequest)(nil),      // 6: user_grpc.getUserByUserNameRequest
	(*SearchUsersByUserNameResponse)(nil), // 7: user_grpc.searchUsersByUserNameResponse
	(*UserName)(nil),                      // 8: user_grpc.UserName
	(*GetExistingUserNamesRequest)(nil),   // 9: user_grpc.getExistingUserNamesRequest
	(*GetExistingUserNamesResponse)(nil),  // 10: user_grpc.getExistingUserNamesResponse
	(*StatusOkResponse)(nil),              // 11: user_grpc.StatusOkResponse
	(*Error)(nil),                         // 12: user_grpc.Error
}
var file_user_user_service_proto_depIdxs = []int32{
	12, // 0: user_grpc.AddUserResponse.err:type_name -> user_grpc.Error
	0,  // 1: user_grpc.UploadProfilePicRequest.user:type_name -> user_grpc.User
	4,  // 2: user_grpc.MyMessages.userMsg:type_name -> user_grpc.UserMsg
	12, // 3: user_grpc.MyMessages.error:type_name -> user_grpc.Error
	0,  // 4: user_grpc.searchUsersByUserNameResponse.user:type_name -> user_grpc.User
	8,  // 5: user_grpc.getExistingUserNamesResponse.userName:type_name -> user_grpc.UserName
	12, // 6: user_grpc.StatusOkResponse.err:type_name -> user_grpc.Error
	0,  // 7: user_grpc.UserService.add:input_type -> user_grpc.User
	2,  // 8: user_grpc.UserService.uploadProfilePic:input_type -> user_grpc.UploadProfilePicRequest
	0,  // 9: user_grpc.UserService.loadMyMessages:input_type -> user_grpc.User
	5,  // 10: user_grpc.UserService.getUserByUserId:input_type -> user_grpc.getUserByUserIdRequest
	0,  // 11: user_grpc.UserService.update:input_type -> user_grpc.User
	6,  // 12: user_grpc.UserService.searchUsersByUserName:input_type -> user_grpc.getUserByUserNameRequest
	9,  // 13: user_grpc.UserService.getExistingUserNames:input_type -> user_grpc.getExistingUserNamesRequest
	1,  // 14: user_grpc.UserService.add:output_type -> user_grpc.AddUserResponse
	11, // 15: user_grpc.UserService.uploadProfilePic:output_type -> user_grpc.StatusOkResponse
	3,  // 16: user_grpc.UserService.loadMyMessages:output_type -> user_grpc.MyMessages
	0,  // 17: user_grpc.UserService.getUserByUserId:output_type -> user_grpc.User
	11, // 18: user_grpc.UserService.update:output_type -> user_grpc.StatusOkResponse
	7,  // 19: user_grpc.UserService.searchUsersByUserName:output_type -> user_grpc.searchUsersByUserNameResponse
	10, // 20: user_grpc.UserService.getExistingUserNames:output_type -> user_grpc.getExistingUserNamesResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_user_user_service_proto_init() }
func file_user_user_service_proto_init() {
	if File_user_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserResponse); i {
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
		file_user_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadProfilePicRequest); i {
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
		file_user_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyMessages); i {
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
		file_user_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMsg); i {
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
		file_user_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUserIdRequest); i {
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
		file_user_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUserNameRequest); i {
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
		file_user_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUsersByUserNameResponse); i {
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
		file_user_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserName); i {
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
		file_user_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExistingUserNamesRequest); i {
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
		file_user_user_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExistingUserNamesResponse); i {
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
		file_user_user_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_user_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_service_proto_goTypes,
		DependencyIndexes: file_user_user_service_proto_depIdxs,
		MessageInfos:      file_user_user_service_proto_msgTypes,
	}.Build()
	File_user_user_service_proto = out.File
	file_user_user_service_proto_rawDesc = nil
	file_user_user_service_proto_goTypes = nil
	file_user_user_service_proto_depIdxs = nil
}
