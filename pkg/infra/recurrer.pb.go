// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: recurrer.proto

package infra

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_elojah_game_01_pkg_ulid "github.com/elojah/game_01/pkg/ulid"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Recurrer struct {
	TokenID              github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=TokenID,json=tokenID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"TokenID"`
	EntityID             github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,2,opt,name=EntityID,json=entityID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"EntityID"`
	Action               QAction                               `protobuf:"varint,3,opt,name=Action,json=action,proto3,enum=QAction" json:"Action,omitempty"`
	Pool                 github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,4,opt,name=Pool,json=pool,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Pool"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Recurrer) Reset()      { *m = Recurrer{} }
func (*Recurrer) ProtoMessage() {}
func (*Recurrer) Descriptor() ([]byte, []int) {
	return fileDescriptor_recurrer_76e8f700da55ecb6, []int{0}
}
func (m *Recurrer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Recurrer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Recurrer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Recurrer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recurrer.Merge(dst, src)
}
func (m *Recurrer) XXX_Size() int {
	return m.Size()
}
func (m *Recurrer) XXX_DiscardUnknown() {
	xxx_messageInfo_Recurrer.DiscardUnknown(m)
}

var xxx_messageInfo_Recurrer proto.InternalMessageInfo

func (m *Recurrer) GetAction() QAction {
	if m != nil {
		return m.Action
	}
	return Open
}

func init() {
	proto.RegisterType((*Recurrer)(nil), "Recurrer")
}
func (this *Recurrer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Recurrer)
	if !ok {
		that2, ok := that.(Recurrer)
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
	if !this.TokenID.Equal(that1.TokenID) {
		return false
	}
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	if !this.Pool.Equal(that1.Pool) {
		return false
	}
	return true
}
func (this *Recurrer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&infra.Recurrer{")
	s = append(s, "TokenID: "+fmt.Sprintf("%#v", this.TokenID)+",\n")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	s = append(s, "Pool: "+fmt.Sprintf("%#v", this.Pool)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRecurrer(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Recurrer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Recurrer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintRecurrer(dAtA, i, uint64(m.TokenID.Size()))
	n1, err := m.TokenID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintRecurrer(dAtA, i, uint64(m.EntityID.Size()))
	n2, err := m.EntityID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Action != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRecurrer(dAtA, i, uint64(m.Action))
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintRecurrer(dAtA, i, uint64(m.Pool.Size()))
	n3, err := m.Pool.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func encodeVarintRecurrer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedRecurrer(r randyRecurrer, easy bool) *Recurrer {
	this := &Recurrer{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.TokenID = *v1
	v2 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v2
	this.Action = QAction([]int32{0, 1}[r.Intn(2)])
	v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Pool = *v3
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyRecurrer interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRecurrer(r randyRecurrer) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringRecurrer(r randyRecurrer) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneRecurrer(r)
	}
	return string(tmps)
}
func randUnrecognizedRecurrer(r randyRecurrer, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldRecurrer(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldRecurrer(dAtA []byte, r randyRecurrer, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateRecurrer(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateRecurrer(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateRecurrer(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateRecurrer(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateRecurrer(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateRecurrer(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateRecurrer(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Recurrer) Size() (n int) {
	var l int
	_ = l
	l = m.TokenID.Size()
	n += 1 + l + sovRecurrer(uint64(l))
	l = m.EntityID.Size()
	n += 1 + l + sovRecurrer(uint64(l))
	if m.Action != 0 {
		n += 1 + sovRecurrer(uint64(m.Action))
	}
	l = m.Pool.Size()
	n += 1 + l + sovRecurrer(uint64(l))
	return n
}

func sovRecurrer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRecurrer(x uint64) (n int) {
	return sovRecurrer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Recurrer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Recurrer{`,
		`TokenID:` + fmt.Sprintf("%v", this.TokenID) + `,`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`Pool:` + fmt.Sprintf("%v", this.Pool) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRecurrer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Recurrer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecurrer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Recurrer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Recurrer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecurrer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecurrer
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecurrer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecurrer
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecurrer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= (QAction(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecurrer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecurrer
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecurrer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecurrer
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
func skipRecurrer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecurrer
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
					return 0, ErrIntOverflowRecurrer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecurrer
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRecurrer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRecurrer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRecurrer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRecurrer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecurrer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("recurrer.proto", fileDescriptor_recurrer_76e8f700da55ecb6) }

var fileDescriptor_recurrer_76e8f700da55ecb6 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0x2e,
	0x2d, 0x2a, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x0b, 0x27, 0x95,
	0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0x55, 0xce, 0x57, 0x18, 0x9f, 0x98, 0x5c, 0x92, 0x99,
	0x9f, 0x07, 0xe1, 0x2b, 0x35, 0x33, 0x71, 0x71, 0x04, 0x41, 0x4d, 0x14, 0x72, 0xe7, 0x62, 0x0f,
	0xc9, 0xcf, 0x4e, 0xcd, 0xf3, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x71, 0xd2, 0x3d, 0x71,
	0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0x55, 0x24, 0x4b, 0x52, 0x73, 0xf2, 0xb3, 0x12, 0x33, 0xf4,
	0xd3, 0x13, 0x73, 0x53, 0xe3, 0x0d, 0x0c, 0xf5, 0x0b, 0xb2, 0xd3, 0xf5, 0x4b, 0x73, 0x32, 0x53,
	0xf4, 0x3c, 0x5d, 0x82, 0xd8, 0x4b, 0x20, 0xba, 0x85, 0x3c, 0xb9, 0x38, 0x5c, 0xf3, 0x4a, 0x32,
	0x4b, 0x2a, 0x3d, 0x5d, 0x24, 0x98, 0xc8, 0x31, 0x89, 0x23, 0x15, 0xaa, 0x5d, 0x48, 0x81, 0x8b,
	0xcd, 0x11, 0xec, 0x60, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x0e, 0xbd, 0x40, 0x08, 0x3f,
	0x88, 0x0d, 0xe2, 0x11, 0x21, 0x47, 0x2e, 0x96, 0x80, 0xfc, 0xfc, 0x1c, 0x09, 0x16, 0x72, 0x2c,
	0x62, 0x29, 0xc8, 0xcf, 0xcf, 0x71, 0xb2, 0xbf, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86,
	0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c,
	0xe3, 0x8e, 0x47, 0x72, 0x8c, 0x07, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0x43, 0x14, 0x6b, 0x66, 0x5e, 0x5a, 0x51, 0x62, 0x12, 0x1b, 0x38, 0x34, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x66, 0xec, 0xe0, 0x6b, 0x9e, 0x01, 0x00, 0x00,
}