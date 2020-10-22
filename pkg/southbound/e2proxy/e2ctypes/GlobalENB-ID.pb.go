// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GlobalENB-ID.proto

package e2ctypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GlobalENB_IDT struct {
	PLMN_Identity        string   `protobuf:"bytes,1,opt,name=pLMN_Identity,json=pLMNIdentity,proto3" json:"pLMN_Identity,omitempty"`
	ENB_ID               *ENB_IDT `protobuf:"bytes,2,opt,name=eNB_ID,json=eNBID,proto3" json:"eNB_ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalENB_IDT) Reset()         { *m = GlobalENB_IDT{} }
func (m *GlobalENB_IDT) String() string { return proto.CompactTextString(m) }
func (*GlobalENB_IDT) ProtoMessage()    {}
func (*GlobalENB_IDT) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c2baa01b30192c, []int{0}
}

func (m *GlobalENB_IDT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalENB_IDT.Unmarshal(m, b)
}
func (m *GlobalENB_IDT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalENB_IDT.Marshal(b, m, deterministic)
}
func (m *GlobalENB_IDT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalENB_IDT.Merge(m, src)
}
func (m *GlobalENB_IDT) XXX_Size() int {
	return xxx_messageInfo_GlobalENB_IDT.Size(m)
}
func (m *GlobalENB_IDT) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalENB_IDT.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalENB_IDT proto.InternalMessageInfo

func (m *GlobalENB_IDT) GetPLMN_Identity() string {
	if m != nil {
		return m.PLMN_Identity
	}
	return ""
}

func (m *GlobalENB_IDT) GetENB_ID() *ENB_IDT {
	if m != nil {
		return m.ENB_ID
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalENB_IDT)(nil), "e2ctypes.GlobalENB_ID_t")
}

func init() { proto.RegisterFile("GlobalENB-ID.proto", fileDescriptor_b5c2baa01b30192c) }

var fileDescriptor_b5c2baa01b30192c = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x72, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x71, 0xf5, 0x73, 0xd2, 0xf5, 0x74, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xe2, 0x41, 0x16, 0x57, 0x4a, 0xe0, 0xe2,
	0x83, 0xab, 0x8e, 0xf7, 0x74, 0x89, 0x2f, 0x11, 0x52, 0xe6, 0xe2, 0x2d, 0xf0, 0xf1, 0xf5, 0x8b,
	0xf7, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2,
	0x01, 0x09, 0xc2, 0xc4, 0x84, 0x34, 0xb9, 0xd8, 0x52, 0xc1, 0x1a, 0x24, 0x98, 0x14, 0x18, 0x35,
	0xb8, 0x8d, 0x84, 0xf4, 0x60, 0xe6, 0xeb, 0xc1, 0x0c, 0x0a, 0x62, 0x4d, 0xf5, 0x73, 0xf2, 0x74,
	0x49, 0x62, 0x03, 0x5b, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x53, 0xc4, 0xc8, 0x96,
	0x00, 0x00, 0x00,
}