// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/stream.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type ID struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_bb5a5d0ef9df88e8, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ID) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Message struct {
	Message *any.Any `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Sender's address and public key.
	Sender *ID `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// Sender's signature of message.
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// Nonce is the request/response ID. Null if ID associated to a message is not a request/response.
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Is this message a response?
	IsResponse           bool     `protobuf:"varint,5,opt,name=is_response,json=isResponse,proto3" json:"is_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_bb5a5d0ef9df88e8, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Message) GetSender() *ID {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Message) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Message) GetIsResponse() bool {
	if m != nil {
		return m.IsResponse
	}
	return false
}

type HandshakeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeRequest) Reset()         { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()    {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_bb5a5d0ef9df88e8, []int{2}
}
func (m *HandshakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeRequest.Unmarshal(m, b)
}
func (m *HandshakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeRequest.Marshal(b, m, deterministic)
}
func (dst *HandshakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeRequest.Merge(dst, src)
}
func (m *HandshakeRequest) XXX_Size() int {
	return xxx_messageInfo_HandshakeRequest.Size(m)
}
func (m *HandshakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeRequest proto.InternalMessageInfo

type HandshakeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeResponse) Reset()         { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()    {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_bb5a5d0ef9df88e8, []int{3}
}
func (m *HandshakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeResponse.Unmarshal(m, b)
}
func (m *HandshakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeResponse.Marshal(b, m, deterministic)
}
func (dst *HandshakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeResponse.Merge(dst, src)
}
func (m *HandshakeResponse) XXX_Size() int {
	return xxx_messageInfo_HandshakeResponse.Size(m)
}
func (m *HandshakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeResponse proto.InternalMessageInfo

type LookupNodeRequest struct {
	Target               *ID      `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupNodeRequest) Reset()         { *m = LookupNodeRequest{} }
func (m *LookupNodeRequest) String() string { return proto.CompactTextString(m) }
func (*LookupNodeRequest) ProtoMessage()    {}
func (*LookupNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_bb5a5d0ef9df88e8, []int{4}
}
func (m *LookupNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupNodeRequest.Unmarshal(m, b)
}
func (m *LookupNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupNodeRequest.Marshal(b, m, deterministic)
}
func (dst *LookupNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupNodeRequest.Merge(dst, src)
}
func (m *LookupNodeRequest) XXX_Size() int {
	return xxx_messageInfo_LookupNodeRequest.Size(m)
}
func (m *LookupNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupNodeRequest proto.InternalMessageInfo

func (m *LookupNodeRequest) GetTarget() *ID {
	if m != nil {
		return m.Target
	}
	return nil
}

type LookupNodeResponse struct {
	Peers                []*ID    `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupNodeResponse) Reset()         { *m = LookupNodeResponse{} }
func (m *LookupNodeResponse) String() string { return proto.CompactTextString(m) }
func (*LookupNodeResponse) ProtoMessage()    {}
func (*LookupNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_bb5a5d0ef9df88e8, []int{5}
}
func (m *LookupNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupNodeResponse.Unmarshal(m, b)
}
func (m *LookupNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupNodeResponse.Marshal(b, m, deterministic)
}
func (dst *LookupNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupNodeResponse.Merge(dst, src)
}
func (m *LookupNodeResponse) XXX_Size() int {
	return xxx_messageInfo_LookupNodeResponse.Size(m)
}
func (m *LookupNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupNodeResponse proto.InternalMessageInfo

func (m *LookupNodeResponse) GetPeers() []*ID {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*ID)(nil), "protobuf.ID")
	proto.RegisterType((*Message)(nil), "protobuf.Message")
	proto.RegisterType((*HandshakeRequest)(nil), "protobuf.HandshakeRequest")
	proto.RegisterType((*HandshakeResponse)(nil), "protobuf.HandshakeResponse")
	proto.RegisterType((*LookupNodeRequest)(nil), "protobuf.LookupNodeRequest")
	proto.RegisterType((*LookupNodeResponse)(nil), "protobuf.LookupNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoiseClient is the client API for Noise service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoiseClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Noise_StreamClient, error)
}

type noiseClient struct {
	cc *grpc.ClientConn
}

func NewNoiseClient(cc *grpc.ClientConn) NoiseClient {
	return &noiseClient{cc}
}

func (c *noiseClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Noise_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Noise_serviceDesc.Streams[0], "/protobuf.Noise/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &noiseStreamClient{stream}
	return x, nil
}

type Noise_StreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type noiseStreamClient struct {
	grpc.ClientStream
}

func (x *noiseStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *noiseStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NoiseServer is the server API for Noise service.
type NoiseServer interface {
	Stream(Noise_StreamServer) error
}

func RegisterNoiseServer(s *grpc.Server, srv NoiseServer) {
	s.RegisterService(&_Noise_serviceDesc, srv)
}

func _Noise_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NoiseServer).Stream(&noiseStreamServer{stream})
}

type Noise_StreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type noiseStreamServer struct {
	grpc.ServerStream
}

func (x *noiseStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *noiseStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Noise_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Noise",
	HandlerType: (*NoiseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Noise_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/stream.proto",
}

func init() { proto.RegisterFile("protobuf/stream.proto", fileDescriptor_stream_bb5a5d0ef9df88e8) }

var fileDescriptor_stream_bb5a5d0ef9df88e8 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0xef, 0x00, 0xe5, 0xcf, 0x81, 0xc5, 0x65, 0x2e, 0xd7, 0x54, 0xa2, 0xb1, 0x69, 0x8c,
	0xe9, 0xaa, 0x18, 0x74, 0xa1, 0x0b, 0x16, 0x12, 0x16, 0x12, 0x95, 0x90, 0xf1, 0x01, 0xc8, 0x40,
	0x8f, 0xb5, 0x01, 0x66, 0xea, 0x4c, 0xbb, 0xe8, 0x7b, 0xf9, 0x80, 0x86, 0x4e, 0x0b, 0x44, 0x76,
	0x73, 0xbe, 0xf3, 0xcd, 0xef, 0x9c, 0x76, 0xe0, 0x7f, 0xac, 0x64, 0x22, 0x97, 0xe9, 0xc7, 0x40,
	0x27, 0x0a, 0xf9, 0xd6, 0xcf, 0x6b, 0xda, 0x2c, 0x71, 0xff, 0x3c, 0x94, 0x32, 0xdc, 0xe0, 0x60,
	0xef, 0x71, 0x91, 0x19, 0xc9, 0x1d, 0x41, 0x65, 0x3a, 0xa1, 0x97, 0x00, 0x71, 0xba, 0xdc, 0x44,
	0xab, 0xc5, 0x1a, 0x33, 0x9b, 0x38, 0xc4, 0xeb, 0xb0, 0x96, 0x21, 0x2f, 0x98, 0x51, 0x1b, 0x1a,
	0x3c, 0x08, 0x14, 0x6a, 0x6d, 0x57, 0x1c, 0xe2, 0xb5, 0x58, 0x59, 0xba, 0xdf, 0x04, 0x1a, 0x6f,
	0xa8, 0x35, 0x0f, 0x91, 0xfa, 0xd0, 0xd8, 0x9a, 0x63, 0x9e, 0xd0, 0x1e, 0xf6, 0x7c, 0x33, 0xd7,
	0x2f, 0xe7, 0xfa, 0x4f, 0x22, 0x63, 0xa5, 0x44, 0xaf, 0xa1, 0xae, 0x51, 0x04, 0xa8, 0xf2, 0xd0,
	0xf6, 0xb0, 0x73, 0xf0, 0xa6, 0x13, 0x56, 0xf4, 0xe8, 0x05, 0xb4, 0x74, 0x14, 0x0a, 0x9e, 0xa4,
	0x0a, 0xed, 0xaa, 0xd9, 0x6c, 0x0f, 0x68, 0x0f, 0x2c, 0x21, 0xc5, 0x0a, 0xed, 0x9a, 0x43, 0xbc,
	0x1a, 0x33, 0x05, 0xbd, 0x82, 0x76, 0xa4, 0x17, 0x0a, 0x75, 0x2c, 0x85, 0x46, 0xdb, 0x72, 0x88,
	0xd7, 0x64, 0x10, 0x69, 0x56, 0x10, 0x97, 0xc2, 0xdf, 0x67, 0x2e, 0x02, 0xfd, 0xc9, 0xd7, 0xc8,
	0xf0, 0x2b, 0x45, 0x9d, 0xb8, 0xff, 0xa0, 0x7b, 0xc4, 0x0a, 0xf1, 0x11, 0xba, 0xaf, 0x52, 0xae,
	0xd3, 0x78, 0x26, 0x83, 0xd2, 0xdc, 0x2d, 0x9e, 0x70, 0x15, 0x62, 0x52, 0x7c, 0xe7, 0xaf, 0xc5,
	0x4d, 0xcf, 0x7d, 0x00, 0x7a, 0x7c, 0xd5, 0x04, 0x52, 0x17, 0xac, 0x18, 0x51, 0x69, 0x9b, 0x38,
	0xd5, 0x93, 0xab, 0xa6, 0x35, 0x1c, 0x81, 0x35, 0x93, 0x91, 0x46, 0x7a, 0x0f, 0xf5, 0xf7, 0xfc,
	0x45, 0x69, 0xf7, 0xe0, 0x15, 0xbf, 0xbb, 0x7f, 0x8a, 0xdc, 0x3f, 0x1e, 0xb9, 0x25, 0xe3, 0x1b,
	0x38, 0x93, 0x2a, 0xf4, 0x63, 0x54, 0x9b, 0x48, 0xf8, 0x62, 0x97, 0x64, 0xd4, 0x31, 0xe4, 0xb1,
	0xf3, 0xdd, 0x79, 0x4e, 0x96, 0xf5, 0x1c, 0xde, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xf4,
	0x2f, 0x58, 0x3f, 0x02, 0x00, 0x00,
}
