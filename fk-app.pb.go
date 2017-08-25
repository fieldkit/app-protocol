// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-app.proto

/*
Package fk_app is a generated protocol buffer package.

It is generated from these files:
	fk-app.proto

It has these top-level messages:
	PingRequest
	PingResponse
	RequestHeader
	ResponseHeader
*/
package fk_app

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

type RequestHeader_MessageType int32

const (
	RequestHeader_PING      RequestHeader_MessageType = 0
	RequestHeader_SAY_HELLO RequestHeader_MessageType = 1
)

var RequestHeader_MessageType_name = map[int32]string{
	0: "PING",
	1: "SAY_HELLO",
}
var RequestHeader_MessageType_value = map[string]int32{
	"PING":      0,
	"SAY_HELLO": 1,
}

func (x RequestHeader_MessageType) String() string {
	return proto.EnumName(RequestHeader_MessageType_name, int32(x))
}
func (RequestHeader_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type PingRequest struct {
	Time int64 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type PingResponse struct {
	Time int64 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type RequestHeader struct {
	Type RequestHeader_MessageType `protobuf:"varint,1,opt,name=type,enum=fk_app.RequestHeader_MessageType" json:"type,omitempty"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestHeader) GetType() RequestHeader_MessageType {
	if m != nil {
		return m.Type
	}
	return RequestHeader_PING
}

type ResponseHeader struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResponseHeader) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResponseHeader) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "fk_app.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "fk_app.PingResponse")
	proto.RegisterType((*RequestHeader)(nil), "fk_app.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "fk_app.ResponseHeader")
	proto.RegisterEnum("fk_app.RequestHeader_MessageType", RequestHeader_MessageType_name, RequestHeader_MessageType_value)
}

func init() { proto.RegisterFile("fk-app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xd6, 0x4d,
	0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcb, 0x8e, 0x4f, 0x2c, 0x28,
	0x50, 0x52, 0xe4, 0xe2, 0x0e, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe2, 0x62, 0x29, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0xb3,
	0x95, 0x94, 0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xb1, 0xaa, 0xc9, 0xe3,
	0xe2, 0x85, 0x1a, 0xe1, 0x91, 0x9a, 0x98, 0x92, 0x5a, 0x24, 0x64, 0xca, 0xc5, 0x52, 0x52, 0x59,
	0x00, 0x51, 0xc4, 0x67, 0xa4, 0xa8, 0x07, 0xb1, 0x4e, 0x0f, 0x45, 0x91, 0x9e, 0x6f, 0x6a, 0x71,
	0x71, 0x62, 0x7a, 0x6a, 0x48, 0x65, 0x41, 0x6a, 0x10, 0x58, 0xb9, 0x92, 0x1a, 0x17, 0x37, 0x92,
	0xa0, 0x10, 0x07, 0x17, 0x4b, 0x80, 0xa7, 0x9f, 0xbb, 0x00, 0x83, 0x10, 0x2f, 0x17, 0x67, 0xb0,
	0x63, 0x64, 0xbc, 0x87, 0xab, 0x8f, 0x8f, 0xbf, 0x00, 0xa3, 0x92, 0x03, 0x17, 0x1f, 0xcc, 0x3d,
	0x50, 0x0b, 0x25, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0xc1, 0x76, 0x72, 0x04,
	0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0x45, 0x45, 0xf9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x10, 0x8e, 0x91, 0x1d, 0x17, 0x7b, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90,
	0x31, 0x17, 0x0b, 0xc8, 0x83, 0x42, 0xc2, 0x30, 0x57, 0x22, 0x85, 0x88, 0x94, 0x08, 0xaa, 0x20,
	0xc4, 0x4e, 0x25, 0x86, 0x24, 0x36, 0x70, 0x38, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x4c, 0x1a, 0x5c, 0x57, 0x01, 0x00, 0x00,
}
