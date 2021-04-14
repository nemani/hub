// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/proof.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	types "github.com/sentinel-official/hub/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Proof struct {
	Channel      uint64          `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Subscription uint64          `protobuf:"varint,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Node         string          `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Duration     time.Duration   `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration"`
	Bandwidth    types.Bandwidth `protobuf:"bytes,5,opt,name=bandwidth,proto3" json:"bandwidth"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7e6824be930eaf, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Proof)(nil), "sentinel.session.v1.Proof")
}

func init() { proto.RegisterFile("sentinel/session/v1/proof.proto", fileDescriptor_5c7e6824be930eaf) }

var fileDescriptor_5c7e6824be930eaf = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x4f, 0xe3, 0x30,
	0x18, 0xc6, 0xe3, 0xbb, 0xf4, 0xae, 0xcd, 0xdd, 0x94, 0xbb, 0x21, 0x57, 0x9d, 0x9c, 0xd2, 0xa9,
	0x0b, 0x36, 0x81, 0x8d, 0x05, 0x14, 0x31, 0x31, 0xa1, 0x8c, 0x6c, 0x71, 0xe2, 0x24, 0x96, 0x52,
	0xbf, 0x51, 0x9c, 0x04, 0xf8, 0x06, 0x8c, 0x8c, 0x8c, 0xfd, 0x38, 0x1d, 0x3b, 0x32, 0x01, 0x6a,
	0x25, 0xc4, 0xc7, 0x40, 0x71, 0x9a, 0x20, 0xb6, 0xf7, 0xcf, 0xf3, 0x3c, 0xfe, 0xd9, 0xb6, 0x5c,
	0xc5, 0x65, 0x25, 0x24, 0xcf, 0xa9, 0xe2, 0x4a, 0x09, 0x90, 0xb4, 0xf1, 0x68, 0x51, 0x02, 0x24,
	0xa4, 0x28, 0xa1, 0x02, 0xfb, 0x4f, 0x2f, 0x20, 0x7b, 0x01, 0x69, 0xbc, 0x29, 0x8e, 0x40, 0x2d,
	0x41, 0x51, 0x16, 0x2a, 0x4e, 0x1b, 0x8f, 0xf1, 0x2a, 0xf4, 0x68, 0x04, 0x42, 0x76, 0xa6, 0xe9,
	0xdf, 0x14, 0x52, 0xd0, 0x25, 0x6d, 0xab, 0xfd, 0x14, 0xa7, 0x00, 0x69, 0xce, 0xa9, 0xee, 0x58,
	0x9d, 0xd0, 0xb8, 0x2e, 0xc3, 0xaa, 0x8d, 0xec, 0xf6, 0x07, 0x03, 0x4b, 0x75, 0x57, 0x70, 0xd5,
	0x92, 0xb0, 0x50, 0xc6, 0x37, 0x22, 0xae, 0xb2, 0x4e, 0x32, 0x7f, 0x43, 0xd6, 0xe8, 0xaa, 0xa5,
	0xb3, 0x1d, 0xeb, 0x67, 0x94, 0x85, 0x52, 0xf2, 0xdc, 0x41, 0x33, 0xb4, 0x30, 0x83, 0xbe, 0xb5,
	0xe7, 0xd6, 0x6f, 0x55, 0x33, 0x15, 0x95, 0xa2, 0x68, 0xc3, 0x9d, 0x6f, 0x7a, 0xfd, 0x65, 0x66,
	0xdb, 0x96, 0x29, 0x21, 0xe6, 0xce, 0xf7, 0x19, 0x5a, 0x4c, 0x02, 0x5d, 0xdb, 0x67, 0xd6, 0xb8,
	0x07, 0x72, 0xcc, 0x19, 0x5a, 0xfc, 0x3a, 0xfe, 0x47, 0x3a, 0x62, 0xd2, 0x13, 0x93, 0x8b, 0xbd,
	0xc0, 0x1f, 0xaf, 0x9f, 0x5d, 0xe3, 0xf1, 0xc5, 0x45, 0xc1, 0x60, 0xb2, 0xcf, 0xad, 0xc9, 0xc0,
	0xeb, 0x8c, 0x74, 0xc2, 0x7f, 0x32, 0x3c, 0x9f, 0xbe, 0x13, 0x69, 0x3c, 0xe2, 0xf7, 0x1a, 0xdf,
	0x6c, 0x43, 0x82, 0x4f, 0xd3, 0xe9, 0xf8, 0x7e, 0xe5, 0x1a, 0xef, 0x2b, 0xd7, 0xf0, 0x2f, 0xd7,
	0x5b, 0x8c, 0x36, 0x5b, 0x8c, 0x5e, 0xb7, 0x18, 0x3d, 0xec, 0xb0, 0xb1, 0xd9, 0x61, 0xe3, 0x69,
	0x87, 0x8d, 0xeb, 0xa3, 0x54, 0x54, 0x59, 0xcd, 0x48, 0x04, 0x4b, 0xda, 0x87, 0x1f, 0x42, 0x92,
	0x88, 0x48, 0x84, 0x39, 0xcd, 0x6a, 0x46, 0x6f, 0x87, 0xbf, 0xd4, 0x67, 0xb2, 0x1f, 0x1a, 0xff,
	0xe4, 0x23, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xeb, 0x3c, 0xb0, 0xec, 0x01, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProof(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProof(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintProof(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Subscription != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x10
	}
	if m.Channel != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Channel))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Channel != 0 {
		n += 1 + sovProof(uint64(m.Channel))
	}
	if m.Subscription != 0 {
		n += 1 + sovProof(uint64(m.Subscription))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovProof(uint64(l))
	l = m.Bandwidth.Size()
	n += 1 + l + sovProof(uint64(l))
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			m.Channel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Channel |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)