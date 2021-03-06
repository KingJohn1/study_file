// Code generated by protoc-gen-go. DO NOT EDIT.
// source: liu.proto

package packet

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Header struct {
	MessageId            *string  `protobuf:"bytes,1,req,name=messageId" json:"messageId,omitempty"`
	Topic                *string  `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_71f36ce712fa1479, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Header) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

type BytesMessage struct {
	Header               *Header  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesMessage) Reset()         { *m = BytesMessage{} }
func (m *BytesMessage) String() string { return proto.CompactTextString(m) }
func (*BytesMessage) ProtoMessage()    {}
func (*BytesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_71f36ce712fa1479, []int{1}
}

func (m *BytesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesMessage.Unmarshal(m, b)
}
func (m *BytesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesMessage.Marshal(b, m, deterministic)
}
func (m *BytesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesMessage.Merge(m, src)
}
func (m *BytesMessage) XXX_Size() int {
	return xxx_messageInfo_BytesMessage.Size(m)
}
func (m *BytesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BytesMessage proto.InternalMessageInfo

func (m *BytesMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BytesMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type StringMessage struct {
	Header               *Header  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body                 *string  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_71f36ce712fa1479, []int{2}
}

func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (m *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(m, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *StringMessage) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Header)(nil), "packet.Header")
	proto.RegisterType((*BytesMessage)(nil), "packet.BytesMessage")
	proto.RegisterType((*StringMessage)(nil), "packet.StringMessage")
}

func init() { proto.RegisterFile("liu.proto", fileDescriptor_71f36ce712fa1479) }

var fileDescriptor_71f36ce712fa1479 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xc9, 0x2c, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x51, 0xb2, 0xe1,
	0x62, 0xf3, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x12, 0x92, 0xe1, 0xe2, 0xcc, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x4f, 0xf5, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x89, 0x70,
	0xb1, 0x96, 0xe4, 0x17, 0x64, 0x26, 0x4b, 0x30, 0x81, 0x65, 0x20, 0x1c, 0x25, 0x2f, 0x2e, 0x1e,
	0xa7, 0xca, 0x92, 0xd4, 0x62, 0x5f, 0x88, 0x3a, 0x21, 0x35, 0x2e, 0xb6, 0x0c, 0xb0, 0x69, 0x60,
	0x03, 0xb8, 0x8d, 0xf8, 0xf4, 0x20, 0xd6, 0xe8, 0x41, 0xec, 0x08, 0x82, 0xca, 0x0a, 0x09, 0x71,
	0xb1, 0x24, 0xe5, 0xa7, 0x54, 0x82, 0x0d, 0xe3, 0x09, 0x02, 0xb3, 0x95, 0xbc, 0xb9, 0x78, 0x83,
	0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x29, 0x31, 0x8c, 0x13, 0x62, 0x18, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x00, 0x38, 0x0e, 0xe8, 0xea, 0x00, 0x00, 0x00,
}
