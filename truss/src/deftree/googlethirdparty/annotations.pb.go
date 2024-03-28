// Code generated by protoc-enum_gen-gogo. DO NOT EDIT.
// source: annotations.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295728,
	Name:          "google.api.http",
	Tag:           "bytes,72295728,opt,name=http",
	Filename:      "annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Http)
}

func init() { proto.RegisterFile("annotations.proto", fileDescriptor_annotations_2b1588018817de26) }

var fileDescriptor_annotations_2b1588018817de26 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0x94, 0xb2, 0x4b, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0x4b, 0x2d, 0x2a, 0x4e, 0xd5,
	0x2f, 0x29, 0x2a, 0x2d, 0x2e, 0xd6, 0x4f, 0x49, 0x4d, 0x2b, 0x29, 0x4a, 0x4d, 0xd5, 0x87, 0x28,
	0x2f, 0xc9, 0xc8, 0x2c, 0x4a, 0x29, 0x48, 0x2c, 0x2a, 0xa9, 0xd4, 0xcf, 0x28, 0x29, 0x29, 0x80,
	0x98, 0x25, 0xe5, 0x42, 0x8e, 0xfe, 0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xfc, 0x22,
	0x88, 0x29, 0x56, 0xde, 0x5c, 0x2c, 0x20, 0x33, 0x85, 0xe4, 0xf4, 0xa0, 0x4e, 0x03, 0x0b, 0x27,
	0x95, 0xa6, 0xe9, 0xf9, 0xa6, 0x96, 0x64, 0xe4, 0xa7, 0xf8, 0x17, 0x80, 0xdd, 0x2f, 0xb1, 0xe1,
	0xd4, 0x1e, 0x25, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x84, 0x1f, 0xf4, 0x3c, 0x4a, 0x4a,
	0x0a, 0x82, 0x4a, 0x73, 0x52, 0x83, 0xc0, 0x86, 0x38, 0x39, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x03, 0x17, 0x5f, 0x72, 0x7e, 0x2e, 0x92, 0x46, 0x27, 0x01, 0x47, 0x44, 0xd8,
	0x04, 0x80, 0x6c, 0x0c, 0x60, 0x5c, 0xc4, 0xc4, 0xe2, 0xee, 0x18, 0xe0, 0x99, 0xc4, 0x06, 0x76,
	0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x2d, 0x6a, 0x40, 0x3f, 0x01, 0x00, 0x00,
}
