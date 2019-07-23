// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracer.proto

package tracerprotobuf

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

// The request message containing the user's name.
type SpanBatch struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanBatch) Reset()         { *m = SpanBatch{} }
func (m *SpanBatch) String() string { return proto.CompactTextString(m) }
func (*SpanBatch) ProtoMessage()    {}
func (*SpanBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{0}
}

func (m *SpanBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanBatch.Unmarshal(m, b)
}
func (m *SpanBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanBatch.Marshal(b, m, deterministic)
}
func (m *SpanBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanBatch.Merge(m, src)
}
func (m *SpanBatch) XXX_Size() int {
	return xxx_messageInfo_SpanBatch.Size(m)
}
func (m *SpanBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanBatch.DiscardUnknown(m)
}

var xxx_messageInfo_SpanBatch proto.InternalMessageInfo

func (m *SpanBatch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type BatchResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchResponse) Reset()         { *m = BatchResponse{} }
func (m *BatchResponse) String() string { return proto.CompactTextString(m) }
func (*BatchResponse) ProtoMessage()    {}
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{1}
}

func (m *BatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchResponse.Unmarshal(m, b)
}
func (m *BatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchResponse.Marshal(b, m, deterministic)
}
func (m *BatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchResponse.Merge(m, src)
}
func (m *BatchResponse) XXX_Size() int {
	return xxx_messageInfo_BatchResponse.Size(m)
}
func (m *BatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchResponse proto.InternalMessageInfo

func (m *BatchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SpanBatch)(nil), "tracerprotobuf.SpanBatch")
	proto.RegisterType((*BatchResponse)(nil), "tracerprotobuf.BatchResponse")
}

func init() { proto.RegisterFile("tracer.proto", fileDescriptor_6d422d7c66fbbd8f) }

var fileDescriptor_6d422d7c66fbbd8f = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x29, 0x4a, 0x4c,
	0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xf0, 0xc0, 0x9c, 0xa4, 0xd2,
	0x34, 0x25, 0x79, 0x2e, 0xce, 0xe0, 0x82, 0xc4, 0x3c, 0xa7, 0xc4, 0x92, 0xe4, 0x0c, 0x21, 0x21,
	0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49,
	0x93, 0x8b, 0x17, 0x2c, 0x19, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc1, 0xc5,
	0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0x53, 0x07, 0xe3, 0x1a, 0x05, 0x72, 0xb1, 0x85, 0x80,
	0x4d, 0x17, 0x72, 0xe7, 0xe2, 0x0c, 0x4e, 0xcd, 0x4b, 0x81, 0x98, 0x2a, 0xa9, 0x87, 0x6a, 0xa7,
	0x1e, 0xdc, 0x42, 0x29, 0x59, 0x74, 0x29, 0x14, 0xab, 0x94, 0x18, 0x9c, 0x0c, 0xb8, 0xa4, 0x33,
	0xf3, 0xf5, 0xd2, 0x8b, 0x0a, 0x92, 0xf5, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x8b, 0xf5,
	0x32, 0x52, 0x73, 0x72, 0xf2, 0xcb, 0xf3, 0x8b, 0x72, 0x52, 0x9c, 0xf8, 0x3d, 0x40, 0xec, 0x70,
	0x10, 0x3b, 0x00, 0x64, 0x44, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x54, 0x0c, 0x51, 0xf7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TracerClient is the client API for Tracer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracerClient interface {
	// Sends a batch
	SendBatch(ctx context.Context, in *SpanBatch, opts ...grpc.CallOption) (*BatchResponse, error)
}

type tracerClient struct {
	cc *grpc.ClientConn
}

func NewTracerClient(cc *grpc.ClientConn) TracerClient {
	return &tracerClient{cc}
}

func (c *tracerClient) SendBatch(ctx context.Context, in *SpanBatch, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, "/tracerprotobuf.Tracer/SendBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracerServer is the server API for Tracer service.
type TracerServer interface {
	// Sends a batch
	SendBatch(context.Context, *SpanBatch) (*BatchResponse, error)
}

// UnimplementedTracerServer can be embedded to have forward compatible implementations.
type UnimplementedTracerServer struct {
}

func (*UnimplementedTracerServer) SendBatch(ctx context.Context, req *SpanBatch) (*BatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBatch not implemented")
}

func RegisterTracerServer(s *grpc.Server, srv TracerServer) {
	s.RegisterService(&_Tracer_serviceDesc, srv)
}

func _Tracer_SendBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).SendBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracerprotobuf.Tracer/SendBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).SendBatch(ctx, req.(*SpanBatch))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tracerprotobuf.Tracer",
	HandlerType: (*TracerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBatch",
			Handler:    _Tracer_SendBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracer.proto",
}
