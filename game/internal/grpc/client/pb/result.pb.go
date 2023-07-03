// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: result.proto

package result

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

// EventType 0表示匹配的结果，1表示游戏的结果，2表示游戏地图
type ResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType   int32        `protobuf:"varint,1,opt,name=EventType,proto3" json:"EventType,omitempty"`
	MatchResult *MatchResult `protobuf:"bytes,2,opt,name=MatchResult,proto3" json:"MatchResult,omitempty"`
	GameResult  *GameResult  `protobuf:"bytes,3,opt,name=GameResult,proto3" json:"GameResult,omitempty"`
	GameMap     *GameMap     `protobuf:"bytes,4,opt,name=GameMap,proto3" json:"GameMap,omitempty"`
}

func (x *ResultReq) Reset() {
	*x = ResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultReq) ProtoMessage() {}

func (x *ResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultReq.ProtoReflect.Descriptor instead.
func (*ResultReq) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *ResultReq) GetEventType() int32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

func (x *ResultReq) GetMatchResult() *MatchResult {
	if x != nil {
		return x.MatchResult
	}
	return nil
}

func (x *ResultReq) GetGameResult() *GameResult {
	if x != nil {
		return x.GameResult
	}
	return nil
}

func (x *ResultReq) GetGameMap() *GameMap {
	if x != nil {
		return x.GameMap
	}
	return nil
}

type MatchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AId    int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	ABotId int32 `protobuf:"varint,2,opt,name=ABotId,proto3" json:"ABotId,omitempty"`
	BId    int32 `protobuf:"varint,3,opt,name=BId,proto3" json:"BId,omitempty"`
	BBotId int32 `protobuf:"varint,4,opt,name=BBotId,proto3" json:"BBotId,omitempty"`
}

func (x *MatchResult) Reset() {
	*x = MatchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResult) ProtoMessage() {}

func (x *MatchResult) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResult.ProtoReflect.Descriptor instead.
func (*MatchResult) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *MatchResult) GetAId() int32 {
	if x != nil {
		return x.AId
	}
	return 0
}

func (x *MatchResult) GetABotId() int32 {
	if x != nil {
		return x.ABotId
	}
	return 0
}

func (x *MatchResult) GetBId() int32 {
	if x != nil {
		return x.BId
	}
	return 0
}

func (x *MatchResult) GetBBotId() int32 {
	if x != nil {
		return x.BBotId
	}
	return 0
}

type GameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Loser string `protobuf:"bytes,1,opt,name=Loser,proto3" json:"Loser,omitempty"`
}

func (x *GameResult) Reset() {
	*x = GameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResult) ProtoMessage() {}

func (x *GameResult) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResult.ProtoReflect.Descriptor instead.
func (*GameResult) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{2}
}

func (x *GameResult) GetLoser() string {
	if x != nil {
		return x.Loser
	}
	return ""
}

type ResultResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ResultResp) Reset() {
	*x = ResultResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResp) ProtoMessage() {}

func (x *ResultResp) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResp.ProtoReflect.Descriptor instead.
func (*ResultResp) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{3}
}

func (x *ResultResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GameMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AId     int32   `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	ASx     int32   `protobuf:"varint,2,opt,name=ASx,proto3" json:"ASx,omitempty"`
	ASy     int32   `protobuf:"varint,3,opt,name=ASy,proto3" json:"ASy,omitempty"`
	BId     int32   `protobuf:"varint,4,opt,name=BId,proto3" json:"BId,omitempty"`
	BSx     int32   `protobuf:"varint,5,opt,name=BSx,proto3" json:"BSx,omitempty"`
	BSy     int32   `protobuf:"varint,6,opt,name=BSy,proto3" json:"BSy,omitempty"`
	GameMap []*Edge `protobuf:"bytes,7,rep,name=GameMap,proto3" json:"GameMap,omitempty"`
	PlayerA *User   `protobuf:"bytes,8,opt,name=playerA,proto3" json:"playerA,omitempty"`
	PlayerB *User   `protobuf:"bytes,9,opt,name=playerB,proto3" json:"playerB,omitempty"`
}

func (x *GameMap) Reset() {
	*x = GameMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMap) ProtoMessage() {}

func (x *GameMap) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMap.ProtoReflect.Descriptor instead.
func (*GameMap) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{4}
}

func (x *GameMap) GetAId() int32 {
	if x != nil {
		return x.AId
	}
	return 0
}

func (x *GameMap) GetASx() int32 {
	if x != nil {
		return x.ASx
	}
	return 0
}

func (x *GameMap) GetASy() int32 {
	if x != nil {
		return x.ASy
	}
	return 0
}

func (x *GameMap) GetBId() int32 {
	if x != nil {
		return x.BId
	}
	return 0
}

func (x *GameMap) GetBSx() int32 {
	if x != nil {
		return x.BSx
	}
	return 0
}

func (x *GameMap) GetBSy() int32 {
	if x != nil {
		return x.BSy
	}
	return 0
}

func (x *GameMap) GetGameMap() []*Edge {
	if x != nil {
		return x.GameMap
	}
	return nil
}

func (x *GameMap) GetPlayerA() *User {
	if x != nil {
		return x.PlayerA
	}
	return nil
}

func (x *GameMap) GetPlayerB() *User {
	if x != nil {
		return x.PlayerB
	}
	return nil
}

type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Edge []int32 `protobuf:"varint,1,rep,packed,name=edge,proto3" json:"edge,omitempty"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{5}
}

func (x *Edge) GetEdge() []int32 {
	if x != nil {
		return x.Edge
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photo    string `protobuf:"bytes,1,opt,name=photo,proto3" json:"photo,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	UserID   int32  `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[6]
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
	return file_result_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62, 0x22, 0xc8, 0x01, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x35, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x61,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x62, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x07, 0x47, 0x61, 0x6d,
	0x65, 0x4d, 0x61, 0x70, 0x22, 0x61, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x41, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x42, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x42, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x42, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x42, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x73, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x41, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x53, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x41, 0x53, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x53, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x41, 0x53, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x42, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x53, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x42, 0x53, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x53, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x42, 0x53, 0x79, 0x12, 0x29, 0x0a, 0x07, 0x47,
	0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x07, 0x47,
	0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x22, 0x1a, 0x0a, 0x04,
	0x45, 0x64, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x04, 0x65, 0x64, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x32, 0x41, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_result_proto_goTypes = []interface{}{
	(*ResultReq)(nil),   // 0: result.pb.ResultReq
	(*MatchResult)(nil), // 1: result.pb.matchResult
	(*GameResult)(nil),  // 2: result.pb.gameResult
	(*ResultResp)(nil),  // 3: result.pb.ResultResp
	(*GameMap)(nil),     // 4: result.pb.gameMap
	(*Edge)(nil),        // 5: result.pb.Edge
	(*User)(nil),        // 6: result.pb.User
}
var file_result_proto_depIdxs = []int32{
	1, // 0: result.pb.ResultReq.MatchResult:type_name -> result.pb.matchResult
	2, // 1: result.pb.ResultReq.GameResult:type_name -> result.pb.gameResult
	4, // 2: result.pb.ResultReq.GameMap:type_name -> result.pb.gameMap
	5, // 3: result.pb.gameMap.GameMap:type_name -> result.pb.Edge
	6, // 4: result.pb.gameMap.playerA:type_name -> result.pb.User
	6, // 5: result.pb.gameMap.playerB:type_name -> result.pb.User
	0, // 6: result.pb.Result.Result:input_type -> result.pb.ResultReq
	3, // 7: result.pb.Result.Result:output_type -> result.pb.ResultResp
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultReq); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResult); i {
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
		file_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResult); i {
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
		file_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResp); i {
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
		file_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMap); i {
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
		file_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge); i {
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
		file_result_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}