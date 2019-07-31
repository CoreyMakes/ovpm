// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type AuthStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthStatusRequest) Reset()         { *m = AuthStatusRequest{} }
func (m *AuthStatusRequest) String() string { return proto.CompactTextString(m) }
func (*AuthStatusRequest) ProtoMessage()    {}
func (*AuthStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_5d9157b05cb925f8, []int{0}
}
func (m *AuthStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatusRequest.Unmarshal(m, b)
}
func (m *AuthStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatusRequest.Marshal(b, m, deterministic)
}
func (dst *AuthStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusRequest.Merge(dst, src)
}
func (m *AuthStatusRequest) XXX_Size() int {
	return xxx_messageInfo_AuthStatusRequest.Size(m)
}
func (m *AuthStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusRequest proto.InternalMessageInfo

type AuthAuthenticateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthAuthenticateRequest) Reset()         { *m = AuthAuthenticateRequest{} }
func (m *AuthAuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthAuthenticateRequest) ProtoMessage()    {}
func (*AuthAuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_5d9157b05cb925f8, []int{1}
}
func (m *AuthAuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthAuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthAuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthAuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthAuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthAuthenticateRequest.Merge(dst, src)
}
func (m *AuthAuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthAuthenticateRequest.Size(m)
}
func (m *AuthAuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthAuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthAuthenticateRequest proto.InternalMessageInfo

func (m *AuthAuthenticateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthAuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthStatusResponse struct {
	User                 *UserResponse_User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	IsRoot               bool               `protobuf:"varint,2,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuthStatusResponse) Reset()         { *m = AuthStatusResponse{} }
func (m *AuthStatusResponse) String() string { return proto.CompactTextString(m) }
func (*AuthStatusResponse) ProtoMessage()    {}
func (*AuthStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_5d9157b05cb925f8, []int{2}
}
func (m *AuthStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatusResponse.Unmarshal(m, b)
}
func (m *AuthStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatusResponse.Marshal(b, m, deterministic)
}
func (dst *AuthStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusResponse.Merge(dst, src)
}
func (m *AuthStatusResponse) XXX_Size() int {
	return xxx_messageInfo_AuthStatusResponse.Size(m)
}
func (m *AuthStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusResponse proto.InternalMessageInfo

func (m *AuthStatusResponse) GetUser() *UserResponse_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AuthStatusResponse) GetIsRoot() bool {
	if m != nil {
		return m.IsRoot
	}
	return false
}

type AuthAuthenticateResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthAuthenticateResponse) Reset()         { *m = AuthAuthenticateResponse{} }
func (m *AuthAuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthAuthenticateResponse) ProtoMessage()    {}
func (*AuthAuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_5d9157b05cb925f8, []int{3}
}
func (m *AuthAuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthAuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthAuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthAuthenticateResponse.Marshal(b, m, deterministic)
}
func (dst *AuthAuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthAuthenticateResponse.Merge(dst, src)
}
func (m *AuthAuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthAuthenticateResponse.Size(m)
}
func (m *AuthAuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthAuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthAuthenticateResponse proto.InternalMessageInfo

func (m *AuthAuthenticateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthStatusRequest)(nil), "pb.AuthStatusRequest")
	proto.RegisterType((*AuthAuthenticateRequest)(nil), "pb.AuthAuthenticateRequest")
	proto.RegisterType((*AuthStatusResponse)(nil), "pb.AuthStatusResponse")
	proto.RegisterType((*AuthAuthenticateResponse)(nil), "pb.AuthAuthenticateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Status(ctx context.Context, in *AuthStatusRequest, opts ...grpc.CallOption) (*AuthStatusResponse, error)
	Authenticate(ctx context.Context, in *AuthAuthenticateRequest, opts ...grpc.CallOption) (*AuthAuthenticateResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Status(ctx context.Context, in *AuthStatusRequest, opts ...grpc.CallOption) (*AuthStatusResponse, error) {
	out := new(AuthStatusResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthAuthenticateRequest, opts ...grpc.CallOption) (*AuthAuthenticateResponse, error) {
	out := new(AuthAuthenticateResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Status(context.Context, *AuthStatusRequest) (*AuthStatusResponse, error)
	Authenticate(context.Context, *AuthAuthenticateRequest) (*AuthAuthenticateResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Status(ctx, req.(*AuthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthAuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthAuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AuthService_Status_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_5d9157b05cb925f8) }

var fileDescriptor_auth_5d9157b05cb925f8 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x22, 0x08, 0x65, 0xcb, 0x05, 0x97, 0xd2, 0x90, 0x56, 0x08, 0x59, 0x1c, 0x80, 0x43,
	0x02, 0xe5, 0xc6, 0x8d, 0x27, 0x10, 0x40, 0xe2, 0x86, 0x9c, 0x62, 0x15, 0x0b, 0xf0, 0x9a, 0xd8,
	0x29, 0x77, 0xbe, 0xc0, 0xd3, 0xe0, 0x09, 0x3c, 0x84, 0xd8, 0x4e, 0x50, 0xa0, 0x3d, 0x44, 0xca,
	0x78, 0x66, 0x67, 0x46, 0xbb, 0x00, 0xac, 0x32, 0x8f, 0xa9, 0x2a, 0xd1, 0x20, 0x09, 0x55, 0x91,
	0x4c, 0xe6, 0x88, 0xf3, 0x67, 0x9e, 0x31, 0x25, 0x32, 0x26, 0x25, 0x1a, 0x66, 0x04, 0x4a, 0xed,
	0x15, 0x09, 0x54, 0x9a, 0x97, 0xfe, 0x9f, 0x0e, 0x60, 0xfb, 0xb2, 0x9e, 0xbd, 0xae, 0x15, 0x95,
	0xce, 0xf9, 0x6b, 0xc5, 0xb5, 0xa1, 0x57, 0x30, 0xb2, 0x8f, 0xf6, 0xe3, 0xd2, 0x88, 0x19, 0x33,
	0xbc, 0xa1, 0x48, 0x02, 0x3d, 0x3b, 0x2d, 0xd9, 0x0b, 0x8f, 0x83, 0x83, 0xe0, 0x68, 0x33, 0xff,
	0xc5, 0x96, 0x53, 0x4c, 0xeb, 0x37, 0x2c, 0x1f, 0xe2, 0xd0, 0x73, 0x2d, 0xa6, 0x77, 0x40, 0xba,
	0x39, 0x5a, 0xd5, 0x75, 0x38, 0x39, 0x86, 0x35, 0x3b, 0xed, 0x9c, 0xfa, 0xd3, 0x61, 0xaa, 0x8a,
	0xf4, 0xb6, 0xc6, 0x2d, 0xef, 0x81, 0x93, 0x90, 0x11, 0x6c, 0x08, 0x7d, 0x5f, 0x22, 0x1a, 0xe7,
	0xdd, 0xcb, 0x23, 0xa1, 0xf3, 0x1a, 0xd1, 0x53, 0x88, 0x97, 0xcb, 0x36, 0xfe, 0x3b, 0xb0, 0x6e,
	0xf0, 0x89, 0xcb, 0xa6, 0xaa, 0x07, 0xd3, 0xaf, 0x00, 0xfa, 0xae, 0x0c, 0x2f, 0x17, 0x62, 0xc6,
	0xc9, 0x0d, 0x44, 0xbe, 0x17, 0x71, 0x0d, 0x96, 0xf6, 0x91, 0xec, 0xfe, 0x7f, 0xf6, 0xf6, 0x74,
	0xfc, 0xfe, 0xf9, 0xfd, 0x11, 0x0e, 0xc9, 0xc0, 0x2d, 0x7a, 0x71, 0x96, 0xd9, 0x33, 0x64, 0xda,
	0x7b, 0x21, 0x6c, 0x75, 0x3b, 0x91, 0x71, 0x6b, 0xb2, 0x62, 0xad, 0xc9, 0x64, 0x35, 0xd9, 0xe4,
	0x1c, 0xba, 0x9c, 0x7d, 0xba, 0xf7, 0x27, 0x87, 0x75, 0xa4, 0x17, 0xc1, 0x49, 0x11, 0xb9, 0x8b,
	0x9e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x87, 0xd8, 0x13, 0x0a, 0x0d, 0x02, 0x00, 0x00,
}