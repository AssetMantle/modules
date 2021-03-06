// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/orders/internal/key/orderID.proto

package key

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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

type OrderID struct {
	ClassificationID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classification_i_d"`
	MakerOwnableID   github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,2,opt,name=maker_ownable_i_d,json=makerOwnableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"maker_ownable_i_d"`
	TakerOwnableID   github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,3,opt,name=taker_ownable_i_d,json=takerOwnableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"taker_ownable_i_d"`
	RateID           github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,4,opt,name=rate_i_d,json=rateID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"rate_i_d"`
	CreationID       github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,5,opt,name=creation_i_d,json=creationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"creation_i_d"`
	MakerID          github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,6,opt,name=maker_i_d,json=makerID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"maker_i_d"`
	HashID           github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,7,opt,name=hash_i_d,json=hashID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"hash_i_d"`
}

func (m *OrderID) Reset()      { *m = OrderID{} }
func (*OrderID) ProtoMessage() {}
func (*OrderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_db3a535cdb38bb7b, []int{0}
}
func (m *OrderID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderID.Merge(m, src)
}
func (m *OrderID) XXX_Size() int {
	return m.Size()
}
func (m *OrderID) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderID.DiscardUnknown(m)
}

var xxx_messageInfo_OrderID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OrderID)(nil), "persistence_sdk.modules.orders.internal.key.OrderID")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/orders/internal/key/orderID.proto", fileDescriptor_db3a535cdb38bb7b)
}

var fileDescriptor_db3a535cdb38bb7b = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x3f, 0x4b, 0xf3, 0x40,
	0x18, 0xc0, 0x93, 0xf7, 0xed, 0x9f, 0xb7, 0xc7, 0x8b, 0x68, 0x70, 0x10, 0x87, 0x54, 0x9c, 0x04,
	0xe1, 0x6e, 0x70, 0x51, 0x27, 0x29, 0xb7, 0x04, 0x87, 0x82, 0x6e, 0x45, 0x0c, 0x97, 0xe4, 0xb1,
	0x39, 0x92, 0xdc, 0x95, 0xbb, 0x2b, 0xd2, 0x6f, 0xe1, 0xe8, 0xe8, 0xc7, 0xe9, 0x58, 0x70, 0x11,
	0x87, 0x22, 0xed, 0x17, 0x91, 0xe4, 0x2c, 0xb6, 0x0a, 0xe2, 0x90, 0x2d, 0x39, 0xf8, 0xfd, 0x7e,
	0xf0, 0x3c, 0x3c, 0xe8, 0x6c, 0x04, 0x4a, 0x73, 0x6d, 0x40, 0xc4, 0x10, 0xea, 0x24, 0x23, 0x85,
	0x4c, 0xc6, 0x39, 0x68, 0x22, 0x55, 0x02, 0x4a, 0x13, 0x2e, 0x0c, 0x28, 0xc1, 0x72, 0x92, 0xc1,
	0xc4, 0xbe, 0x05, 0x14, 0x8f, 0x94, 0x34, 0xd2, 0x3b, 0xfe, 0x82, 0xe2, 0x0f, 0x14, 0x5b, 0x14,
	0xaf, 0x50, 0x9c, 0xc1, 0x64, 0x7f, 0x77, 0x28, 0x87, 0xb2, 0xe2, 0x48, 0xf9, 0x65, 0x15, 0x87,
	0xcf, 0x4d, 0xd4, 0xee, 0x5b, 0xa9, 0x27, 0x90, 0x17, 0xe7, 0x4c, 0x6b, 0x7e, 0xc7, 0x63, 0x66,
	0xb8, 0x14, 0x21, 0x0f, 0x93, 0x3d, 0xf7, 0xc0, 0x3d, 0xea, 0xf4, 0x2e, 0xa6, 0xf3, 0xae, 0xf3,
	0x3a, 0xef, 0x9e, 0x0e, 0xb9, 0x49, 0xc7, 0x11, 0x8e, 0x65, 0x41, 0xd6, 0xea, 0x7d, 0x01, 0xeb,
	0xbf, 0xd7, 0xf4, 0x92, 0xe8, 0x38, 0x85, 0x82, 0x11, 0x33, 0x19, 0x81, 0xc6, 0x01, 0xbd, 0xda,
	0xde, 0x74, 0x07, 0xd4, 0xcb, 0xd0, 0x4e, 0xc1, 0x32, 0x50, 0xa1, 0xbc, 0x17, 0x2c, 0xca, 0xa1,
	0xca, 0xfd, 0xa9, 0x29, 0xb7, 0x55, 0xa9, 0xfb, 0xd6, 0x6c, 0x63, 0xe6, 0x5b, 0xec, 0x6f, 0x5d,
	0x31, 0xb3, 0x19, 0x1b, 0xa0, 0x7f, 0x8a, 0x19, 0xdb, 0x68, 0xd4, 0xd4, 0x68, 0x95, 0xc6, 0x80,
	0x7a, 0x11, 0xfa, 0x1f, 0x2b, 0xf8, 0xdc, 0x4f, 0xb3, 0x26, 0x3f, 0x5a, 0x59, 0x03, 0xea, 0xdd,
	0xa0, 0x8e, 0xdd, 0x4c, 0x19, 0x68, 0xd5, 0x14, 0x68, 0x57, 0x4a, 0x3b, 0x9d, 0x94, 0xe9, 0xb4,
	0x92, 0xb7, 0xeb, 0x9a, 0x4e, 0x69, 0x0c, 0xe8, 0x79, 0xe3, 0xf1, 0xa9, 0xeb, 0xf4, 0x6e, 0xa7,
	0x0b, 0xdf, 0x9d, 0x2d, 0x7c, 0xf7, 0x6d, 0xe1, 0xbb, 0x0f, 0x4b, 0xdf, 0x99, 0x2d, 0x7d, 0xe7,
	0x65, 0xe9, 0x3b, 0x03, 0xfa, 0xeb, 0xc2, 0x0f, 0x67, 0x18, 0xb5, 0xaa, 0xe3, 0x39, 0x79, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0xe6, 0x46, 0x64, 0xbc, 0x03, 0x00, 0x00,
}

func (m *OrderID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.HashID.Size()
		i -= size
		if _, err := m.HashID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.MakerID.Size()
		i -= size
		if _, err := m.MakerID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.CreationID.Size()
		i -= size
		if _, err := m.CreationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.RateID.Size()
		i -= size
		if _, err := m.RateID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TakerOwnableID.Size()
		i -= size
		if _, err := m.TakerOwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MakerOwnableID.Size()
		i -= size
		if _, err := m.MakerOwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ClassificationID.Size()
		i -= size
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintOrderID(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClassificationID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	l = m.MakerOwnableID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	l = m.TakerOwnableID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	l = m.RateID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	l = m.CreationID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	l = m.MakerID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	l = m.HashID.Size()
	n += 1 + l + sovOrderID(uint64(l))
	return n
}

func sovOrderID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderID(x uint64) (n int) {
	return sovOrderID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderID
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
			return fmt.Errorf("proto: OrderID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CreationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderID
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
				return ErrInvalidLengthOrderID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrderID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderID
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
func skipOrderID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderID
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
					return 0, ErrIntOverflowOrderID
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
					return 0, ErrIntOverflowOrderID
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
				return 0, ErrInvalidLengthOrderID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderID = fmt.Errorf("proto: unexpected end of group")
)
