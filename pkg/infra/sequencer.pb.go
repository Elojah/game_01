// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sequencer.proto

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

type Sequencer struct {
	ID                   github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,json=iD,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"ID"`
	Action               QAction                               `protobuf:"varint,2,opt,name=Action,json=action,proto3,enum=infra.QAction" json:"Action,omitempty"`
	Pool                 github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,3,opt,name=Pool,json=pool,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Pool"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Sequencer) Reset()      { *m = Sequencer{} }
func (*Sequencer) ProtoMessage() {}
func (*Sequencer) Descriptor() ([]byte, []int) {
	return fileDescriptor_sequencer_e8158d0cd4b6e4b8, []int{0}
}
func (m *Sequencer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequencer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequencer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Sequencer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequencer.Merge(dst, src)
}
func (m *Sequencer) XXX_Size() int {
	return m.Size()
}
func (m *Sequencer) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequencer.DiscardUnknown(m)
}

var xxx_messageInfo_Sequencer proto.InternalMessageInfo

func (m *Sequencer) GetAction() QAction {
	if m != nil {
		return m.Action
	}
	return Open
}

func init() {
	proto.RegisterType((*Sequencer)(nil), "infra.Sequencer")
}
func (this *Sequencer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Sequencer)
	if !ok {
		that2, ok := that.(Sequencer)
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
	if !this.ID.Equal(that1.ID) {
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
func (this *Sequencer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&infra.Sequencer{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	s = append(s, "Pool: "+fmt.Sprintf("%#v", this.Pool)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSequencer(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Sequencer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequencer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintSequencer(dAtA, i, uint64(m.ID.Size()))
	n1, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Action != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSequencer(dAtA, i, uint64(m.Action))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSequencer(dAtA, i, uint64(m.Pool.Size()))
	n2, err := m.Pool.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func encodeVarintSequencer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedSequencer(r randySequencer, easy bool) *Sequencer {
	this := &Sequencer{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	this.Action = QAction([]int32{0, 1}[r.Intn(2)])
	v2 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Pool = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySequencer interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSequencer(r randySequencer) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSequencer(r randySequencer) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneSequencer(r)
	}
	return string(tmps)
}
func randUnrecognizedSequencer(r randySequencer, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSequencer(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSequencer(dAtA []byte, r randySequencer, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSequencer(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateSequencer(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateSequencer(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSequencer(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSequencer(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSequencer(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSequencer(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Sequencer) Size() (n int) {
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovSequencer(uint64(l))
	if m.Action != 0 {
		n += 1 + sovSequencer(uint64(m.Action))
	}
	l = m.Pool.Size()
	n += 1 + l + sovSequencer(uint64(l))
	return n
}

func sovSequencer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSequencer(x uint64) (n int) {
	return sovSequencer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Sequencer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Sequencer{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`Pool:` + fmt.Sprintf("%v", this.Pool) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSequencer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Sequencer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencer
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
			return fmt.Errorf("proto: Sequencer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequencer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
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
			skippy, err := skipSequencer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSequencer
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
func skipSequencer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
				return 0, ErrInvalidLengthSequencer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSequencer
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
				next, err := skipSequencer(dAtA[start:])
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
	ErrInvalidLengthSequencer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSequencer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sequencer.proto", fileDescriptor_sequencer_e8158d0cd4b6e4b8) }

var fileDescriptor_sequencer_e8158d0cd4b6e4b8 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4e, 0x2d, 0x2c,
	0x4d, 0xcd, 0x4b, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0x4b,
	0x2b, 0x4a, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1,
	0x25, 0xc5, 0x57, 0x18, 0x9f, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x07, 0xe1, 0x2b, 0x6d, 0x65, 0xe4,
	0xe2, 0x0c, 0x86, 0x99, 0x2c, 0x64, 0xcb, 0xc5, 0xe4, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0xe3, 0xa4, 0x7b, 0xe2, 0x9e, 0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0xaa, 0x48, 0x16, 0xa4, 0xe6, 0xe4,
	0x67, 0x25, 0x66, 0xe8, 0xa7, 0x27, 0xe6, 0xa6, 0xc6, 0x1b, 0x18, 0xea, 0x17, 0x64, 0xa7, 0xeb,
	0x97, 0xe6, 0x64, 0xa6, 0xe8, 0x79, 0xba, 0x04, 0x31, 0x65, 0xba, 0x08, 0xa9, 0x71, 0xb1, 0x39,
	0x82, 0x0d, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x33, 0xe2, 0xd3, 0x03, 0xbb, 0x51, 0x2f, 0x10,
	0x22, 0x1a, 0xc4, 0x06, 0xb1, 0x5a, 0xc8, 0x91, 0x8b, 0x25, 0x20, 0x3f, 0x3f, 0x47, 0x82, 0x99,
	0x1c, 0x8b, 0x58, 0x0a, 0xf2, 0xf3, 0x73, 0x9c, 0x2c, 0x2e, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50,
	0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x1f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78,
	0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x81, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31,
	0x4e, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xf6, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x03, 0x62, 0x10, 0x51, 0x01, 0x00, 0x00,
}
