// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloud_dmx.proto

package clouddmxpb

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

type CloudDmx_Type int32

const (
	CloudDmx_PING     CloudDmx_Type = 0
	CloudDmx_DELTA    CloudDmx_Type = 1
	CloudDmx_COMPLETE CloudDmx_Type = 2
)

var CloudDmx_Type_name = map[int32]string{
	0: "PING",
	1: "DELTA",
	2: "COMPLETE",
}

var CloudDmx_Type_value = map[string]int32{
	"PING":     0,
	"DELTA":    1,
	"COMPLETE": 2,
}

func (x CloudDmx_Type) Enum() *CloudDmx_Type {
	p := new(CloudDmx_Type)
	*p = x
	return p
}

func (x CloudDmx_Type) String() string {
	return proto.EnumName(CloudDmx_Type_name, int32(x))
}

func (x *CloudDmx_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CloudDmx_Type_value, data, "CloudDmx_Type")
	if err != nil {
		return err
	}
	*x = CloudDmx_Type(value)
	return nil
}

func (CloudDmx_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7c7f7f1616c4b6e, []int{0, 0}
}

type CloudDmx struct {
	Type                 *CloudDmx_Type `protobuf:"varint,1,req,name=type,enum=CloudDmx_Type" json:"type,omitempty"`
	Channels             []uint32       `protobuf:"varint,2,rep,name=channels" json:"channels,omitempty"`
	Values               []uint32       `protobuf:"varint,3,rep,name=values" json:"values,omitempty"`
	Slots                []byte         `protobuf:"bytes,4,opt,name=slots" json:"slots,omitempty"`
	Start                *uint32        `protobuf:"varint,5,opt,name=start" json:"start,omitempty"`
	Id                   *string        `protobuf:"bytes,6,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CloudDmx) Reset()         { *m = CloudDmx{} }
func (m *CloudDmx) String() string { return proto.CompactTextString(m) }
func (*CloudDmx) ProtoMessage()    {}
func (*CloudDmx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7f7f1616c4b6e, []int{0}
}

func (m *CloudDmx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudDmx.Unmarshal(m, b)
}
func (m *CloudDmx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudDmx.Marshal(b, m, deterministic)
}
func (m *CloudDmx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudDmx.Merge(m, src)
}
func (m *CloudDmx) XXX_Size() int {
	return xxx_messageInfo_CloudDmx.Size(m)
}
func (m *CloudDmx) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudDmx.DiscardUnknown(m)
}

var xxx_messageInfo_CloudDmx proto.InternalMessageInfo

func (m *CloudDmx) GetType() CloudDmx_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CloudDmx_PING
}

func (m *CloudDmx) GetChannels() []uint32 {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *CloudDmx) GetValues() []uint32 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *CloudDmx) GetSlots() []byte {
	if m != nil {
		return m.Slots
	}
	return nil
}

func (m *CloudDmx) GetStart() uint32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *CloudDmx) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("CloudDmx_Type", CloudDmx_Type_name, CloudDmx_Type_value)
	proto.RegisterType((*CloudDmx)(nil), "cloudDmx")
}

func init() { proto.RegisterFile("cloud_dmx.proto", fileDescriptor_f7c7f7f1616c4b6e) }

var fileDescriptor_f7c7f7f1616c4b6e = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xc9, 0x2f,
	0x4d, 0x89, 0x4f, 0xc9, 0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x9a, 0xcf, 0xc8, 0xc5,
	0x01, 0x16, 0x73, 0xc9, 0xad, 0x10, 0x92, 0xe1, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54,
	0x60, 0xd2, 0xe0, 0x33, 0xe2, 0xd3, 0x83, 0x49, 0xe8, 0x85, 0x54, 0x16, 0xa4, 0x0a, 0x09, 0x70,
	0x71, 0x24, 0x67, 0x24, 0xe6, 0xe5, 0xa5, 0xe6, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0xf0, 0x0a,
	0xf1, 0x71, 0xb1, 0x95, 0x25, 0xe6, 0x94, 0xa6, 0x16, 0x4b, 0x30, 0x83, 0xf9, 0xbc, 0x5c, 0xac,
	0xc5, 0x39, 0xf9, 0x25, 0xc5, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x3c, 0x60, 0x6e, 0x49, 0x62, 0x51,
	0x89, 0x04, 0xab, 0x02, 0xa3, 0x06, 0xaf, 0x10, 0x17, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x9b, 0x02,
	0xa3, 0x06, 0xa7, 0x92, 0x26, 0x17, 0x0b, 0xd8, 0x4c, 0x0e, 0x2e, 0x96, 0x00, 0x4f, 0x3f, 0x77,
	0x01, 0x06, 0x21, 0x4e, 0x2e, 0x56, 0x17, 0x57, 0x9f, 0x10, 0x47, 0x01, 0x46, 0x21, 0x1e, 0x2e,
	0x0e, 0x67, 0x7f, 0xdf, 0x00, 0x1f, 0xd7, 0x10, 0x57, 0x01, 0x26, 0x27, 0x9e, 0x28, 0x2e, 0xb0,
	0x3b, 0x52, 0x72, 0x2b, 0x0a, 0x92, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf1, 0x7c, 0x42,
	0xc1, 0x00, 0x00, 0x00,
}
