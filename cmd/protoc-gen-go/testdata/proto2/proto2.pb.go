// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/proto2.proto

package proto2

import proto "github.com/golang/protobuf/proto"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	I32                  *int32   `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	M                    *Message `protobuf:"bytes,2,opt,name=m" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d756bbe8817c03c1, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() { proto.RegisterFile("proto2/proto2.proto", fileDescriptor_d756bbe8817c03c1) }

var fileDescriptor_d756bbe8817c03c1 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd2, 0x87, 0x50, 0x7a, 0x60, 0x4a, 0x48, 0x34, 0x3d, 0x1f, 0xcc, 0x80, 0x70, 0x93,
	0x21, 0x94, 0x91, 0x92, 0x27, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x00,
	0x17, 0x73, 0xa6, 0xb1, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88, 0x29, 0xa4, 0xc3,
	0xc5, 0x98, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa7, 0x87, 0x55, 0xbf, 0x1e, 0x54,
	0x73, 0x10, 0x63, 0xae, 0x93, 0x75, 0x94, 0x65, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x5e, 0x7a,
	0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc4, 0x0d, 0xfa, 0xc9, 0xb9, 0x29, 0x10,
	0x56, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0xbe, 0x7e, 0x49, 0x6a, 0x71, 0x49, 0x4a, 0x62,
	0x49, 0x22, 0xd4, 0x91, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x68, 0x66, 0x37, 0xb4, 0x00,
	0x00, 0x00,
}
