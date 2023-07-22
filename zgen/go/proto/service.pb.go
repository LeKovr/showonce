// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/service.proto

package showonce

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ItemStatus int32

const (
	// A Standard tournament
	ItemStatus_UNKNOWN ItemStatus = 0
	// Item was not read
	ItemStatus_WAIT ItemStatus = 1
	// Item was read
	ItemStatus_READ ItemStatus = 2
	// Item was expired
	ItemStatus_EXPIRED ItemStatus = 3
	// Item was cleared
	ItemStatus_CLEARED ItemStatus = 4
)

// Enum value maps for ItemStatus.
var (
	ItemStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "WAIT",
		2: "READ",
		3: "EXPIRED",
		4: "CLEARED",
	}
	ItemStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"WAIT":    1,
		"READ":    2,
		"EXPIRED": 3,
		"CLEARED": 4,
	}
)

func (x ItemStatus) Enum() *ItemStatus {
	p := new(ItemStatus)
	*p = x
	return p
}

func (x ItemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_service_proto_enumTypes[0].Descriptor()
}

func (ItemStatus) Type() protoreflect.EnumType {
	return &file_proto_service_proto_enumTypes[0]
}

func (x ItemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemStatus.Descriptor instead.
func (ItemStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

// Item ULID
type ItemId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ItemId) Reset() {
	*x = ItemId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemId) ProtoMessage() {}

func (x *ItemId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemId.ProtoReflect.Descriptor instead.
func (*ItemId) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *ItemId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NewItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Group      string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Expire     string `protobuf:"bytes,3,opt,name=expire,proto3" json:"expire,omitempty"`
	ExpireUnit string `protobuf:"bytes,4,opt,name=expire_unit,json=expireUnit,proto3" json:"expire_unit,omitempty"`
	Data       string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NewItemRequest) Reset() {
	*x = NewItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewItemRequest) ProtoMessage() {}

func (x *NewItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewItemRequest.ProtoReflect.Descriptor instead.
func (*NewItemRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewItemRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewItemRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *NewItemRequest) GetExpire() string {
	if x != nil {
		return x.Expire
	}
	return ""
}

func (x *NewItemRequest) GetExpireUnit() string {
	if x != nil {
		return x.ExpireUnit
	}
	return ""
}

func (x *NewItemRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ItemData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ItemData) Reset() {
	*x = ItemData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemData) ProtoMessage() {}

func (x *ItemData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemData.ProtoReflect.Descriptor instead.
func (*ItemData) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *ItemData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ItemMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Group      string                 `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Owner      string                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Status     ItemStatus             `protobuf:"varint,4,opt,name=status,proto3,enum=api.showonce.v1.ItemStatus" json:"status,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"` // >now() means Expire (and Status=1)
}

func (x *ItemMeta) Reset() {
	*x = ItemMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemMeta) ProtoMessage() {}

func (x *ItemMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemMeta.ProtoReflect.Descriptor instead.
func (*ItemMeta) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *ItemMeta) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemMeta) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ItemMeta) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ItemMeta) GetStatus() ItemStatus {
	if x != nil {
		return x.Status
	}
	return ItemStatus_UNKNOWN
}

func (x *ItemMeta) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ItemMeta) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

type ItemMetaWithId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Meta *ItemMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ItemMetaWithId) Reset() {
	*x = ItemMetaWithId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemMetaWithId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemMetaWithId) ProtoMessage() {}

func (x *ItemMetaWithId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemMetaWithId.ProtoReflect.Descriptor instead.
func (*ItemMetaWithId) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *ItemMetaWithId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemMetaWithId) GetMeta() *ItemMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemMetaWithId `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ItemList) Reset() {
	*x = ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemList) ProtoMessage() {}

func (x *ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemList.ProtoReflect.Descriptor instead.
func (*ItemList) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *ItemList) GetItems() []*ItemMetaWithId {
	if x != nil {
		return x.Items
	}
	return nil
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Wait    int32 `protobuf:"varint,2,opt,name=wait,proto3" json:"wait,omitempty"`
	Read    int32 `protobuf:"varint,3,opt,name=read,proto3" json:"read,omitempty"`
	Expired int32 `protobuf:"varint,4,opt,name=expired,proto3" json:"expired,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *Stats) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Stats) GetWait() int32 {
	if x != nil {
		return x.Wait
	}
	return 0
}

func (x *Stats) GetRead() int32 {
	if x != nil {
		return x.Read
	}
	return 0
}

func (x *Stats) GetExpired() int32 {
	if x != nil {
		return x.Expired
	}
	return 0
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	My    *Stats `protobuf:"bytes,1,opt,name=my,proto3" json:"my,omitempty"`
	Other *Stats `protobuf:"bytes,2,opt,name=other,proto3" json:"other,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *StatsResponse) GetMy() *Stats {
	if x != nil {
		return x.My
	}
	return nil
}

func (x *StatsResponse) GetOther() *Stats {
	if x != nil {
		return x.Other
	}
	return nil
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x89, 0x01, 0x0a,
	0x0e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf9, 0x01, 0x0a, 0x08, 0x49, 0x74, 0x65,
	0x6d, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x4f, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61,
	0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x61, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x61, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x02, 0x6d, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f,
	0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x02,
	0x6d, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x2a, 0x47, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57,
	0x41, 0x49, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x4c, 0x45, 0x41, 0x52, 0x45, 0x44, 0x10, 0x04, 0x32, 0xbc, 0x01, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x11,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x65,
	0x6d, 0x12, 0x55, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77,
	0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x32, 0xbe, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0a, 0x4e,
	0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x22, 0x34, 0x92, 0x41, 0x1b, 0x62, 0x19, 0x0a, 0x17, 0x0a, 0x06, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x32, 0x12, 0x0d, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x05, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x6d, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x6d, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x58, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x6d, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x4e, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x68,
	0x6f, 0x77, 0x4f, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x69, 0x74,
	0x65, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x73, 0x68, 0x6f, 0x77, 0x6f, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_service_proto_goTypes = []interface{}{
	(ItemStatus)(0),               // 0: api.showonce.v1.ItemStatus
	(*ItemId)(nil),                // 1: api.showonce.v1.ItemId
	(*NewItemRequest)(nil),        // 2: api.showonce.v1.NewItemRequest
	(*ItemData)(nil),              // 3: api.showonce.v1.ItemData
	(*ItemMeta)(nil),              // 4: api.showonce.v1.ItemMeta
	(*ItemMetaWithId)(nil),        // 5: api.showonce.v1.ItemMetaWithId
	(*ItemList)(nil),              // 6: api.showonce.v1.ItemList
	(*Stats)(nil),                 // 7: api.showonce.v1.Stats
	(*StatsResponse)(nil),         // 8: api.showonce.v1.StatsResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_proto_service_proto_depIdxs = []int32{
	0,  // 0: api.showonce.v1.ItemMeta.status:type_name -> api.showonce.v1.ItemStatus
	9,  // 1: api.showonce.v1.ItemMeta.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: api.showonce.v1.ItemMeta.modified_at:type_name -> google.protobuf.Timestamp
	4,  // 3: api.showonce.v1.ItemMetaWithId.meta:type_name -> api.showonce.v1.ItemMeta
	5,  // 4: api.showonce.v1.ItemList.items:type_name -> api.showonce.v1.ItemMetaWithId
	7,  // 5: api.showonce.v1.StatsResponse.my:type_name -> api.showonce.v1.Stats
	7,  // 6: api.showonce.v1.StatsResponse.other:type_name -> api.showonce.v1.Stats
	1,  // 7: api.showonce.v1.PublicService.GetMetadata:input_type -> api.showonce.v1.ItemId
	1,  // 8: api.showonce.v1.PublicService.GetData:input_type -> api.showonce.v1.ItemId
	2,  // 9: api.showonce.v1.PrivateService.NewMessage:input_type -> api.showonce.v1.NewItemRequest
	10, // 10: api.showonce.v1.PrivateService.GetItems:input_type -> google.protobuf.Empty
	10, // 11: api.showonce.v1.PrivateService.GetStats:input_type -> google.protobuf.Empty
	4,  // 12: api.showonce.v1.PublicService.GetMetadata:output_type -> api.showonce.v1.ItemMeta
	3,  // 13: api.showonce.v1.PublicService.GetData:output_type -> api.showonce.v1.ItemData
	1,  // 14: api.showonce.v1.PrivateService.NewMessage:output_type -> api.showonce.v1.ItemId
	6,  // 15: api.showonce.v1.PrivateService.GetItems:output_type -> api.showonce.v1.ItemList
	8,  // 16: api.showonce.v1.PrivateService.GetStats:output_type -> api.showonce.v1.StatsResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemId); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewItemRequest); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemData); i {
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
		file_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemMeta); i {
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
		file_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemMetaWithId); i {
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
		file_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemList); i {
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
		file_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_proto_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		EnumInfos:         file_proto_service_proto_enumTypes,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
