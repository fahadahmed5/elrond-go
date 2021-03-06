// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: esdt.proto

package builtInFunctions

import (
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_data "github.com/ElrondNetwork/elrond-go/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// ESDigitalToken holds the data for a elrond standard digital token transaction
type ESDigitalToken struct {
	Value *math_big.Int `protobuf:"bytes,1,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"value"`
}

func (m *ESDigitalToken) Reset()      { *m = ESDigitalToken{} }
func (*ESDigitalToken) ProtoMessage() {}
func (*ESDigitalToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e413e402abc6a34c, []int{0}
}
func (m *ESDigitalToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ESDigitalToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ESDigitalToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ESDigitalToken.Merge(m, src)
}
func (m *ESDigitalToken) XXX_Size() int {
	return m.Size()
}
func (m *ESDigitalToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ESDigitalToken.DiscardUnknown(m)
}

var xxx_messageInfo_ESDigitalToken proto.InternalMessageInfo

func (m *ESDigitalToken) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ESDigitalToken)(nil), "protoBuiltInFunctions.ESDigitalToken")
}

func init() { proto.RegisterFile("esdt.proto", fileDescriptor_e413e402abc6a34c) }

var fileDescriptor_e413e402abc6a34c = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2d, 0x4e, 0x29,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x05, 0x53, 0x4e, 0xa5, 0x99, 0x39, 0x25, 0x9e,
	0x79, 0x6e, 0xa5, 0x79, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0x65, 0x49, 0xa5,
	0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x4c, 0x51, 0x2a, 0xe3, 0xe2, 0x73, 0x0d, 0x76, 0xc9,
	0x4c, 0xcf, 0x2c, 0x49, 0xcc, 0x09, 0xc9, 0xcf, 0x4e, 0xcd, 0x13, 0x4a, 0xe1, 0x62, 0x0d, 0x4b,
	0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x71, 0xf2, 0x7b, 0x75, 0x4f, 0x9e, 0xb5,
	0x0c, 0x24, 0xb0, 0xea, 0xbe, 0xbc, 0x63, 0x6e, 0x62, 0x49, 0x86, 0x7e, 0x52, 0x66, 0xba, 0x9e,
	0x67, 0x5e, 0x89, 0x35, 0x92, 0x4d, 0xae, 0x39, 0x45, 0xf9, 0x79, 0x29, 0x7e, 0xa9, 0x25, 0xe5,
	0xf9, 0x45, 0xd9, 0xfa, 0xa9, 0x60, 0x9e, 0x6e, 0x7a, 0xbe, 0x7e, 0x4a, 0x62, 0x49, 0xa2, 0x9e,
	0x53, 0x66, 0xba, 0x67, 0x5e, 0x89, 0x73, 0x62, 0x71, 0x49, 0x6a, 0x51, 0x10, 0xc4, 0x70, 0x27,
	0xaf, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91,
	0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x09, 0x24, 0xa1,
	0x79, 0x39, 0x89, 0x0d, 0xec, 0x15, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x85, 0x25,
	0x21, 0x1e, 0x01, 0x00, 0x00,
}

func (this *ESDigitalToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ESDigitalToken)
	if !ok {
		that2, ok := that.(ESDigitalToken)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	return true
}
func (this *ESDigitalToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&builtInFunctions.ESDigitalToken{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEsdt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ESDigitalToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ESDigitalToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ESDigitalToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEsdt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEsdt(dAtA []byte, offset int, v uint64) int {
	offset -= sovEsdt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ESDigitalToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovEsdt(uint64(l))
	}
	return n
}

func sovEsdt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEsdt(x uint64) (n int) {
	return sovEsdt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ESDigitalToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ESDigitalToken{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEsdt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ESDigitalToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEsdt
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
			return fmt.Errorf("proto: ESDigitalToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ESDigitalToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEsdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEsdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEsdt
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
func skipEsdt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEsdt
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
					return 0, ErrIntOverflowEsdt
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
					return 0, ErrIntOverflowEsdt
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
				return 0, ErrInvalidLengthEsdt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEsdt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEsdt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEsdt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEsdt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEsdt = fmt.Errorf("proto: unexpected end of group")
)
