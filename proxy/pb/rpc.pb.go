// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package pb

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

type RecipientNodeInfo struct {
	RecipientNodePeerId  string   `protobuf:"bytes,1,opt,name=recipient_node_peer_id,json=recipientNodePeerId,proto3" json:"recipient_node_peer_id,omitempty"`
	RecipientUserPeerId  string   `protobuf:"bytes,2,opt,name=recipient_user_peer_id,json=recipientUserPeerId,proto3" json:"recipient_user_peer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecipientNodeInfo) Reset()         { *m = RecipientNodeInfo{} }
func (m *RecipientNodeInfo) String() string { return proto.CompactTextString(m) }
func (*RecipientNodeInfo) ProtoMessage()    {}
func (*RecipientNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *RecipientNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipientNodeInfo.Unmarshal(m, b)
}
func (m *RecipientNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipientNodeInfo.Marshal(b, m, deterministic)
}
func (m *RecipientNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipientNodeInfo.Merge(m, src)
}
func (m *RecipientNodeInfo) XXX_Size() int {
	return xxx_messageInfo_RecipientNodeInfo.Size(m)
}
func (m *RecipientNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipientNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecipientNodeInfo proto.InternalMessageInfo

func (m *RecipientNodeInfo) GetRecipientNodePeerId() string {
	if m != nil {
		return m.RecipientNodePeerId
	}
	return ""
}

func (m *RecipientNodeInfo) GetRecipientUserPeerId() string {
	if m != nil {
		return m.RecipientUserPeerId
	}
	return ""
}

type OpenChannelRequest struct {
	NodePubkeyString     string             `protobuf:"bytes,1,opt,name=node_pubkey_string,json=nodePubkeyString,proto3" json:"node_pubkey_string,omitempty"`
	NodePubkeyIndex      int32              `protobuf:"varint,2,opt,name=node_pubkey_index,json=nodePubkeyIndex,proto3" json:"node_pubkey_index,omitempty"`
	Private              bool               `protobuf:"varint,3,opt,name=private,proto3" json:"private,omitempty"`
	RecipientInfo        *RecipientNodeInfo `protobuf:"bytes,4,opt,name=recipientInfo,proto3" json:"recipientInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OpenChannelRequest) Reset()         { *m = OpenChannelRequest{} }
func (m *OpenChannelRequest) String() string { return proto.CompactTextString(m) }
func (*OpenChannelRequest) ProtoMessage()    {}
func (*OpenChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *OpenChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelRequest.Unmarshal(m, b)
}
func (m *OpenChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelRequest.Marshal(b, m, deterministic)
}
func (m *OpenChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelRequest.Merge(m, src)
}
func (m *OpenChannelRequest) XXX_Size() int {
	return xxx_messageInfo_OpenChannelRequest.Size(m)
}
func (m *OpenChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelRequest proto.InternalMessageInfo

func (m *OpenChannelRequest) GetNodePubkeyString() string {
	if m != nil {
		return m.NodePubkeyString
	}
	return ""
}

func (m *OpenChannelRequest) GetNodePubkeyIndex() int32 {
	if m != nil {
		return m.NodePubkeyIndex
	}
	return 0
}

func (m *OpenChannelRequest) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

func (m *OpenChannelRequest) GetRecipientInfo() *RecipientNodeInfo {
	if m != nil {
		return m.RecipientInfo
	}
	return nil
}

type OpenChannelResponse struct {
	TemplateChannelId    string   `protobuf:"bytes,1,opt,name=template_channel_id,json=templateChannelId,proto3" json:"template_channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenChannelResponse) Reset()         { *m = OpenChannelResponse{} }
func (m *OpenChannelResponse) String() string { return proto.CompactTextString(m) }
func (*OpenChannelResponse) ProtoMessage()    {}
func (*OpenChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *OpenChannelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelResponse.Unmarshal(m, b)
}
func (m *OpenChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelResponse.Marshal(b, m, deterministic)
}
func (m *OpenChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelResponse.Merge(m, src)
}
func (m *OpenChannelResponse) XXX_Size() int {
	return xxx_messageInfo_OpenChannelResponse.Size(m)
}
func (m *OpenChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelResponse proto.InternalMessageInfo

func (m *OpenChannelResponse) GetTemplateChannelId() string {
	if m != nil {
		return m.TemplateChannelId
	}
	return ""
}

type FundChannelRequest struct {
	TemplateChannelId    string             `protobuf:"bytes,1,opt,name=template_channel_id,json=templateChannelId,proto3" json:"template_channel_id,omitempty"`
	BtcAmount            float64            `protobuf:"fixed64,2,opt,name=btc_amount,json=btcAmount,proto3" json:"btc_amount,omitempty"`
	PropertyId           int64              `protobuf:"varint,3,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	AssetAmount          float64            `protobuf:"fixed64,4,opt,name=asset_amount,json=assetAmount,proto3" json:"asset_amount,omitempty"`
	RecipientInfo        *RecipientNodeInfo `protobuf:"bytes,5,opt,name=recipientInfo,proto3" json:"recipientInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FundChannelRequest) Reset()         { *m = FundChannelRequest{} }
func (m *FundChannelRequest) String() string { return proto.CompactTextString(m) }
func (*FundChannelRequest) ProtoMessage()    {}
func (*FundChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *FundChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundChannelRequest.Unmarshal(m, b)
}
func (m *FundChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundChannelRequest.Marshal(b, m, deterministic)
}
func (m *FundChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundChannelRequest.Merge(m, src)
}
func (m *FundChannelRequest) XXX_Size() int {
	return xxx_messageInfo_FundChannelRequest.Size(m)
}
func (m *FundChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FundChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FundChannelRequest proto.InternalMessageInfo

func (m *FundChannelRequest) GetTemplateChannelId() string {
	if m != nil {
		return m.TemplateChannelId
	}
	return ""
}

func (m *FundChannelRequest) GetBtcAmount() float64 {
	if m != nil {
		return m.BtcAmount
	}
	return 0
}

func (m *FundChannelRequest) GetPropertyId() int64 {
	if m != nil {
		return m.PropertyId
	}
	return 0
}

func (m *FundChannelRequest) GetAssetAmount() float64 {
	if m != nil {
		return m.AssetAmount
	}
	return 0
}

func (m *FundChannelRequest) GetRecipientInfo() *RecipientNodeInfo {
	if m != nil {
		return m.RecipientInfo
	}
	return nil
}

type FundChannelResponse struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundChannelResponse) Reset()         { *m = FundChannelResponse{} }
func (m *FundChannelResponse) String() string { return proto.CompactTextString(m) }
func (*FundChannelResponse) ProtoMessage()    {}
func (*FundChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *FundChannelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundChannelResponse.Unmarshal(m, b)
}
func (m *FundChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundChannelResponse.Marshal(b, m, deterministic)
}
func (m *FundChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundChannelResponse.Merge(m, src)
}
func (m *FundChannelResponse) XXX_Size() int {
	return xxx_messageInfo_FundChannelResponse.Size(m)
}
func (m *FundChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundChannelResponse proto.InternalMessageInfo

func (m *FundChannelResponse) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

type HelloRequest struct {
	Sayhi                string   `protobuf:"bytes,1,opt,name=sayhi,proto3" json:"sayhi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetSayhi() string {
	if m != nil {
		return m.Sayhi
	}
	return ""
}

type HelloResponse struct {
	Resp                 string   `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func init() {
	proto.RegisterType((*RecipientNodeInfo)(nil), "proxy.RecipientNodeInfo")
	proto.RegisterType((*OpenChannelRequest)(nil), "proxy.OpenChannelRequest")
	proto.RegisterType((*OpenChannelResponse)(nil), "proxy.OpenChannelResponse")
	proto.RegisterType((*FundChannelRequest)(nil), "proxy.FundChannelRequest")
	proto.RegisterType((*FundChannelResponse)(nil), "proxy.FundChannelResponse")
	proto.RegisterType((*HelloRequest)(nil), "proxy.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "proxy.HelloResponse")
}

func init() {
	proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1)
}

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x36, 0x06, 0x3c, 0x69, 0x05, 0x19, 0x57, 0xc8, 0x44, 0xaa, 0x08, 0x86, 0x43,
	0x84, 0x50, 0x0e, 0x29, 0x67, 0x24, 0xfe, 0x0a, 0x4b, 0x08, 0x2a, 0x23, 0x2e, 0x5c, 0x2c, 0xff,
	0x19, 0x1a, 0x8b, 0x74, 0x77, 0xd9, 0x5d, 0xa3, 0xe6, 0xc0, 0xfb, 0x21, 0xf1, 0x20, 0xbc, 0x46,
	0xe5, 0xcd, 0xda, 0xb1, 0xeb, 0x5e, 0x72, 0xf3, 0xce, 0x7c, 0xdf, 0x37, 0x9e, 0x9f, 0xd7, 0xe0,
	0x49, 0x91, 0x2f, 0x84, 0xe4, 0x9a, 0xa3, 0x2b, 0x24, 0xbf, 0xda, 0x84, 0x7f, 0x60, 0x12, 0x53,
	0x5e, 0x8a, 0x92, 0x98, 0xfe, 0xcc, 0x0b, 0x8a, 0xd8, 0x0f, 0x8e, 0x67, 0xf0, 0x50, 0x36, 0xc5,
	0x84, 0xf1, 0x82, 0x12, 0x41, 0x24, 0x93, 0xb2, 0x08, 0x9c, 0x99, 0x33, 0xf7, 0x62, 0x5f, 0x76,
	0x2d, 0xe7, 0x44, 0x32, 0x2a, 0xfa, 0xa6, 0x4a, 0x91, 0x6c, 0x4d, 0x07, 0x37, 0x4c, 0xdf, 0x14,
	0xc9, 0xad, 0x29, 0xfc, 0xeb, 0x00, 0x7e, 0x11, 0xc4, 0xde, 0xae, 0x52, 0xc6, 0x68, 0x1d, 0xd3,
	0xaf, 0x8a, 0x94, 0xc6, 0x17, 0x80, 0xdb, 0xb1, 0x55, 0xf6, 0x93, 0x36, 0x89, 0xd2, 0xb2, 0x64,
	0x17, 0x76, 0xf8, 0x83, 0xba, 0x73, 0x6e, 0x1a, 0x5f, 0x4d, 0x1d, 0x9f, 0xc3, 0xa4, 0xab, 0x2e,
	0x59, 0x41, 0x57, 0x66, 0xa8, 0x1b, 0xdf, 0xdf, 0x89, 0xa3, 0xba, 0x8c, 0x01, 0xdc, 0x15, 0xb2,
	0xfc, 0x9d, 0x6a, 0x0a, 0x0e, 0x67, 0xce, 0xfc, 0x5e, 0xdc, 0x1c, 0xf1, 0x15, 0x1c, 0xb7, 0x6f,
	0x58, 0x53, 0x08, 0x46, 0x33, 0x67, 0x3e, 0x5e, 0x06, 0x0b, 0x03, 0x6a, 0x31, 0xa0, 0x14, 0xf7,
	0xe5, 0xe1, 0x7b, 0xf0, 0x7b, 0x9b, 0x28, 0xc1, 0x99, 0x22, 0x5c, 0x80, 0xaf, 0xe9, 0x52, 0xac,
	0x53, 0x4d, 0x49, 0xbe, 0xed, 0xed, 0x40, 0x4e, 0x9a, 0x96, 0x75, 0x45, 0x45, 0xf8, 0xdf, 0x01,
	0xfc, 0x50, 0xb1, 0xe2, 0x06, 0x91, 0x3d, 0x63, 0xf0, 0x14, 0x20, 0xd3, 0x79, 0x92, 0x5e, 0xf2,
	0x8a, 0x69, 0x03, 0xc3, 0x89, 0xbd, 0x4c, 0xe7, 0xaf, 0x4d, 0x01, 0x1f, 0xc3, 0x58, 0x48, 0x2e,
	0x48, 0xea, 0x4d, 0x1d, 0x53, 0xa3, 0x38, 0x8c, 0xa1, 0x29, 0x45, 0x05, 0x3e, 0x81, 0xa3, 0x54,
	0x29, 0xd2, 0x4d, 0xc2, 0xc8, 0x24, 0x8c, 0x4d, 0xcd, 0x66, 0x0c, 0x80, 0xb9, 0xfb, 0x01, 0x7b,
	0x09, 0x7e, 0x6f, 0x51, 0x0b, 0xec, 0x14, 0x60, 0xb0, 0xa0, 0x97, 0xb7, 0x7c, 0x9e, 0xc1, 0xd1,
	0x47, 0x5a, 0xaf, 0x79, 0x03, 0xe6, 0x04, 0x5c, 0x95, 0x6e, 0x56, 0xa5, 0x55, 0x6e, 0x0f, 0xe1,
	0x53, 0x38, 0xb6, 0x2a, 0x9b, 0x8a, 0x30, 0x92, 0xa4, 0x84, 0x55, 0x99, 0xe7, 0xe5, 0x3f, 0x07,
	0xbc, 0x4f, 0xe5, 0xc5, 0x4a, 0xb3, 0xfa, 0x16, 0x2d, 0xc1, 0x35, 0x16, 0xf4, 0xed, 0x02, 0xdd,
	0x31, 0xd3, 0x93, 0x7e, 0xd1, 0xa6, 0xbe, 0x83, 0x71, 0xe7, 0x9b, 0xe3, 0x23, 0x2b, 0x1a, 0xde,
	0xe8, 0xe9, 0xf4, 0xb6, 0xd6, 0x2e, 0xa5, 0x03, 0xa2, 0x4d, 0x19, 0xde, 0x82, 0x36, 0xe5, 0x16,
	0x6e, 0x6f, 0x46, 0xdf, 0x0f, 0x44, 0x96, 0xdd, 0x31, 0x7f, 0xf7, 0xd9, 0x75, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0x38, 0x50, 0xd1, 0xea, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LightningClient is the client API for Lightning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LightningClient interface {
	// obdcli: `hello`
	//hello is a `say hi` gRPC.
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelResponse, error)
	FundChannel(ctx context.Context, in *FundChannelRequest, opts ...grpc.CallOption) (*FundChannelResponse, error)
}

type lightningClient struct {
	cc grpc.ClientConnInterface
}

func NewLightningClient(cc grpc.ClientConnInterface) LightningClient {
	return &lightningClient{cc}
}

func (c *lightningClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proxy.Lightning/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningClient) OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelResponse, error) {
	out := new(OpenChannelResponse)
	err := c.cc.Invoke(ctx, "/proxy.Lightning/OpenChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningClient) FundChannel(ctx context.Context, in *FundChannelRequest, opts ...grpc.CallOption) (*FundChannelResponse, error) {
	out := new(FundChannelResponse)
	err := c.cc.Invoke(ctx, "/proxy.Lightning/FundChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightningServer is the server API for Lightning service.
type LightningServer interface {
	// obdcli: `hello`
	//hello is a `say hi` gRPC.
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	OpenChannel(context.Context, *OpenChannelRequest) (*OpenChannelResponse, error)
	FundChannel(context.Context, *FundChannelRequest) (*FundChannelResponse, error)
}

// UnimplementedLightningServer can be embedded to have forward compatible implementations.
type UnimplementedLightningServer struct {
}

func (*UnimplementedLightningServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedLightningServer) OpenChannel(ctx context.Context, req *OpenChannelRequest) (*OpenChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenChannel not implemented")
}
func (*UnimplementedLightningServer) FundChannel(ctx context.Context, req *FundChannelRequest) (*FundChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundChannel not implemented")
}

func RegisterLightningServer(s *grpc.Server, srv LightningServer) {
	s.RegisterService(&_Lightning_serviceDesc, srv)
}

func _Lightning_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.Lightning/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lightning_OpenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).OpenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.Lightning/OpenChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).OpenChannel(ctx, req.(*OpenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lightning_FundChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).FundChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.Lightning/FundChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).FundChannel(ctx, req.(*FundChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lightning_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.Lightning",
	HandlerType: (*LightningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Lightning_Hello_Handler,
		},
		{
			MethodName: "OpenChannel",
			Handler:    _Lightning_OpenChannel_Handler,
		},
		{
			MethodName: "FundChannel",
			Handler:    _Lightning_FundChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
