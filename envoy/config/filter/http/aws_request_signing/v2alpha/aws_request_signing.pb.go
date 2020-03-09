// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/aws_request_signing/v2alpha/aws_request_signing.proto

package envoy_config_filter_http_aws_request_signing_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AwsRequestSigning struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	HostRewrite          string   `protobuf:"bytes,3,opt,name=host_rewrite,json=hostRewrite,proto3" json:"host_rewrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsRequestSigning) Reset()         { *m = AwsRequestSigning{} }
func (m *AwsRequestSigning) String() string { return proto.CompactTextString(m) }
func (*AwsRequestSigning) ProtoMessage()    {}
func (*AwsRequestSigning) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebccdae368886250, []int{0}
}

func (m *AwsRequestSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsRequestSigning.Unmarshal(m, b)
}
func (m *AwsRequestSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsRequestSigning.Marshal(b, m, deterministic)
}
func (m *AwsRequestSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsRequestSigning.Merge(m, src)
}
func (m *AwsRequestSigning) XXX_Size() int {
	return xxx_messageInfo_AwsRequestSigning.Size(m)
}
func (m *AwsRequestSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsRequestSigning.DiscardUnknown(m)
}

var xxx_messageInfo_AwsRequestSigning proto.InternalMessageInfo

func (m *AwsRequestSigning) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AwsRequestSigning) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AwsRequestSigning) GetHostRewrite() string {
	if m != nil {
		return m.HostRewrite
	}
	return ""
}

func init() {
	proto.RegisterType((*AwsRequestSigning)(nil), "envoy.config.filter.http.aws_request_signing.v2alpha.AwsRequestSigning")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/aws_request_signing/v2alpha/aws_request_signing.proto", fileDescriptor_ebccdae368886250)
}

var fileDescriptor_ebccdae368886250 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x02, 0x45, 0xb8, 0x5d, 0xc8, 0x00, 0x51, 0x07, 0x28, 0x4c, 0x88, 0xc1, 0x96,
	0x28, 0x62, 0x27, 0x0f, 0x50, 0x55, 0xe1, 0x01, 0x22, 0xd3, 0x5e, 0x93, 0x93, 0x12, 0x3b, 0xd8,
	0xd7, 0xa4, 0x9d, 0x79, 0x00, 0x56, 0xc4, 0x63, 0x32, 0x32, 0x20, 0x94, 0x38, 0x1d, 0x10, 0x15,
	0x43, 0xb7, 0xd3, 0xff, 0xdf, 0x7d, 0xd2, 0x7d, 0x7c, 0x0a, 0xba, 0x32, 0x1b, 0x39, 0x37, 0x7a,
	0x89, 0xa9, 0x5c, 0x62, 0x4e, 0x60, 0x65, 0x46, 0x54, 0x4a, 0x55, 0xbb, 0xc4, 0xc2, 0xcb, 0x0a,
	0x1c, 0x25, 0x0e, 0x53, 0x8d, 0x3a, 0x95, 0xd5, 0x9d, 0xca, 0xcb, 0x4c, 0xed, 0xea, 0x44, 0x69,
	0x0d, 0x99, 0xe0, 0xbe, 0xe5, 0x09, 0xcf, 0x13, 0x9e, 0x27, 0x1a, 0x9e, 0xd8, 0x75, 0xd3, 0xf1,
	0x46, 0x17, 0xab, 0x45, 0xa9, 0xa4, 0xd2, 0xda, 0x90, 0x22, 0x34, 0xda, 0xc9, 0x02, 0x53, 0xab,
	0x08, 0x3c, 0x75, 0x74, 0x5e, 0xa9, 0x1c, 0x17, 0x8a, 0x40, 0x6e, 0x07, 0x5f, 0x5c, 0xbf, 0x32,
	0x7e, 0xfa, 0x58, 0xbb, 0xd8, 0x73, 0x9f, 0x3c, 0x36, 0xb8, 0xe5, 0x43, 0x07, 0xb6, 0xc2, 0x39,
	0x24, 0x5a, 0x15, 0x10, 0xb2, 0x31, 0xbb, 0x39, 0x89, 0x8e, 0xbf, 0xa2, 0x43, 0xdb, 0x1b, 0xb3,
	0x78, 0xd0, 0x95, 0x53, 0x55, 0x40, 0x70, 0xc9, 0xfb, 0x16, 0x52, 0x34, 0x3a, 0xec, 0xfd, 0xde,
	0xea, 0xe2, 0xe0, 0x8a, 0x0f, 0x33, 0xe3, 0x28, 0xb1, 0x50, 0x5b, 0x24, 0x08, 0x0f, 0x9a, 0xb5,
	0x78, 0xd0, 0x64, 0xb1, 0x8f, 0xa2, 0x0f, 0xf6, 0xf9, 0xfe, 0xfd, 0x76, 0xf4, 0xb0, 0xfd, 0x1e,
	0xd6, 0x04, 0xda, 0x35, 0x7f, 0x74, 0x06, 0xdc, 0x3f, 0x0a, 0x26, 0x3c, 0x42, 0x23, 0xda, 0xc3,
	0xd2, 0x9a, 0xf5, 0x46, 0xec, 0x63, 0x30, 0x3a, 0xfb, 0x63, 0x61, 0xd6, 0x08, 0x9a, 0xb1, 0xe7,
	0x7e, 0x6b, 0x6a, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x54, 0x3a, 0xd0, 0xbb, 0xea, 0x01, 0x00,
	0x00,
}
