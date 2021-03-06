// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greetings.proto

/*
Package greetings is a generated protocol buffer package.

It is generated from these files:
	greetings.proto

It has these top-level messages:
	MsgRequest
	MsgResponse
*/
package greetings

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

type MsgRequest struct {
	Count int32  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *MsgRequest) Reset()                    { *m = MsgRequest{} }
func (m *MsgRequest) String() string            { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()               {}
func (*MsgRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MsgRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MsgRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MsgResponse struct {
	Msg     string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgResponse) Reset()                    { *m = MsgResponse{} }
func (m *MsgResponse) String() string            { return proto.CompactTextString(m) }
func (*MsgResponse) ProtoMessage()               {}
func (*MsgResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MsgResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MsgResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgRequest)(nil), "greetings.MsgRequest")
	proto.RegisterType((*MsgResponse)(nil), "greetings.MsgResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GreetingsService service

type GreetingsServiceClient interface {
	Talk(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (GreetingsService_TalkClient, error)
}

type greetingsServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetingsServiceClient(cc *grpc.ClientConn) GreetingsServiceClient {
	return &greetingsServiceClient{cc}
}

func (c *greetingsServiceClient) Talk(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (GreetingsService_TalkClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GreetingsService_serviceDesc.Streams[0], c.cc, "/greetings.GreetingsService/Talk", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingsServiceTalkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetingsService_TalkClient interface {
	Recv() (*MsgResponse, error)
	grpc.ClientStream
}

type greetingsServiceTalkClient struct {
	grpc.ClientStream
}

func (x *greetingsServiceTalkClient) Recv() (*MsgResponse, error) {
	m := new(MsgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GreetingsService service

type GreetingsServiceServer interface {
	Talk(*MsgRequest, GreetingsService_TalkServer) error
}

func RegisterGreetingsServiceServer(s *grpc.Server, srv GreetingsServiceServer) {
	s.RegisterService(&_GreetingsService_serviceDesc, srv)
}

func _GreetingsService_Talk_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MsgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetingsServiceServer).Talk(m, &greetingsServiceTalkServer{stream})
}

type GreetingsService_TalkServer interface {
	Send(*MsgResponse) error
	grpc.ServerStream
}

type greetingsServiceTalkServer struct {
	grpc.ServerStream
}

func (x *greetingsServiceTalkServer) Send(m *MsgResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greetings.GreetingsService",
	HandlerType: (*GreetingsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Talk",
			Handler:       _GreetingsService_Talk_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greetings.proto",
}

func init() { proto.RegisterFile("greetings.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x4a, 0x4d,
	0x2d, 0xc9, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x99, 0x70, 0x71, 0xf9, 0x16, 0xa7, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70,
	0xb1, 0x26, 0xe7, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x42,
	0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92,
	0x25, 0x17, 0x37, 0x58, 0x57, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x4c, 0x01, 0x23, 0x5c, 0x81,
	0x90, 0x04, 0x17, 0x7b, 0x41, 0x62, 0x65, 0x4e, 0x7e, 0x62, 0x0a, 0x58, 0x1b, 0x4f, 0x10, 0x8c,
	0x6b, 0xe4, 0xcd, 0x25, 0xe0, 0x0e, 0xb3, 0x3d, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8,
	0x9c, 0x8b, 0x25, 0x24, 0x31, 0x27, 0x5b, 0x48, 0x54, 0x0f, 0xe1, 0x52, 0x84, 0xab, 0xa4, 0xc4,
	0xd0, 0x85, 0x21, 0xd6, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0xfd, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xec, 0x17, 0x72, 0x9c, 0xe2, 0x00, 0x00, 0x00,
}
