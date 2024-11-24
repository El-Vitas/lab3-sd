// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: hex-s.proto

package hex_s

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

type AddRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddRecordRequest) Reset() {
	*x = AddRecordRequest{}
	mi := &file_hex_s_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordRequest) ProtoMessage() {}

func (x *AddRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordRequest.ProtoReflect.Descriptor instead.
func (*AddRecordRequest) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{0}
}

func (x *AddRecordRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AddRecordRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *AddRecordRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AddRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *AddRecordResponse) Reset() {
	*x = AddRecordResponse{}
	mi := &file_hex_s_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordResponse) ProtoMessage() {}

func (x *AddRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordResponse.ProtoReflect.Descriptor instead.
func (*AddRecordResponse) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{1}
}

func (x *AddRecordResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type DeleteRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *DeleteRecordRequest) Reset() {
	*x = DeleteRecordRequest{}
	mi := &file_hex_s_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecordRequest) ProtoMessage() {}

func (x *DeleteRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteRecordRequest) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRecordRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DeleteRecordRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

type DeleteRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *DeleteRecordResponse) Reset() {
	*x = DeleteRecordResponse{}
	mi := &file_hex_s_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecordResponse) ProtoMessage() {}

func (x *DeleteRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecordResponse.ProtoReflect.Descriptor instead.
func (*DeleteRecordResponse) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRecordResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type RenameRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region     string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	OldProduct string `protobuf:"bytes,2,opt,name=OldProduct,proto3" json:"OldProduct,omitempty"`
	NewProduct string `protobuf:"bytes,3,opt,name=NewProduct,proto3" json:"NewProduct,omitempty"`
}

func (x *RenameRecordRequest) Reset() {
	*x = RenameRecordRequest{}
	mi := &file_hex_s_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenameRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameRecordRequest) ProtoMessage() {}

func (x *RenameRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameRecordRequest.ProtoReflect.Descriptor instead.
func (*RenameRecordRequest) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{4}
}

func (x *RenameRecordRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RenameRecordRequest) GetOldProduct() string {
	if x != nil {
		return x.OldProduct
	}
	return ""
}

func (x *RenameRecordRequest) GetNewProduct() string {
	if x != nil {
		return x.NewProduct
	}
	return ""
}

type RenameRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *RenameRecordResponse) Reset() {
	*x = RenameRecordResponse{}
	mi := &file_hex_s_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenameRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameRecordResponse) ProtoMessage() {}

func (x *RenameRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameRecordResponse.ProtoReflect.Descriptor instead.
func (*RenameRecordResponse) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{5}
}

func (x *RenameRecordResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type UpdateRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product  string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	NewValue string `protobuf:"bytes,3,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
}

func (x *UpdateRecordRequest) Reset() {
	*x = UpdateRecordRequest{}
	mi := &file_hex_s_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordRequest) ProtoMessage() {}

func (x *UpdateRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateRecordRequest) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRecordRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpdateRecordRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *UpdateRecordRequest) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

type UpdateRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *UpdateRecordResponse) Reset() {
	*x = UpdateRecordResponse{}
	mi := &file_hex_s_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordResponse) ProtoMessage() {}

func (x *UpdateRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordResponse.ProtoReflect.Descriptor instead.
func (*UpdateRecordResponse) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRecordResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type GetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *GetDataRequest) Reset() {
	*x = GetDataRequest{}
	mi := &file_hex_s_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataRequest) ProtoMessage() {}

func (x *GetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataRequest.ProtoReflect.Descriptor instead.
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{8}
}

func (x *GetDataRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GetDataRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

type GetDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
	Quantity    int32   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *GetDataResponse) Reset() {
	*x = GetDataResponse{}
	mi := &file_hex_s_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResponse) ProtoMessage() {}

func (x *GetDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hex_s_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResponse.ProtoReflect.Descriptor instead.
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return file_hex_s_proto_rawDescGZIP(), []int{9}
}

func (x *GetDataResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

func (x *GetDataResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_hex_s_proto protoreflect.FileDescriptor

var file_hex_s_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x78, 0x2d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x65, 0x78, 0x5f, 0x73, 0x22, 0x5a, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x36, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x47, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x39, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x6d, 0x0a, 0x13,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x4f,
	0x6c, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x4f, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4e,
	0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x39, 0x0a, 0x14, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x64, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x50, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xe0, 0x02,
	0x0a, 0x04, 0x48, 0x65, 0x78, 0x53, 0x12, 0x3e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x78, 0x5f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68,
	0x65, 0x78, 0x5f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x78, 0x5f, 0x73, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x78, 0x5f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x1a, 0x2e, 0x68, 0x65, 0x78, 0x5f, 0x73, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65,
	0x78, 0x5f, 0x73, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x2e,
	0x68, 0x65, 0x78, 0x5f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x78, 0x5f,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x78, 0x5f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x78, 0x5f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x68, 0x65,
	0x78, 0x5f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hex_s_proto_rawDescOnce sync.Once
	file_hex_s_proto_rawDescData = file_hex_s_proto_rawDesc
)

func file_hex_s_proto_rawDescGZIP() []byte {
	file_hex_s_proto_rawDescOnce.Do(func() {
		file_hex_s_proto_rawDescData = protoimpl.X.CompressGZIP(file_hex_s_proto_rawDescData)
	})
	return file_hex_s_proto_rawDescData
}

var file_hex_s_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_hex_s_proto_goTypes = []any{
	(*AddRecordRequest)(nil),     // 0: hex_s.AddRecordRequest
	(*AddRecordResponse)(nil),    // 1: hex_s.AddRecordResponse
	(*DeleteRecordRequest)(nil),  // 2: hex_s.DeleteRecordRequest
	(*DeleteRecordResponse)(nil), // 3: hex_s.DeleteRecordResponse
	(*RenameRecordRequest)(nil),  // 4: hex_s.RenameRecordRequest
	(*RenameRecordResponse)(nil), // 5: hex_s.RenameRecordResponse
	(*UpdateRecordRequest)(nil),  // 6: hex_s.UpdateRecordRequest
	(*UpdateRecordResponse)(nil), // 7: hex_s.UpdateRecordResponse
	(*GetDataRequest)(nil),       // 8: hex_s.GetDataRequest
	(*GetDataResponse)(nil),      // 9: hex_s.GetDataResponse
}
var file_hex_s_proto_depIdxs = []int32{
	0, // 0: hex_s.HexS.AddRecord:input_type -> hex_s.AddRecordRequest
	2, // 1: hex_s.HexS.DeleteRecord:input_type -> hex_s.DeleteRecordRequest
	4, // 2: hex_s.HexS.RenameRecord:input_type -> hex_s.RenameRecordRequest
	6, // 3: hex_s.HexS.UpdateRecordValue:input_type -> hex_s.UpdateRecordRequest
	8, // 4: hex_s.HexS.GetData:input_type -> hex_s.GetDataRequest
	1, // 5: hex_s.HexS.AddRecord:output_type -> hex_s.AddRecordResponse
	3, // 6: hex_s.HexS.DeleteRecord:output_type -> hex_s.DeleteRecordResponse
	5, // 7: hex_s.HexS.RenameRecord:output_type -> hex_s.RenameRecordResponse
	7, // 8: hex_s.HexS.UpdateRecordValue:output_type -> hex_s.UpdateRecordResponse
	9, // 9: hex_s.HexS.GetData:output_type -> hex_s.GetDataResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hex_s_proto_init() }
func file_hex_s_proto_init() {
	if File_hex_s_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hex_s_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hex_s_proto_goTypes,
		DependencyIndexes: file_hex_s_proto_depIdxs,
		MessageInfos:      file_hex_s_proto_msgTypes,
	}.Build()
	File_hex_s_proto = out.File
	file_hex_s_proto_rawDesc = nil
	file_hex_s_proto_goTypes = nil
	file_hex_s_proto_depIdxs = nil
}
