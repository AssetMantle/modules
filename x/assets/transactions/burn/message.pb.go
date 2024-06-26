// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/transactions/burn/message.proto

package burn

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
	proto "github.com/cosmos/gogoproto/proto"
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
	From    string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID  *base.IdentityID `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	AssetID *base.AssetID    `protobuf:"bytes,3,opt,name=asset_i_d,json=assetID,proto3" json:"asset_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb7afe7fb7f3af4, []int{0}
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

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetFromID() *base.IdentityID {
	if m != nil {
		return m.FromID
	}
	return nil
}

func (m *Message) GetAssetID() *base.AssetID {
	if m != nil {
		return m.AssetID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "AssetMantle.modules.x.assets.transactions.burn.Message")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/assets/transactions/burn/message.proto", fileDescriptor_9bb7afe7fb7f3af4)
}

var fileDescriptor_9bb7afe7fb7f3af4 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x6b, 0xe3, 0x30,
	0x14, 0xc7, 0x23, 0xdf, 0x91, 0x5c, 0x7c, 0x37, 0x79, 0x0a, 0x19, 0x4c, 0x38, 0x0e, 0x2e, 0x14,
	0x2a, 0x41, 0x32, 0x55, 0x74, 0xa8, 0x4d, 0xa0, 0x78, 0x30, 0x84, 0x90, 0xc1, 0x14, 0x83, 0x91,
	0x2d, 0x35, 0x11, 0xc4, 0x76, 0xb1, 0x64, 0x48, 0xbf, 0x45, 0x3f, 0x43, 0xc6, 0x7e, 0x92, 0xd2,
	0x29, 0x63, 0xc7, 0xe2, 0x6c, 0xdd, 0xbb, 0x17, 0x59, 0xa6, 0x38, 0x4b, 0x20, 0x93, 0x1f, 0xcf,
	0xef, 0xf7, 0xff, 0xff, 0xf5, 0x24, 0xf3, 0xda, 0x11, 0x82, 0x49, 0x9f, 0x64, 0x72, 0xc3, 0x50,
	0x9a, 0xd3, 0x72, 0xc3, 0x04, 0xda, 0x22, 0xa2, 0xba, 0x02, 0xc9, 0x82, 0x64, 0x82, 0x24, 0x92,
	0xe7, 0x99, 0x40, 0x71, 0x59, 0x64, 0x28, 0x65, 0x42, 0x90, 0x15, 0x83, 0x0f, 0x45, 0x2e, 0x73,
	0x0b, 0xb6, 0x68, 0xd8, 0xd0, 0x70, 0x0b, 0x35, 0x0d, 0xdb, 0x34, 0x54, 0xf4, 0xf0, 0xa2, 0xed,
	0x26, 0x92, 0x35, 0x4b, 0x09, 0xe2, 0x54, 0xa0, 0x98, 0x08, 0xa6, 0x3d, 0x23, 0x4e, 0xb5, 0xf6,
	0xf0, 0xf2, 0xd4, 0x2c, 0xa7, 0x2c, 0x93, 0x5c, 0x3e, 0x7e, 0x8f, 0xff, 0xdd, 0x01, 0xb3, 0xe7,
	0xeb, 0x70, 0x96, 0x65, 0xfe, 0xbc, 0x2f, 0xf2, 0x74, 0x00, 0x46, 0x60, 0xdc, 0x5f, 0xd4, 0xb5,
	0xe5, 0x98, 0xbf, 0xd4, 0x37, 0xe2, 0x11, 0x1d, 0x18, 0x23, 0x30, 0xfe, 0x3d, 0xf9, 0x7f, 0x94,
	0x5e, 0x3b, 0x40, 0x4e, 0x05, 0x54, 0x0e, 0xd0, 0x6b, 0x1c, 0xbc, 0xd9, 0xa2, 0xab, 0x40, 0x6f,
	0x66, 0xdd, 0x98, 0xfd, 0x26, 0x63, 0x44, 0x07, 0x3f, 0x6a, 0x8d, 0x7f, 0x27, 0x35, 0xea, 0x7f,
	0xde, 0x6c, 0xd1, 0x23, 0xba, 0x70, 0x3f, 0x8d, 0x97, 0xca, 0x06, 0xfb, 0xca, 0x06, 0xef, 0x95,
	0x0d, 0x9e, 0x0e, 0x76, 0x67, 0x7f, 0xb0, 0x3b, 0x6f, 0x07, 0xbb, 0x63, 0x4e, 0x92, 0x3c, 0x3d,
	0x73, 0x9d, 0xee, 0x9f, 0xe6, 0xc0, 0x73, 0xb5, 0x81, 0x39, 0xb8, 0xbb, 0x5a, 0x71, 0xb9, 0x2e,
	0x63, 0x98, 0xe4, 0x29, 0x3a, 0xef, 0x5e, 0x77, 0x46, 0xd7, 0xf1, 0x03, 0x67, 0xe9, 0x3e, 0x1b,
	0x47, 0x21, 0xfc, 0x26, 0x44, 0xa0, 0xbb, 0x02, 0x2e, 0xdb, 0x21, 0xdc, 0xb2, 0xc8, 0x5e, 0x8f,
	0x80, 0xb0, 0x01, 0xc2, 0x20, 0xd4, 0x40, 0xd8, 0x06, 0x42, 0x05, 0x54, 0x06, 0x3e, 0x0f, 0x08,
	0x6f, 0xe7, 0xae, 0xcf, 0x24, 0xa1, 0x44, 0x92, 0x0f, 0x63, 0xda, 0x82, 0x31, 0x6e, 0x68, 0x8c,
	0x03, 0xac, 0x55, 0x05, 0xc6, 0x6d, 0x01, 0x8c, 0x95, 0x42, 0xdc, 0xad, 0xdf, 0xc8, 0xf4, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xd1, 0x93, 0xb4, 0x45, 0xee, 0x02, 0x00, 0x00,
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
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FromID != nil {
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
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
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
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.FromID != nil {
		l = m.FromID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.AssetID != nil {
		l = m.AssetID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
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
			m.From = string(dAtA[iNdEx:postIndex])
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
			if m.FromID == nil {
				m.FromID = &base.IdentityID{}
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
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
			if m.AssetID == nil {
				m.AssetID = &base.AssetID{}
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
