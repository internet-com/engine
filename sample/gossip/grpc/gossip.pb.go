// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gossip.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	gossip.proto

It has these top-level messages:
	Empty
	GossipTable
	PeerInfo
	PeersInfo
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GossipTable struct {
	MyID     string      `protobuf:"bytes,1,opt,name=myID" json:"myID,omitempty"`
	PeerInfo []*PeerInfo `protobuf:"bytes,2,rep,name=peerInfo" json:"peerInfo,omitempty"`
}

func (m *GossipTable) Reset()                    { *m = GossipTable{} }
func (m *GossipTable) String() string            { return proto.CompactTextString(m) }
func (*GossipTable) ProtoMessage()               {}
func (*GossipTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GossipTable) GetMyID() string {
	if m != nil {
		return m.MyID
	}
	return ""
}

func (m *GossipTable) GetPeerInfo() []*PeerInfo {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type PeerInfo struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ipAddress" json:"ipAddress,omitempty"`
	PeerID    string `protobuf:"bytes,2,opt,name=peerID" json:"peerID,omitempty"`
	Counter   int64  `protobuf:"varint,3,opt,name=counter" json:"counter,omitempty"`
}

func (m *PeerInfo) Reset()                    { *m = PeerInfo{} }
func (m *PeerInfo) String() string            { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()               {}
func (*PeerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PeerInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *PeerInfo) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *PeerInfo) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

type PeersInfo struct {
}

func (m *PeersInfo) Reset()                    { *m = PeersInfo{} }
func (m *PeersInfo) String() string            { return proto.CompactTextString(m) }
func (*PeersInfo) ProtoMessage()               {}
func (*PeersInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Empty)(nil), "grpc.Empty")
	proto.RegisterType((*GossipTable)(nil), "grpc.GossipTable")
	proto.RegisterType((*PeerInfo)(nil), "grpc.PeerInfo")
	proto.RegisterType((*PeersInfo)(nil), "grpc.PeersInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Gossip service

type GossipClient interface {
	PushGossip(ctx context.Context, in *GossipTable, opts ...grpc1.CallOption) (*Empty, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*Empty, error)
	GetGossipTable(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*GossipTable, error)
	PullGossip(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*GossipTable, error)
}

type gossipClient struct {
	cc *grpc1.ClientConn
}

func NewGossipClient(cc *grpc1.ClientConn) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) PushGossip(ctx context.Context, in *GossipTable, opts ...grpc1.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc1.Invoke(ctx, "/grpc.Gossip/PushGossip", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) Ping(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc1.Invoke(ctx, "/grpc.Gossip/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) GetGossipTable(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*GossipTable, error) {
	out := new(GossipTable)
	err := grpc1.Invoke(ctx, "/grpc.Gossip/GetGossipTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) PullGossip(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*GossipTable, error) {
	out := new(GossipTable)
	err := grpc1.Invoke(ctx, "/grpc.Gossip/PullGossip", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gossip service

type GossipServer interface {
	PushGossip(context.Context, *GossipTable) (*Empty, error)
	Ping(context.Context, *Empty) (*Empty, error)
	GetGossipTable(context.Context, *Empty) (*GossipTable, error)
	PullGossip(context.Context, *Empty) (*GossipTable, error)
}

func RegisterGossipServer(s *grpc1.Server, srv GossipServer) {
	s.RegisterService(&_Gossip_serviceDesc, srv)
}

func _Gossip_PushGossip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(GossipTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).PushGossip(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Gossip/PushGossip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).PushGossip(ctx, req.(*GossipTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).Ping(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Gossip/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_GetGossipTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).GetGossipTable(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Gossip/GetGossipTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).GetGossipTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_PullGossip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).PullGossip(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Gossip/PullGossip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).PullGossip(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gossip_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "PushGossip",
			Handler:    _Gossip_PushGossip_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Gossip_Ping_Handler,
		},
		{
			MethodName: "GetGossipTable",
			Handler:    _Gossip_GetGossipTable_Handler,
		},
		{
			MethodName: "PullGossip",
			Handler:    _Gossip_PullGossip_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "gossip.proto",
}

func init() { proto.RegisterFile("gossip.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0x2f, 0x2e,
	0xce, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x56, 0x62,
	0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0xf2, 0xe5, 0xe2, 0x76, 0x07, 0x4b, 0x87, 0x24,
	0x26, 0xe5, 0xa4, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x56, 0x7a, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x42, 0x5a, 0x5c, 0x1c, 0x05, 0xa9, 0xa9, 0x45, 0x9e, 0x79, 0x69, 0xf9,
	0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x7c, 0x7a, 0x20, 0x43, 0xf4, 0x02, 0xa0, 0xa2, 0x41,
	0x70, 0x79, 0xa5, 0x28, 0x2e, 0x0e, 0x98, 0xa8, 0x90, 0x0c, 0x17, 0x67, 0x66, 0x81, 0x63, 0x4a,
	0x4a, 0x51, 0x6a, 0x71, 0x31, 0xd4, 0x40, 0x84, 0x80, 0x90, 0x18, 0x17, 0x1b, 0x58, 0x97, 0x8b,
	0x04, 0x13, 0x58, 0x0a, 0xca, 0x13, 0x92, 0xe0, 0x62, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d,
	0x92, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x71, 0x95, 0xb8, 0xb9, 0x38, 0x41, 0x66, 0x17,
	0x83, 0x0c, 0x37, 0x3a, 0xc0, 0xc8, 0xc5, 0x06, 0x71, 0xb8, 0x90, 0x1e, 0x17, 0x57, 0x40, 0x69,
	0x71, 0x06, 0x94, 0x27, 0x08, 0x71, 0x1b, 0x92, 0xa7, 0xa4, 0xb8, 0x21, 0x42, 0x10, 0x0f, 0x33,
	0x08, 0x29, 0x71, 0xb1, 0x04, 0x64, 0xe6, 0xa5, 0x0b, 0x21, 0x0b, 0xa3, 0xab, 0x31, 0xe2, 0xe2,
	0x73, 0x4f, 0x2d, 0x41, 0x0e, 0x19, 0x14, 0xd5, 0x98, 0x96, 0x28, 0x31, 0x40, 0xdc, 0x91, 0x93,
	0x03, 0x75, 0x07, 0x41, 0xf5, 0x49, 0x6c, 0xe0, 0x08, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0xe5, 0x3a, 0xde, 0xa0, 0x01, 0x00, 0x00,
}
