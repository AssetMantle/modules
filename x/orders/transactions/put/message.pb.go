// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/put/message.proto

package put

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
	base1 "github.com/AssetMantle/schema/go/types/base"
	proto "github.com/gogo/protobuf/proto"
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
	From              string             `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID            *base.IdentityID   `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	MakerOwnableID    *base.AnyOwnableID `protobuf:"bytes,3,opt,name=maker_ownable_i_d,json=makerOwnableID,proto3" json:"maker_ownable_i_d,omitempty"`
	TakerOwnableID    *base.AnyOwnableID `protobuf:"bytes,4,opt,name=taker_ownable_i_d,json=takerOwnableID,proto3" json:"taker_ownable_i_d,omitempty"`
	MakerOwnableSplit string             `protobuf:"bytes,5,opt,name=maker_ownable_split,json=makerOwnableSplit,proto3" json:"maker_ownable_split,omitempty"`
	TakerOwnableSplit string             `protobuf:"bytes,6,opt,name=taker_ownable_split,json=takerOwnableSplit,proto3" json:"taker_ownable_split,omitempty"`
	ExpiresOn         *base1.Height      `protobuf:"bytes,7,opt,name=expires_on,json=expiresOn,proto3" json:"expires_on,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_46c9f863123241d9, []int{0}
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

func (m *Message) GetMakerOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.MakerOwnableID
	}
	return nil
}

func (m *Message) GetTakerOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.TakerOwnableID
	}
	return nil
}

func (m *Message) GetMakerOwnableSplit() string {
	if m != nil {
		return m.MakerOwnableSplit
	}
	return ""
}

func (m *Message) GetTakerOwnableSplit() string {
	if m != nil {
		return m.TakerOwnableSplit
	}
	return ""
}

func (m *Message) GetExpiresOn() *base1.Height {
	if m != nil {
		return m.ExpiresOn
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.orders.transactions.put.Message")
}

func init() {
	proto.RegisterFile("orders/transactions/put/message.proto", fileDescriptor_46c9f863123241d9)
}

var fileDescriptor_46c9f863123241d9 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x9b, 0xec, 0x6e, 0xd7, 0x1d, 0x45, 0xd8, 0x78, 0x30, 0x14, 0x0c, 0x8b, 0xb0, 0xb8,
	0xb2, 0x30, 0x03, 0x7a, 0x50, 0x72, 0x4b, 0x29, 0x68, 0x0f, 0xa1, 0xa1, 0xf6, 0x24, 0x81, 0x30,
	0xed, 0x8c, 0xed, 0x60, 0x33, 0x13, 0x32, 0x5f, 0x70, 0xfb, 0x2f, 0xfc, 0x0d, 0x1e, 0x3d, 0xfa,
	0x2b, 0xc4, 0xd3, 0x1e, 0x3d, 0x4a, 0x7b, 0xf3, 0x57, 0x48, 0x66, 0xc2, 0x6e, 0xca, 0xd6, 0x85,
	0x9e, 0xda, 0xf0, 0xbd, 0xef, 0xf3, 0xbe, 0xe4, 0xfb, 0x82, 0xce, 0x55, 0xc9, 0x78, 0xa9, 0x09,
	0x94, 0x54, 0x6a, 0x3a, 0x03, 0xa1, 0xa4, 0x26, 0x45, 0x05, 0x24, 0xe7, 0x5a, 0xd3, 0x39, 0xc7,
	0x45, 0xa9, 0x40, 0x79, 0x97, 0x54, 0x6b, 0x0e, 0x39, 0x95, 0xb0, 0xe4, 0x38, 0x57, 0xac, 0x5a,
	0x72, 0x8d, 0xad, 0x15, 0xb7, 0xad, 0xb8, 0xa8, 0xa0, 0xf7, 0x4c, 0x30, 0x4d, 0xa6, 0x54, 0x73,
	0x42, 0xe5, 0x2a, 0x53, 0x5f, 0x24, 0x9d, 0x2e, 0x79, 0x26, 0x98, 0x65, 0xf5, 0x7a, 0x37, 0x63,
	0xc1, 0xb8, 0x04, 0x01, 0xab, 0xdb, 0xd9, 0x53, 0x58, 0x15, 0xbc, 0x99, 0x2e, 0xb8, 0x98, 0x2f,
	0xc0, 0x0e, 0x9e, 0xff, 0x38, 0x40, 0xc7, 0xb1, 0xad, 0xe4, 0x79, 0xe8, 0xf0, 0x53, 0xa9, 0x72,
	0xdf, 0x39, 0x73, 0x2e, 0x4e, 0xc6, 0xe6, 0xbf, 0x17, 0xa1, 0x07, 0xf5, 0x6f, 0x26, 0x32, 0xe6,
	0xbb, 0x67, 0xce, 0xc5, 0xc3, 0x57, 0x2f, 0x70, 0xbb, 0xb3, 0x9e, 0x2d, 0x78, 0x4e, 0xb1, 0x60,
	0x1a, 0xd7, 0x70, 0x3c, 0x6c, 0xa2, 0x87, 0x83, 0x71, 0xb7, 0x36, 0x0e, 0x07, 0xde, 0x04, 0x9d,
	0xe6, 0xf4, 0x33, 0x2f, 0x6f, 0x1b, 0x67, 0xcc, 0x3f, 0x30, 0xac, 0x97, 0xf7, 0xb2, 0x22, 0xb9,
	0x1a, 0x59, 0xcb, 0x70, 0x30, 0x7e, 0x6c, 0x18, 0x37, 0xcf, 0x35, 0x15, 0xee, 0x50, 0x0f, 0xf7,
	0xa6, 0xc2, 0x36, 0x15, 0xa3, 0x27, 0xdb, 0x5d, 0x75, 0xb1, 0x14, 0xe0, 0x1f, 0x99, 0x37, 0x72,
	0xda, 0xae, 0xf0, 0xa1, 0x1e, 0xd4, 0x7a, 0xd8, 0xa1, 0xef, 0x5a, 0x3d, 0xdc, 0xd1, 0x0f, 0x10,
	0xe2, 0x57, 0x85, 0x28, 0xb9, 0xce, 0x94, 0xf4, 0x8f, 0x4d, 0xdd, 0xf3, 0x5d, 0x75, 0xcd, 0xbe,
	0x6c, 0xe1, 0xf7, 0x66, 0x5f, 0xe3, 0x93, 0xc6, 0x38, 0x92, 0xfd, 0x8d, 0xfb, 0x73, 0x1d, 0x38,
	0xd7, 0xeb, 0xc0, 0xf9, 0xb3, 0x0e, 0x9c, 0xaf, 0x9b, 0xa0, 0x73, 0xbd, 0x09, 0x3a, 0xbf, 0x37,
	0x41, 0x07, 0x91, 0x99, 0xca, 0xf1, 0x1e, 0x47, 0xd5, 0x7f, 0xd4, 0x6c, 0x3f, 0xa9, 0xcf, 0x21,
	0x71, 0x3e, 0xbe, 0x9d, 0x0b, 0x58, 0x54, 0x53, 0x3c, 0x53, 0x39, 0x89, 0x6a, 0x4e, 0x6c, 0x38,
	0xa4, 0xe1, 0x90, 0x2b, 0xf2, 0x9f, 0xcb, 0xfe, 0xe6, 0x1e, 0x45, 0xf1, 0x68, 0x92, 0x7c, 0x77,
	0x2f, 0xa3, 0x56, 0x7e, 0xdc, 0xe4, 0x8f, 0x6c, 0xfe, 0xa4, 0x9d, 0x9f, 0x54, 0xf0, 0x6b, 0x4b,
	0x9d, 0x36, 0xea, 0xd4, 0xaa, 0xd3, 0xb6, 0x3a, 0x4d, 0x2a, 0x58, 0xbb, 0x6f, 0xf6, 0x50, 0xa7,
	0xef, 0x92, 0x7e, 0xcc, 0x81, 0x32, 0x0a, 0xf4, 0xaf, 0x4b, 0x5a, 0xce, 0x30, 0x6c, 0xac, 0x61,
	0x68, 0xbd, 0x61, 0xd8, 0x36, 0x87, 0x61, 0x52, 0xc1, 0xb4, 0x6b, 0xbe, 0x90, 0xd7, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x05, 0xca, 0x14, 0xdd, 0xcb, 0x03, 0x00, 0x00,
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
	if m.ExpiresOn != nil {
		{
			size, err := m.ExpiresOn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TakerOwnableSplit) > 0 {
		i -= len(m.TakerOwnableSplit)
		copy(dAtA[i:], m.TakerOwnableSplit)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.TakerOwnableSplit)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MakerOwnableSplit) > 0 {
		i -= len(m.MakerOwnableSplit)
		copy(dAtA[i:], m.MakerOwnableSplit)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.MakerOwnableSplit)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TakerOwnableID != nil {
		{
			size, err := m.TakerOwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.MakerOwnableID != nil {
		{
			size, err := m.MakerOwnableID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.MakerOwnableID != nil {
		l = m.MakerOwnableID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.TakerOwnableID != nil {
		l = m.TakerOwnableID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.MakerOwnableSplit)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.TakerOwnableSplit)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ExpiresOn != nil {
		l = m.ExpiresOn.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableID", wireType)
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
			if m.MakerOwnableID == nil {
				m.MakerOwnableID = &base.AnyOwnableID{}
			}
			if err := m.MakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableID", wireType)
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
			if m.TakerOwnableID == nil {
				m.TakerOwnableID = &base.AnyOwnableID{}
			}
			if err := m.TakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableSplit", wireType)
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
			m.MakerOwnableSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableSplit", wireType)
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
			m.TakerOwnableSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresOn", wireType)
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
			if m.ExpiresOn == nil {
				m.ExpiresOn = &base1.Height{}
			}
			if err := m.ExpiresOn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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