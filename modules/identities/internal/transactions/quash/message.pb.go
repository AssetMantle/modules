// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/identities/internal/transactions/quash/message.proto

package quash

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/types/base"
	github_com_persistenceOne_persistenceSDK_schema_types_base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type Message struct {
	From       github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required From missing"`
	FromID     base.ID                                                               `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d" valid:"required~required FromID missing"`
	IdentityID base.ID                                                               `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3" json:"identity_i_d" valid:"required~required IdentityID missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0477d664af2a02, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type TransactionResponse struct {
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0477d664af2a02, []int{1}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "persistence_sdk.modules.identities.internal.transactions.quash.Message")
	proto.RegisterType((*TransactionResponse)(nil), "persistence_sdk.modules.identities.internal.transactions.quash.TransactionResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/identities/internal/transactions/quash/message.proto", fileDescriptor_7a0477d664af2a02)
}

var fileDescriptor_7a0477d664af2a02 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x77, 0x8c, 0xd6, 0x3a, 0x7a, 0x5a, 0x15, 0x4a, 0x0e, 0xbb, 0x21, 0xa0, 0x04, 0x91,
	0x19, 0xa9, 0x9e, 0x7a, 0x10, 0x1a, 0xa2, 0x12, 0xb4, 0x88, 0xad, 0x20, 0x78, 0x09, 0x93, 0x9d,
	0xb7, 0x9b, 0xc1, 0xec, 0x4c, 0x32, 0xef, 0x44, 0xe9, 0x41, 0xcf, 0x82, 0x17, 0xc1, 0x2f, 0xd0,
	0x8b, 0x47, 0xbf, 0x47, 0x8f, 0x3d, 0x8a, 0x87, 0x20, 0xc9, 0xc5, 0x73, 0x3f, 0x81, 0xec, 0xec,
	0xd6, 0x2c, 0xf1, 0x0f, 0x96, 0x9c, 0xf6, 0xdd, 0x61, 0xf9, 0xfd, 0x1e, 0xe6, 0x7d, 0x96, 0x3e,
	0x19, 0x81, 0x45, 0x85, 0x0e, 0x74, 0x02, 0x3d, 0x94, 0xaf, 0x78, 0x66, 0xe4, 0x64, 0x08, 0xc8,
	0x95, 0x04, 0xed, 0x94, 0x53, 0xf9, 0xa8, 0x1d, 0x58, 0x2d, 0x86, 0xdc, 0x59, 0xa1, 0x51, 0x24,
	0x4e, 0x19, 0x8d, 0x7c, 0x3c, 0x11, 0x38, 0xe0, 0x19, 0x20, 0x8a, 0x14, 0xd8, 0xc8, 0x1a, 0x67,
	0xc2, 0xfb, 0x4b, 0x34, 0x56, 0xd2, 0xd8, 0x82, 0xc6, 0x4e, 0x69, 0xac, 0x4a, 0x63, 0x9e, 0x56,
	0xbf, 0x96, 0x9a, 0xd4, 0x78, 0x14, 0xcf, 0xa7, 0x82, 0x5a, 0xbf, 0xb5, 0x9c, 0x11, 0x93, 0x01,
	0x64, 0x82, 0xbb, 0x83, 0x11, 0x20, 0xef, 0x0b, 0x04, 0xae, 0x64, 0xf1, 0x6d, 0xf3, 0x53, 0x8d,
	0x5e, 0xdc, 0x29, 0x32, 0x85, 0x1f, 0x08, 0x3d, 0xbf, 0x6f, 0x4d, 0xb6, 0x41, 0x1a, 0xa4, 0x75,
	0xa9, 0xfd, 0xe6, 0x68, 0x1a, 0x07, 0xdf, 0xa6, 0xf1, 0x83, 0x54, 0xb9, 0xc1, 0xa4, 0xcf, 0x12,
	0x93, 0xf1, 0x0a, 0xf9, 0xa9, 0x86, 0xea, 0xeb, 0x5e, 0xe7, 0xf1, 0xef, 0x1e, 0xb6, 0x9d, 0x24,
	0xdb, 0x52, 0x5a, 0x40, 0x3c, 0x99, 0xc6, 0x37, 0x5f, 0x8b, 0xa1, 0x92, 0x5b, 0x4d, 0x0b, 0xe3,
	0x89, 0xb2, 0x20, 0xdf, 0x9d, 0x0e, 0x8d, 0x87, 0xd6, 0x64, 0x8d, 0x4c, 0x21, 0x2a, 0x9d, 0x36,
	0x77, 0x7d, 0x88, 0x70, 0x4c, 0xd7, 0xf3, 0x67, 0x4f, 0xf5, 0xe4, 0xc6, 0xb9, 0x06, 0x69, 0x5d,
	0xde, 0xbc, 0xc1, 0x96, 0xaf, 0xab, 0x10, 0x32, 0x2f, 0x64, 0x5e, 0xd8, 0xed, 0xb4, 0xef, 0xe4,
	0xb9, 0x4f, 0xa6, 0x71, 0xeb, 0x9f, 0xba, 0x6e, 0x67, 0x21, 0x5c, 0xdb, 0xf7, 0x07, 0xe1, 0x5b,
	0x7a, 0xa5, 0xbc, 0xf8, 0x03, 0xaf, 0xad, 0x9d, 0x45, 0x7b, 0xaf, 0xd4, 0xde, 0xfe, 0xab, 0xb6,
	0x5b, 0xb2, 0xab, 0x6a, 0xaa, 0x7e, 0x1d, 0x6e, 0xad, 0xbf, 0x3f, 0x8c, 0x83, 0x1f, 0x87, 0x71,
	0xd0, 0xbc, 0x4e, 0xaf, 0x3e, 0x5f, 0x6c, 0x7b, 0x17, 0x70, 0x64, 0x34, 0xc2, 0xe6, 0x17, 0x42,
	0x6b, 0x3b, 0x98, 0x86, 0x9f, 0x09, 0xbd, 0xf0, 0x2c, 0x2f, 0x40, 0xf8, 0x88, 0xad, 0xd6, 0x20,
	0x56, 0xee, 0xbe, 0xbe, 0xb7, 0x2a, 0xe8, 0x0f, 0x79, 0xdb, 0xe3, 0xa3, 0x59, 0x44, 0x8e, 0x67,
	0x11, 0xf9, 0x3e, 0x8b, 0xc8, 0xc7, 0x79, 0x14, 0x1c, 0xcf, 0xa3, 0xe0, 0xeb, 0x3c, 0x0a, 0x5e,
	0xbe, 0xf8, 0xef, 0x4e, 0x9d, 0xed, 0xff, 0xea, 0xaf, 0xf9, 0x5a, 0xdf, 0xfd, 0x19, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0xf9, 0x1c, 0x5a, 0xa8, 0x03, 0x00, 0x00,
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
	Quash(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Quash(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.identities.internal.transactions.quash.Msg/Quash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Quash(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Quash(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quash not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Quash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Quash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.identities.internal.transactions.quash.Msg/Quash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Quash(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.identities.internal.transactions.quash.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Quash",
			Handler:    _Msg_Quash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/identities/internal/transactions/quash/message.proto",
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.IdentityID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.FromID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.From.Size()
		i -= size
		if _, err := m.From.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.From.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.IdentityID.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *TransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: TransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
