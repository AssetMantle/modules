// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/queries/assets/service.proto

package assets

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("x/assets/internal/queries/assets/service.proto", fileDescriptor_9a172b94948eebbe)
}

var fileDescriptor_9a172b94948eebbe = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xab, 0xd0, 0x4f, 0x2c,
	0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c,
	0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x86, 0x89, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0xea, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x42, 0x44, 0xf5, 0xa0, 0x8a, 0xf4, 0x20, 0x5c, 0x29, 0x99,
	0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92,
	0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x26, 0x29, 0x63, 0x82, 0x96, 0x80, 0xb8, 0x95, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x50, 0x4d, 0x26, 0xc4, 0x6a, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b,
	0x86, 0xba, 0xcf, 0xa8, 0x95, 0x91, 0x8b, 0x3d, 0x18, 0xe2, 0x62, 0xa1, 0x2a, 0x2e, 0x36, 0x8f,
	0xc4, 0xbc, 0x94, 0x9c, 0x54, 0x21, 0x65, 0x3d, 0xac, 0xce, 0xd6, 0x0b, 0x44, 0xb2, 0x56, 0x4a,
	0x05, 0xbf, 0x22, 0x88, 0x35, 0x4a, 0xaa, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x92, 0x17, 0x92, 0xd5,
	0xcf, 0x4d, 0xcc, 0x2b, 0x01, 0x79, 0x19, 0xe2, 0x98, 0x32, 0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43,
	0x28, 0xd7, 0xe9, 0x07, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0x49,
	0x26, 0xe7, 0xe7, 0x62, 0xb7, 0xca, 0x89, 0x07, 0xea, 0xf4, 0x00, 0x90, 0x5f, 0x02, 0x18, 0xa3,
	0xec, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x1d, 0x41, 0x4a, 0x7c,
	0x21, 0x76, 0xe6, 0xe6, 0xa7, 0x94, 0xe6, 0xa4, 0x16, 0xeb, 0x13, 0x0a, 0xa2, 0x45, 0x4c, 0xcc,
	0x8e, 0x81, 0x8e, 0xab, 0x98, 0x44, 0x1d, 0x11, 0x5e, 0x01, 0x59, 0x07, 0xe1, 0x9e, 0x82, 0x89,
	0xc7, 0x40, 0xc5, 0x63, 0x20, 0xdc, 0x47, 0x4c, 0x8a, 0x58, 0xc5, 0x63, 0xdc, 0x03, 0x9c, 0x7c,
	0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f, 0x31, 0x89, 0x43, 0x04, 0xad, 0xac, 0xa0, 0x8a,
	0xac, 0xac, 0x20, 0x02, 0x49, 0x6c, 0xe0, 0x98, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x38, 0x83, 0x83, 0x5b, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assets.queries.assets.Service/Handle", in, out, opts...)
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
		FullMethod: "/assets.queries.assets.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.queries.assets.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/assets/internal/queries/assets/service.proto",
}
