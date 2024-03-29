// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package signaref

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type UserDetails struct {
	Email                string               `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Passcode             string               `protobuf:"bytes,3,opt,name=passcode,proto3" json:"passcode,omitempty"`
	UserType             string               `protobuf:"bytes,4,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	Status               string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserDetails) Reset()         { *m = UserDetails{} }
func (m *UserDetails) String() string { return proto.CompactTextString(m) }
func (*UserDetails) ProtoMessage()    {}
func (*UserDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetails.Unmarshal(m, b)
}
func (m *UserDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetails.Marshal(b, m, deterministic)
}
func (m *UserDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetails.Merge(m, src)
}
func (m *UserDetails) XXX_Size() int {
	return xxx_messageInfo_UserDetails.Size(m)
}
func (m *UserDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetails.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetails proto.InternalMessageInfo

func (m *UserDetails) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDetails) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserDetails) GetPasscode() string {
	if m != nil {
		return m.Passcode
	}
	return ""
}

func (m *UserDetails) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *UserDetails) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UserDetails) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type LoginDetails struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Passcode             string   `protobuf:"bytes,3,opt,name=passcode,proto3" json:"passcode,omitempty"`
	UserType             string   `protobuf:"bytes,4,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginDetails) Reset()         { *m = LoginDetails{} }
func (m *LoginDetails) String() string { return proto.CompactTextString(m) }
func (*LoginDetails) ProtoMessage()    {}
func (*LoginDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *LoginDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginDetails.Unmarshal(m, b)
}
func (m *LoginDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginDetails.Marshal(b, m, deterministic)
}
func (m *LoginDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginDetails.Merge(m, src)
}
func (m *LoginDetails) XXX_Size() int {
	return xxx_messageInfo_LoginDetails.Size(m)
}
func (m *LoginDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginDetails.DiscardUnknown(m)
}

var xxx_messageInfo_LoginDetails proto.InternalMessageInfo

func (m *LoginDetails) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginDetails) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *LoginDetails) GetPasscode() string {
	if m != nil {
		return m.Passcode
	}
	return ""
}

func (m *LoginDetails) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

type UserChangePassword struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	CurrentPasscode      string   `protobuf:"bytes,3,opt,name=currentPasscode,proto3" json:"currentPasscode,omitempty"`
	NewPasscode          string   `protobuf:"bytes,4,opt,name=newPasscode,proto3" json:"newPasscode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserChangePassword) Reset()         { *m = UserChangePassword{} }
func (m *UserChangePassword) String() string { return proto.CompactTextString(m) }
func (*UserChangePassword) ProtoMessage()    {}
func (*UserChangePassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserChangePassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserChangePassword.Unmarshal(m, b)
}
func (m *UserChangePassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserChangePassword.Marshal(b, m, deterministic)
}
func (m *UserChangePassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserChangePassword.Merge(m, src)
}
func (m *UserChangePassword) XXX_Size() int {
	return xxx_messageInfo_UserChangePassword.Size(m)
}
func (m *UserChangePassword) XXX_DiscardUnknown() {
	xxx_messageInfo_UserChangePassword.DiscardUnknown(m)
}

var xxx_messageInfo_UserChangePassword proto.InternalMessageInfo

func (m *UserChangePassword) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserChangePassword) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserChangePassword) GetCurrentPasscode() string {
	if m != nil {
		return m.CurrentPasscode
	}
	return ""
}

func (m *UserChangePassword) GetNewPasscode() string {
	if m != nil {
		return m.NewPasscode
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDetails)(nil), "signaref.UserDetails")
	proto.RegisterType((*LoginDetails)(nil), "signaref.LoginDetails")
	proto.RegisterType((*UserChangePassword)(nil), "signaref.UserChangePassword")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x3d, 0xaf, 0xd3, 0x30,
	0x14, 0x95, 0x1f, 0xef, 0x55, 0xaf, 0x4e, 0x04, 0xc2, 0x40, 0x95, 0x17, 0x9e, 0x44, 0x95, 0xa9,
	0x7a, 0x43, 0x22, 0xca, 0x56, 0x89, 0x09, 0x24, 0x16, 0x86, 0x2a, 0x6a, 0x07, 0x16, 0x2a, 0xb7,
	0xb9, 0x0d, 0x96, 0x12, 0xdb, 0xf2, 0x75, 0xa8, 0xba, 0xf2, 0x03, 0x58, 0xf8, 0x45, 0xcc, 0x8c,
	0xfc, 0x05, 0x7e, 0x08, 0xb2, 0x9d, 0x7e, 0xc0, 0x80, 0xc4, 0xf2, 0xc6, 0x73, 0xee, 0xc9, 0xc9,
	0x39, 0xf7, 0x9a, 0xd2, 0x0e, 0xc1, 0xe4, 0xda, 0x28, 0xab, 0xd8, 0x35, 0x8a, 0x5a, 0x72, 0x03,
	0xdb, 0xf4, 0x45, 0xad, 0x54, 0xdd, 0x40, 0xe1, 0xf9, 0x75, 0xb7, 0x2d, 0xac, 0x68, 0x01, 0x2d,
	0x6f, 0x75, 0x90, 0xa6, 0xb7, 0xbd, 0x80, 0x6b, 0x51, 0x70, 0x29, 0x95, 0xe5, 0x56, 0x28, 0x89,
	0xfd, 0x34, 0xde, 0xa8, 0xb6, 0x55, 0x32, 0xa0, 0xec, 0x07, 0xa1, 0xd1, 0x12, 0xc1, 0xbc, 0x05,
	0xcb, 0x45, 0x83, 0xec, 0x29, 0xbd, 0x82, 0x96, 0x8b, 0x26, 0x21, 0x63, 0x32, 0x19, 0x96, 0x01,
	0xb0, 0x11, 0x1d, 0xb4, 0x6a, 0x2d, 0x1a, 0x48, 0x2e, 0x3c, 0xdd, 0x23, 0x96, 0xd2, 0x6b, 0xcd,
	0x11, 0x37, 0xaa, 0x82, 0xe4, 0x81, 0x9f, 0x1c, 0x31, 0x7b, 0x4e, 0x87, 0x2e, 0xfe, 0xca, 0xee,
	0x35, 0x24, 0x97, 0x61, 0xe8, 0x88, 0xc5, 0x5e, 0x83, 0x33, 0x44, 0xcb, 0x6d, 0x87, 0xc9, 0x55,
	0x30, 0x0c, 0x88, 0xbd, 0xa6, 0x71, 0xc3, 0xd1, 0xae, 0x3a, 0x5d, 0x71, 0x0b, 0x55, 0x32, 0x18,
	0x93, 0x49, 0x34, 0x4d, 0xf3, 0xd0, 0x28, 0x3f, 0x54, 0xce, 0x17, 0x87, 0xca, 0x65, 0xe4, 0xf4,
	0xcb, 0x20, 0xcf, 0x3a, 0x1a, 0xbf, 0x57, 0xb5, 0x90, 0xf7, 0xdb, 0x26, 0xfb, 0x4a, 0x28, 0x73,
	0x4b, 0x7c, 0xf3, 0x89, 0xcb, 0x1a, 0xe6, 0x1c, 0x71, 0xa7, 0x4c, 0xf5, 0x9f, 0x7f, 0x9f, 0xd0,
	0x47, 0x9b, 0xce, 0x18, 0x90, 0x76, 0xfe, 0x67, 0x88, 0xbf, 0x69, 0x36, 0xa6, 0x91, 0x84, 0xdd,
	0x51, 0x15, 0xd2, 0x9c, 0x53, 0xd3, 0xef, 0x84, 0x5e, 0xba, 0x40, 0xec, 0x23, 0x8d, 0x4b, 0xa8,
	0x05, 0x5a, 0x30, 0x1e, 0x3f, 0xcb, 0x0f, 0xcf, 0x28, 0x3f, 0xbb, 0x7a, 0x7a, 0x73, 0xa2, 0xdf,
	0x81, 0x04, 0xc3, 0x9b, 0x12, 0x50, 0x2b, 0x89, 0x90, 0xdd, 0x7e, 0xf9, 0xf9, 0xeb, 0xdb, 0xc5,
	0x28, 0x7b, 0x5c, 0x7c, 0x7e, 0x59, 0xb8, 0xc6, 0x85, 0xe9, 0x0d, 0x67, 0xe4, 0x8e, 0x7d, 0xa0,
	0x43, 0xbf, 0x70, 0x6f, 0x3e, 0x3a, 0xb9, 0x9c, 0x5f, 0xe1, 0x5f, 0xee, 0x37, 0xde, 0xfd, 0x49,
	0xf6, 0xf0, 0xe8, 0xde, 0xb8, 0x2f, 0x67, 0xe4, 0x6e, 0x3d, 0xf0, 0xc7, 0x7e, 0xf5, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0xfb, 0xf6, 0x58, 0x05, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	RegisterUser(ctx context.Context, in *UserDetails, opts ...grpc.CallOption) (*GeneralResponse, error)
	LoginUser(ctx context.Context, in *LoginDetails, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) RegisterUser(ctx context.Context, in *UserDetails, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/signaref.User/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginUser(ctx context.Context, in *LoginDetails, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/signaref.User/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	RegisterUser(context.Context, *UserDetails) (*GeneralResponse, error)
	LoginUser(context.Context, *LoginDetails) (*GeneralResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signaref.User/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterUser(ctx, req.(*UserDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signaref.User/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginUser(ctx, req.(*LoginDetails))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "signaref.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _User_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _User_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
