// Code generated by protoc-gen-go.
// source: members.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Member struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Age      int64  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	EyeColor string `protobuf:"bytes,3,opt,name=eyeColor" json:"eyeColor,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Gender   string `protobuf:"bytes,5,opt,name=gender" json:"gender,omitempty"`
	Company  string `protobuf:"bytes,6,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,8,opt,name=phone" json:"phone,omitempty"`
	Address  string `protobuf:"bytes,9,opt,name=address" json:"address,omitempty"`
	About    string `protobuf:"bytes,10,opt,name=about" json:"about,omitempty"`
	Greeting string `protobuf:"bytes,11,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto1.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Member) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Member) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Member) GetEyeColor() string {
	if m != nil {
		return m.EyeColor
	}
	return ""
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Member) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Member) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Member) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Member) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Member) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *Member) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type GetMembersResponse struct {
	Members []*Member `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
}

func (m *GetMembersResponse) Reset()                    { *m = GetMembersResponse{} }
func (m *GetMembersResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetMembersResponse) ProtoMessage()               {}
func (*GetMembersResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetMembersResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto1.RegisterType((*Member)(nil), "Members.Member")
	proto1.RegisterType((*GetMembersResponse)(nil), "Members.GetMembersResponse")
}

func init() { proto1.RegisterFile("members.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0x87, 0xb1, 0x1d, 0xdb, 0xc9, 0x85, 0xfe, 0xe1, 0x28, 0xe5, 0xc8, 0x64, 0x32, 0xb9, 0x50,
	0x3c, 0xb4, 0x0f, 0x50, 0x48, 0x87, 0x4e, 0x5d, 0x3c, 0x76, 0x93, 0xa3, 0xc3, 0x35, 0x8d, 0x24,
	0x23, 0xd9, 0x43, 0xf6, 0x3e, 0x78, 0xb1, 0x24, 0x77, 0xba, 0xfb, 0xbe, 0x3b, 0x8e, 0xe3, 0x07,
	0x37, 0x8a, 0x55, 0xc7, 0xd6, 0x35, 0xa3, 0x35, 0x93, 0xc1, 0xf2, 0x33, 0xe0, 0xf1, 0x37, 0x85,
	0x22, 0xf4, 0x78, 0x0b, 0xe9, 0x20, 0x29, 0xa9, 0x92, 0x3a, 0x6b, 0xd3, 0x41, 0xe2, 0x3d, 0x64,
	0xa2, 0x67, 0x4a, 0xbd, 0x58, 0x5a, 0x3c, 0xc0, 0x96, 0xaf, 0xfc, 0x6e, 0x2e, 0xc6, 0x52, 0x56,
	0x25, 0xf5, 0xae, 0xfd, 0x67, 0x44, 0xd8, 0x68, 0xa1, 0x98, 0x36, 0xde, 0xfb, 0x1e, 0x1f, 0xa1,
	0xe8, 0x59, 0x4b, 0xb6, 0x94, 0x7b, 0x1b, 0x09, 0x09, 0xca, 0xb3, 0x51, 0xa3, 0xd0, 0x57, 0x2a,
	0xfc, 0x60, 0x45, 0x7c, 0x80, 0x9c, 0x95, 0x18, 0x2e, 0x54, 0x7a, 0x1f, 0x60, 0xb1, 0xe3, 0xb7,
	0xd1, 0x4c, 0xdb, 0x60, 0x3d, 0x2c, 0x57, 0x84, 0x94, 0x96, 0x9d, 0xa3, 0x5d, 0xb8, 0x12, 0x71,
	0xd9, 0x17, 0x9d, 0x99, 0x27, 0x82, 0xb0, 0xef, 0x61, 0xf9, 0xbe, 0xb7, 0xcc, 0xd3, 0xa0, 0x7b,
	0xda, 0x87, 0xef, 0x57, 0x3e, 0xbe, 0x01, 0x7e, 0xf0, 0x14, 0x43, 0x69, 0xd9, 0x8d, 0x46, 0x3b,
	0xc6, 0x27, 0x28, 0x63, 0x6c, 0x94, 0x54, 0x59, 0xbd, 0x7f, 0xb9, 0x6b, 0xe2, 0x4a, 0xac, 0xed,
	0x3a, 0x3f, 0x3d, 0xc3, 0xe1, 0x6c, 0x54, 0x33, 0x89, 0x9f, 0xd9, 0xb1, 0xea, 0x44, 0xd3, 0x9b,
	0x59, 0x8a, 0x90, 0xb6, 0x3b, 0xad, 0x71, 0x7f, 0xe5, 0x5e, 0x74, 0x85, 0x2f, 0xaf, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0x73, 0xb9, 0x74, 0x96, 0x01, 0x00, 0x00,
}