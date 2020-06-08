// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package grpctesting

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type EchoRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{2}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{3}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "grpctesting.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "grpctesting.EmptyResponse")
	proto.RegisterType((*EchoRequest)(nil), "grpctesting.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "grpctesting.EchoResponse")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x2f, 0x2a, 0x48, 0x2e, 0x49, 0x2d, 0x2e,
	0xc9, 0xcc, 0x4b, 0x57, 0xe2, 0xe3, 0xe2, 0x71, 0xcd, 0x2d, 0x28, 0xa9, 0x0c, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x51, 0xe2, 0xe7, 0xe2, 0x85, 0xf2, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95,
	0xd4, 0xb9, 0xb8, 0x5d, 0x93, 0x33, 0xf2, 0xa1, 0xf2, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x06, 0x17,
	0x0f, 0x44, 0x21, 0x44, 0x23, 0x6e, 0x95, 0x46, 0xaf, 0x99, 0x20, 0x66, 0x06, 0xa7, 0x16, 0x95,
	0x65, 0x26, 0xa7, 0x0a, 0x59, 0x73, 0xb1, 0x80, 0xb8, 0x42, 0x12, 0x7a, 0x48, 0x2e, 0xd3, 0x43,
	0xb2, 0x55, 0x4a, 0x12, 0x8b, 0x0c, 0xd4, 0x1a, 0x3b, 0x2e, 0x56, 0xb0, 0x83, 0x85, 0xd0, 0xd4,
	0x20, 0x79, 0x4a, 0x4a, 0x0a, 0x9b, 0x14, 0x54, 0xbf, 0x27, 0x97, 0x00, 0xc8, 0x3c, 0xe7, 0x9c,
	0xcc, 0xd4, 0xbc, 0x92, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0xb2, 0x1c, 0xa2, 0xc1, 0x08, 0x33,
	0x0a, 0xe4, 0xad, 0xd4, 0x22, 0x0a, 0x8c, 0x32, 0x00, 0x19, 0xc5, 0x07, 0x12, 0x71, 0xca, 0x4c,
	0xc9, 0xa4, 0xc8, 0x4d, 0x06, 0x8c, 0x49, 0x6c, 0xe0, 0x58, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x96, 0xba, 0xb3, 0xf0, 0x03, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	Empty(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	EchoClientStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClientStreamClient, error)
	EchoServerStream(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (EchoService_EchoServerStreamClient, error)
	EchoBidiStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoBidiStreamClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/grpctesting.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) Empty(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/grpctesting.EchoService/Empty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoClientStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[0], "/grpctesting.EchoService/EchoClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoClientStreamClient{stream}
	return x, nil
}

type EchoService_EchoClientStreamClient interface {
	Send(*EchoRequest) error
	CloseAndRecv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceEchoClientStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoClientStreamClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoClientStreamClient) CloseAndRecv() (*EchoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) EchoServerStream(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (EchoService_EchoServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[1], "/grpctesting.EchoService/EchoServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_EchoServerStreamClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceEchoServerStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoServerStreamClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) EchoBidiStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[2], "/grpctesting.EchoService/EchoBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoBidiStreamClient{stream}
	return x, nil
}

type EchoService_EchoBidiStreamClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceEchoBidiStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoBidiStreamClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoBidiStreamClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	Empty(context.Context, *EmptyRequest) (*EmptyResponse, error)
	EchoClientStream(EchoService_EchoClientStreamServer) error
	EchoServerStream(*EchoRequest, EchoService_EchoServerStreamServer) error
	EchoBidiStream(EchoService_EchoBidiStreamServer) error
}

// UnimplementedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (*UnimplementedEchoServiceServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedEchoServiceServer) Empty(ctx context.Context, req *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Empty not implemented")
}
func (*UnimplementedEchoServiceServer) EchoClientStream(srv EchoService_EchoClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoClientStream not implemented")
}
func (*UnimplementedEchoServiceServer) EchoServerStream(req *EchoRequest, srv EchoService_EchoServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoServerStream not implemented")
}
func (*UnimplementedEchoServiceServer) EchoBidiStream(srv EchoService_EchoBidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoBidiStream not implemented")
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctesting.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_Empty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Empty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctesting.EchoService/Empty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Empty(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).EchoClientStream(&echoServiceEchoClientStreamServer{stream})
}

type EchoService_EchoClientStreamServer interface {
	SendAndClose(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoServiceEchoClientStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoClientStreamServer) SendAndClose(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoClientStreamServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoService_EchoServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).EchoServerStream(m, &echoServiceEchoServerStreamServer{stream})
}

type EchoService_EchoServerStreamServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type echoServiceEchoServerStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoServerStreamServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoService_EchoBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).EchoBidiStream(&echoServiceEchoBidiStreamServer{stream})
}

type EchoService_EchoBidiStreamServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoServiceEchoBidiStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoBidiStreamServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoBidiStreamServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctesting.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "Empty",
			Handler:    _EchoService_Empty_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoClientStream",
			Handler:       _EchoService_EchoClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EchoServerStream",
			Handler:       _EchoService_EchoServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EchoBidiStream",
			Handler:       _EchoService_EchoBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}