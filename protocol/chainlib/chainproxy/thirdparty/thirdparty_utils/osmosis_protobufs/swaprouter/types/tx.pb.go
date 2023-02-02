// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/swaprouter/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// ===================== MsgSwapExactAmountIn
type MsgSwapExactAmountIn struct {
	Sender            string                                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	Routes            []SwapAmountInRoute                    `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes"`
	TokenIn           types.Coin                             `protobuf:"bytes,3,opt,name=token_in,json=tokenIn,proto3" json:"token_in" yaml:"token_in"`
	TokenOutMinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=token_out_min_amount,json=tokenOutMinAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_out_min_amount" yaml:"token_out_min_amount"`
}

func (m *MsgSwapExactAmountIn) Reset()         { *m = MsgSwapExactAmountIn{} }
func (m *MsgSwapExactAmountIn) String() string { return proto.CompactTextString(m) }
func (*MsgSwapExactAmountIn) ProtoMessage()    {}
func (*MsgSwapExactAmountIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a4da63b1afc25d, []int{0}
}
func (m *MsgSwapExactAmountIn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapExactAmountIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapExactAmountIn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapExactAmountIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapExactAmountIn.Merge(m, src)
}
func (m *MsgSwapExactAmountIn) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapExactAmountIn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapExactAmountIn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapExactAmountIn proto.InternalMessageInfo

func (m *MsgSwapExactAmountIn) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSwapExactAmountIn) GetRoutes() []SwapAmountInRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *MsgSwapExactAmountIn) GetTokenIn() types.Coin {
	if m != nil {
		return m.TokenIn
	}
	return types.Coin{}
}

type MsgSwapExactAmountInResponse struct {
	TokenOutAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=token_out_amount,json=tokenOutAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_out_amount" yaml:"token_out_amount"`
}

func (m *MsgSwapExactAmountInResponse) Reset()         { *m = MsgSwapExactAmountInResponse{} }
func (m *MsgSwapExactAmountInResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapExactAmountInResponse) ProtoMessage()    {}
func (*MsgSwapExactAmountInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a4da63b1afc25d, []int{1}
}
func (m *MsgSwapExactAmountInResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapExactAmountInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapExactAmountInResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapExactAmountInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapExactAmountInResponse.Merge(m, src)
}
func (m *MsgSwapExactAmountInResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapExactAmountInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapExactAmountInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapExactAmountInResponse proto.InternalMessageInfo

// ===================== MsgSwapExactAmountOut
type MsgSwapExactAmountOut struct {
	Sender           string                                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	Routes           []SwapAmountOutRoute                   `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes"`
	TokenInMaxAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=token_in_max_amount,json=tokenInMaxAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_in_max_amount" yaml:"token_in_max_amount"`
	TokenOut         types.Coin                             `protobuf:"bytes,4,opt,name=token_out,json=tokenOut,proto3" json:"token_out" yaml:"token_out"`
}

func (m *MsgSwapExactAmountOut) Reset()         { *m = MsgSwapExactAmountOut{} }
func (m *MsgSwapExactAmountOut) String() string { return proto.CompactTextString(m) }
func (*MsgSwapExactAmountOut) ProtoMessage()    {}
func (*MsgSwapExactAmountOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a4da63b1afc25d, []int{2}
}
func (m *MsgSwapExactAmountOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapExactAmountOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapExactAmountOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapExactAmountOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapExactAmountOut.Merge(m, src)
}
func (m *MsgSwapExactAmountOut) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapExactAmountOut) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapExactAmountOut.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapExactAmountOut proto.InternalMessageInfo

func (m *MsgSwapExactAmountOut) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSwapExactAmountOut) GetRoutes() []SwapAmountOutRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *MsgSwapExactAmountOut) GetTokenOut() types.Coin {
	if m != nil {
		return m.TokenOut
	}
	return types.Coin{}
}

type MsgSwapExactAmountOutResponse struct {
	TokenInAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=token_in_amount,json=tokenInAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_in_amount" yaml:"token_in_amount"`
}

func (m *MsgSwapExactAmountOutResponse) Reset()         { *m = MsgSwapExactAmountOutResponse{} }
func (m *MsgSwapExactAmountOutResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapExactAmountOutResponse) ProtoMessage()    {}
func (*MsgSwapExactAmountOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a4da63b1afc25d, []int{3}
}
func (m *MsgSwapExactAmountOutResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapExactAmountOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapExactAmountOutResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapExactAmountOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapExactAmountOutResponse.Merge(m, src)
}
func (m *MsgSwapExactAmountOutResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapExactAmountOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapExactAmountOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapExactAmountOutResponse proto.InternalMessageInfo

func init() {
	//proto.RegisterType((*MsgSwapExactAmountIn)(nil), "osmosis.swaprouter.v1beta1.MsgSwapExactAmountIn")
	//proto.RegisterType((*MsgSwapExactAmountInResponse)(nil), "osmosis.swaprouter.v1beta1.MsgSwapExactAmountInResponse")
	//proto.RegisterType((*MsgSwapExactAmountOut)(nil), "osmosis.swaprouter.v1beta1.MsgSwapExactAmountOut")
	//proto.RegisterType((*MsgSwapExactAmountOutResponse)(nil), "osmosis.swaprouter.v1beta1.MsgSwapExactAmountOutResponse")
}

func init() {
	//proto.RegisterFile("osmosis/swaprouter/v1beta1/tx.proto", fileDescriptor_05a4da63b1afc25d)
}

var fileDescriptor_05a4da63b1afc25d = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6b, 0xd3, 0x50,
	0x14, 0x6f, 0xda, 0x31, 0xb7, 0x3b, 0xe6, 0xda, 0x58, 0x5d, 0x8d, 0x9a, 0x96, 0x08, 0x52, 0x91,
	0xde, 0xd8, 0x0e, 0x44, 0x7d, 0xb3, 0x22, 0x58, 0x5c, 0xe8, 0x88, 0x6f, 0xbe, 0x94, 0xdb, 0x2e,
	0xd4, 0xb0, 0xe5, 0xde, 0xd0, 0x7b, 0xb3, 0x75, 0x08, 0x0a, 0xbe, 0xfa, 0xa2, 0xf8, 0xa5, 0xf6,
	0xb8, 0x47, 0x11, 0xac, 0xd2, 0x7e, 0x83, 0x7e, 0x02, 0xc9, 0xfd, 0x93, 0xb5, 0x5d, 0x2d, 0xe6,
	0x69, 0xd9, 0xc9, 0x39, 0xbf, 0xf3, 0xfb, 0x73, 0x1a, 0x70, 0x9f, 0xd0, 0x80, 0x50, 0x9f, 0xda,
	0xf4, 0x14, 0x85, 0x03, 0x12, 0x31, 0x6f, 0x60, 0x9f, 0xd4, 0xbb, 0x1e, 0x43, 0x75, 0x9b, 0x0d,
	0x61, 0x38, 0x20, 0x8c, 0xe8, 0x86, 0x6c, 0x82, 0x97, 0x4d, 0x50, 0x36, 0x19, 0xc5, 0x3e, 0xe9,
	0x13, 0xde, 0x66, 0xc7, 0x4f, 0x62, 0xc2, 0x30, 0x7b, 0x7c, 0xc4, 0xee, 0x22, 0xea, 0x25, 0x78,
	0x3d, 0xe2, 0x63, 0xf9, 0xfe, 0xd1, 0x8a, 0xb5, 0x71, 0xa9, 0xc3, 0x6b, 0xa2, 0xd9, 0xfa, 0x95,
	0x05, 0x45, 0x87, 0xf6, 0xdf, 0x9e, 0xa2, 0xf0, 0xd5, 0x10, 0xf5, 0xd8, 0x8b, 0x80, 0x44, 0x98,
	0xb5, 0xb0, 0xfe, 0x10, 0xac, 0x53, 0x0f, 0x1f, 0x7a, 0x83, 0x92, 0x56, 0xd1, 0xaa, 0x9b, 0xcd,
	0xc2, 0x74, 0x54, 0xde, 0x3e, 0x43, 0xc1, 0xf1, 0x73, 0x4b, 0xd4, 0x2d, 0x57, 0x36, 0xe8, 0x6f,
	0xc0, 0x3a, 0x87, 0xa4, 0xa5, 0x6c, 0x25, 0x57, 0xdd, 0x6a, 0xd4, 0xe0, 0xbf, 0x35, 0xc1, 0x78,
	0x93, 0x5a, 0xe2, 0xc6, 0xaf, 0x9a, 0x6b, 0xe7, 0xa3, 0x72, 0xc6, 0x95, 0x10, 0xba, 0x03, 0x36,
	0x18, 0x39, 0xf2, 0x70, 0xc7, 0xc7, 0xa5, 0x5c, 0x45, 0xab, 0x6e, 0x35, 0x6e, 0x43, 0x21, 0x18,
	0xc6, 0x82, 0x13, 0x9c, 0x97, 0xc4, 0xc7, 0xcd, 0xdd, 0x78, 0x74, 0x3a, 0x2a, 0xef, 0x08, 0x62,
	0x6a, 0xd0, 0x72, 0xaf, 0xf1, 0xc7, 0x16, 0xd6, 0x3f, 0x82, 0xa2, 0xa8, 0x92, 0x88, 0x75, 0x02,
	0x1f, 0x77, 0x10, 0xdf, 0x5d, 0x5a, 0xe3, 0xa2, 0x9c, 0x78, 0xfe, 0xe7, 0xa8, 0xfc, 0xa0, 0xef,
	0xb3, 0xf7, 0x51, 0x17, 0xf6, 0x48, 0x60, 0x4b, 0x77, 0xc5, 0x9f, 0x1a, 0x3d, 0x3c, 0xb2, 0xd9,
	0x59, 0xe8, 0x51, 0xd8, 0xc2, 0x6c, 0x3a, 0x2a, 0xdf, 0x99, 0xdd, 0x34, 0x8f, 0x69, 0xb9, 0x05,
	0x5e, 0x6e, 0x47, 0xcc, 0xf1, 0xb1, 0xd0, 0x68, 0x7d, 0xd7, 0xc0, 0xdd, 0x65, 0xfe, 0xba, 0x1e,
	0x0d, 0x09, 0xa6, 0x9e, 0x4e, 0x41, 0xfe, 0x12, 0x4c, 0x92, 0x13, 0x8e, 0xb7, 0x52, 0x93, 0xdb,
	0x5d, 0x24, 0xa7, 0x88, 0x5d, 0x57, 0xc4, 0x24, 0xab, 0xdf, 0x59, 0x70, 0xf3, 0x2a, 0xab, 0x76,
	0xc4, 0xd2, 0xc4, 0xbe, 0xbf, 0x10, 0x3b, 0xfc, 0xbf, 0xd8, 0xdb, 0x11, 0x5b, 0x96, 0xfb, 0x07,
	0x70, 0x43, 0xc5, 0xd7, 0x09, 0xd0, 0x50, 0x59, 0x91, 0xe3, 0x2c, 0xf6, 0x53, 0x5b, 0x61, 0xcc,
	0x5f, 0xc4, 0x0c, 0xa4, 0xe5, 0xe6, 0xe5, 0x71, 0x38, 0x68, 0x28, 0x28, 0xe9, 0x07, 0x60, 0x33,
	0x31, 0x8d, 0x9f, 0xc6, 0xca, 0xab, 0x2b, 0xc9, 0xab, 0xcb, 0x2f, 0xd8, 0x6d, 0xb9, 0x1b, 0xca,
	0x67, 0xeb, 0x9b, 0x06, 0xee, 0x2d, 0x75, 0x38, 0x09, 0x3e, 0x04, 0x3b, 0x09, 0xbb, 0xb9, 0xdc,
	0x5f, 0xa7, 0x16, 0x7b, 0x6b, 0x41, 0xac, 0x12, 0xba, 0x2d, 0x85, 0x8a, 0xe5, 0x8d, 0x2f, 0x59,
	0x90, 0x73, 0x68, 0x5f, 0xff, 0x04, 0x0a, 0x57, 0x7f, 0xef, 0x8f, 0x57, 0xa5, 0xb7, 0xec, 0x82,
	0x8d, 0xa7, 0x69, 0x27, 0x12, 0xe9, 0x9f, 0x35, 0xa0, 0x2f, 0xb9, 0xbd, 0x7a, 0x3a, 0xc0, 0x76,
	0xc4, 0x8c, 0x67, 0xa9, 0x47, 0x14, 0x89, 0xe6, 0xc1, 0xf9, 0xd8, 0xd4, 0x2e, 0xc6, 0xa6, 0xf6,
	0x67, 0x6c, 0x6a, 0x5f, 0x27, 0x66, 0xe6, 0x62, 0x62, 0x66, 0x7e, 0x4c, 0xcc, 0xcc, 0xbb, 0x27,
	0x33, 0xc6, 0x4b, 0xf8, 0xda, 0x31, 0xea, 0x52, 0xf5, 0x8f, 0x7d, 0x52, 0xdf, 0xb3, 0x87, 0xb3,
	0x9f, 0x57, 0x1e, 0x46, 0x77, 0x9d, 0x7f, 0x52, 0xf7, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0x56, 0xac, 0xa7, 0xf8, 0x05, 0x00, 0x00,
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
	SwapExactAmountIn(ctx context.Context, in *MsgSwapExactAmountIn, opts ...grpc.CallOption) (*MsgSwapExactAmountInResponse, error)
	SwapExactAmountOut(ctx context.Context, in *MsgSwapExactAmountOut, opts ...grpc.CallOption) (*MsgSwapExactAmountOutResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SwapExactAmountIn(ctx context.Context, in *MsgSwapExactAmountIn, opts ...grpc.CallOption) (*MsgSwapExactAmountInResponse, error) {
	out := new(MsgSwapExactAmountInResponse)
	err := c.cc.Invoke(ctx, "/osmosis.swaprouter.v1beta1.Msg/SwapExactAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapExactAmountOut(ctx context.Context, in *MsgSwapExactAmountOut, opts ...grpc.CallOption) (*MsgSwapExactAmountOutResponse, error) {
	out := new(MsgSwapExactAmountOutResponse)
	err := c.cc.Invoke(ctx, "/osmosis.swaprouter.v1beta1.Msg/SwapExactAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SwapExactAmountIn(context.Context, *MsgSwapExactAmountIn) (*MsgSwapExactAmountInResponse, error)
	SwapExactAmountOut(context.Context, *MsgSwapExactAmountOut) (*MsgSwapExactAmountOutResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SwapExactAmountIn(ctx context.Context, req *MsgSwapExactAmountIn) (*MsgSwapExactAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapExactAmountIn not implemented")
}
func (*UnimplementedMsgServer) SwapExactAmountOut(ctx context.Context, req *MsgSwapExactAmountOut) (*MsgSwapExactAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapExactAmountOut not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SwapExactAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapExactAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapExactAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.swaprouter.v1beta1.Msg/SwapExactAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapExactAmountIn(ctx, req.(*MsgSwapExactAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapExactAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapExactAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapExactAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.swaprouter.v1beta1.Msg/SwapExactAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapExactAmountOut(ctx, req.(*MsgSwapExactAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.swaprouter.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SwapExactAmountIn",
			Handler:    _Msg_SwapExactAmountIn_Handler,
		},
		{
			MethodName: "SwapExactAmountOut",
			Handler:    _Msg_SwapExactAmountOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/swaprouter/v1beta1/tx.proto",
}

func (m *MsgSwapExactAmountIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapExactAmountIn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapExactAmountIn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenOutMinAmount.Size()
		i -= size
		if _, err := m.TokenOutMinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TokenIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Routes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapExactAmountInResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapExactAmountInResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapExactAmountInResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenOutAmount.Size()
		i -= size
		if _, err := m.TokenOutAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgSwapExactAmountOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapExactAmountOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapExactAmountOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TokenInMaxAmount.Size()
		i -= size
		if _, err := m.TokenInMaxAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Routes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapExactAmountOutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapExactAmountOutResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapExactAmountOutResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenInAmount.Size()
		i -= size
		if _, err := m.TokenInAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *MsgSwapExactAmountIn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = m.TokenIn.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.TokenOutMinAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapExactAmountInResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenOutAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapExactAmountOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = m.TokenInMaxAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.TokenOut.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapExactAmountOutResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenInAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwapExactAmountIn) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapExactAmountIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapExactAmountIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
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
			m.Routes = append(m.Routes, SwapAmountInRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
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
			if err := m.TokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutMinAmount", wireType)
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
			if err := m.TokenOutMinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSwapExactAmountInResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapExactAmountInResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapExactAmountInResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutAmount", wireType)
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
			if err := m.TokenOutAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSwapExactAmountOut) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapExactAmountOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapExactAmountOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
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
			m.Routes = append(m.Routes, SwapAmountOutRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInMaxAmount", wireType)
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
			if err := m.TokenInMaxAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
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
			if err := m.TokenOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSwapExactAmountOutResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapExactAmountOutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapExactAmountOutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInAmount", wireType)
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
			if err := m.TokenInAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
