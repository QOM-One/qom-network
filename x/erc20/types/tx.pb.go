// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qom/erc20/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgConvertCoin defines a Msg to convert a native Cosmos coin to a ERC20 token
type MsgConvertCoin struct {
	// Cosmos coin which denomination is registered in a token pair. The coin
	// amount defines the amount of coins to convert.
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin"`
	// recipient hex address to receive ERC20 token
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// cosmos bech32 address from the owner of the given Cosmos coins
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertCoin) Reset()         { *m = MsgConvertCoin{} }
func (m *MsgConvertCoin) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCoin) ProtoMessage()    {}
func (*MsgConvertCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_7537931906c3702f, []int{0}
}
func (m *MsgConvertCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCoin.Merge(m, src)
}
func (m *MsgConvertCoin) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCoin proto.InternalMessageInfo

func (m *MsgConvertCoin) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *MsgConvertCoin) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertCoin) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertCoinResponse returns no fields
type MsgConvertCoinResponse struct {
}

func (m *MsgConvertCoinResponse) Reset()         { *m = MsgConvertCoinResponse{} }
func (m *MsgConvertCoinResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCoinResponse) ProtoMessage()    {}
func (*MsgConvertCoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7537931906c3702f, []int{1}
}
func (m *MsgConvertCoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCoinResponse.Merge(m, src)
}
func (m *MsgConvertCoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCoinResponse proto.InternalMessageInfo

// MsgConvertERC20 defines a Msg to convert a ERC20 token to a native Cosmos
// coin.
type MsgConvertERC20 struct {
	// ERC20 token contract address registered in a token pair
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// amount of ERC20 tokens to convert
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	// bech32 address to receive native Cosmos coins
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender hex address from the owner of the given ERC20 tokens
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertERC20) Reset()         { *m = MsgConvertERC20{} }
func (m *MsgConvertERC20) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC20) ProtoMessage()    {}
func (*MsgConvertERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_7537931906c3702f, []int{2}
}
func (m *MsgConvertERC20) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC20.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC20.Merge(m, src)
}
func (m *MsgConvertERC20) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC20.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC20 proto.InternalMessageInfo

func (m *MsgConvertERC20) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgConvertERC20) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertERC20) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertERC20Response returns no fields
type MsgConvertERC20Response struct {
}

func (m *MsgConvertERC20Response) Reset()         { *m = MsgConvertERC20Response{} }
func (m *MsgConvertERC20Response) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC20Response) ProtoMessage()    {}
func (*MsgConvertERC20Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7537931906c3702f, []int{3}
}
func (m *MsgConvertERC20Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC20Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC20Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC20Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC20Response.Merge(m, src)
}
func (m *MsgConvertERC20Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC20Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC20Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC20Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertCoin)(nil), "qom.erc20.v1.MsgConvertCoin")
	proto.RegisterType((*MsgConvertCoinResponse)(nil), "qom.erc20.v1.MsgConvertCoinResponse")
	proto.RegisterType((*MsgConvertERC20)(nil), "qom.erc20.v1.MsgConvertERC20")
	proto.RegisterType((*MsgConvertERC20Response)(nil), "qom.erc20.v1.MsgConvertERC20Response")
}

func init() { proto.RegisterFile("qom/erc20/v1/tx.proto", fileDescriptor_7537931906c3702f) }

var fileDescriptor_7537931906c3702f = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0x62, 0x63, 0x9a, 0x4d, 0x68, 0xca, 0xd2, 0xa6, 0x8e, 0x48, 0x64, 0x63, 0x9a, 0x90,
	0x16, 0xb2, 0x13, 0x3b, 0x87, 0x9e, 0x63, 0xb7, 0x85, 0x1e, 0x4c, 0x88, 0x8e, 0xbd, 0x04, 0x79,
	0xb5, 0xa8, 0xa2, 0xd5, 0x8e, 0xb2, 0xbb, 0x11, 0x0e, 0xf4, 0xd4, 0x2f, 0x28, 0xf4, 0x67, 0xfa,
	0x09, 0x39, 0x06, 0x7a, 0x29, 0x3d, 0x84, 0x62, 0xf7, 0x37, 0x0a, 0x45, 0x2b, 0xc5, 0xb1, 0x02,
	0xce, 0x49, 0xbb, 0xef, 0x8d, 0xde, 0xcc, 0xbc, 0xb7, 0xe4, 0xd9, 0x39, 0x26, 0x20, 0x14, 0xef,
	0x1f, 0x42, 0xd6, 0x03, 0x33, 0x61, 0xa9, 0x42, 0x83, 0x74, 0xfd, 0x1c, 0x13, 0x66, 0x61, 0x96,
	0xf5, 0xdc, 0xed, 0x08, 0x31, 0xfa, 0x2c, 0x20, 0x48, 0x63, 0x08, 0xa4, 0x44, 0x13, 0x98, 0x18,
	0xa5, 0x2e, 0x6a, 0xdd, 0xa7, 0x11, 0x46, 0x68, 0x8f, 0x90, 0x9f, 0x4a, 0xd4, 0xe3, 0xa8, 0x13,
	0xd4, 0x30, 0x0e, 0xb4, 0x80, 0xac, 0x37, 0x16, 0x26, 0xe8, 0x01, 0xc7, 0x58, 0x16, 0x7c, 0xf7,
	0x92, 0x3c, 0x1e, 0xe9, 0x68, 0x88, 0x32, 0x13, 0xca, 0x0c, 0x31, 0x96, 0xf4, 0x88, 0x34, 0x72,
	0xbe, 0xe5, 0x74, 0x9c, 0xfd, 0xb5, 0xfe, 0x16, 0x2b, 0x04, 0x58, 0x2e, 0xc0, 0x4a, 0x01, 0x96,
	0x17, 0x0e, 0x1a, 0x57, 0x37, 0xed, 0x9a, 0x6f, 0x8b, 0xa9, 0x4b, 0x1e, 0x29, 0xc1, 0x45, 0x9c,
	0x09, 0xd5, 0x5a, 0xe9, 0x38, 0xfb, 0xab, 0xfe, 0xfc, 0x4e, 0x37, 0x49, 0x53, 0x0b, 0x19, 0x0a,
	0xd5, 0xaa, 0x5b, 0xa6, 0xbc, 0x75, 0x5b, 0x64, 0xb3, 0xda, 0xda, 0x17, 0x3a, 0x45, 0xa9, 0x45,
	0xf7, 0x87, 0x43, 0x36, 0xee, 0xa8, 0xb7, 0xfe, 0xb0, 0x7f, 0x48, 0x5f, 0x92, 0x27, 0x1c, 0xa5,
	0x51, 0x01, 0x37, 0x67, 0x41, 0x18, 0x2a, 0xa1, 0xb5, 0x1d, 0x71, 0xd5, 0xdf, 0xb8, 0xc5, 0x8f,
	0x0b, 0x98, 0xbe, 0x23, 0xcd, 0x20, 0xc1, 0x0b, 0x69, 0x8a, 0x51, 0x06, 0x2c, 0x1f, 0xf4, 0xf7,
	0x4d, 0x7b, 0x2f, 0x8a, 0xcd, 0xc7, 0x8b, 0x31, 0xe3, 0x98, 0x40, 0x69, 0x4b, 0xf1, 0x39, 0xd0,
	0xe1, 0x27, 0x30, 0x97, 0xa9, 0xd0, 0xec, 0xbd, 0x34, 0x7e, 0xf9, 0x77, 0x65, 0xa9, 0xfa, 0xd2,
	0xa5, 0x1a, 0x95, 0xa5, 0xb6, 0xc8, 0xf3, 0x7b, 0x93, 0xdf, 0x6e, 0xd5, 0xff, 0xe7, 0x90, 0xfa,
	0x48, 0x47, 0x74, 0x42, 0xd6, 0x16, 0xfd, 0xde, 0x66, 0x8b, 0x21, 0xb3, 0xaa, 0x25, 0xee, 0x8b,
	0x87, 0xd8, 0xb9, 0x61, 0xbb, 0x5f, 0x7f, 0xfe, 0xfd, 0xbe, 0xd2, 0xa6, 0x3b, 0x70, 0xef, 0x1d,
	0x01, 0x2f, 0xaa, 0xcf, 0x6c, 0x4a, 0x5f, 0xc8, 0x7a, 0xc5, 0xd3, 0x9d, 0x65, 0xe2, 0x96, 0x76,
	0x77, 0x1f, 0xa4, 0xe7, 0xcd, 0xf7, 0x6c, 0xf3, 0x0e, 0xf5, 0x96, 0x36, 0xb7, 0xd8, 0xe0, 0xcd,
	0xd5, 0xd4, 0x73, 0xae, 0xa7, 0x9e, 0xf3, 0x67, 0xea, 0x39, 0xdf, 0x66, 0x5e, 0xed, 0x7a, 0xe6,
	0xd5, 0x7e, 0xcd, 0xbc, 0xda, 0x87, 0x57, 0x0b, 0xc1, 0x9c, 0x9e, 0x8c, 0x0e, 0x4e, 0xa4, 0x80,
	0x53, 0x4c, 0x8e, 0xd3, 0x14, 0xb2, 0xd7, 0x30, 0x29, 0x35, 0x6d, 0x40, 0xe3, 0xa6, 0x7d, 0xb7,
	0x47, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x67, 0x48, 0x1b, 0x64, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertCoin mints a ERC20 representation of the native Cosmos coin denom
	// that is registered on the token mapping.
	ConvertCoin(ctx context.Context, in *MsgConvertCoin, opts ...grpc.CallOption) (*MsgConvertCoinResponse, error)
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertCoin(ctx context.Context, in *MsgConvertCoin, opts ...grpc.CallOption) (*MsgConvertCoinResponse, error) {
	out := new(MsgConvertCoinResponse)
	err := c.cc.Invoke(ctx, "/qom.erc20.v1.Msg/ConvertCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error) {
	out := new(MsgConvertERC20Response)
	err := c.cc.Invoke(ctx, "/qom.erc20.v1.Msg/ConvertERC20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertCoin mints a ERC20 representation of the native Cosmos coin denom
	// that is registered on the token mapping.
	ConvertCoin(context.Context, *MsgConvertCoin) (*MsgConvertCoinResponse, error)
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(context.Context, *MsgConvertERC20) (*MsgConvertERC20Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertCoin(ctx context.Context, req *MsgConvertCoin) (*MsgConvertCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCoin not implemented")
}
func (*UnimplementedMsgServer) ConvertERC20(ctx context.Context, req *MsgConvertERC20) (*MsgConvertERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertERC20 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertCoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qom.erc20.v1.Msg/ConvertCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertCoin(ctx, req.(*MsgConvertCoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConvertERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qom.erc20.v1.Msg/ConvertERC20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertERC20(ctx, req.(*MsgConvertERC20))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qom.erc20.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertCoin",
			Handler:    _Msg_ConvertCoin_Handler,
		},
		{
			MethodName: "ConvertERC20",
			Handler:    _Msg_ConvertERC20_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qom/erc20/v1/tx.proto",
}

func (m *MsgConvertCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgConvertCoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC20) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC20) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC20) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC20Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC20Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC20Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvertCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertCoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConvertERC20) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertERC20Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvertCoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConvertCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConvertCoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConvertCoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConvertERC20) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConvertERC20: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC20: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConvertERC20Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConvertERC20Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC20Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
