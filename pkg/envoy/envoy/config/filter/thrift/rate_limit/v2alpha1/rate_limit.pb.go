// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto

package v2alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

// [#comment:next free field: 5]
type RateLimit struct {
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configuration stage. Each configured rate limit filter performs a
	// rate limit check using descriptors configured in the
	// :ref:`envoy_api_msg_config.filter.network.thrift_proxy.v2alpha1.RouteAction` for the request.
	// Only those entries with a matching stage number are used for a given filter. If not set, the
	// default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny      bool     `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_961acdee13c1bd42, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.thrift.rate_limit.v2alpha1.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto", fileDescriptor_961acdee13c1bd42)
}

var fileDescriptor_961acdee13c1bd42 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0xe5, 0xb6, 0x94, 0xd6, 0x08, 0x21, 0x22, 0x24, 0x42, 0x07, 0x1a, 0x98, 0xa2, 0x0e,
	0xb6, 0x28, 0x13, 0x13, 0x52, 0xd5, 0x11, 0x96, 0x8c, 0x2c, 0x95, 0x8b, 0xff, 0xb8, 0x96, 0x1c,
	0xff, 0x95, 0xf9, 0x53, 0xa9, 0x6f, 0xc2, 0x93, 0x30, 0x30, 0xf1, 0x26, 0xcc, 0xbc, 0x05, 0x6a,
	0x9c, 0x48, 0xdd, 0x4e, 0x77, 0xe7, 0xcf, 0xa7, 0x9f, 0x3f, 0x83, 0xdf, 0xe1, 0x5e, 0xbe, 0xa3,
	0x2f, 0xad, 0x91, 0xa5, 0x75, 0x04, 0x41, 0xd2, 0x26, 0xd8, 0x92, 0x64, 0x50, 0x04, 0x2b, 0x67,
	0x2b, 0x4b, 0x72, 0x37, 0x57, 0x6e, 0xbb, 0x51, 0x0f, 0x47, 0x9e, 0xd8, 0x06, 0x24, 0x4c, 0x44,
	0x03, 0x10, 0x11, 0x20, 0x22, 0x40, 0x44, 0x80, 0x38, 0x2a, 0x77, 0x80, 0xc9, 0xad, 0x41, 0x34,
	0x0e, 0x64, 0xf3, 0x7a, 0x5d, 0x97, 0x52, 0xd7, 0x41, 0x91, 0x45, 0x1f, 0x79, 0x93, 0xeb, 0x9d,
	0x72, 0x56, 0x2b, 0x02, 0xd9, 0x89, 0x36, 0xb8, 0x32, 0x68, 0xb0, 0x91, 0xf2, 0xa0, 0xa2, 0x7b,
	0xff, 0xc5, 0xf8, 0xb8, 0x50, 0x04, 0x2f, 0x87, 0x5f, 0x92, 0x3b, 0x3e, 0xd4, 0x58, 0x29, 0xeb,
	0x53, 0x96, 0xb1, 0x7c, 0xbc, 0x18, 0x7f, 0xff, 0xfd, 0xf4, 0x07, 0xa1, 0x97, 0xb1, 0xa2, 0x0d,
	0x92, 0x29, 0x3f, 0xf9, 0x20, 0x65, 0x20, 0xed, 0x65, 0x2c, 0x3f, 0x6f, 0x1b, 0xb3, 0x5e, 0xca,
	0x8b, 0xe8, 0x27, 0x4f, 0xfc, 0x94, 0x6c, 0x05, 0x58, 0x53, 0xda, 0xcf, 0x58, 0x7e, 0x36, 0xbf,
	0x11, 0x71, 0xb2, 0xe8, 0x26, 0x8b, 0x65, 0x3b, 0x79, 0x31, 0xf8, 0xfc, 0x9d, 0xb2, 0xa2, 0xeb,
	0x27, 0x33, 0x7e, 0x59, 0x2a, 0xeb, 0xea, 0x00, 0xab, 0x0a, 0x35, 0xac, 0x34, 0xf8, 0x7d, 0x3a,
	0xc8, 0x58, 0x3e, 0x2a, 0x2e, 0xda, 0xe0, 0x15, 0x35, 0x2c, 0xc1, 0xef, 0x17, 0xfc, 0x6d, 0xd4,
	0xdd, 0x64, 0x3d, 0x6c, 0xc8, 0x8f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0xb6, 0x10, 0x58,
	0x8d, 0x01, 0x00, 0x00,
}
