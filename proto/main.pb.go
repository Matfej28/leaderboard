// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: main.proto

package __

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

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username        string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegistrationRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LogInRequest) Reset() {
	*x = LogInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInRequest) ProtoMessage() {}

func (x *LogInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInRequest.ProtoReflect.Descriptor instead.
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{2}
}

func (x *LogInRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LogInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LogInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogInResponse) Reset() {
	*x = LogInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInResponse) ProtoMessage() {}

func (x *LogInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInResponse.ProtoReflect.Descriptor instead.
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{3}
}

func (x *LogInResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Score uint32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *RankRequest) Reset() {
	*x = RankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRequest) ProtoMessage() {}

func (x *RankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRequest.ProtoReflect.Descriptor instead.
func (*RankRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{4}
}

func (x *RankRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RankRequest) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type RankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank uint32 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *RankResponse) Reset() {
	*x = RankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankResponse) ProtoMessage() {}

func (x *RankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankResponse.ProtoReflect.Descriptor instead.
func (*RankResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{5}
}

func (x *RankResponse) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type LeaderboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Page      uint32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Monthly   bool    `protobuf:"varint,3,opt,name=monthly,proto3" json:"monthly,omitempty"` // if false => all time
	ResAmount uint32  `protobuf:"varint,4,opt,name=resAmount,proto3" json:"resAmount,omitempty"`
}

func (x *LeaderboardRequest) Reset() {
	*x = LeaderboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardRequest) ProtoMessage() {}

func (x *LeaderboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardRequest.ProtoReflect.Descriptor instead.
func (*LeaderboardRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{6}
}

func (x *LeaderboardRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LeaderboardRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *LeaderboardRequest) GetMonthly() bool {
	if x != nil {
		return x.Monthly
	}
	return false
}

func (x *LeaderboardRequest) GetResAmount() uint32 {
	if x != nil {
		return x.ResAmount
	}
	return 0
}

type LeaderboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results  []*Player `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	AroundMe []*Player `protobuf:"bytes,2,rep,name=around_me,json=aroundMe,proto3" json:"around_me,omitempty"`
	NextPage uint32    `protobuf:"varint,3,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *LeaderboardResponse) Reset() {
	*x = LeaderboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardResponse) ProtoMessage() {}

func (x *LeaderboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardResponse.ProtoReflect.Descriptor instead.
func (*LeaderboardResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{7}
}

func (x *LeaderboardResponse) GetResults() []*Player {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *LeaderboardResponse) GetAroundMe() []*Player {
	if x != nil {
		return x.AroundMe
	}
	return nil
}

func (x *LeaderboardResponse) GetNextPage() uint32 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Score uint32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Rank  uint32 `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{8}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Player) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x22, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x30, 0x0a, 0x09, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x08, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x46,
	0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x32, 0xc0, 0x02, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x55, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x12, 0x19, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_main_proto_goTypes = []interface{}{
	(*RegistrationRequest)(nil),  // 0: leaderboard.RegistrationRequest
	(*RegistrationResponse)(nil), // 1: leaderboard.RegistrationResponse
	(*LogInRequest)(nil),         // 2: leaderboard.LogInRequest
	(*LogInResponse)(nil),        // 3: leaderboard.LogInResponse
	(*RankRequest)(nil),          // 4: leaderboard.RankRequest
	(*RankResponse)(nil),         // 5: leaderboard.RankResponse
	(*LeaderboardRequest)(nil),   // 6: leaderboard.LeaderboardRequest
	(*LeaderboardResponse)(nil),  // 7: leaderboard.LeaderboardResponse
	(*Player)(nil),               // 8: leaderboard.Player
}
var file_main_proto_depIdxs = []int32{
	8, // 0: leaderboard.LeaderboardResponse.results:type_name -> leaderboard.Player
	8, // 1: leaderboard.LeaderboardResponse.around_me:type_name -> leaderboard.Player
	0, // 2: leaderboard.Leaderboard.Registration:input_type -> leaderboard.RegistrationRequest
	2, // 3: leaderboard.Leaderboard.LogIn:input_type -> leaderboard.LogInRequest
	4, // 4: leaderboard.Leaderboard.Find:input_type -> leaderboard.RankRequest
	6, // 5: leaderboard.Leaderboard.GetLeaderboard:input_type -> leaderboard.LeaderboardRequest
	1, // 6: leaderboard.Leaderboard.Registration:output_type -> leaderboard.RegistrationResponse
	3, // 7: leaderboard.Leaderboard.LogIn:output_type -> leaderboard.LogInResponse
	5, // 8: leaderboard.Leaderboard.Find:output_type -> leaderboard.RankResponse
	7, // 9: leaderboard.Leaderboard.GetLeaderboard:output_type -> leaderboard.LeaderboardResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationResponse); i {
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
		file_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInRequest); i {
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
		file_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInResponse); i {
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
		file_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRequest); i {
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
		file_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankResponse); i {
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
		file_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardRequest); i {
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
		file_main_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardResponse); i {
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
		file_main_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
	file_main_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}
