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

	Led   int32  `protobuf:"varint,1,opt,name=led,proto3" json:"led,omitempty"`
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

func (x *SetLedRequest) GetLed() int32 {
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

type SetLedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetLedResponse) Reset() {
	*x = SetLedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLedResponse) ProtoMessage() {}

func (x *SetLedResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetLedResponse.ProtoReflect.Descriptor instead.
func (*SetLedResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{1}
}

type SetStaticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color uint32 `protobuf:"varint,1,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *SetStaticRequest) Reset() {
	*x = SetStaticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStaticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStaticRequest) ProtoMessage() {}

func (x *SetStaticRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetStaticRequest.ProtoReflect.Descriptor instead.
func (*SetStaticRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{2}
}

func (x *SetStaticRequest) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

type SetStaticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetStaticResponse) Reset() {
	*x = SetStaticResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStaticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStaticResponse) ProtoMessage() {}

func (x *SetStaticResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetStaticResponse.ProtoReflect.Descriptor instead.
func (*SetStaticResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{3}
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
		mi := &file_xmaspi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAnimationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAnimationRequest) ProtoMessage() {}

func (x *SetAnimationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetAnimationRequest.ProtoReflect.Descriptor instead.
func (*SetAnimationRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{4}
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
		mi := &file_xmaspi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAnimationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAnimationResponse) ProtoMessage() {}

func (x *SetAnimationResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetAnimationResponse.ProtoReflect.Descriptor instead.
func (*SetAnimationResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{5}
}

type ControllerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ControllerInfoRequest) Reset() {
	*x = ControllerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmaspi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerInfoRequest) ProtoMessage() {}

func (x *ControllerInfoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ControllerInfoRequest.ProtoReflect.Descriptor instead.
func (*ControllerInfoRequest) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{6}
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
		mi := &file_xmaspi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerInfoResponse) ProtoMessage() {}

func (x *ControllerInfoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ControllerInfoResponse.ProtoReflect.Descriptor instead.
func (*ControllerInfoResponse) Descriptor() ([]byte, []int) {
	return file_xmaspi_proto_rawDescGZIP(), []int{7}
}

func (x *ControllerInfoResponse) GetLedCount() int64 {
	if x != nil {
		return x.LedCount
	}
	return 0
}

var File_xmaspi_proto protoreflect.FileDescriptor

var file_xmaspi_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x78, 0x6d, 0x61, 0x73, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4c, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xea, 0x01, 0x0a, 0x06, 0x58,
	0x6d, 0x61, 0x73, 0x50, 0x49, 0x12, 0x29, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x12,
	0x0e, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x53, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x6e, 0x44, 0x65, 0x76,
	0x2f, 0x78, 0x6d, 0x61, 0x73, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_xmaspi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_xmaspi_proto_goTypes = []interface{}{
	(*SetLedRequest)(nil),          // 0: SetLedRequest
	(*SetLedResponse)(nil),         // 1: SetLedResponse
	(*SetStaticRequest)(nil),       // 2: SetStaticRequest
	(*SetStaticResponse)(nil),      // 3: SetStaticResponse
	(*SetAnimationRequest)(nil),    // 4: SetAnimationRequest
	(*SetAnimationResponse)(nil),   // 5: SetAnimationResponse
	(*ControllerInfoRequest)(nil),  // 6: ControllerInfoRequest
	(*ControllerInfoResponse)(nil), // 7: ControllerInfoResponse
}
var file_xmaspi_proto_depIdxs = []int32{
	0, // 0: XmasPI.SetLed:input_type -> SetLedRequest
	4, // 1: XmasPI.SetAnimation:input_type -> SetAnimationRequest
	2, // 2: XmasPI.SetStatic:input_type -> SetStaticRequest
	6, // 3: XmasPI.GetControllerInfo:input_type -> ControllerInfoRequest
	1, // 4: XmasPI.SetLed:output_type -> SetLedResponse
	5, // 5: XmasPI.SetAnimation:output_type -> SetAnimationResponse
	3, // 6: XmasPI.SetStatic:output_type -> SetStaticResponse
	7, // 7: XmasPI.GetControllerInfo:output_type -> ControllerInfoResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
		file_xmaspi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_xmaspi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_xmaspi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_xmaspi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_xmaspi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_xmaspi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xmaspi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
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
