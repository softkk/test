// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CalcRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{0}
}

func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (m *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(m, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *CalcRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type CalcReply struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcReply) Reset()         { *m = CalcReply{} }
func (m *CalcReply) String() string { return proto.CompactTextString(m) }
func (*CalcReply) ProtoMessage()    {}
func (*CalcReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{1}
}

func (m *CalcReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcReply.Unmarshal(m, b)
}
func (m *CalcReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcReply.Marshal(b, m, deterministic)
}
func (m *CalcReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcReply.Merge(m, src)
}
func (m *CalcReply) XXX_Size() int {
	return xxx_messageInfo_CalcReply.Size(m)
}
func (m *CalcReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcReply.DiscardUnknown(m)
}

var xxx_messageInfo_CalcReply proto.InternalMessageInfo

func (m *CalcReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalcRequest)(nil), "pb.CalcRequest")
	proto.RegisterType((*CalcReply)(nil), "pb.CalcReply")
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor_a2b9900dc883ea68) }

var fileDescriptor_a2b9900dc883ea68 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe4, 0xe2, 0x76, 0x4e,
	0xcc, 0x49, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25,
	0x65, 0x2e, 0x4e, 0x88, 0xd2, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2,
	0x9c, 0x12, 0xa8, 0x6a, 0x28, 0xcf, 0xc8, 0x8c, 0x8b, 0x0b, 0xa4, 0xa8, 0x34, 0x27, 0xb1, 0x24,
	0xbf, 0x48, 0x48, 0x83, 0x8b, 0x25, 0x20, 0xa7, 0xb4, 0x58, 0x88, 0x5f, 0xaf, 0x20, 0x49, 0x0f,
	0xc9, 0x1e, 0x29, 0x5e, 0x84, 0x40, 0x41, 0x4e, 0xa5, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x49, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x3f, 0x87, 0xbc, 0xa0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error) {
	out := new(CalcReply)
	err := c.cc.Invoke(ctx, "/pb.Calculator/Plus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Plus(context.Context, *CalcRequest) (*CalcReply, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Calculator/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Plus(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plus",
			Handler:    _Calculator_Plus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}
