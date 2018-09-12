// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/geo/proto/geo.proto

package geo

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

// The latitude and longitude of the current location.
type Request struct {
	Lat                  float32  `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float32  `protobuf:"fixed32,2,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_geo_e8a14454c7648c01, []int{0}
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

func (m *Request) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Request) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Result struct {
	HotelIds             []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_geo_e8a14454c7648c01, []int{1}
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

func (m *Result) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "geo.Request")
	proto.RegisterType((*Result)(nil), "geo.Result")
}

func init() { proto.RegisterFile("srv/geo/proto/geo.proto", fileDescriptor_geo_e8a14454c7648c01) }

var fileDescriptor_geo_e8a14454c7648c01 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2e, 0x2a, 0xd3,
	0x4f, 0x4f, 0xcd, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0xb1, 0xf4, 0xc0, 0x2c, 0x21, 0xe6,
	0xf4, 0xd4, 0x7c, 0x25, 0x5d, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01,
	0x2e, 0xe6, 0x9c, 0xc4, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xa6, 0x20, 0x10, 0x13, 0x2c, 0x92,
	0x9f, 0x27, 0xc1, 0x04, 0x15, 0xc9, 0xcf, 0x53, 0x52, 0xe1, 0x62, 0x0b, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0xc8, 0x2f, 0x49, 0xcd, 0xf1, 0x4c, 0x29, 0x96, 0x60, 0x54,
	0x60, 0xd6, 0xe0, 0x0c, 0x82, 0xf3, 0x8d, 0xb4, 0xb8, 0x98, 0xdd, 0x53, 0xf3, 0x85, 0x94, 0xb9,
	0xd8, 0xfc, 0x52, 0x13, 0x8b, 0x92, 0x2a, 0x85, 0x78, 0xf4, 0x40, 0xd6, 0x42, 0x2d, 0x92, 0xe2,
	0x86, 0xf2, 0x40, 0xe6, 0x24, 0xb1, 0x81, 0x1d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xa0, 0xba, 0x03, 0xa7, 0x00, 0x00, 0x00,
}
