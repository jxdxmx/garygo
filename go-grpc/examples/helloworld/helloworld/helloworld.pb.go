// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	CustomerFilter
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Avatar    string `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	CompanyId int64  `protobuf:"varint,4,opt,name=company_id,json=companyId" json:"company_id,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HelloReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *HelloReply) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *HelloReply) GetCompanyId() int64 {
	if m != nil {
		return m.CompanyId
	}
	return 0
}

type CustomerFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *CustomerFilter) Reset()                    { *m = CustomerFilter{} }
func (m *CustomerFilter) String() string            { return proto.CompactTextString(m) }
func (*CustomerFilter) ProtoMessage()               {}
func (*CustomerFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomerFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*CustomerFilter)(nil), "helloworld.CustomerFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Greeter_GetCustomersClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Greeter_GetCustomersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/helloworld.Greeter/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_GetCustomersClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterGetCustomersClient struct {
	grpc.ClientStream
}

func (x *greeterGetCustomersClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	GetCustomers(*CustomerFilter, Greeter_GetCustomersServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GetCustomers(m, &greeterGetCustomersServer{stream})
}

type Greeter_GetCustomersServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterGetCustomersServer struct {
	grpc.ServerStream
}

func (x *greeterGetCustomersServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _Greeter_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xdf, 0x6c, 0x5f, 0xda, 0xee, 0x50, 0xaa, 0xe4, 0x50, 0xc2, 0x8a, 0x52, 0xf6, 0x54,
	0x3c, 0x2c, 0x45, 0xef, 0x1e, 0x2a, 0xb4, 0x7a, 0x2b, 0xeb, 0xc1, 0xa3, 0xc4, 0x66, 0xd0, 0x60,
	0xb6, 0x89, 0x49, 0xd6, 0x9a, 0x8f, 0xe1, 0x37, 0x96, 0x8d, 0x5b, 0x5d, 0x45, 0x6f, 0xf3, 0xcc,
	0xf3, 0x63, 0xfe, 0xf0, 0xc0, 0xe1, 0x23, 0x2a, 0xa5, 0x77, 0xda, 0x2a, 0x51, 0x18, 0xab, 0xbd,
	0xa6, 0xf0, 0xd5, 0xc9, 0x4f, 0x60, 0x74, 0xd5, 0xa8, 0x12, 0x9f, 0x6b, 0x74, 0x9e, 0x8e, 0x21,
	0x91, 0x82, 0x91, 0x29, 0x99, 0xf5, 0xca, 0x44, 0x8a, 0x5c, 0x03, 0xb4, 0xbe, 0x51, 0xe1, 0xa7,
	0x4b, 0x33, 0x18, 0xd6, 0x0e, 0xed, 0x96, 0x57, 0xc8, 0x92, 0x29, 0x99, 0xa5, 0xe5, 0xa7, 0xa6,
	0x13, 0xe8, 0xf3, 0x17, 0xee, 0xb9, 0x65, 0xbd, 0xe8, 0xb4, 0x8a, 0x1e, 0x03, 0x6c, 0x74, 0x65,
	0xf8, 0x36, 0xdc, 0x49, 0xc1, 0xfe, 0xc7, 0x59, 0x69, 0xdb, 0xb9, 0x16, 0xf9, 0x29, 0x8c, 0x2f,
	0x6b, 0xe7, 0x75, 0x85, 0x76, 0x29, 0x95, 0x47, 0x4b, 0x19, 0x0c, 0x9e, 0x30, 0xec, 0xb4, 0xfd,
	0xd8, 0x9c, 0x96, 0x7b, 0x79, 0xf6, 0x46, 0x60, 0xb0, 0xb2, 0x88, 0x0d, 0x75, 0x01, 0xc3, 0x1b,
	0x1e, 0xe2, 0xad, 0x94, 0x15, 0x9d, 0x9f, 0xbb, 0xef, 0x65, 0x93, 0x5f, 0x1c, 0xa3, 0x42, 0xfe,
	0x8f, 0x2e, 0x61, 0xb4, 0x42, 0xbf, 0x5f, 0xed, 0x68, 0xd6, 0x25, 0xbf, 0x5f, 0xf4, 0xf7, 0x94,
	0x39, 0x59, 0xcc, 0xe1, 0x48, 0xea, 0xe2, 0xc1, 0x9a, 0x4d, 0x81, 0xaf, 0xbc, 0x32, 0x0a, 0x5d,
	0x87, 0x5e, 0x1c, 0x44, 0xfc, 0xb6, 0xa9, 0xd7, 0x4d, 0x18, 0x6b, 0x72, 0xdf, 0x8f, 0xa9, 0x9c,
	0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xf5, 0xd9, 0x1e, 0xa9, 0x01, 0x00, 0x00,
}
