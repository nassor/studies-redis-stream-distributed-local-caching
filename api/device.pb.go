// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Device struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 string               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	InstalledAt          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=installed_at,json=installedAt,proto3" json:"installed_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ArchivedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=archived_at,json=archivedAt,proto3" json:"archived_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_dde08fe9c9cf9d0d, []int{0}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Device) GetInstalledAt() *timestamp.Timestamp {
	if m != nil {
		return m.InstalledAt
	}
	return nil
}

func (m *Device) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Device) GetArchivedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ArchivedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "api.Device")
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_device_dde08fe9c9cf9d0d) }

var fileDescriptor_device_dde08fe9c9cf9d0d = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x4f,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6,
	0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x54, 0x29, 0x3d, 0x66, 0xe4, 0x62, 0x73, 0x01, 0x6b, 0x13,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0xb6, 0x5c,
	0x3c, 0x99, 0x79, 0xc5, 0x25, 0x89, 0x39, 0x39, 0xa9, 0x29, 0xf1, 0x89, 0x25, 0x12, 0xcc, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0x6b, 0xf4, 0x60, 0xd6, 0xe8, 0x85, 0xc0, 0xac, 0x09,
	0xe2, 0x86, 0xab, 0x77, 0x2c, 0x11, 0xb2, 0xe4, 0xe2, 0x2a, 0x2d, 0x48, 0x49, 0x2c, 0x81, 0x68,
	0x66, 0x21, 0xa8, 0x99, 0x13, 0xaa, 0xda, 0xb1, 0x44, 0xc8, 0x9a, 0x8b, 0x3b, 0xb1, 0x28, 0x39,
	0x23, 0xb3, 0x0c, 0xa2, 0x97, 0x95, 0xa0, 0x5e, 0x2e, 0x98, 0x72, 0xc7, 0x12, 0x27, 0x76, 0x0f,
	0xc6, 0x28, 0x50, 0x78, 0x24, 0xb1, 0x81, 0x15, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76,
	0xaa, 0x92, 0x65, 0x2b, 0x01, 0x00, 0x00,
}
