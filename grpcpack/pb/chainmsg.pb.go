// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chainmsg.proto

package rpc_btc

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

type AddressRequest struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressRequest) Reset()         { *m = AddressRequest{} }
func (m *AddressRequest) String() string { return proto.CompactTextString(m) }
func (*AddressRequest) ProtoMessage()    {}
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8380b8d99192274c, []int{0}
}

func (m *AddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressRequest.Unmarshal(m, b)
}
func (m *AddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressRequest.Marshal(b, m, deterministic)
}
func (m *AddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressRequest.Merge(m, src)
}
func (m *AddressRequest) XXX_Size() int {
	return xxx_messageInfo_AddressRequest.Size(m)
}
func (m *AddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressRequest proto.InternalMessageInfo

func (m *AddressRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type AddressReply struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressReply) Reset()         { *m = AddressReply{} }
func (m *AddressReply) String() string { return proto.CompactTextString(m) }
func (*AddressReply) ProtoMessage()    {}
func (*AddressReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8380b8d99192274c, []int{1}
}

func (m *AddressReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressReply.Unmarshal(m, b)
}
func (m *AddressReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressReply.Marshal(b, m, deterministic)
}
func (m *AddressReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressReply.Merge(m, src)
}
func (m *AddressReply) XXX_Size() int {
	return xxx_messageInfo_AddressReply.Size(m)
}
func (m *AddressReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddressReply proto.InternalMessageInfo

func (m *AddressReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8380b8d99192274c, []int{2}
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

type BlockCountReply struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockCountReply) Reset()         { *m = BlockCountReply{} }
func (m *BlockCountReply) String() string { return proto.CompactTextString(m) }
func (*BlockCountReply) ProtoMessage()    {}
func (*BlockCountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8380b8d99192274c, []int{3}
}

func (m *BlockCountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockCountReply.Unmarshal(m, b)
}
func (m *BlockCountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockCountReply.Marshal(b, m, deterministic)
}
func (m *BlockCountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockCountReply.Merge(m, src)
}
func (m *BlockCountReply) XXX_Size() int {
	return xxx_messageInfo_BlockCountReply.Size(m)
}
func (m *BlockCountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockCountReply.DiscardUnknown(m)
}

var xxx_messageInfo_BlockCountReply proto.InternalMessageInfo

func (m *BlockCountReply) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type MiningInfoReply struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MiningInfoReply) Reset()         { *m = MiningInfoReply{} }
func (m *MiningInfoReply) String() string { return proto.CompactTextString(m) }
func (*MiningInfoReply) ProtoMessage()    {}
func (*MiningInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8380b8d99192274c, []int{4}
}

func (m *MiningInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MiningInfoReply.Unmarshal(m, b)
}
func (m *MiningInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MiningInfoReply.Marshal(b, m, deterministic)
}
func (m *MiningInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiningInfoReply.Merge(m, src)
}
func (m *MiningInfoReply) XXX_Size() int {
	return xxx_messageInfo_MiningInfoReply.Size(m)
}
func (m *MiningInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MiningInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_MiningInfoReply proto.InternalMessageInfo

func (m *MiningInfoReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*AddressRequest)(nil), "rpc.btc.AddressRequest")
	proto.RegisterType((*AddressReply)(nil), "rpc.btc.AddressReply")
	proto.RegisterType((*EmptyRequest)(nil), "rpc.btc.EmptyRequest")
	proto.RegisterType((*BlockCountReply)(nil), "rpc.btc.BlockCountReply")
	proto.RegisterType((*MiningInfoReply)(nil), "rpc.btc.MiningInfoReply")
}

func init() { proto.RegisterFile("chainmsg.proto", fileDescriptor_8380b8d99192274c) }

var fileDescriptor_8380b8d99192274c = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xce, 0x48, 0xcc,
	0xcc, 0xcb, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2a, 0x48, 0xd6,
	0x4b, 0x2a, 0x49, 0x56, 0x52, 0xe3, 0xe2, 0x73, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x0e, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0x49, 0x4c, 0x4a, 0xcd, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x34, 0xb8, 0x78, 0xe0, 0xea, 0x0a, 0x72, 0x2a,
	0x85, 0x24, 0xb8, 0xd8, 0x13, 0x21, 0x7c, 0xa8, 0x3a, 0x18, 0x57, 0x89, 0x8f, 0x8b, 0xc7, 0x35,
	0xb7, 0xa0, 0xa4, 0x12, 0x6a, 0x9e, 0x92, 0x3a, 0x17, 0xbf, 0x53, 0x4e, 0x7e, 0x72, 0xb6, 0x73,
	0x7e, 0x69, 0x5e, 0x09, 0x44, 0xb3, 0x08, 0x17, 0x6b, 0x32, 0x88, 0x07, 0xd6, 0xca, 0x1a, 0x04,
	0xe1, 0x28, 0xa9, 0x72, 0xf1, 0xfb, 0x66, 0xe6, 0x65, 0xe6, 0xa5, 0x7b, 0xe6, 0xa5, 0xe5, 0x43,
	0x14, 0x0a, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x42, 0xad, 0x00, 0xb3, 0x8d, 0xae, 0x33, 0x72,
	0x71, 0x39, 0x95, 0x24, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x39, 0x72, 0xf1, 0xba,
	0xa7, 0x96, 0xf8, 0xa5, 0x96, 0x43, 0x9d, 0x27, 0x24, 0xae, 0x07, 0xf5, 0x9b, 0x1e, 0xaa, 0xc7,
	0xa4, 0x44, 0x31, 0x25, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x84, 0x9c, 0xc0, 0x46, 0x20, 0x1c, 0x29,
	0x84, 0x50, 0x89, 0xec, 0x13, 0x29, 0x09, 0xb8, 0x30, 0x9a, 0x87, 0xe0, 0x66, 0x20, 0xdc, 0x4f,
	0xd8, 0x0c, 0x34, 0xbf, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xe3, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x0b, 0xfa, 0x85, 0xad, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BtcServiceClient is the client API for BtcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BtcServiceClient interface {
	GetNewAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressReply, error)
	GetBlockCount(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*BlockCountReply, error)
	GetMiningInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MiningInfoReply, error)
}

type btcServiceClient struct {
	cc *grpc.ClientConn
}

func NewBtcServiceClient(cc *grpc.ClientConn) BtcServiceClient {
	return &btcServiceClient{cc}
}

func (c *btcServiceClient) GetNewAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressReply, error) {
	out := new(AddressReply)
	err := c.cc.Invoke(ctx, "/rpc.btc.BtcService/GetNewAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *btcServiceClient) GetBlockCount(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*BlockCountReply, error) {
	out := new(BlockCountReply)
	err := c.cc.Invoke(ctx, "/rpc.btc.BtcService/GetBlockCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *btcServiceClient) GetMiningInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MiningInfoReply, error) {
	out := new(MiningInfoReply)
	err := c.cc.Invoke(ctx, "/rpc.btc.BtcService/GetMiningInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BtcServiceServer is the server API for BtcService service.
type BtcServiceServer interface {
	GetNewAddress(context.Context, *AddressRequest) (*AddressReply, error)
	GetBlockCount(context.Context, *EmptyRequest) (*BlockCountReply, error)
	GetMiningInfo(context.Context, *EmptyRequest) (*MiningInfoReply, error)
}

// UnimplementedBtcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBtcServiceServer struct {
}

func (*UnimplementedBtcServiceServer) GetNewAddress(ctx context.Context, req *AddressRequest) (*AddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewAddress not implemented")
}
func (*UnimplementedBtcServiceServer) GetBlockCount(ctx context.Context, req *EmptyRequest) (*BlockCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockCount not implemented")
}
func (*UnimplementedBtcServiceServer) GetMiningInfo(ctx context.Context, req *EmptyRequest) (*MiningInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMiningInfo not implemented")
}

func RegisterBtcServiceServer(s *grpc.Server, srv BtcServiceServer) {
	s.RegisterService(&_BtcService_serviceDesc, srv)
}

func _BtcService_GetNewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BtcServiceServer).GetNewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.btc.BtcService/GetNewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BtcServiceServer).GetNewAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BtcService_GetBlockCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BtcServiceServer).GetBlockCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.btc.BtcService/GetBlockCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BtcServiceServer).GetBlockCount(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BtcService_GetMiningInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BtcServiceServer).GetMiningInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.btc.BtcService/GetMiningInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BtcServiceServer).GetMiningInfo(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BtcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.btc.BtcService",
	HandlerType: (*BtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewAddress",
			Handler:    _BtcService_GetNewAddress_Handler,
		},
		{
			MethodName: "GetBlockCount",
			Handler:    _BtcService_GetBlockCount_Handler,
		},
		{
			MethodName: "GetMiningInfo",
			Handler:    _BtcService_GetMiningInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chainmsg.proto",
}
