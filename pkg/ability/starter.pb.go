// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: starter.proto

package ability

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

type Starter struct {
	EntityID             github_com_elojah_game_01_pkg_ulid.ID   `protobuf:"bytes,1,opt,name=EntityID,json=entityID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"EntityID"`
	AbilityIDs           []github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,2,rep,name=AbilityIDs,json=abilityIDs,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"AbilityIDs"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *Starter) Reset()      { *m = Starter{} }
func (*Starter) ProtoMessage() {}
func (*Starter) Descriptor() ([]byte, []int) {
	return fileDescriptor_starter_c9ae51362579aea6, []int{0}
}
func (m *Starter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Starter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Starter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Starter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Starter.Merge(dst, src)
}
func (m *Starter) XXX_Size() int {
	return m.Size()
}
func (m *Starter) XXX_DiscardUnknown() {
	xxx_messageInfo_Starter.DiscardUnknown(m)
}

var xxx_messageInfo_Starter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Starter)(nil), "ability.Starter")
}
func (this *Starter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Starter)
	if !ok {
		that2, ok := that.(Starter)
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
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	if len(this.AbilityIDs) != len(that1.AbilityIDs) {
		return false
	}
	for i := range this.AbilityIDs {
		if !this.AbilityIDs[i].Equal(that1.AbilityIDs[i]) {
			return false
		}
	}
	return true
}
func (this *Starter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&ability.Starter{")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "AbilityIDs: "+fmt.Sprintf("%#v", this.AbilityIDs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStarter(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Starter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Starter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStarter(dAtA, i, uint64(m.EntityID.Size()))
	n1, err := m.EntityID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.AbilityIDs) > 0 {
		for _, msg := range m.AbilityIDs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintStarter(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintStarter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedStarter(r randyStarter, easy bool) *Starter {
	this := &Starter{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v1
	v2 := r.Intn(10)
	this.AbilityIDs = make([]github_com_elojah_game_01_pkg_ulid.ID, v2)
	for i := 0; i < v2; i++ {
		v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
		this.AbilityIDs[i] = *v3
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyStarter interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneStarter(r randyStarter) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringStarter(r randyStarter) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneStarter(r)
	}
	return string(tmps)
}
func randUnrecognizedStarter(r randyStarter, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldStarter(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldStarter(dAtA []byte, r randyStarter, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateStarter(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateStarter(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateStarter(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateStarter(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateStarter(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateStarter(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateStarter(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Starter) Size() (n int) {
	var l int
	_ = l
	l = m.EntityID.Size()
	n += 1 + l + sovStarter(uint64(l))
	if len(m.AbilityIDs) > 0 {
		for _, e := range m.AbilityIDs {
			l = e.Size()
			n += 1 + l + sovStarter(uint64(l))
		}
	}
	return n
}

func sovStarter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStarter(x uint64) (n int) {
	return sovStarter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Starter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Starter{`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`AbilityIDs:` + fmt.Sprintf("%v", this.AbilityIDs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStarter(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Starter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStarter
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
			return fmt.Errorf("proto: Starter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Starter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStarter
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
				return ErrInvalidLengthStarter
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbilityIDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStarter
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
				return ErrInvalidLengthStarter
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_01_pkg_ulid.ID
			m.AbilityIDs = append(m.AbilityIDs, v)
			if err := m.AbilityIDs[len(m.AbilityIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStarter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStarter
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
func skipStarter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStarter
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
					return 0, ErrIntOverflowStarter
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
					return 0, ErrIntOverflowStarter
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
				return 0, ErrInvalidLengthStarter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStarter
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
				next, err := skipStarter(dAtA[start:])
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
	ErrInvalidLengthStarter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStarter   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("starter.proto", fileDescriptor_starter_c9ae51362579aea6) }

var fileDescriptor_starter_c9ae51362579aea6 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x49, 0x2c,
	0x2a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4c, 0xca, 0xcc, 0xc9,
	0x2c, 0xa9, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1,
	0xa7, 0xb4, 0x98, 0x91, 0x8b, 0x3d, 0x18, 0x62, 0x92, 0x90, 0x27, 0x17, 0x87, 0x6b, 0x5e, 0x49,
	0x66, 0x49, 0xa5, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xee, 0x89, 0x7b, 0xf2,
	0x0c, 0xb7, 0xee, 0xc9, 0xab, 0x22, 0x19, 0x9a, 0x9a, 0x93, 0x9f, 0x95, 0x98, 0xa1, 0x9f, 0x9e,
	0x98, 0x9b, 0x1a, 0x6f, 0x60, 0xa8, 0x5f, 0x90, 0x9d, 0xae, 0x5f, 0x9a, 0x93, 0x99, 0xa2, 0xe7,
	0xe9, 0x12, 0xc4, 0x91, 0x0a, 0xd5, 0x2e, 0xe4, 0xcb, 0xc5, 0xe5, 0x08, 0x71, 0x90, 0xa7, 0x4b,
	0xb1, 0x04, 0x93, 0x02, 0x33, 0xe9, 0x86, 0x71, 0x25, 0xc2, 0x0d, 0x70, 0xb2, 0xb8, 0xf0, 0x50,
	0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0x6c, 0x78,
	0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x8e, 0x47, 0x72, 0x8c, 0x07, 0x1e, 0xc9, 0x31, 0x9e,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31,
	0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0x9b, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0x08, 0xe1, 0x47, 0x2f, 0x01, 0x00, 0x00,
}
