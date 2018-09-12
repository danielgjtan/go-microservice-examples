// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/auth/proto/auth.proto

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_57e0e2f691f4a467, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type Result struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_57e0e2f691f4a467, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type Customer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthToken            string   `protobuf:"bytes,2,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_57e0e2f691f4a467, []int{2}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Result)(nil), "auth.Result")
	proto.RegisterType((*Customer)(nil), "auth.Customer")
}

func init() { proto.RegisterFile("srv/auth/proto/auth.proto", fileDescriptor_auth_57e0e2f691f4a467) }

var fileDescriptor_auth_57e0e2f691f4a467 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x2e, 0x2a, 0xd3,
	0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0x33, 0xf5, 0xc0, 0x4c, 0x21,
	0x16, 0x10, 0x5b, 0x49, 0x9d, 0x8b, 0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x86,
	0x8b, 0x13, 0x24, 0x14, 0x92, 0x9f, 0x9d, 0x9a, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84,
	0x10, 0x50, 0x32, 0xe1, 0x62, 0x0b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0xd2, 0xe2, 0xe2, 0x48,
	0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2d, 0x02, 0x2b, 0xe3, 0x36, 0xe2, 0xd3, 0x03, 0x9b, 0xeb,
	0x0c, 0x15, 0x0d, 0x82, 0xcb, 0x2b, 0x59, 0x70, 0x71, 0xc0, 0x44, 0x85, 0xf8, 0xb8, 0x98, 0x32,
	0x53, 0xc0, 0x3a, 0x58, 0x83, 0x98, 0x32, 0x53, 0x50, 0xed, 0x63, 0x42, 0xb3, 0xcf, 0xc8, 0x84,
	0x8b, 0xc5, 0xb1, 0xb4, 0x24, 0x43, 0x48, 0x87, 0x8b, 0x3b, 0x2c, 0xb5, 0x28, 0x33, 0xad, 0x12,
	0x2c, 0x2c, 0xc4, 0x0b, 0xb1, 0x0a, 0xea, 0x66, 0x29, 0x1e, 0x18, 0x17, 0xe4, 0x32, 0x25, 0x86,
	0x24, 0x36, 0xb0, 0xdf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x53, 0x1b, 0x1d, 0xf8,
	0x00, 0x00, 0x00,
}
