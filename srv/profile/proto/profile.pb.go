// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/profile/proto/profile.proto

package profile

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
	HotelIds             []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	Locale               string   `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_e9734da63d526c41, []int{0}
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

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type Result struct {
	Hotels               []*Hotel `protobuf:"bytes,1,rep,name=hotels,proto3" json:"hotels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_e9734da63d526c41, []int{1}
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

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Address              *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Images               []*Image `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hotel) Reset()         { *m = Hotel{} }
func (m *Hotel) String() string { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()    {}
func (*Hotel) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_e9734da63d526c41, []int{2}
}
func (m *Hotel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hotel.Unmarshal(m, b)
}
func (m *Hotel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hotel.Marshal(b, m, deterministic)
}
func (dst *Hotel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hotel.Merge(dst, src)
}
func (m *Hotel) XXX_Size() int {
	return xxx_messageInfo_Hotel.Size(m)
}
func (m *Hotel) XXX_DiscardUnknown() {
	xxx_messageInfo_Hotel.DiscardUnknown(m)
}

var xxx_messageInfo_Hotel proto.InternalMessageInfo

func (m *Hotel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hotel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hotel) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Hotel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber         string   `protobuf:"bytes,1,opt,name=streetNumber,proto3" json:"streetNumber,omitempty"`
	StreetName           string   `protobuf:"bytes,2,opt,name=streetName,proto3" json:"streetName,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	PostalCode           string   `protobuf:"bytes,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_e9734da63d526c41, []int{3}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetStreetNumber() string {
	if m != nil {
		return m.StreetNumber
	}
	return ""
}

func (m *Address) GetStreetName() string {
	if m != nil {
		return m.StreetName
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

type Image struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Default              bool     `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_e9734da63d526c41, []int{4}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
}

func init() {
	proto.RegisterFile("srv/profile/proto/profile.proto", fileDescriptor_profile_e9734da63d526c41)
}

var fileDescriptor_profile_e9734da63d526c41 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0xae, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0xb6, 0x49, 0xda, 0x09, 0x2a, 0xd5, 0x08, 0x21, 0xab, 0x0b, 0x88, 0xb2, 0x40,
	0x15, 0x8b, 0x52, 0xb5, 0x4b, 0xc4, 0x02, 0xb1, 0x80, 0x6e, 0x10, 0xf2, 0x0d, 0xd2, 0x78, 0x4a,
	0x23, 0xb9, 0x71, 0xb0, 0x1d, 0xa4, 0x1e, 0x8b, 0x33, 0x70, 0xb1, 0x27, 0x3b, 0x4e, 0x9b, 0xf7,
	0x56, 0x9e, 0xff, 0x9b, 0xb1, 0x67, 0x7e, 0x6b, 0xe0, 0xbd, 0xd1, 0x7f, 0x3f, 0xb5, 0x5a, 0x9d,
	0x6b, 0x49, 0xee, 0xb4, 0x6a, 0x50, 0x5b, 0xaf, 0x30, 0x0d, 0xb2, 0xf8, 0x02, 0x29, 0xa7, 0x3f,
	0x1d, 0x19, 0x8b, 0x6b, 0x98, 0x5f, 0x94, 0x25, 0x79, 0x14, 0x86, 0x45, 0xf9, 0x74, 0xb3, 0xe0,
	0x77, 0x8d, 0x6f, 0x21, 0x91, 0xaa, 0x2a, 0x25, 0xb1, 0x49, 0x1e, 0x6d, 0x16, 0x3c, 0xa8, 0x62,
	0x07, 0x09, 0x27, 0xd3, 0x49, 0x8b, 0x1f, 0x20, 0xf1, 0xd5, 0xfd, 0xdd, 0x6c, 0xbf, 0xdc, 0x0e,
	0x1d, 0x7f, 0x38, 0xcc, 0x43, 0xb6, 0xf8, 0x1f, 0x41, 0xec, 0x09, 0x2e, 0x61, 0x52, 0x0b, 0x16,
	0xf9, 0xf7, 0x26, 0xb5, 0x40, 0x84, 0x59, 0x53, 0x5e, 0x87, 0x0e, 0x3e, 0xc6, 0x1c, 0xb2, 0xf6,
	0xa2, 0x1a, 0xfa, 0xd9, 0x5d, 0x4f, 0xa4, 0xd9, 0xd4, 0xa7, 0xc6, 0xc8, 0x55, 0x08, 0x32, 0x95,
	0xae, 0x5b, 0x5b, 0xab, 0x86, 0xcd, 0xfa, 0x8a, 0x11, 0xc2, 0x8f, 0x90, 0x96, 0x42, 0x68, 0x32,
	0x86, 0xc5, 0x79, 0xb4, 0xc9, 0xf6, 0xab, 0xfb, 0x68, 0x5f, 0x7b, 0xce, 0x87, 0x02, 0xe7, 0xa2,
	0xbe, 0x96, 0xbf, 0xc9, 0xb0, 0xe4, 0x85, 0x8b, 0xa3, 0xc3, 0x3c, 0x64, 0x8b, 0x7f, 0x11, 0xa4,
	0xe1, 0x32, 0x16, 0xf0, 0xca, 0x58, 0x4d, 0x64, 0xc3, 0x90, 0xbd, 0xa3, 0x67, 0x0c, 0xdf, 0x01,
	0x04, 0xfd, 0x70, 0x38, 0x22, 0xce, 0x7b, 0x55, 0xdb, 0x5b, 0x30, 0xe8, 0x63, 0x7c, 0x03, 0xb1,
	0xb1, 0xa5, 0xa5, 0xe0, 0xa9, 0x17, 0xc8, 0x20, 0xad, 0x54, 0xd7, 0x58, 0x7d, 0xf3, 0x6e, 0x16,
	0x7c, 0x90, 0xae, 0x47, 0xab, 0x8c, 0x2d, 0xe5, 0x37, 0x25, 0x88, 0x25, 0x7d, 0x8f, 0x07, 0x29,
	0x0e, 0x10, 0x7b, 0x13, 0xb8, 0x82, 0x69, 0xa7, 0x65, 0x98, 0xd3, 0x85, 0xee, 0x51, 0x41, 0xe7,
	0xb2, 0x93, 0xd6, 0xcf, 0x36, 0xe7, 0x83, 0xdc, 0x7f, 0x86, 0xf4, 0x57, 0xff, 0x03, 0xb8, 0x83,
	0xec, 0x3b, 0xd9, 0xa0, 0x0c, 0x3e, 0x7e, 0x31, 0x2c, 0xd0, 0xfa, 0xf5, 0x88, 0xb8, 0x9d, 0x38,
	0x25, 0x7e, 0xd9, 0x0e, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x6c, 0xd9, 0x83, 0x8f, 0x02,
	0x00, 0x00,
}
