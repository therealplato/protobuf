// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test/test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	proto/test/test.proto

It has these top-level messages:
	ExtraStuff
	PingRequest
	PingResponse
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/johanbrandhorst/protobuf/proto"

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

type PingRequest_FailureType int32

const (
	PingRequest_NONE PingRequest_FailureType = 0
	PingRequest_CODE PingRequest_FailureType = 1
	PingRequest_DROP PingRequest_FailureType = 2
)

var PingRequest_FailureType_name = map[int32]string{
	0: "NONE",
	1: "CODE",
	2: "DROP",
}
var PingRequest_FailureType_value = map[string]int32{
	"NONE": 0,
	"CODE": 1,
	"DROP": 2,
}

func (x PingRequest_FailureType) String() string {
	return proto.EnumName(PingRequest_FailureType_name, int32(x))
}
func (PingRequest_FailureType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type ExtraStuff struct {
	Addresses map[int32]string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Title:
	//	*ExtraStuff_FirstName
	//	*ExtraStuff_IdNumber
	Title       isExtraStuff_Title `protobuf_oneof:"title"`
	CardNumbers []uint32           `protobuf:"varint,4,rep,packed,name=card_numbers,json=cardNumbers" json:"card_numbers,omitempty"`
}

func (m *ExtraStuff) Reset()                    { *m = ExtraStuff{} }
func (m *ExtraStuff) String() string            { return proto.CompactTextString(m) }
func (*ExtraStuff) ProtoMessage()               {}
func (*ExtraStuff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isExtraStuff_Title interface {
	isExtraStuff_Title()
}

type ExtraStuff_FirstName struct {
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,oneof"`
}
type ExtraStuff_IdNumber struct {
	IdNumber int32 `protobuf:"varint,3,opt,name=id_number,json=idNumber,oneof"`
}

func (*ExtraStuff_FirstName) isExtraStuff_Title() {}
func (*ExtraStuff_IdNumber) isExtraStuff_Title()  {}

func (m *ExtraStuff) GetTitle() isExtraStuff_Title {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *ExtraStuff) GetAddresses() map[int32]string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ExtraStuff) GetFirstName() string {
	if x, ok := m.GetTitle().(*ExtraStuff_FirstName); ok {
		return x.FirstName
	}
	return ""
}

func (m *ExtraStuff) GetIdNumber() int32 {
	if x, ok := m.GetTitle().(*ExtraStuff_IdNumber); ok {
		return x.IdNumber
	}
	return 0
}

func (m *ExtraStuff) GetCardNumbers() []uint32 {
	if m != nil {
		return m.CardNumbers
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtraStuff) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtraStuff_OneofMarshaler, _ExtraStuff_OneofUnmarshaler, _ExtraStuff_OneofSizer, []interface{}{
		(*ExtraStuff_FirstName)(nil),
		(*ExtraStuff_IdNumber)(nil),
	}
}

func _ExtraStuff_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtraStuff)
	// title
	switch x := m.Title.(type) {
	case *ExtraStuff_FirstName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FirstName)
	case *ExtraStuff_IdNumber:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IdNumber))
	case nil:
	default:
		return fmt.Errorf("ExtraStuff.Title has unexpected type %T", x)
	}
	return nil
}

func _ExtraStuff_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtraStuff)
	switch tag {
	case 2: // title.first_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Title = &ExtraStuff_FirstName{x}
		return true, err
	case 3: // title.id_number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Title = &ExtraStuff_IdNumber{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _ExtraStuff_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtraStuff)
	// title
	switch x := m.Title.(type) {
	case *ExtraStuff_FirstName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FirstName)))
		n += len(x.FirstName)
	case *ExtraStuff_IdNumber:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IdNumber))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PingRequest struct {
	Value             string                  `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	ResponseCount     int32                   `protobuf:"varint,2,opt,name=response_count,json=responseCount" json:"response_count,omitempty"`
	ErrorCodeReturned uint32                  `protobuf:"varint,3,opt,name=error_code_returned,json=errorCodeReturned" json:"error_code_returned,omitempty"`
	FailureType       PingRequest_FailureType `protobuf:"varint,4,opt,name=failure_type,json=failureType,enum=test.PingRequest_FailureType" json:"failure_type,omitempty"`
	CheckMetadata     bool                    `protobuf:"varint,5,opt,name=check_metadata,json=checkMetadata" json:"check_metadata,omitempty"`
	SendHeaders       bool                    `protobuf:"varint,6,opt,name=send_headers,json=sendHeaders" json:"send_headers,omitempty"`
	SendTrailers      bool                    `protobuf:"varint,7,opt,name=send_trailers,json=sendTrailers" json:"send_trailers,omitempty"`
	MessageLatencyMs  int32                   `protobuf:"varint,8,opt,name=message_latency_ms,json=messageLatencyMs" json:"message_latency_ms,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

func (m *PingRequest) GetFailureType() PingRequest_FailureType {
	if m != nil {
		return m.FailureType
	}
	return PingRequest_NONE
}

func (m *PingRequest) GetCheckMetadata() bool {
	if m != nil {
		return m.CheckMetadata
	}
	return false
}

func (m *PingRequest) GetSendHeaders() bool {
	if m != nil {
		return m.SendHeaders
	}
	return false
}

func (m *PingRequest) GetSendTrailers() bool {
	if m != nil {
		return m.SendTrailers
	}
	return false
}

func (m *PingRequest) GetMessageLatencyMs() int32 {
	if m != nil {
		return m.MessageLatencyMs
	}
	return 0
}

type PingResponse struct {
	Value   string `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
	Counter int32  `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtraStuff)(nil), "test.ExtraStuff")
	proto.RegisterType((*PingRequest)(nil), "test.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "test.PingResponse")
	proto.RegisterEnum("test.PingRequest_FailureType", PingRequest_FailureType_name, PingRequest_FailureType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.TestService/PingEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.TestService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/test.TestService/PingError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/test.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	PingEmpty(context.Context, *google_protobuf.Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*google_protobuf.Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/test/test.proto",
}

func init() { proto.RegisterFile("proto/test/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x8d, 0x9b, 0xa4, 0x4d, 0xc6, 0x4d, 0x95, 0xee, 0xf7, 0x81, 0xac, 0xa0, 0xaa, 0x21, 0x08,
	0x29, 0x12, 0xc8, 0x41, 0x45, 0x48, 0xa5, 0x02, 0x54, 0xda, 0x06, 0xf5, 0xd0, 0xa6, 0xd5, 0xb6,
	0xe2, 0xc0, 0xc5, 0xda, 0xd8, 0x93, 0xc4, 0xad, 0xed, 0x0d, 0xbb, 0xeb, 0x0a, 0xff, 0x01, 0x7e,
	0x2c, 0x77, 0x2e, 0x9c, 0xd0, 0xee, 0x3a, 0x4a, 0x10, 0x54, 0xea, 0xc5, 0x9a, 0x79, 0xf3, 0xde,
	0xcc, 0xec, 0x8c, 0xc6, 0xf0, 0x68, 0x2e, 0xb8, 0xe2, 0x03, 0x85, 0x52, 0x99, 0x8f, 0x6f, 0x7c,
	0x52, 0xd3, 0x76, 0xe7, 0xc9, 0x94, 0xf3, 0x69, 0x82, 0x03, 0x83, 0x8d, 0xf3, 0xc9, 0x00, 0xd3,
	0xb9, 0x2a, 0x2c, 0xa5, 0xb3, 0x3f, 0x8d, 0xd5, 0x2c, 0x1f, 0xfb, 0x21, 0x4f, 0x07, 0x37, 0x7c,
	0xc6, 0xb2, 0xb1, 0x60, 0x59, 0x34, 0xe3, 0x42, 0xaa, 0xa5, 0xc0, 0x66, 0x9f, 0xf2, 0xf9, 0x0c,
	0xc5, 0x8d, 0xb4, 0xca, 0xde, 0x4f, 0x07, 0x60, 0xf8, 0x4d, 0x09, 0x76, 0xa5, 0xf2, 0xc9, 0x84,
	0xbc, 0x87, 0x26, 0x8b, 0x22, 0x81, 0x52, 0xa2, 0xf4, 0x9c, 0x6e, 0xb5, 0xef, 0xee, 0xed, 0xfa,
	0xa6, 0x97, 0x25, 0xc9, 0xff, 0xb8, 0x60, 0x0c, 0x33, 0x25, 0x0a, 0xba, 0x54, 0x90, 0x5d, 0x80,
	0x49, 0x2c, 0xa4, 0x0a, 0x32, 0x96, 0xa2, 0xb7, 0xd6, 0x75, 0xfa, 0xcd, 0xd3, 0x0a, 0x6d, 0x1a,
	0x6c, 0xc4, 0x52, 0x24, 0x3b, 0xd0, 0x8c, 0xa3, 0x20, 0xcb, 0xd3, 0x31, 0x0a, 0xaf, 0xda, 0x75,
	0xfa, 0xf5, 0xd3, 0x0a, 0x6d, 0xc4, 0xd1, 0xc8, 0x20, 0xe4, 0x29, 0x6c, 0x86, 0x4c, 0x2c, 0x08,
	0xd2, 0xab, 0x75, 0xab, 0xfd, 0x16, 0x75, 0x35, 0x66, 0x19, 0xb2, 0xf3, 0x0e, 0xb6, 0xfe, 0xac,
	0x4f, 0xda, 0x50, 0xbd, 0xc5, 0xc2, 0x73, 0x74, 0x36, 0xaa, 0x4d, 0xf2, 0x3f, 0xd4, 0xef, 0x58,
	0x92, 0x97, 0x1d, 0x50, 0xeb, 0x1c, 0xac, 0xed, 0x3b, 0x47, 0x1b, 0x50, 0x57, 0xb1, 0x4a, 0xb0,
	0xf7, 0xbd, 0x0a, 0xee, 0x65, 0x9c, 0x4d, 0x29, 0x7e, 0xcd, 0x51, 0xaa, 0xa5, 0xc4, 0x59, 0x91,
	0x90, 0xe7, 0xb0, 0x25, 0x50, 0xce, 0x79, 0x26, 0x31, 0x08, 0x79, 0x9e, 0x29, 0x93, 0xb1, 0x4e,
	0x5b, 0x0b, 0xf4, 0x58, 0x83, 0xc4, 0x87, 0xff, 0x50, 0x08, 0x2e, 0x82, 0x90, 0x47, 0x18, 0x08,
	0x54, 0xb9, 0xc8, 0x30, 0x32, 0xef, 0x6b, 0xd1, 0x6d, 0x13, 0x3a, 0xe6, 0x11, 0xd2, 0x32, 0x40,
	0x0e, 0x61, 0x73, 0xc2, 0xe2, 0x24, 0x17, 0x18, 0xa8, 0x62, 0x8e, 0x5e, 0xad, 0xeb, 0xf4, 0xb7,
	0xf6, 0x76, 0xec, 0xa0, 0x57, 0xba, 0xf2, 0x3f, 0x59, 0xd6, 0x75, 0x31, 0x47, 0xea, 0x4e, 0x96,
	0x8e, 0x6e, 0x2c, 0x9c, 0x61, 0x78, 0x1b, 0xa4, 0xa8, 0x58, 0xc4, 0x14, 0xf3, 0xea, 0x5d, 0xa7,
	0xdf, 0xa0, 0x2d, 0x83, 0x9e, 0x97, 0xa0, 0x9e, 0xa7, 0xc4, 0x2c, 0x0a, 0x66, 0xc8, 0x22, 0x3d,
	0xcf, 0x75, 0x43, 0x72, 0x35, 0x76, 0x6a, 0x21, 0xf2, 0x0c, 0x5a, 0x86, 0xa2, 0x04, 0x8b, 0x13,
	0xcd, 0xd9, 0x30, 0x1c, 0xa3, 0xbb, 0x2e, 0x31, 0xf2, 0x12, 0x48, 0x8a, 0x52, 0xb2, 0x29, 0x06,
	0x09, 0x53, 0x98, 0x85, 0x45, 0x90, 0x4a, 0xaf, 0x61, 0x66, 0xd1, 0x2e, 0x23, 0x67, 0x36, 0x70,
	0x2e, 0x7b, 0x2f, 0xc0, 0x5d, 0x69, 0x9c, 0x34, 0xa0, 0x36, 0xba, 0x18, 0x0d, 0xdb, 0x15, 0x6d,
	0x1d, 0x5f, 0x9c, 0x0c, 0xdb, 0x8e, 0xb6, 0x4e, 0xe8, 0xc5, 0x65, 0x7b, 0xad, 0xf7, 0x01, 0x36,
	0xed, 0x8b, 0xed, 0x40, 0xf5, 0x22, 0x3e, 0xaf, 0x2e, 0xc2, 0x38, 0xc4, 0x83, 0x0d, 0x33, 0x7f,
	0x14, 0xe5, 0x06, 0x16, 0xee, 0xde, 0x0f, 0x07, 0xdc, 0x6b, 0x94, 0xea, 0x0a, 0xc5, 0x5d, 0x1c,
	0x22, 0x79, 0x0b, 0x4d, 0x9d, 0x6f, 0xa8, 0xaf, 0x83, 0x3c, 0xf6, 0xed, 0xd5, 0xf8, 0x8b, 0x23,
	0xf0, 0x0d, 0xde, 0x21, 0xab, 0xa3, 0xb6, 0x85, 0x7b, 0x15, 0x32, 0x80, 0x9a, 0x46, 0xc8, 0xf6,
	0x5f, 0x8b, 0xb8, 0x47, 0xb0, 0x5f, 0xd6, 0xd2, 0x0b, 0xfe, 0x97, 0xea, 0x9e, 0xf2, 0xbd, 0x0a,
	0x79, 0x03, 0x0d, 0x4d, 0x3c, 0x8b, 0xa5, 0x7a, 0x70, 0xb9, 0x57, 0xce, 0x51, 0xf1, 0xeb, 0xf0,
	0xe0, 0x21, 0xa7, 0x6e, 0x7e, 0x21, 0x61, 0x12, 0x63, 0x56, 0x82, 0x06, 0xf9, 0xf2, 0x70, 0xad,
	0x44, 0x71, 0x87, 0x62, 0x45, 0x3b, 0x5e, 0x37, 0xf6, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0x47, 0xa4, 0x76, 0xa5, 0x04, 0x00, 0x00,
}
