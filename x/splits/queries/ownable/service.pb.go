// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: splits/queries/ownable/service.proto

package ownable

import (
	context "context"
	fmt "fmt"
	math "math"

	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("splits/queries/ownable/service.proto", fileDescriptor_73ae5fa6e0722b8d)
}

var fileDescriptor_73ae5fa6e0722b8d = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x9b, 0xc0, 0xed, 0x85, 0x70, 0x57, 0x59, 0x96, 0xcb, 0x20, 0x22, 0x2e, 0x2a, 0xcc,
	0x50, 0x05, 0xff, 0xcc, 0xae, 0xdd, 0xe8, 0x26, 0xb4, 0xb5, 0x3b, 0x09, 0xc8, 0xa4, 0x3d, 0xd4,
	0x40, 0x3a, 0x93, 0x66, 0x26, 0x55, 0xb7, 0x3e, 0x81, 0xe0, 0x1b, 0x08, 0x6e, 0xdc, 0xf9, 0x16,
	0xe2, 0xaa, 0xe0, 0xc6, 0x8d, 0x20, 0xa9, 0x2b, 0x9f, 0x42, 0xe2, 0x1c, 0xa1, 0x82, 0x62, 0x75,
	0x19, 0xf8, 0xfd, 0xbe, 0xef, 0x3b, 0x43, 0xbc, 0x15, 0x9d, 0x26, 0xb1, 0xd1, 0x6c, 0x9c, 0x43,
	0x16, 0x83, 0x66, 0xea, 0x58, 0x8a, 0x28, 0x01, 0xa6, 0x21, 0x9b, 0xc4, 0x7d, 0xa0, 0x69, 0xa6,
	0x8c, 0xf2, 0xeb, 0x42, 0x6b, 0x30, 0x23, 0x21, 0x4d, 0x02, 0x74, 0xa4, 0x06, 0x79, 0x02, 0x9a,
	0x5a, 0x93, 0xa2, 0x49, 0xd1, 0xac, 0xfd, 0x1f, 0x2a, 0x35, 0x4c, 0x80, 0x89, 0x34, 0x66, 0x42,
	0x4a, 0x65, 0x84, 0x89, 0x95, 0xd4, 0x36, 0xa9, 0x56, 0xff, 0xa2, 0xaf, 0xfc, 0x3e, 0x3d, 0xcc,
	0x60, 0x9c, 0x83, 0x36, 0xc8, 0xae, 0x7d, 0xc3, 0xea, 0x54, 0x49, 0x8d, 0x13, 0xd7, 0x6f, 0x1c,
	0xef, 0x6f, 0xcf, 0x8e, 0xf6, 0xaf, 0x1c, 0xaf, 0xba, 0x27, 0xe4, 0x20, 0x01, 0x7f, 0x9b, 0x2e,
	0x3e, 0x9d, 0x76, 0xcb, 0xe0, 0x7d, 0xbb, 0xa1, 0xb6, 0xf3, 0x0b, 0xd3, 0x2e, 0x5a, 0x5e, 0x3d,
	0xbb, 0x7f, 0xbe, 0x70, 0x97, 0x7c, 0xc2, 0xac, 0xcd, 0xf0, 0x90, 0x49, 0x23, 0x02, 0x23, 0x1a,
	0xef, 0x87, 0xb4, 0x1e, 0xdd, 0xdb, 0x82, 0x38, 0xd3, 0x82, 0x38, 0x4f, 0x05, 0x71, 0xce, 0x67,
	0xa4, 0x32, 0x9d, 0x91, 0xca, 0xc3, 0x8c, 0x54, 0x3c, 0xda, 0x57, 0xa3, 0x1f, 0x0c, 0x68, 0xfd,
	0xc3, 0xdb, 0x3b, 0xe5, 0x63, 0x74, 0x9c, 0x83, 0xad, 0x61, 0x6c, 0x8e, 0xf2, 0xa8, 0x8c, 0x61,
	0xcd, 0x32, 0x26, 0xb0, 0x4b, 0x30, 0x86, 0x9d, 0xb0, 0xcf, 0x1f, 0xf7, 0xd2, 0xfd, 0xd3, 0x0c,
	0x7a, 0xdd, 0xf6, 0xb5, 0x5b, 0x6f, 0xce, 0xb5, 0x07, 0xd8, 0xde, 0xb3, 0xed, 0x5d, 0x6c, 0x6f,
	0x5b, 0xe9, 0xee, 0x03, 0x1c, 0x22, 0x1c, 0x5a, 0x38, 0x44, 0x38, 0x44, 0xb8, 0x70, 0x37, 0x17,
	0x87, 0xc3, 0xdd, 0x4e, 0x2b, 0x00, 0x23, 0x06, 0xc2, 0x88, 0x17, 0x97, 0xce, 0x89, 0x9c, 0xa3,
	0xc9, 0xb9, 0x55, 0x39, 0x47, 0x97, 0x73, 0x94, 0xa3, 0xea, 0xdb, 0xaf, 0xb1, 0xf1, 0x1a, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x91, 0xbb, 0x67, 0xe5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.splits.queries.ownable.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.splits.queries.ownable.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.splits.queries.ownable.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "splits/queries/ownable/service.proto",
}