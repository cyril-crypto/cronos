// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cronos/cronos.proto

package types

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

// Params defines the parameters for the cronos module.
type Params struct {
	IbcCroDenom string `protobuf:"bytes,1,opt,name=ibc_cro_denom,json=ibcCroDenom,proto3" json:"ibc_cro_denom,omitempty" yaml:"ibc_cro_denom,omitempty"`
	IbcTimeout  uint64 `protobuf:"varint,2,opt,name=ibc_timeout,json=ibcTimeout,proto3" json:"ibc_timeout,omitempty"`
	// the admin address who can update token mapping
	CronosAdmin          string `protobuf:"bytes,3,opt,name=cronos_admin,json=cronosAdmin,proto3" json:"cronos_admin,omitempty"`
	EnableAutoDeployment bool   `protobuf:"varint,4,opt,name=enable_auto_deployment,json=enableAutoDeployment,proto3" json:"enable_auto_deployment,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bc54992a93db2d2, []int{0}
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

func (m *Params) GetIbcCroDenom() string {
	if m != nil {
		return m.IbcCroDenom
	}
	return ""
}

func (m *Params) GetIbcTimeout() uint64 {
	if m != nil {
		return m.IbcTimeout
	}
	return 0
}

func (m *Params) GetCronosAdmin() string {
	if m != nil {
		return m.CronosAdmin
	}
	return ""
}

func (m *Params) GetEnableAutoDeployment() bool {
	if m != nil {
		return m.EnableAutoDeployment
	}
	return false
}

// TokenMappingChangeProposal defines a proposal to change one token mapping.
type TokenMappingChangeProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Denom       string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	Contract    string `protobuf:"bytes,4,opt,name=contract,proto3" json:"contract,omitempty"`
	// only when updating cronos (source) tokens
	Symbol  string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal uint32 `protobuf:"varint,6,opt,name=decimal,proto3" json:"decimal,omitempty"`
}

func (m *TokenMappingChangeProposal) Reset()      { *m = TokenMappingChangeProposal{} }
func (*TokenMappingChangeProposal) ProtoMessage() {}
func (*TokenMappingChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bc54992a93db2d2, []int{1}
}
func (m *TokenMappingChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMappingChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMappingChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMappingChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMappingChangeProposal.Merge(m, src)
}
func (m *TokenMappingChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *TokenMappingChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMappingChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMappingChangeProposal proto.InternalMessageInfo

// TokenMapping defines a mapping between native denom and contract
type TokenMapping struct {
	Denom    string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *TokenMapping) Reset()         { *m = TokenMapping{} }
func (m *TokenMapping) String() string { return proto.CompactTextString(m) }
func (*TokenMapping) ProtoMessage()    {}
func (*TokenMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bc54992a93db2d2, []int{2}
}
func (m *TokenMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMapping.Merge(m, src)
}
func (m *TokenMapping) XXX_Size() int {
	return m.Size()
}
func (m *TokenMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMapping.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMapping proto.InternalMessageInfo

func (m *TokenMapping) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *TokenMapping) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "cronos.Params")
	proto.RegisterType((*TokenMappingChangeProposal)(nil), "cronos.TokenMappingChangeProposal")
	proto.RegisterType((*TokenMapping)(nil), "cronos.TokenMapping")
}

func init() { proto.RegisterFile("cronos/cronos.proto", fileDescriptor_8bc54992a93db2d2) }

var fileDescriptor_8bc54992a93db2d2 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0xc5, 0xd4, 0x51, 0x6d, 0x3a, 0x59, 0x58, 0x23, 0x10, 0x3c, 0xc8, 0xaa, 0x26, 0x0f,
	0x4d, 0x04, 0xb4, 0x99, 0x3c, 0x35, 0x3f, 0xe8, 0xd6, 0x22, 0x20, 0x32, 0x75, 0x11, 0x28, 0x8a,
	0x90, 0x89, 0x8a, 0xbc, 0x04, 0x45, 0x15, 0xd5, 0x1b, 0x74, 0xec, 0xd8, 0x31, 0xcf, 0xd2, 0xa9,
	0xa3, 0xc7, 0x4e, 0x45, 0x61, 0xbf, 0x41, 0x9f, 0xa0, 0x90, 0x98, 0xa4, 0xc9, 0x90, 0x49, 0xf7,
	0x3b, 0x07, 0xc2, 0x39, 0x97, 0xb8, 0xf8, 0x05, 0xb7, 0xa0, 0xa1, 0xc9, 0xfc, 0xe7, 0xc4, 0x58,
	0x70, 0x40, 0x42, 0x4f, 0xf3, 0x59, 0x05, 0x15, 0x0c, 0x52, 0xd6, 0x4f, 0xde, 0x4d, 0x37, 0x08,
	0x87, 0x57, 0xcc, 0x32, 0xd5, 0x90, 0x77, 0xf8, 0x50, 0x16, 0x3c, 0xe7, 0x16, 0xf2, 0x52, 0x68,
	0x50, 0x11, 0x4a, 0xd0, 0x72, 0x72, 0x9e, 0xfe, 0xfd, 0xbd, 0x88, 0x3b, 0xa6, 0xea, 0x55, 0xfa,
	0xc8, 0x7e, 0x05, 0x4a, 0x3a, 0xa1, 0x8c, 0xeb, 0x52, 0x3a, 0x95, 0x05, 0xbf, 0xb0, 0x70, 0xd9,
	0xeb, 0x64, 0x81, 0x7b, 0xcc, 0x9d, 0x54, 0x02, 0x5a, 0x17, 0xed, 0x25, 0x68, 0x39, 0xa2, 0x58,
	0x16, 0xfc, 0xda, 0x2b, 0xe4, 0x25, 0x3e, 0xf0, 0x9d, 0x72, 0x56, 0x2a, 0xa9, 0xa3, 0x67, 0x7d,
	0x0e, 0x9d, 0x7a, 0xed, 0xac, 0x97, 0xc8, 0x29, 0x3e, 0x12, 0x9a, 0x15, 0xb5, 0xc8, 0x59, 0xeb,
	0xfa, 0x40, 0x53, 0x43, 0xa7, 0x84, 0x76, 0xd1, 0x28, 0x41, 0xcb, 0x31, 0x9d, 0x79, 0xf7, 0xac,
	0x75, 0x70, 0x79, 0xef, 0xad, 0x46, 0xdf, 0x6f, 0x16, 0x41, 0xfa, 0x03, 0xe1, 0xf9, 0x35, 0x7c,
	0x12, 0xfa, 0x3d, 0x33, 0x46, 0xea, 0xea, 0x62, 0xcd, 0x74, 0x25, 0xae, 0x2c, 0x18, 0x68, 0x58,
	0x4d, 0x66, 0x78, 0xdf, 0x49, 0x57, 0x0b, 0xbf, 0x1e, 0xf5, 0x40, 0x12, 0x3c, 0x2d, 0x45, 0xc3,
	0xad, 0x34, 0x4e, 0x82, 0x1e, 0x4a, 0x4f, 0xe8, 0x43, 0xa9, 0xff, 0xcf, 0x3f, 0x8b, 0xaf, 0xeb,
	0x81, 0xcc, 0xf1, 0x98, 0x83, 0x76, 0x96, 0x71, 0x5f, 0x6d, 0x42, 0xef, 0x99, 0x1c, 0xe1, 0xb0,
	0xe9, 0x54, 0x01, 0x75, 0xb4, 0x3f, 0x38, 0xb7, 0x44, 0x22, 0xfc, 0xbc, 0x14, 0x5c, 0x2a, 0x56,
	0x47, 0x61, 0x82, 0x96, 0x87, 0xf4, 0x0e, 0x57, 0xe3, 0xaf, 0x37, 0x8b, 0x60, 0x58, 0xe2, 0x2d,
	0x3e, 0x78, 0xb8, 0xc3, 0xff, 0x74, 0xf4, 0x54, 0xfa, 0xde, 0xe3, 0xf4, 0xf3, 0x0f, 0x3f, 0xb7,
	0x31, 0xda, 0x6c, 0x63, 0xf4, 0x67, 0x1b, 0xa3, 0x6f, 0xbb, 0x38, 0xd8, 0xec, 0xe2, 0xe0, 0xd7,
	0x2e, 0x0e, 0x3e, 0x9e, 0x56, 0xd2, 0xad, 0xdb, 0xe2, 0x84, 0x83, 0xca, 0xb8, 0xed, 0x8c, 0x83,
	0x63, 0xb0, 0xd5, 0x31, 0x5f, 0x33, 0xa9, 0x6f, 0x6f, 0x27, 0xfb, 0xfc, 0x3a, 0xfb, 0x72, 0x37,
	0xbb, 0xce, 0x88, 0xa6, 0x08, 0x87, 0x83, 0x79, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x74,
	0x94, 0x69, 0x65, 0x02, 0x00, 0x00,
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
	if m.EnableAutoDeployment {
		i--
		if m.EnableAutoDeployment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.CronosAdmin) > 0 {
		i -= len(m.CronosAdmin)
		copy(dAtA[i:], m.CronosAdmin)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.CronosAdmin)))
		i--
		dAtA[i] = 0x1a
	}
	if m.IbcTimeout != 0 {
		i = encodeVarintCronos(dAtA, i, uint64(m.IbcTimeout))
		i--
		dAtA[i] = 0x10
	}
	if len(m.IbcCroDenom) > 0 {
		i -= len(m.IbcCroDenom)
		copy(dAtA[i:], m.IbcCroDenom)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.IbcCroDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenMappingChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMappingChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMappingChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimal != 0 {
		i = encodeVarintCronos(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintCronos(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCronos(dAtA []byte, offset int, v uint64) int {
	offset -= sovCronos(v)
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
	l = len(m.IbcCroDenom)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	if m.IbcTimeout != 0 {
		n += 1 + sovCronos(uint64(m.IbcTimeout))
	}
	l = len(m.CronosAdmin)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	if m.EnableAutoDeployment {
		n += 2
	}
	return n
}

func (m *TokenMappingChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovCronos(uint64(m.Decimal))
	}
	return n
}

func (m *TokenMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovCronos(uint64(l))
	}
	return n
}

func sovCronos(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCronos(x uint64) (n int) {
	return sovCronos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCronos
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcCroDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcCroDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTimeout", wireType)
			}
			m.IbcTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CronosAdmin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CronosAdmin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableAutoDeployment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableAutoDeployment = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCronos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCronos
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
func (m *TokenMappingChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCronos
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
			return fmt.Errorf("proto: TokenMappingChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMappingChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCronos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCronos
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
func (m *TokenMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCronos
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
			return fmt.Errorf("proto: TokenMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCronos
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
				return ErrInvalidLengthCronos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCronos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCronos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCronos
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
func skipCronos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCronos
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
					return 0, ErrIntOverflowCronos
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
					return 0, ErrIntOverflowCronos
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
				return 0, ErrInvalidLengthCronos
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCronos
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCronos
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCronos        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCronos          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCronos = fmt.Errorf("proto: unexpected end of group")
)
