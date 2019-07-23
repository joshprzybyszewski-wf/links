// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository.proto

package repository

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

type NewRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{0}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type IDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{1}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Link struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{2}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{3}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewRequest)(nil), "NewRequest")
	proto.RegisterType((*IDRequest)(nil), "IDRequest")
	proto.RegisterType((*Link)(nil), "Link")
	proto.RegisterType((*Nothing)(nil), "Nothing")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor_10d86afa5a89ec9d) }

var fileDescriptor_10d86afa5a89ec9d = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2d, 0xc8,
	0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe3, 0xe2,
	0xf2, 0x4b, 0x2d, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d,
	0xca, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0xa4, 0xb9, 0x38, 0x3d, 0x5d,
	0x60, 0xd2, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x50, 0x59, 0xa6, 0xcc, 0x14, 0x25, 0x3b, 0x2e, 0x16,
	0x9f, 0xcc, 0xbc, 0x6c, 0x74, 0x71, 0x98, 0x31, 0x4c, 0x70, 0x63, 0x84, 0x44, 0xb8, 0x58, 0x93,
	0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x98, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x25, 0x4e, 0x2e,
	0x76, 0xbf, 0xfc, 0x92, 0x8c, 0xcc, 0xbc, 0x74, 0xa3, 0x6c, 0x2e, 0xae, 0x20, 0xb8, 0xdb, 0x84,
	0xa4, 0xb9, 0x98, 0xfd, 0x52, 0xcb, 0x85, 0xb8, 0xf5, 0x10, 0x6e, 0x93, 0x62, 0xd5, 0x03, 0xd9,
	0xa5, 0xc4, 0x20, 0x24, 0xc5, 0xc5, 0xec, 0x9e, 0x5a, 0x22, 0xc4, 0xa5, 0x07, 0x77, 0x18, 0x42,
	0x4e, 0x85, 0x8b, 0xcb, 0x19, 0x64, 0x74, 0x58, 0x66, 0x71, 0x26, 0xaa, 0x12, 0x0e, 0x3d, 0xa8,
	0x55, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xbf, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x3b,
	0x3f, 0x7f, 0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*Link, error)
	Get(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Link, error)
	CountVisit(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type repositoryClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryClient(cc *grpc.ClientConn) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/Repository/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) Get(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/Repository/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) CountVisit(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/Repository/CountVisit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
type RepositoryServer interface {
	New(context.Context, *NewRequest) (*Link, error)
	Get(context.Context, *IDRequest) (*Link, error)
	CountVisit(context.Context, *IDRequest) (*Nothing, error)
}

// UnimplementedRepositoryServer can be embedded to have forward compatible implementations.
type UnimplementedRepositoryServer struct {
}

func (*UnimplementedRepositoryServer) New(ctx context.Context, req *NewRequest) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (*UnimplementedRepositoryServer) Get(ctx context.Context, req *IDRequest) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRepositoryServer) CountVisit(ctx context.Context, req *IDRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountVisit not implemented")
}

func RegisterRepositoryServer(s *grpc.Server, srv RepositoryServer) {
	s.RegisterService(&_Repository_serviceDesc, srv)
}

func _Repository_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Repository/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Repository/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).Get(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_CountVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).CountVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Repository/CountVisit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).CountVisit(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Repository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Repository_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Repository_Get_Handler,
		},
		{
			MethodName: "CountVisit",
			Handler:    _Repository_CountVisit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repository.proto",
}
