// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: activation.proto

package ws

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ActivationKind int32

const (
	ActivationKind_UNKNOWN      ActivationKind = 0
	ActivationKind_MOL          ActivationKind = 1
	ActivationKind_PRIORITY     ActivationKind = 2
	ActivationKind_PROPORTIONAL ActivationKind = 3
)

// Enum value maps for ActivationKind.
var (
	ActivationKind_name = map[int32]string{
		0: "UNKNOWN",
		1: "MOL",
		2: "PRIORITY",
		3: "PROPORTIONAL",
	}
	ActivationKind_value = map[string]int32{
		"UNKNOWN":      0,
		"MOL":          1,
		"PRIORITY":     2,
		"PROPORTIONAL": 3,
	}
)

func (x ActivationKind) Enum() *ActivationKind {
	p := new(ActivationKind)
	*p = x
	return p
}

func (x ActivationKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivationKind) Descriptor() protoreflect.EnumDescriptor {
	return file_activation_proto_enumTypes[0].Descriptor()
}

func (ActivationKind) Type() protoreflect.EnumType {
	return &file_activation_proto_enumTypes[0]
}

func (x ActivationKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivationKind.Descriptor instead.
func (ActivationKind) EnumDescriptor() ([]byte, []int) {
	return file_activation_proto_rawDescGZIP(), []int{0}
}

type Activation_ActivationState int32

const (
	Activation_UNKNOWN   Activation_ActivationState = 0
	Activation_PENDING   Activation_ActivationState = 1
	Activation_ACCEPTED  Activation_ActivationState = 2
	Activation_DECLINED  Activation_ActivationState = 3
	Activation_FINISHED  Activation_ActivationState = 4
	Activation_PASSED_ON Activation_ActivationState = 5
)

// Enum value maps for Activation_ActivationState.
var (
	Activation_ActivationState_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "ACCEPTED",
		3: "DECLINED",
		4: "FINISHED",
		5: "PASSED_ON",
	}
	Activation_ActivationState_value = map[string]int32{
		"UNKNOWN":   0,
		"PENDING":   1,
		"ACCEPTED":  2,
		"DECLINED":  3,
		"FINISHED":  4,
		"PASSED_ON": 5,
	}
)

func (x Activation_ActivationState) Enum() *Activation_ActivationState {
	p := new(Activation_ActivationState)
	*p = x
	return p
}

func (x Activation_ActivationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activation_ActivationState) Descriptor() protoreflect.EnumDescriptor {
	return file_activation_proto_enumTypes[1].Descriptor()
}

func (Activation_ActivationState) Type() protoreflect.EnumType {
	return &file_activation_proto_enumTypes[1]
}

func (x Activation_ActivationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activation_ActivationState.Descriptor instead.
func (Activation_ActivationState) EnumDescriptor() ([]byte, []int) {
	return file_activation_proto_rawDescGZIP(), []int{0, 0}
}

type Activation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion                int32                      `protobuf:"varint,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	ActivationId              int32                      `protobuf:"varint,2,opt,name=activationId,proto3" json:"activationId,omitempty"`
	GenerationId              int32                      `protobuf:"varint,3,opt,name=generationId,proto3" json:"generationId,omitempty"`
	GenerationName            string                     `protobuf:"bytes,4,opt,name=generationName,proto3" json:"generationName,omitempty"`
	State                     Activation_ActivationState `protobuf:"varint,5,opt,name=state,proto3,enum=ws.Activation_ActivationState" json:"state,omitempty"`
	ActivationDatetime        *timestamppb.Timestamp     `protobuf:"bytes,6,opt,name=activationDatetime,proto3" json:"activationDatetime,omitempty"`
	GenerationWCode           string                     `protobuf:"bytes,7,opt,name=generationWCode,proto3" json:"generationWCode,omitempty"`
	ActivationKind            ActivationKind             `protobuf:"varint,8,opt,name=activationKind,proto3,enum=ws.ActivationKind" json:"activationKind,omitempty"`
	MarketAreaId              int32                      `protobuf:"varint,9,opt,name=marketAreaId,proto3" json:"marketAreaId,omitempty"`
	IsActivation              bool                       `protobuf:"varint,10,opt,name=isActivation,proto3" json:"isActivation,omitempty"`
	Volume                    int32                      `protobuf:"varint,11,opt,name=volume,proto3" json:"volume,omitempty"`
	DsoId                     int32                      `protobuf:"varint,12,opt,name=dsoId,proto3" json:"dsoId,omitempty"`
	HasReplacedActivation     bool                       `protobuf:"varint,15,opt,name=hasReplacedActivation,proto3" json:"hasReplacedActivation,omitempty"`
	PreviousVolume            int32                      `protobuf:"varint,16,opt,name=previousVolume,proto3" json:"previousVolume,omitempty"`
	DeltaVolume               int32                      `protobuf:"varint,17,opt,name=deltaVolume,proto3" json:"deltaVolume,omitempty"`
	InitialActivationDatetime *timestamppb.Timestamp     `protobuf:"bytes,18,opt,name=initialActivationDatetime,proto3" json:"initialActivationDatetime,omitempty"`
	Comment                   string                     `protobuf:"bytes,19,opt,name=comment,proto3" json:"comment,omitempty"`
	SeenByResponsibleUser     bool                       `protobuf:"varint,20,opt,name=seenByResponsibleUser,proto3" json:"seenByResponsibleUser,omitempty"`
}

func (x *Activation) Reset() {
	*x = Activation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activation) ProtoMessage() {}

func (x *Activation) ProtoReflect() protoreflect.Message {
	mi := &file_activation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activation.ProtoReflect.Descriptor instead.
func (*Activation) Descriptor() ([]byte, []int) {
	return file_activation_proto_rawDescGZIP(), []int{0}
}

func (x *Activation) GetApiVersion() int32 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *Activation) GetActivationId() int32 {
	if x != nil {
		return x.ActivationId
	}
	return 0
}

func (x *Activation) GetGenerationId() int32 {
	if x != nil {
		return x.GenerationId
	}
	return 0
}

func (x *Activation) GetGenerationName() string {
	if x != nil {
		return x.GenerationName
	}
	return ""
}

func (x *Activation) GetState() Activation_ActivationState {
	if x != nil {
		return x.State
	}
	return Activation_UNKNOWN
}

func (x *Activation) GetActivationDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.ActivationDatetime
	}
	return nil
}

func (x *Activation) GetGenerationWCode() string {
	if x != nil {
		return x.GenerationWCode
	}
	return ""
}

func (x *Activation) GetActivationKind() ActivationKind {
	if x != nil {
		return x.ActivationKind
	}
	return ActivationKind_UNKNOWN
}

func (x *Activation) GetMarketAreaId() int32 {
	if x != nil {
		return x.MarketAreaId
	}
	return 0
}

func (x *Activation) GetIsActivation() bool {
	if x != nil {
		return x.IsActivation
	}
	return false
}

func (x *Activation) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Activation) GetDsoId() int32 {
	if x != nil {
		return x.DsoId
	}
	return 0
}

func (x *Activation) GetHasReplacedActivation() bool {
	if x != nil {
		return x.HasReplacedActivation
	}
	return false
}

func (x *Activation) GetPreviousVolume() int32 {
	if x != nil {
		return x.PreviousVolume
	}
	return 0
}

func (x *Activation) GetDeltaVolume() int32 {
	if x != nil {
		return x.DeltaVolume
	}
	return 0
}

func (x *Activation) GetInitialActivationDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.InitialActivationDatetime
	}
	return nil
}

func (x *Activation) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Activation) GetSeenByResponsibleUser() bool {
	if x != nil {
		return x.SeenByResponsibleUser
	}
	return false
}

var File_activation_proto protoreflect.FileDescriptor

var file_activation_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x77, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x07, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x77, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x12,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x69, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x77, 0x73, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x73, 0x6f, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64,
	0x73, 0x6f, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x68, 0x61, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x68, 0x61, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x65, 0x65, 0x6e,
	0x42, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x73, 0x65, 0x65, 0x6e, 0x42, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x64,
	0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x43,
	0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x5f,
	0x4f, 0x4e, 0x10, 0x05, 0x2a, 0x46, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52,
	0x4f, 0x50, 0x4f, 0x52, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x42, 0x10, 0x5a, 0x0e,
	0x6e, 0x69, 0x6b, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x77, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activation_proto_rawDescOnce sync.Once
	file_activation_proto_rawDescData = file_activation_proto_rawDesc
)

func file_activation_proto_rawDescGZIP() []byte {
	file_activation_proto_rawDescOnce.Do(func() {
		file_activation_proto_rawDescData = protoimpl.X.CompressGZIP(file_activation_proto_rawDescData)
	})
	return file_activation_proto_rawDescData
}

var file_activation_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_activation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_activation_proto_goTypes = []interface{}{
	(ActivationKind)(0),             // 0: ws.ActivationKind
	(Activation_ActivationState)(0), // 1: ws.Activation.ActivationState
	(*Activation)(nil),              // 2: ws.Activation
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_activation_proto_depIdxs = []int32{
	1, // 0: ws.Activation.state:type_name -> ws.Activation.ActivationState
	3, // 1: ws.Activation.activationDatetime:type_name -> google.protobuf.Timestamp
	0, // 2: ws.Activation.activationKind:type_name -> ws.ActivationKind
	3, // 3: ws.Activation.initialActivationDatetime:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_activation_proto_init() }
func file_activation_proto_init() {
	if File_activation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activation); i {
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
			RawDescriptor: file_activation_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_activation_proto_goTypes,
		DependencyIndexes: file_activation_proto_depIdxs,
		EnumInfos:         file_activation_proto_enumTypes,
		MessageInfos:      file_activation_proto_msgTypes,
	}.Build()
	File_activation_proto = out.File
	file_activation_proto_rawDesc = nil
	file_activation_proto_goTypes = nil
	file_activation_proto_depIdxs = nil
}