// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: xmaspi.proto

package proto

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

type SetLedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Led   int64  `protobuf:"varint,1,opt,name=led,proto3" json:"led,omitempty"`
	Color uint32 `protobuf:"varint,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *SetLedRequest) Reset() {
	*x = SetLedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLedRequest) ProtoMessage() {}

func (x *SetLedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLedRequest.ProtoReflect.Descriptor instead.
func (*SetLedRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{0}
}

func (x *SetLedRequest) GetLed() int64 {
	if x != nil {
		return x.Led
	}
	return 0
}

func (x *SetLedRequest) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

type GetLedCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLedCountRequest) Reset() {
	*x = GetLedCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedCountRequest) ProtoMessage() {}

func (x *GetLedCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedCountRequest.ProtoReflect.Descriptor instead.
func (*GetLedCountRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{1}
}

type GetLedCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetLedCountResponse) Reset() {
	*x = GetLedCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedCountResponse) ProtoMessage() {}

func (x *GetLedCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedCountResponse.ProtoReflect.Descriptor instead.
func (*GetLedCountResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{2}
}

func (x *GetLedCountResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAnimationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAnimationsRequest) Reset() {
	*x = GetAnimationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimationsRequest) ProtoMessage() {}

func (x *GetAnimationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimationsRequest.ProtoReflect.Descriptor instead.
func (*GetAnimationsRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{3}
}

type GetAnimationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Animations []*GetAnimationsResponse_Animation `protobuf:"bytes,1,rep,name=animations,proto3" json:"animations,omitempty"`
}

func (x *GetAnimationsResponse) Reset() {
	*x = GetAnimationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimationsResponse) ProtoMessage() {}

func (x *GetAnimationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimationsResponse.ProtoReflect.Descriptor instead.
func (*GetAnimationsResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{4}
}

func (x *GetAnimationsResponse) GetAnimations() []*GetAnimationsResponse_Animation {
	if x != nil {
		return x.Animations
	}
	return nil
}

type GetStaticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStaticsRequest) Reset() {
	*x = GetStaticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsRequest) ProtoMessage() {}

func (x *GetStaticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsRequest.ProtoReflect.Descriptor instead.
func (*GetStaticsRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{5}
}

type GetStaticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statics []*GetStaticsResponse_Static `protobuf:"bytes,1,rep,name=statics,proto3" json:"statics,omitempty"`
}

func (x *GetStaticsResponse) Reset() {
	*x = GetStaticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsResponse) ProtoMessage() {}

func (x *GetStaticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsResponse.ProtoReflect.Descriptor instead.
func (*GetStaticsResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{6}
}

func (x *GetStaticsResponse) GetStatics() []*GetStaticsResponse_Static {
	if x != nil {
		return x.Statics
	}
	return nil
}

type SetLedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetLedResponse) Reset() {
	*x = SetLedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLedResponse) ProtoMessage() {}

func (x *SetLedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLedResponse.ProtoReflect.Descriptor instead.
func (*SetLedResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{7}
}

type SetStaticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetStaticRequest) Reset() {
	*x = SetStaticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStaticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStaticRequest) ProtoMessage() {}

func (x *SetStaticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStaticRequest.ProtoReflect.Descriptor instead.
func (*SetStaticRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{8}
}

func (x *SetStaticRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetStaticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetStaticResponse) Reset() {
	*x = SetStaticResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStaticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStaticResponse) ProtoMessage() {}

func (x *SetStaticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStaticResponse.ProtoReflect.Descriptor instead.
func (*SetStaticResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{9}
}

type SetAnimationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetAnimationRequest) Reset() {
	*x = SetAnimationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAnimationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAnimationRequest) ProtoMessage() {}

func (x *SetAnimationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAnimationRequest.ProtoReflect.Descriptor instead.
func (*SetAnimationRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{10}
}

func (x *SetAnimationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetAnimationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetAnimationResponse) Reset() {
	*x = SetAnimationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAnimationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAnimationResponse) ProtoMessage() {}

func (x *SetAnimationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAnimationResponse.ProtoReflect.Descriptor instead.
func (*SetAnimationResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{11}
}

type ControllerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ControllerInfoRequest) Reset() {
	*x = ControllerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerInfoRequest) ProtoMessage() {}

func (x *ControllerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerInfoRequest.ProtoReflect.Descriptor instead.
func (*ControllerInfoRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{12}
}

type ControllerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LedCount int64 `protobuf:"varint,1,opt,name=ledCount,proto3" json:"ledCount,omitempty"`
}

func (x *ControllerInfoResponse) Reset() {
	*x = ControllerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerInfoResponse) ProtoMessage() {}

func (x *ControllerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerInfoResponse.ProtoReflect.Descriptor instead.
func (*ControllerInfoResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{13}
}

func (x *ControllerInfoResponse) GetLedCount() int64 {
	if x != nil {
		return x.LedCount
	}
	return 0
}

type RenderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenderRequest) Reset() {
	*x = RenderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderRequest) ProtoMessage() {}

func (x *RenderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderRequest.ProtoReflect.Descriptor instead.
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{14}
}

type RenderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenderResponse) Reset() {
	*x = RenderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderResponse) ProtoMessage() {}

func (x *RenderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderResponse.ProtoReflect.Descriptor instead.
func (*RenderResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{15}
}

type GetAnimationsResponse_Animation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetAnimationsResponse_Animation) Reset() {
	*x = GetAnimationsResponse_Animation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimationsResponse_Animation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimationsResponse_Animation) ProtoMessage() {}

func (x *GetAnimationsResponse_Animation) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimationsResponse_Animation.ProtoReflect.Descriptor instead.
func (*GetAnimationsResponse_Animation) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetAnimationsResponse_Animation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAnimationsResponse_Animation) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GetAnimationsResponse_Animation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetStaticsResponse_Static struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetStaticsResponse_Static) Reset() {
	*x = GetStaticsResponse_Static{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticsResponse_Static) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsResponse_Static) ProtoMessage() {}

func (x *GetStaticsResponse_Static) ProtoReflect() protoreflect.Message {
	mi := &file_xmaspi_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsResponse_Static.ProtoReflect.Descriptor instead.
func (*GetStaticsResponse_Static) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetStaticsResponse_Static) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStaticsResponse_Static) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GetStaticsResponse_Static) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_xmaspi_proto protoreflect.FileDescriptor

var file_xmaspi_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x78, 0x6d, 0x61, 0x73, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a,
	0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x63,
	0x0a, 0x09, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x60, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4c, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x34, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc6, 0x03, 0x0a, 0x06,
	0x58, 0x6d, 0x61, 0x73, 0x50, 0x49, 0x12, 0x29, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64,
	0x12, 0x0e, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x53, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x12, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x6e, 0x44, 0x65, 0x76, 0x2f, 0x78, 0x6d,
	0x61, 0x73, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xmaspi_proto_rawDescOnce sync.Once
	file_xmaspi_proto_rawDescData = file_xmaspi_proto_rawDesc
)

func file_xmaspi_proto_rawDescGZIP() []byte {
	file_xmaspi_proto_rawDescOnce.Do(func() {
		file_xmaspi_proto_rawDescData = protoimpl.X.CompressGZIP(file_xmaspi_proto_rawDescData)
	})
	return file_xmaspi_proto_rawDescData
}

var file_xmaspi_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_xmaspi_proto_goTypes = []interface{}{
	(*SetLedRequest)(nil),                   // 0: SetLedRequest
	(*GetLedCountRequest)(nil),              // 1: GetLedCountRequest
	(*GetLedCountResponse)(nil),             // 2: GetLedCountResponse
	(*GetAnimationsRequest)(nil),            // 3: GetAnimationsRequest
	(*GetAnimationsResponse)(nil),           // 4: GetAnimationsResponse
	(*GetStaticsRequest)(nil),               // 5: GetStaticsRequest
	(*GetStaticsResponse)(nil),              // 6: GetStaticsResponse
	(*SetLedResponse)(nil),                  // 7: SetLedResponse
	(*SetStaticRequest)(nil),                // 8: SetStaticRequest
	(*SetStaticResponse)(nil),               // 9: SetStaticResponse
	(*SetAnimationRequest)(nil),             // 10: SetAnimationRequest
	(*SetAnimationResponse)(nil),            // 11: SetAnimationResponse
	(*ControllerInfoRequest)(nil),           // 12: ControllerInfoRequest
	(*ControllerInfoResponse)(nil),          // 13: ControllerInfoResponse
	(*RenderRequest)(nil),                   // 14: RenderRequest
	(*RenderResponse)(nil),                  // 15: RenderResponse
	(*GetAnimationsResponse_Animation)(nil), // 16: GetAnimationsResponse.Animation
	(*GetStaticsResponse_Static)(nil),       // 17: GetStaticsResponse.Static
}
var file_xmaspi_proto_depIdxs = []int32{
	16, // 0: GetAnimationsResponse.animations:type_name -> GetAnimationsResponse.Animation
	17, // 1: GetStaticsResponse.statics:type_name -> GetStaticsResponse.Static
	0,  // 2: XmasPI.SetLed:input_type -> SetLedRequest
	10, // 3: XmasPI.SetAnimation:input_type -> SetAnimationRequest
	8,  // 4: XmasPI.SetStatic:input_type -> SetStaticRequest
	12, // 5: XmasPI.GetControllerInfo:input_type -> ControllerInfoRequest
	14, // 6: XmasPI.Render:input_type -> RenderRequest
	1,  // 7: XmasPI.GetLedCount:input_type -> GetLedCountRequest
	3,  // 8: XmasPI.GetAnimations:input_type -> GetAnimationsRequest
	5,  // 9: XmasPI.GetStatics:input_type -> GetStaticsRequest
	7,  // 10: XmasPI.SetLed:output_type -> SetLedResponse
	11, // 11: XmasPI.SetAnimation:output_type -> SetAnimationResponse
	9,  // 12: XmasPI.SetStatic:output_type -> SetStaticResponse
	13, // 13: XmasPI.GetControllerInfo:output_type -> ControllerInfoResponse
	15, // 14: XmasPI.Render:output_type -> RenderResponse
	2,  // 15: XmasPI.GetLedCount:output_type -> GetLedCountResponse
	4,  // 16: XmasPI.GetAnimations:output_type -> GetAnimationsResponse
	6,  // 17: XmasPI.GetStatics:output_type -> GetStaticsResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_xmaspi_proto_init() }
func file_xmaspi_proto_init() {
	if File_xmaspi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xmaspi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLedRequest); i {
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
		file_xmaspi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLedCountRequest); i {
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
		file_xmaspi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLedCountResponse); i {
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
		file_xmaspi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimationsRequest); i {
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
		file_xmaspi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimationsResponse); i {
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
		file_xmaspi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticsRequest); i {
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
		file_xmaspi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticsResponse); i {
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
		file_xmaspi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLedResponse); i {
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
		file_xmaspi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStaticRequest); i {
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
		file_xmaspi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStaticResponse); i {
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
		file_xmaspi_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAnimationRequest); i {
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
		file_xmaspi_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAnimationResponse); i {
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
		file_xmaspi_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerInfoRequest); i {
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
		file_xmaspi_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerInfoResponse); i {
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
		file_xmaspi_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderRequest); i {
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
		file_xmaspi_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderResponse); i {
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
		file_xmaspi_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimationsResponse_Animation); i {
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
		file_xmaspi_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticsResponse_Static); i {
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
			RawDescriptor: file_xmaspi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xmaspi_proto_goTypes,
		DependencyIndexes: file_xmaspi_proto_depIdxs,
		MessageInfos:      file_xmaspi_proto_msgTypes,
	}.Build()
	File_xmaspi_proto = out.File
	file_xmaspi_proto_rawDesc = nil
	file_xmaspi_proto_goTypes = nil
	file_xmaspi_proto_depIdxs = nil
}
