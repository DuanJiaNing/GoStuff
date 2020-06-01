// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2.proto

package example

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

type Person2 struct {
	Value                *float32 `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person2) Reset()         { *m = Person2{} }
func (m *Person2) String() string { return proto.CompactTextString(m) }
func (*Person2) ProtoMessage()    {}
func (*Person2) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f509089572db8e7, []int{0}
}

func (m *Person2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person2.Unmarshal(m, b)
}
func (m *Person2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person2.Marshal(b, m, deterministic)
}
func (m *Person2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person2.Merge(m, src)
}
func (m *Person2) XXX_Size() int {
	return xxx_messageInfo_Person2.Size(m)
}
func (m *Person2) XXX_DiscardUnknown() {
	xxx_messageInfo_Person2.DiscardUnknown(m)
}

var xxx_messageInfo_Person2 proto.InternalMessageInfo

func (m *Person2) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

// p.value=nil -> json: {"p":{}}
// p.value=0   -> json: {"p":{"value":0}}
type Data2 struct {
	P                    *Person2 `protobuf:"bytes,1,req,name=p" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data2) Reset()         { *m = Data2{} }
func (m *Data2) String() string { return proto.CompactTextString(m) }
func (*Data2) ProtoMessage()    {}
func (*Data2) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f509089572db8e7, []int{1}
}

func (m *Data2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data2.Unmarshal(m, b)
}
func (m *Data2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data2.Marshal(b, m, deterministic)
}
func (m *Data2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data2.Merge(m, src)
}
func (m *Data2) XXX_Size() int {
	return xxx_messageInfo_Data2.Size(m)
}
func (m *Data2) XXX_DiscardUnknown() {
	xxx_messageInfo_Data2.DiscardUnknown(m)
}

var xxx_messageInfo_Data2 proto.InternalMessageInfo

func (m *Data2) GetP() *Person2 {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*Person2)(nil), "example.Person2")
	proto.RegisterType((*Data2)(nil), "example.Data2")
}

func init() { proto.RegisterFile("proto2.proto", fileDescriptor_1f509089572db8e7) }

var fileDescriptor_1f509089572db8e7 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd2, 0x03, 0x53, 0x42, 0xec, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0x4a, 0xf2,
	0x5c, 0xec, 0x01, 0xa9, 0x45, 0xc5, 0xf9, 0x79, 0x46, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39,
	0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x4c, 0x41, 0x10, 0x8e, 0x92, 0x3a, 0x17, 0xab, 0x4b,
	0x62, 0x49, 0xa2, 0x91, 0x90, 0x1c, 0x17, 0x63, 0x81, 0x04, 0xa3, 0x02, 0x93, 0x06, 0xb7, 0x91,
	0x80, 0x1e, 0x54, 0xbb, 0x1e, 0x54, 0x6f, 0x10, 0x63, 0x01, 0x20, 0x00, 0x00, 0xff, 0xff, 0x93,
	0xe9, 0x03, 0xb0, 0x61, 0x00, 0x00, 0x00,
}
