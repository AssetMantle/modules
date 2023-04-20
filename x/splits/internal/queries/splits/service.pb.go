// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/queries/splits/service.proto

package splits

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
	proto.RegisterFile("x/splits/internal/queries/splits/service.proto", fileDescriptor_bb2c294bc4f8afa6)
}

var fileDescriptor_bb2c294bc4f8afa6 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xab, 0xd0, 0x2f, 0x2e,
	0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c,
	0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x86, 0x89, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0xea, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x42, 0x44, 0xf5, 0xa0, 0x8a, 0xf4, 0x20, 0x5c, 0x29, 0x99,
	0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92,
	0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x26, 0x29, 0x63, 0x82, 0x96, 0x80, 0xb8, 0x95, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x50, 0x4d, 0x26, 0xc4, 0x6a, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b,
	0x86, 0xba, 0xcf, 0xa8, 0x95, 0x91, 0x8b, 0x3d, 0x18, 0xe2, 0x62, 0xa1, 0x2a, 0x2e, 0x36, 0x8f,
	0xc4, 0xbc, 0x94, 0x9c, 0x54, 0x21, 0x65, 0x3d, 0xac, 0xce, 0xd6, 0x0b, 0x44, 0xb2, 0x56, 0x4a,
	0x05, 0xbf, 0x22, 0x88, 0x35, 0x4a, 0xaa, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x92, 0x17, 0x92, 0xd5,
	0xcf, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0x85, 0x39, 0xa6, 0xcc, 0x30, 0x29, 0xb5, 0x24, 0xd1, 0x10,
	0xca, 0x75, 0xfa, 0xc1, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0x92,
	0xc9, 0xf9, 0xb9, 0xd8, 0xad, 0x72, 0xe2, 0x81, 0x3a, 0x3d, 0x00, 0xe4, 0x97, 0x00, 0xc6, 0x28,
	0xfb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4,
	0x12, 0x5f, 0x88, 0x9d, 0xb9, 0xf9, 0x29, 0xa5, 0x39, 0xa9, 0xc5, 0xfa, 0x84, 0x82, 0x68, 0x11,
	0x13, 0x73, 0x70, 0x60, 0xf0, 0x2a, 0x26, 0xd1, 0x60, 0x84, 0x57, 0x40, 0xd6, 0x41, 0xb8, 0xa7,
	0x60, 0xe2, 0x31, 0x50, 0xf1, 0x18, 0x08, 0xf7, 0x11, 0x93, 0x22, 0x56, 0xf1, 0x18, 0xf7, 0x00,
	0x27, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x57, 0x4c, 0xe2, 0x10, 0x41, 0x2b, 0x2b,
	0xa8, 0x22, 0x2b, 0x2b, 0x88, 0x40, 0x12, 0x1b, 0x38, 0x26, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x6c, 0xae, 0x16, 0x5b, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/splits.queries.splits.Service/Handle", in, out, opts...)
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
		FullMethod: "/splits.queries.splits.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "splits.queries.splits.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/splits/internal/queries/splits/service.proto",
}
