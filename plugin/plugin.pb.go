// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin/plugin.proto

package plugin

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

type Info struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Author               string   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	IconClass            string   `protobuf:"bytes,5,opt,name=icon_class,json=iconClass,proto3" json:"icon_class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f7429f2a742b54b, []int{0}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Info) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Info) GetIconClass() string {
	if m != nil {
		return m.IconClass
	}
	return ""
}

type Metadata struct {
	Identifier           string                   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Info                 *Info                    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	AdminMenuItems       []*MetadataAdminMenuItem `protobuf:"bytes,3,rep,name=admin_menu_items,json=adminMenuItems,proto3" json:"admin_menu_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f7429f2a742b54b, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Metadata) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Metadata) GetAdminMenuItems() []*MetadataAdminMenuItem {
	if m != nil {
		return m.AdminMenuItems
	}
	return nil
}

type MetadataAdminMenuItem struct {
	MenuName             string   `protobuf:"bytes,1,opt,name=menu_name,json=menuName,proto3" json:"menu_name,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataAdminMenuItem) Reset()         { *m = MetadataAdminMenuItem{} }
func (m *MetadataAdminMenuItem) String() string { return proto.CompactTextString(m) }
func (*MetadataAdminMenuItem) ProtoMessage()    {}
func (*MetadataAdminMenuItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f7429f2a742b54b, []int{1, 0}
}

func (m *MetadataAdminMenuItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataAdminMenuItem.Unmarshal(m, b)
}
func (m *MetadataAdminMenuItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataAdminMenuItem.Marshal(b, m, deterministic)
}
func (m *MetadataAdminMenuItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataAdminMenuItem.Merge(m, src)
}
func (m *MetadataAdminMenuItem) XXX_Size() int {
	return xxx_messageInfo_MetadataAdminMenuItem.Size(m)
}
func (m *MetadataAdminMenuItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataAdminMenuItem.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataAdminMenuItem proto.InternalMessageInfo

func (m *MetadataAdminMenuItem) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *MetadataAdminMenuItem) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "plugin.info")
	proto.RegisterType((*Metadata)(nil), "plugin.metadata")
	proto.RegisterType((*MetadataAdminMenuItem)(nil), "plugin.metadata.admin_menu_item")
}

func init() {
	proto.RegisterFile("plugin/plugin.proto", fileDescriptor_5f7429f2a742b54b)
}

var fileDescriptor_5f7429f2a742b54b = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xd1, 0x4a, 0x84, 0x40,
	0x14, 0x86, 0xb1, 0x35, 0xd3, 0x63, 0x54, 0x9c, 0x20, 0x86, 0x8d, 0x42, 0xf6, 0x6a, 0xaf, 0x0c,
	0xb6, 0x47, 0xe8, 0xaa, 0x85, 0xba, 0xf0, 0x05, 0x64, 0xd2, 0xb3, 0x75, 0x60, 0x3d, 0x23, 0x3a,
	0xf6, 0x16, 0xbd, 0x65, 0x0f, 0x12, 0x33, 0x6a, 0xc8, 0x5e, 0xe9, 0xff, 0xfd, 0x67, 0x98, 0xff,
	0xfc, 0x03, 0xb7, 0xed, 0x71, 0xf8, 0x64, 0x79, 0x1a, 0x3f, 0x79, 0xdb, 0x19, 0x6b, 0x30, 0x1a,
	0xd5, 0xe6, 0x27, 0x80, 0x90, 0xe5, 0x60, 0x10, 0x21, 0x14, 0xdd, 0x90, 0x0a, 0xb2, 0x60, 0x9b,
	0x14, 0xfe, 0x1f, 0xef, 0x20, 0xd2, 0x83, 0xfd, 0x32, 0x9d, 0x3a, 0xf3, 0x74, 0x52, 0x98, 0x41,
	0x5a, 0x53, 0x5f, 0x75, 0xdc, 0x5a, 0x36, 0xa2, 0x56, 0xde, 0x5c, 0x22, 0x54, 0x70, 0xf1, 0x4d,
	0x5d, 0xef, 0xdc, 0xd0, 0xbb, 0xb3, 0xc4, 0x07, 0x00, 0xae, 0x8c, 0x94, 0xd5, 0x51, 0xf7, 0xbd,
	0x3a, 0xf7, 0x66, 0xe2, 0xc8, 0x8b, 0x03, 0x9b, 0xdf, 0x00, 0xe2, 0x86, 0xac, 0xae, 0xb5, 0xd5,
	0xf8, 0x08, 0xc0, 0x35, 0x89, 0xe5, 0x03, 0x53, 0x37, 0x25, 0x5b, 0x10, 0xcc, 0xc6, 0xec, 0x3e,
	0x5d, 0xba, 0xbb, 0xcc, 0xa7, 0x0d, 0x1d, 0x2b, 0xc6, 0xad, 0xf6, 0x70, 0xa3, 0xeb, 0x86, 0xa5,
	0x6c, 0x48, 0x86, 0x92, 0x2d, 0x35, 0xbd, 0x5a, 0x65, 0xab, 0x6d, 0xba, 0xcb, 0xe6, 0xe9, 0xf9,
	0xb6, 0xfc, 0x64, 0xb0, 0xb8, 0xf2, 0xe0, 0x8d, 0x64, 0x78, 0x75, 0xe7, 0xd6, 0x7b, 0xb8, 0x3e,
	0x19, 0xc1, 0x7b, 0x48, 0xbc, 0x58, 0x34, 0x17, 0x3b, 0xf0, 0xee, 0xda, 0x5b, 0x43, 0x4c, 0x52,
	0xb7, 0x86, 0xc5, 0x4e, 0xfd, 0xfd, 0xeb, 0x8f, 0xc8, 0xbf, 0xc2, 0xf3, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x26, 0x29, 0xf1, 0x76, 0x9c, 0x01, 0x00, 0x00,
}
