// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pairing/epoch_payments.proto

package types

import (
	fmt "fmt"
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

type EpochPayments struct {
	Index                      string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ProviderPaymentStorageKeys []string `protobuf:"bytes,2,rep,name=providerPaymentStorageKeys,proto3" json:"providerPaymentStorageKeys,omitempty"`
}

func (m *EpochPayments) Reset()         { *m = EpochPayments{} }
func (m *EpochPayments) String() string { return proto.CompactTextString(m) }
func (*EpochPayments) ProtoMessage()    {}
func (*EpochPayments) Descriptor() ([]byte, []int) {
	return fileDescriptor_747a4457843f82da, []int{0}
}
func (m *EpochPayments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochPayments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochPayments.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochPayments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochPayments.Merge(m, src)
}
func (m *EpochPayments) XXX_Size() int {
	return m.Size()
}
func (m *EpochPayments) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochPayments.DiscardUnknown(m)
}

var xxx_messageInfo_EpochPayments proto.InternalMessageInfo

func (m *EpochPayments) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EpochPayments) GetProviderPaymentStorageKeys() []string {
	if m != nil {
		return m.ProviderPaymentStorageKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*EpochPayments)(nil), "lavanet.lava.pairing.EpochPayments")
}

func init() { proto.RegisterFile("pairing/epoch_payments.proto", fileDescriptor_747a4457843f82da) }

var fileDescriptor_747a4457843f82da = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x48, 0xcc, 0x2c,
	0xca, 0xcc, 0x4b, 0xd7, 0x4f, 0x2d, 0xc8, 0x4f, 0xce, 0x88, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd,
	0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc9, 0x49, 0x2c, 0x4b, 0xcc, 0x4b,
	0x2d, 0xd1, 0x03, 0xd1, 0x7a, 0x50, 0xa5, 0x52, 0x6a, 0x30, 0x3d, 0x05, 0x45, 0xf9, 0x65, 0x99,
	0x29, 0xa9, 0x45, 0x30, 0x6d, 0xf1, 0xc5, 0x25, 0xf9, 0x45, 0x89, 0xe9, 0xa9, 0x10, 0xdd, 0x4a,
	0xa9, 0x5c, 0xbc, 0xae, 0x20, 0x53, 0x03, 0xa0, 0x86, 0x0a, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5,
	0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x76, 0x5c, 0x52, 0x30,
	0x83, 0xa0, 0x2a, 0x83, 0x21, 0xc6, 0x78, 0xa7, 0x56, 0x16, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0x70,
	0x06, 0xe1, 0x51, 0xe1, 0xe4, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0xea, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x9f, 0x80, 0x69,
	0xfd, 0x0a, 0x7d, 0x98, 0x17, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x0e, 0x36, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x1a, 0x8c, 0xd4, 0x0e, 0x01, 0x00, 0x00,
}

func (m *EpochPayments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochPayments) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochPayments) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProviderPaymentStorageKeys) > 0 {
		for iNdEx := len(m.ProviderPaymentStorageKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ProviderPaymentStorageKeys[iNdEx])
			copy(dAtA[i:], m.ProviderPaymentStorageKeys[iNdEx])
			i = encodeVarintEpochPayments(dAtA, i, uint64(len(m.ProviderPaymentStorageKeys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEpochPayments(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpochPayments(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochPayments(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochPayments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEpochPayments(uint64(l))
	}
	if len(m.ProviderPaymentStorageKeys) > 0 {
		for _, s := range m.ProviderPaymentStorageKeys {
			l = len(s)
			n += 1 + l + sovEpochPayments(uint64(l))
		}
	}
	return n
}

func sovEpochPayments(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochPayments(x uint64) (n int) {
	return sovEpochPayments(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochPayments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochPayments
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
			return fmt.Errorf("proto: EpochPayments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochPayments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochPayments
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
				return ErrInvalidLengthEpochPayments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochPayments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderPaymentStorageKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochPayments
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
				return ErrInvalidLengthEpochPayments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochPayments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderPaymentStorageKeys = append(m.ProviderPaymentStorageKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEpochPayments(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochPayments
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
func skipEpochPayments(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochPayments
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
					return 0, ErrIntOverflowEpochPayments
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
					return 0, ErrIntOverflowEpochPayments
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
				return 0, ErrInvalidLengthEpochPayments
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochPayments
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochPayments
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochPayments        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochPayments          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochPayments = fmt.Errorf("proto: unexpected end of group")
)
