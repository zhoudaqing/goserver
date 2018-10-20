// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package webs is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	Token
	UserBase
*/
package webs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	Ua    string `protobuf:"bytes,2,opt,name=Ua" json:"Ua,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetUa() string {
	if m != nil {
		return m.Ua
	}
	return ""
}

type UserBase struct {
	Result     bool   `protobuf:"varint,1,opt,name=Result" json:"Result,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	AppendJson string `protobuf:"bytes,3,opt,name=AppendJson" json:"AppendJson,omitempty"`
}

func (m *UserBase) Reset()                    { *m = UserBase{} }
func (m *UserBase) String() string            { return proto.CompactTextString(m) }
func (*UserBase) ProtoMessage()               {}
func (*UserBase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserBase) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *UserBase) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserBase) GetAppendJson() string {
	if m != nil {
		return m.AppendJson
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "webs.Token")
	proto.RegisterType((*UserBase)(nil), "webs.UserBase")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RpcUser service

type RpcUserClient interface {
	Verify(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserBase, error)
}

type rpcUserClient struct {
	cc *grpc.ClientConn
}

func NewRpcUserClient(cc *grpc.ClientConn) RpcUserClient {
	return &rpcUserClient{cc}
}

func (c *rpcUserClient) Verify(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserBase, error) {
	out := new(UserBase)
	err := grpc.Invoke(ctx, "/webs.RpcUser/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RpcUser service

type RpcUserServer interface {
	Verify(context.Context, *Token) (*UserBase, error)
}

func RegisterRpcUserServer(s *grpc.Server, srv RpcUserServer) {
	s.RegisterService(&_RpcUser_serviceDesc, srv)
}

func _RpcUser_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcUserServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webs.RpcUser/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcUserServer).Verify(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "webs.RpcUser",
	HandlerType: (*RpcUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _RpcUser_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x4f, 0x4d, 0x2a, 0x56, 0xd2, 0xe5, 0x62,
	0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x13, 0x12, 0x81, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xa0, 0xa2, 0x7c, 0x5c, 0x4c, 0xa1, 0x89, 0x12, 0x4c, 0x60, 0x21, 0xa6, 0xd0, 0x44, 0xa5, 0x28,
	0x2e, 0x8e, 0xd0, 0xe2, 0xd4, 0x22, 0xa7, 0xc4, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa0, 0xd4,
	0xe2, 0xd2, 0x9c, 0x12, 0xb0, 0x16, 0x8e, 0x20, 0x28, 0x0f, 0x24, 0x0e, 0x52, 0xe3, 0x99, 0x02,
	0xd6, 0xc7, 0x1c, 0x04, 0xe5, 0x09, 0xc9, 0x71, 0x71, 0x39, 0x16, 0x14, 0xa4, 0xe6, 0xa5, 0x78,
	0x15, 0xe7, 0xe7, 0x49, 0x30, 0x83, 0xcd, 0x44, 0x12, 0x31, 0x32, 0xe0, 0x62, 0x0f, 0x2a, 0x48,
	0x06, 0x29, 0x16, 0x52, 0xe5, 0x62, 0x0b, 0x4b, 0x2d, 0xca, 0x4c, 0xab, 0x14, 0xe2, 0xd6, 0x03,
	0x39, 0x53, 0x0f, 0xec, 0x1a, 0x29, 0x3e, 0x08, 0x07, 0xe6, 0x82, 0x24, 0x36, 0xb0, 0x4f, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x3d, 0x6f, 0x7f, 0xd7, 0x00, 0x00, 0x00,
}
