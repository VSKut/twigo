// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/grpc/proto/tweet.proto

package grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateTweetRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTweetRequest) Reset()         { *m = CreateTweetRequest{} }
func (m *CreateTweetRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTweetRequest) ProtoMessage()    {}
func (*CreateTweetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce250c0854693d1b, []int{0}
}

func (m *CreateTweetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTweetRequest.Unmarshal(m, b)
}
func (m *CreateTweetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTweetRequest.Marshal(b, m, deterministic)
}
func (m *CreateTweetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTweetRequest.Merge(m, src)
}
func (m *CreateTweetRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTweetRequest.Size(m)
}
func (m *CreateTweetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTweetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTweetRequest proto.InternalMessageInfo

func (m *CreateTweetRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateTweetResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTweetResponse) Reset()         { *m = CreateTweetResponse{} }
func (m *CreateTweetResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTweetResponse) ProtoMessage()    {}
func (*CreateTweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce250c0854693d1b, []int{1}
}

func (m *CreateTweetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTweetResponse.Unmarshal(m, b)
}
func (m *CreateTweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTweetResponse.Marshal(b, m, deterministic)
}
func (m *CreateTweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTweetResponse.Merge(m, src)
}
func (m *CreateTweetResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTweetResponse.Size(m)
}
func (m *CreateTweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTweetResponse proto.InternalMessageInfo

func (m *CreateTweetResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateTweetResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListTweetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTweetRequest) Reset()         { *m = ListTweetRequest{} }
func (m *ListTweetRequest) String() string { return proto.CompactTextString(m) }
func (*ListTweetRequest) ProtoMessage()    {}
func (*ListTweetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce250c0854693d1b, []int{2}
}

func (m *ListTweetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTweetRequest.Unmarshal(m, b)
}
func (m *ListTweetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTweetRequest.Marshal(b, m, deterministic)
}
func (m *ListTweetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTweetRequest.Merge(m, src)
}
func (m *ListTweetRequest) XXX_Size() int {
	return xxx_messageInfo_ListTweetRequest.Size(m)
}
func (m *ListTweetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTweetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTweetRequest proto.InternalMessageInfo

type ListTweetResponse struct {
	Tweets               []*ListTweetResponse_Tweet `protobuf:"bytes,1,rep,name=tweets,proto3" json:"tweets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListTweetResponse) Reset()         { *m = ListTweetResponse{} }
func (m *ListTweetResponse) String() string { return proto.CompactTextString(m) }
func (*ListTweetResponse) ProtoMessage()    {}
func (*ListTweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce250c0854693d1b, []int{3}
}

func (m *ListTweetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTweetResponse.Unmarshal(m, b)
}
func (m *ListTweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTweetResponse.Marshal(b, m, deterministic)
}
func (m *ListTweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTweetResponse.Merge(m, src)
}
func (m *ListTweetResponse) XXX_Size() int {
	return xxx_messageInfo_ListTweetResponse.Size(m)
}
func (m *ListTweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTweetResponse proto.InternalMessageInfo

func (m *ListTweetResponse) GetTweets() []*ListTweetResponse_Tweet {
	if m != nil {
		return m.Tweets
	}
	return nil
}

type ListTweetResponse_Tweet struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTweetResponse_Tweet) Reset()         { *m = ListTweetResponse_Tweet{} }
func (m *ListTweetResponse_Tweet) String() string { return proto.CompactTextString(m) }
func (*ListTweetResponse_Tweet) ProtoMessage()    {}
func (*ListTweetResponse_Tweet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce250c0854693d1b, []int{3, 0}
}

func (m *ListTweetResponse_Tweet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTweetResponse_Tweet.Unmarshal(m, b)
}
func (m *ListTweetResponse_Tweet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTweetResponse_Tweet.Marshal(b, m, deterministic)
}
func (m *ListTweetResponse_Tweet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTweetResponse_Tweet.Merge(m, src)
}
func (m *ListTweetResponse_Tweet) XXX_Size() int {
	return xxx_messageInfo_ListTweetResponse_Tweet.Size(m)
}
func (m *ListTweetResponse_Tweet) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTweetResponse_Tweet.DiscardUnknown(m)
}

var xxx_messageInfo_ListTweetResponse_Tweet proto.InternalMessageInfo

func (m *ListTweetResponse_Tweet) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListTweetResponse_Tweet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTweetRequest)(nil), "tweet.CreateTweetRequest")
	proto.RegisterType((*CreateTweetResponse)(nil), "tweet.CreateTweetResponse")
	proto.RegisterType((*ListTweetRequest)(nil), "tweet.ListTweetRequest")
	proto.RegisterType((*ListTweetResponse)(nil), "tweet.ListTweetResponse")
	proto.RegisterType((*ListTweetResponse_Tweet)(nil), "tweet.ListTweetResponse.Tweet")
}

func init() { proto.RegisterFile("pkg/grpc/proto/tweet.proto", fileDescriptor_ce250c0854693d1b) }

var fileDescriptor_ce250c0854693d1b = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc8, 0x4e, 0xd7,
	0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x29, 0x4f, 0x4d, 0x2d,
	0xd1, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0x29, 0xf1, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92,
	0x54, 0x7d, 0x18, 0x03, 0x22, 0x2f, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58,
	0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x55,
	0xb2, 0xe2, 0x12, 0x72, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0x01, 0x99, 0x12, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc2, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xf5, 0xcb, 0x89, 0xbd, 0x88, 0x55, 0x80, 0x51, 0x62, 0x06,
	0x53, 0x10, 0x4c, 0x4a, 0xc9, 0x9e, 0x4b, 0x18, 0x45, 0x6f, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x10, 0x1f, 0x17, 0x53, 0x66, 0x0a, 0x58, 0x1f, 0x4b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x04, 0xc2,
	0x30, 0x26, 0x90, 0x61, 0x08, 0x03, 0x84, 0xb8, 0x04, 0x7c, 0x32, 0x8b, 0x4b, 0x90, 0xad, 0x56,
	0xaa, 0xe3, 0x12, 0x44, 0x12, 0x83, 0x1a, 0x69, 0xc6, 0xc5, 0x06, 0xf6, 0x65, 0xb1, 0x04, 0xa3,
	0x02, 0xb3, 0x06, 0xb7, 0x91, 0x9c, 0x1e, 0x24, 0x04, 0x30, 0x54, 0xea, 0x41, 0x78, 0x50, 0xd5,
	0x52, 0x86, 0x5c, 0xac, 0x60, 0x01, 0xe2, 0xdd, 0x64, 0xb4, 0x93, 0x91, 0x8b, 0x07, 0xac, 0x27,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x82, 0x8b, 0x1b, 0xc9, 0x97, 0x42, 0x92, 0x50,
	0xab, 0x31, 0x43, 0x4d, 0x4a, 0x0a, 0x9b, 0x14, 0xc4, 0x5d, 0x4a, 0x42, 0x4d, 0x97, 0x9f, 0x4c,
	0x66, 0xe2, 0x51, 0x62, 0x87, 0xc4, 0x5d, 0xb1, 0x15, 0xa3, 0x96, 0x90, 0x3f, 0x17, 0x27, 0xdc,
	0x03, 0x42, 0xe2, 0x98, 0x5e, 0x82, 0x98, 0x2a, 0x81, 0xcb, 0xaf, 0x4a, 0xfc, 0x60, 0x33, 0x39,
	0x85, 0x60, 0x66, 0x3a, 0x71, 0x45, 0x71, 0xc0, 0x12, 0x4a, 0x12, 0x1b, 0x38, 0x7e, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x60, 0x70, 0xa3, 0x3b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TweetServiceClient is the client API for TweetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TweetServiceClient interface {
	CreateTweet(ctx context.Context, in *CreateTweetRequest, opts ...grpc.CallOption) (*CreateTweetResponse, error)
	ListTweet(ctx context.Context, in *ListTweetRequest, opts ...grpc.CallOption) (*ListTweetResponse, error)
}

type tweetServiceClient struct {
	cc *grpc.ClientConn
}

func NewTweetServiceClient(cc *grpc.ClientConn) TweetServiceClient {
	return &tweetServiceClient{cc}
}

func (c *tweetServiceClient) CreateTweet(ctx context.Context, in *CreateTweetRequest, opts ...grpc.CallOption) (*CreateTweetResponse, error) {
	out := new(CreateTweetResponse)
	err := c.cc.Invoke(ctx, "/tweet.TweetService/CreateTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) ListTweet(ctx context.Context, in *ListTweetRequest, opts ...grpc.CallOption) (*ListTweetResponse, error) {
	out := new(ListTweetResponse)
	err := c.cc.Invoke(ctx, "/tweet.TweetService/ListTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetServiceServer is the server API for TweetService service.
type TweetServiceServer interface {
	CreateTweet(context.Context, *CreateTweetRequest) (*CreateTweetResponse, error)
	ListTweet(context.Context, *ListTweetRequest) (*ListTweetResponse, error)
}

// UnimplementedTweetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTweetServiceServer struct {
}

func (*UnimplementedTweetServiceServer) CreateTweet(ctx context.Context, req *CreateTweetRequest) (*CreateTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTweet not implemented")
}
func (*UnimplementedTweetServiceServer) ListTweet(ctx context.Context, req *ListTweetRequest) (*ListTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTweet not implemented")
}

func RegisterTweetServiceServer(s *grpc.Server, srv TweetServiceServer) {
	s.RegisterService(&_TweetService_serviceDesc, srv)
}

func _TweetService_CreateTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).CreateTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tweet.TweetService/CreateTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).CreateTweet(ctx, req.(*CreateTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_ListTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).ListTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tweet.TweetService/ListTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).ListTweet(ctx, req.(*ListTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TweetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tweet.TweetService",
	HandlerType: (*TweetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTweet",
			Handler:    _TweetService_CreateTweet_Handler,
		},
		{
			MethodName: "ListTweet",
			Handler:    _TweetService_ListTweet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/proto/tweet.proto",
}