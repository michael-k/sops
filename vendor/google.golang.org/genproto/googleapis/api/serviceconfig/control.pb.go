// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/control.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

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

// Selects and configures the service controller used by the service.  The
// service controller handles features like abuse, quota, billing, logging,
// monitoring, etc.
type Control struct {
	// The service control environment to use. If empty, no control plane
	// feature (like quota and billing) will be enabled.
	Environment          string   `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Control) Reset()         { *m = Control{} }
func (m *Control) String() string { return proto.CompactTextString(m) }
func (*Control) ProtoMessage()    {}
func (*Control) Descriptor() ([]byte, []int) {
	return fileDescriptor_74b55b5694b7f0a5, []int{0}
}
func (m *Control) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Control.Unmarshal(m, b)
}
func (m *Control) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Control.Marshal(b, m, deterministic)
}
func (m *Control) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Control.Merge(m, src)
}
func (m *Control) XXX_Size() int {
	return xxx_messageInfo_Control.Size(m)
}
func (m *Control) XXX_DiscardUnknown() {
	xxx_messageInfo_Control.DiscardUnknown(m)
}

var xxx_messageInfo_Control proto.InternalMessageInfo

func (m *Control) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func init() {
	proto.RegisterType((*Control)(nil), "google.api.Control")
}

func init() { proto.RegisterFile("google/api/control.proto", fileDescriptor_74b55b5694b7f0a5) }

var fileDescriptor_74b55b5694b7f0a5 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xe8, 0x25, 0x16, 0x64, 0x2a, 0x69, 0x73, 0xb1,
	0x3b, 0x43, 0x24, 0x85, 0x14, 0xb8, 0xb8, 0x53, 0xf3, 0xca, 0x32, 0x8b, 0xf2, 0xf3, 0x72, 0x53,
	0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x85, 0x9c, 0xf2, 0xb8, 0xf8, 0x92,
	0xf3, 0x73, 0xf5, 0x10, 0xda, 0x9d, 0x78, 0xa0, 0x9a, 0x03, 0x40, 0x06, 0x07, 0x30, 0x46, 0xb9,
	0x42, 0xe5, 0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3,
	0xc0, 0xd6, 0xea, 0x43, 0xa4, 0x12, 0x0b, 0x32, 0x8b, 0xc1, 0x6e, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xb7, 0x46, 0xe1, 0x2d, 0x62, 0x62, 0x71, 0x77,
	0x0c, 0xf0, 0x4c, 0x62, 0x03, 0x6b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x44, 0x6e, 0x78,
	0xbd, 0xcb, 0x00, 0x00, 0x00,
}