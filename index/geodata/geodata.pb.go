// Code generated by protoc-gen-go.
// source: geodata.proto
// DO NOT EDIT!

/*
Package geodata is a generated protocol buffer package.

It is generated from these files:
	geodata.proto

It has these top-level messages:
	Geometry
	GeoData
*/
package geodata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Geometry_Type int32

const (
	Geometry_POINT        Geometry_Type = 0
	Geometry_POLYGON      Geometry_Type = 1
	Geometry_MULTIPOLYGON Geometry_Type = 2
)

var Geometry_Type_name = map[int32]string{
	0: "POINT",
	1: "POLYGON",
	2: "MULTIPOLYGON",
}
var Geometry_Type_value = map[string]int32{
	"POINT":        0,
	"POLYGON":      1,
	"MULTIPOLYGON": 2,
}

func (x Geometry_Type) String() string {
	return proto.EnumName(Geometry_Type_name, int32(x))
}
func (Geometry_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Geometry struct {
	Type        Geometry_Type `protobuf:"varint,1,opt,name=type,enum=geodata.Geometry_Type" json:"type,omitempty"`
	Geometries  []*Geometry   `protobuf:"bytes,2,rep,name=geometries" json:"geometries,omitempty"`
	Coordinates []float64     `protobuf:"fixed64,3,rep,packed,name=coordinates" json:"coordinates,omitempty"`
}

func (m *Geometry) Reset()                    { *m = Geometry{} }
func (m *Geometry) String() string            { return proto.CompactTextString(m) }
func (*Geometry) ProtoMessage()               {}
func (*Geometry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Geometry) GetType() Geometry_Type {
	if m != nil {
		return m.Type
	}
	return Geometry_POINT
}

func (m *Geometry) GetGeometries() []*Geometry {
	if m != nil {
		return m.Geometries
	}
	return nil
}

func (m *Geometry) GetCoordinates() []float64 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type GeoData struct {
	Geometry   *Geometry                         `protobuf:"bytes,1,opt,name=geometry" json:"geometry,omitempty"`
	Properties map[string]*google_protobuf.Value `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GeoData) Reset()                    { *m = GeoData{} }
func (m *GeoData) String() string            { return proto.CompactTextString(m) }
func (*GeoData) ProtoMessage()               {}
func (*GeoData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GeoData) GetGeometry() *Geometry {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *GeoData) GetProperties() map[string]*google_protobuf.Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*Geometry)(nil), "geodata.Geometry")
	proto.RegisterType((*GeoData)(nil), "geodata.GeoData")
	proto.RegisterEnum("geodata.Geometry_Type", Geometry_Type_name, Geometry_Type_value)
}

func init() { proto.RegisterFile("geodata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xdd, 0x16, 0x04, 0x06, 0xff, 0xd4, 0x3d, 0x90, 0x86, 0x78, 0x68, 0x38, 0x35, 0x46,
	0x17, 0xc5, 0x8b, 0xf1, 0xe4, 0x41, 0xd3, 0x90, 0x20, 0x6d, 0x9a, 0x62, 0xe2, 0x71, 0x81, 0xb1,
	0x21, 0x22, 0xdb, 0x6c, 0xb7, 0x26, 0x7d, 0x38, 0x1f, 0xc2, 0x37, 0x32, 0xdd, 0x52, 0x68, 0xd4,
	0xdb, 0xee, 0x37, 0xbf, 0x99, 0x6f, 0xbe, 0x0c, 0x1c, 0xc7, 0x28, 0x96, 0x5c, 0x71, 0x96, 0x48,
	0xa1, 0x04, 0x6d, 0x6d, 0xbf, 0xfd, 0xf3, 0x58, 0x88, 0x78, 0x8d, 0x43, 0x2d, 0xcf, 0xb3, 0xb7,
	0x61, 0xaa, 0x64, 0xb6, 0x50, 0x25, 0x36, 0xf8, 0x22, 0xd0, 0xf6, 0x50, 0x7c, 0xa0, 0x92, 0x39,
	0xbd, 0x80, 0x86, 0xca, 0x13, 0xb4, 0x89, 0x43, 0xdc, 0x93, 0x51, 0x8f, 0x55, 0x13, 0x2b, 0x80,
	0x45, 0x79, 0x82, 0xa1, 0x66, 0xe8, 0x0d, 0x40, 0x5c, 0xca, 0x2b, 0x4c, 0x6d, 0xc3, 0x31, 0xdd,
	0xee, 0xe8, 0xec, 0x4f, 0x47, 0x58, 0x83, 0xa8, 0x03, 0xdd, 0x85, 0x10, 0x72, 0xb9, 0xda, 0x70,
	0x85, 0xa9, 0x6d, 0x3a, 0xa6, 0x4b, 0xc2, 0xba, 0x34, 0xb8, 0x86, 0x46, 0x61, 0x41, 0x3b, 0xd0,
	0x0c, 0xfc, 0xf1, 0x34, 0xb2, 0x0e, 0x68, 0x17, 0x5a, 0x81, 0x3f, 0x79, 0xf5, 0xfc, 0xa9, 0x45,
	0xa8, 0x05, 0x47, 0xcf, 0xb3, 0x49, 0x34, 0xae, 0x14, 0x63, 0xf0, 0x4d, 0xa0, 0xe5, 0xa1, 0x78,
	0xe4, 0x8a, 0xd3, 0x2b, 0x68, 0x6f, 0xdd, 0x72, 0x1d, 0xe1, 0xdf, 0x85, 0x76, 0x08, 0x7d, 0x00,
	0x48, 0xa4, 0x48, 0x50, 0xaa, 0x7d, 0x02, 0xa7, 0xde, 0x50, 0x0c, 0x65, 0xc1, 0x0e, 0x79, 0xda,
	0xe8, 0x40, 0xfb, 0x9e, 0xfe, 0x0c, 0x4e, 0x7f, 0x95, 0xa9, 0x05, 0xe6, 0x3b, 0x96, 0xf6, 0x9d,
	0xb0, 0x78, 0xd2, 0x4b, 0x68, 0x7e, 0xf2, 0x75, 0x86, 0xb6, 0xa1, 0x57, 0xea, 0xb1, 0xf2, 0x1e,
	0xac, 0xba, 0x07, 0x7b, 0x29, 0xaa, 0x61, 0x09, 0xdd, 0x1b, 0x77, 0x64, 0x7e, 0xa8, 0x4b, 0xb7,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xc4, 0x36, 0xa9, 0xd2, 0x01, 0x00, 0x00,
}