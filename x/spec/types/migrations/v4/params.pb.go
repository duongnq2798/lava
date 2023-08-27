// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/spec/params.proto

package v4

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	GeolocationCount int32  `protobuf:"varint,1,opt,name=geolocationCount,proto3" json:"geolocationCount,omitempty" yaml:"geolocation_count"`
	MaxCU            uint64 `protobuf:"varint,2,opt,name=maxCU,proto3" json:"maxCU,omitempty" yaml:"max_compute_units"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2f67d2fb6b4afa, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetGeolocationCount() int32 {
	if m != nil {
		return m.GeolocationCount
	}
	return 0
}

func (m *Params) GetMaxCU() uint64 {
	if m != nil {
		return m.MaxCU
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.spec.ParamsV4")
}

func init() { proto.RegisterFile("lavanet/lava/spec/params.proto", fileDescriptor_ad2f67d2fb6b4afa) }

var fileDescriptor_ad2f67d2fb6b4afa = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xc5, 0x05, 0xa9, 0xc9, 0xfa, 0x05, 0x89, 0x45, 0x89,
	0xb9, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0x50, 0x79, 0x3d, 0x10, 0xad, 0x07,
	0x92, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83, 0x58, 0x10, 0x85, 0x4a, 0x3d,
	0x8c, 0x5c, 0x6c, 0x01, 0x60, 0x9d, 0x42, 0x1e, 0x5c, 0x02, 0xe9, 0xa9, 0xf9, 0x39, 0xf9, 0xc9,
	0x89, 0x25, 0x99, 0xf9, 0x79, 0xce, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac,
	0x4e, 0x32, 0x9f, 0xee, 0xc9, 0x4b, 0x54, 0x26, 0xe6, 0xe6, 0x58, 0x29, 0x21, 0xa9, 0x88, 0x4f,
	0x06, 0x29, 0x51, 0x0a, 0xc2, 0xd0, 0x25, 0x64, 0xc4, 0xc5, 0x9a, 0x9b, 0x58, 0xe1, 0x1c, 0x2a,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x82, 0xac, 0x3d, 0x37, 0xb1, 0x22, 0x3e, 0x39, 0x3f, 0xb7, 0xa0,
	0xb4, 0x24, 0x35, 0xbe, 0x34, 0x2f, 0xb3, 0xa4, 0x58, 0x29, 0x08, 0xa2, 0xd4, 0x8a, 0x65, 0xc6,
	0x02, 0x79, 0x06, 0x27, 0xbb, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52,
	0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xf1, 0x7c, 0x05, 0xc4,
	0xfb, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x5f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x59, 0xea, 0x70, 0x20, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxCU != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxCU))
		i--
		dAtA[i] = 0x10
	}
	if m.GeolocationCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GeolocationCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GeolocationCount != 0 {
		n += 1 + sovParams(uint64(m.GeolocationCount))
	}
	if m.MaxCU != 0 {
		n += 1 + sovParams(uint64(m.MaxCU))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeolocationCount", wireType)
			}
			m.GeolocationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GeolocationCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCU", wireType)
			}
			m.MaxCU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCU |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
