// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Salt                 string   `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	Permissions          []string `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *User) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *User) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.user")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x6a, 0xb8, 0xc0, 0xb4, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90,
	0x14, 0x17, 0x07, 0x48, 0x3c, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x0a, 0xe7, 0x0b, 0x09,
	0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x83, 0xc5, 0xc1, 0x6c, 0x90, 0x58, 0x71, 0x62,
	0x4e, 0x89, 0x04, 0x0b, 0x44, 0x0c, 0xc4, 0x16, 0x52, 0xe0, 0xe2, 0x2e, 0x48, 0x2d, 0xca, 0xcd,
	0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0x96, 0x60, 0x55, 0x60, 0xd6, 0xe0, 0x0c, 0x42, 0x16, 0x4a,
	0x62, 0x03, 0x3b, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x16, 0x9d, 0x9f, 0x81, 0x98, 0x00,
	0x00, 0x00,
}
