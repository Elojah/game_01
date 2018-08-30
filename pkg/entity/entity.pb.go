// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity.proto

package entity

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import geometry "github.com/elojah/game_01/pkg/geometry"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import github_com_elojah_game_01_pkg_ulid "github.com/elojah/game_01/pkg/ulid"
import time "time"

import strings "strings"
import reflect "reflect"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Cast struct {
	AbilityID            github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=AbilityID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"AbilityID"`
	TS                   time.Time                             `protobuf:"bytes,2,opt,name=TS,stdtime" json:"TS"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Cast) Reset()      { *m = Cast{} }
func (*Cast) ProtoMessage() {}
func (*Cast) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_52795d6082e99de4, []int{0}
}
func (m *Cast) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cast.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Cast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cast.Merge(dst, src)
}
func (m *Cast) XXX_Size() int {
	return m.Size()
}
func (m *Cast) XXX_DiscardUnknown() {
	xxx_messageInfo_Cast.DiscardUnknown(m)
}

var xxx_messageInfo_Cast proto.InternalMessageInfo

func (m *Cast) GetTS() time.Time {
	if m != nil {
		return m.TS
	}
	return time.Time{}
}

type E struct {
	ID                   github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"ID"`
	Type                 github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,2,opt,name=Type,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Type"`
	Name                 string                                `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Dead                 bool                                  `protobuf:"varint,4,opt,name=Dead,proto3" json:"Dead,omitempty"`
	HP                   uint64                                `protobuf:"varint,5,opt,name=HP,proto3" json:"HP,omitempty"`
	MaxHP                uint64                                `protobuf:"varint,6,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	MP                   uint64                                `protobuf:"varint,7,opt,name=MP,proto3" json:"MP,omitempty"`
	MaxMP                uint64                                `protobuf:"varint,8,opt,name=MaxMP,proto3" json:"MaxMP,omitempty"`
	Position             geometry.Position                     `protobuf:"bytes,9,opt,name=Position" json:"Position"`
	Cast                 *Cast                                 `protobuf:"bytes,10,opt,name=Cast" json:"Cast,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *E) Reset()      { *m = E{} }
func (*E) ProtoMessage() {}
func (*E) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_52795d6082e99de4, []int{1}
}
func (m *E) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *E) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_E.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *E) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E.Merge(dst, src)
}
func (m *E) XXX_Size() int {
	return m.Size()
}
func (m *E) XXX_DiscardUnknown() {
	xxx_messageInfo_E.DiscardUnknown(m)
}

var xxx_messageInfo_E proto.InternalMessageInfo

func (m *E) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *E) GetDead() bool {
	if m != nil {
		return m.Dead
	}
	return false
}

func (m *E) GetHP() uint64 {
	if m != nil {
		return m.HP
	}
	return 0
}

func (m *E) GetMaxHP() uint64 {
	if m != nil {
		return m.MaxHP
	}
	return 0
}

func (m *E) GetMP() uint64 {
	if m != nil {
		return m.MP
	}
	return 0
}

func (m *E) GetMaxMP() uint64 {
	if m != nil {
		return m.MaxMP
	}
	return 0
}

func (m *E) GetPosition() geometry.Position {
	if m != nil {
		return m.Position
	}
	return geometry.Position{}
}

func (m *E) GetCast() *Cast {
	if m != nil {
		return m.Cast
	}
	return nil
}

func init() {
	proto.RegisterType((*Cast)(nil), "Cast")
	proto.RegisterType((*E)(nil), "E")
}
func (this *Cast) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cast)
	if !ok {
		that2, ok := that.(Cast)
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
	if !this.AbilityID.Equal(that1.AbilityID) {
		return false
	}
	if !this.TS.Equal(that1.TS) {
		return false
	}
	return true
}
func (this *E) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*E)
	if !ok {
		that2, ok := that.(E)
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
	if !this.Type.Equal(that1.Type) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Dead != that1.Dead {
		return false
	}
	if this.HP != that1.HP {
		return false
	}
	if this.MaxHP != that1.MaxHP {
		return false
	}
	if this.MP != that1.MP {
		return false
	}
	if this.MaxMP != that1.MaxMP {
		return false
	}
	if !this.Position.Equal(&that1.Position) {
		return false
	}
	if !this.Cast.Equal(that1.Cast) {
		return false
	}
	return true
}
func (this *Cast) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&entity.Cast{")
	s = append(s, "AbilityID: "+fmt.Sprintf("%#v", this.AbilityID)+",\n")
	s = append(s, "TS: "+fmt.Sprintf("%#v", this.TS)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *E) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&entity.E{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Dead: "+fmt.Sprintf("%#v", this.Dead)+",\n")
	s = append(s, "HP: "+fmt.Sprintf("%#v", this.HP)+",\n")
	s = append(s, "MaxHP: "+fmt.Sprintf("%#v", this.MaxHP)+",\n")
	s = append(s, "MP: "+fmt.Sprintf("%#v", this.MP)+",\n")
	s = append(s, "MaxMP: "+fmt.Sprintf("%#v", this.MaxMP)+",\n")
	s = append(s, "Position: "+strings.Replace(this.Position.GoString(), `&`, ``, 1)+",\n")
	if this.Cast != nil {
		s = append(s, "Cast: "+fmt.Sprintf("%#v", this.Cast)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEntity(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Cast) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cast) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEntity(dAtA, i, uint64(m.AbilityID.Size()))
	n1, err := m.AbilityID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintEntity(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.TS)))
	n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.TS, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *E) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *E) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEntity(dAtA, i, uint64(m.ID.Size()))
	n3, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintEntity(dAtA, i, uint64(m.Type.Size()))
	n4, err := m.Type.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEntity(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Dead {
		dAtA[i] = 0x20
		i++
		if m.Dead {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.HP != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEntity(dAtA, i, uint64(m.HP))
	}
	if m.MaxHP != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintEntity(dAtA, i, uint64(m.MaxHP))
	}
	if m.MP != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintEntity(dAtA, i, uint64(m.MP))
	}
	if m.MaxMP != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintEntity(dAtA, i, uint64(m.MaxMP))
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintEntity(dAtA, i, uint64(m.Position.Size()))
	n5, err := m.Position.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.Cast != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintEntity(dAtA, i, uint64(m.Cast.Size()))
		n6, err := m.Cast.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func encodeVarintEntity(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedCast(r randyEntity, easy bool) *Cast {
	this := &Cast{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.AbilityID = *v1
	v2 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.TS = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedE(r randyEntity, easy bool) *E {
	this := &E{}
	v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.ID = *v3
	v4 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Type = *v4
	this.Name = string(randStringEntity(r))
	this.Dead = bool(bool(r.Intn(2) == 0))
	this.HP = uint64(uint64(r.Uint32()))
	this.MaxHP = uint64(uint64(r.Uint32()))
	this.MP = uint64(uint64(r.Uint32()))
	this.MaxMP = uint64(uint64(r.Uint32()))
	v5 := geometry.NewPopulatedPosition(r, easy)
	this.Position = *v5
	if r.Intn(10) != 0 {
		this.Cast = NewPopulatedCast(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyEntity interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEntity(r randyEntity) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEntity(r randyEntity) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneEntity(r)
	}
	return string(tmps)
}
func randUnrecognizedEntity(r randyEntity, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEntity(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEntity(dAtA []byte, r randyEntity, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEntity(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Cast) Size() (n int) {
	var l int
	_ = l
	l = m.AbilityID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.TS)
	n += 1 + l + sovEntity(uint64(l))
	return n
}

func (m *E) Size() (n int) {
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = m.Type.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	if m.Dead {
		n += 2
	}
	if m.HP != 0 {
		n += 1 + sovEntity(uint64(m.HP))
	}
	if m.MaxHP != 0 {
		n += 1 + sovEntity(uint64(m.MaxHP))
	}
	if m.MP != 0 {
		n += 1 + sovEntity(uint64(m.MP))
	}
	if m.MaxMP != 0 {
		n += 1 + sovEntity(uint64(m.MaxMP))
	}
	l = m.Position.Size()
	n += 1 + l + sovEntity(uint64(l))
	if m.Cast != nil {
		l = m.Cast.Size()
		n += 1 + l + sovEntity(uint64(l))
	}
	return n
}

func sovEntity(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEntity(x uint64) (n int) {
	return sovEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Cast) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Cast{`,
		`AbilityID:` + fmt.Sprintf("%v", this.AbilityID) + `,`,
		`TS:` + strings.Replace(strings.Replace(this.TS.String(), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *E) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&E{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Dead:` + fmt.Sprintf("%v", this.Dead) + `,`,
		`HP:` + fmt.Sprintf("%v", this.HP) + `,`,
		`MaxHP:` + fmt.Sprintf("%v", this.MaxHP) + `,`,
		`MP:` + fmt.Sprintf("%v", this.MP) + `,`,
		`MaxMP:` + fmt.Sprintf("%v", this.MaxMP) + `,`,
		`Position:` + strings.Replace(strings.Replace(this.Position.String(), "Position", "geometry.Position", 1), `&`, ``, 1) + `,`,
		`Cast:` + strings.Replace(fmt.Sprintf("%v", this.Cast), "Cast", "Cast", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEntity(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Cast) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
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
			return fmt.Errorf("proto: Cast: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cast: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbilityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AbilityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.TS, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
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
func (m *E) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
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
			return fmt.Errorf("proto: E: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: E: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Type.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dead", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Dead = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HP", wireType)
			}
			m.HP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HP |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHP", wireType)
			}
			m.MaxHP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxHP |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MP", wireType)
			}
			m.MP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MP |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMP", wireType)
			}
			m.MaxMP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMP |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cast", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cast == nil {
				m.Cast = &Cast{}
			}
			if err := m.Cast.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
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
func skipEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntity
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
					return 0, ErrIntOverflowEntity
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
					return 0, ErrIntOverflowEntity
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
				return 0, ErrInvalidLengthEntity
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEntity
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
				next, err := skipEntity(dAtA[start:])
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
	ErrInvalidLengthEntity = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntity   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("entity.proto", fileDescriptor_entity_52795d6082e99de4) }

var fileDescriptor_entity_52795d6082e99de4 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xfd, 0x0e, 0x27, 0x24, 0x47, 0xc5, 0x70, 0x62, 0x38, 0x75, 0x38, 0x5b, 0x95, 0x90,
	0x2c, 0xa1, 0x9e, 0xf9, 0x39, 0x22, 0xd1, 0x10, 0xa4, 0x44, 0xc8, 0xc8, 0x72, 0x33, 0xb1, 0x20,
	0x87, 0x1e, 0xd7, 0x83, 0xb8, 0x67, 0x35, 0x17, 0x09, 0x6f, 0xac, 0x6c, 0xfd, 0x33, 0xf8, 0x13,
	0x18, 0xd9, 0xe8, 0xd8, 0x11, 0x31, 0x14, 0x7c, 0x2c, 0x8c, 0x1d, 0x19, 0x91, 0xcf, 0x71, 0xc3,
	0x84, 0x04, 0xdb, 0xfb, 0x7e, 0xfd, 0x79, 0xef, 0x7d, 0xed, 0x67, 0xbc, 0x25, 0x8e, 0x8c, 0x32,
	0x15, 0x2f, 0x8f, 0xb5, 0xd1, 0xdb, 0xbb, 0x52, 0x99, 0xc3, 0xd5, 0x9c, 0xbf, 0xd4, 0x45, 0x2c,
	0xb5, 0xd4, 0xb1, 0xb3, 0xe7, 0xab, 0x57, 0x4e, 0x39, 0xe1, 0xaa, 0x35, 0xfe, 0xe0, 0x0f, 0x5c,
	0x2c, 0xf4, 0xeb, 0xfc, 0x30, 0x96, 0x79, 0x21, 0x5e, 0xdc, 0xbe, 0x13, 0x97, 0x6f, 0x64, 0x2c,
	0x85, 0x2e, 0x84, 0x39, 0xae, 0xe2, 0x52, 0x2f, 0x95, 0x51, 0xfa, 0x68, 0xdd, 0x16, 0x48, 0xad,
	0xe5, 0x42, 0x6c, 0x86, 0x1b, 0x55, 0x88, 0xa5, 0xc9, 0x8b, 0xb2, 0x05, 0x76, 0xde, 0x03, 0xf6,
	0x1f, 0xe7, 0x4b, 0x43, 0x9e, 0xe2, 0xe1, 0xde, 0x5c, 0x2d, 0x94, 0xa9, 0xa6, 0x63, 0x0a, 0x21,
	0x44, 0x5b, 0xa3, 0xdd, 0xd3, 0xf3, 0xc0, 0xfb, 0x7a, 0x1e, 0xdc, 0xfc, 0xfb, 0xee, 0xd5, 0x42,
	0x1d, 0xf0, 0xe9, 0x38, 0xdb, 0xf4, 0x93, 0xfb, 0x18, 0xcd, 0xf6, 0x29, 0x0a, 0x21, 0xba, 0x76,
	0x77, 0x9b, 0xb7, 0x19, 0x78, 0x97, 0x81, 0xcf, 0xba, 0x0c, 0xa3, 0x41, 0xb3, 0xe1, 0xe4, 0x5b,
	0x00, 0x19, 0x9a, 0xed, 0xef, 0x7c, 0x46, 0x18, 0x9e, 0x90, 0x87, 0x18, 0xfd, 0x6f, 0x02, 0x34,
	0x1d, 0x93, 0x3d, 0xec, 0xcf, 0xaa, 0x52, 0xb8, 0xe5, 0xff, 0x3c, 0xc0, 0xb5, 0x12, 0x82, 0xfd,
	0x67, 0x79, 0x21, 0xe8, 0x95, 0x10, 0xa2, 0x61, 0xe6, 0xea, 0xc6, 0x1b, 0x8b, 0xfc, 0x80, 0xfa,
	0x21, 0x44, 0x83, 0xcc, 0xd5, 0xe4, 0x3a, 0x46, 0x93, 0x94, 0xf6, 0x42, 0x88, 0xfc, 0x0c, 0x4d,
	0x52, 0x72, 0x03, 0xf7, 0x92, 0xfc, 0xed, 0x24, 0xa5, 0x7d, 0x67, 0xb5, 0xa2, 0xa1, 0x92, 0x94,
	0x5e, 0x6d, 0xa9, 0xa4, 0xa3, 0x92, 0x94, 0x0e, 0x2e, 0xa9, 0x24, 0x25, 0xb7, 0xf0, 0x20, 0x5d,
	0x9f, 0x8e, 0x0e, 0xdd, 0x77, 0x1b, 0xf2, 0xce, 0x18, 0xf9, 0xcd, 0x5b, 0x64, 0x97, 0x00, 0x09,
	0xda, 0x9b, 0x51, 0xec, 0xc0, 0x1e, 0x6f, 0x84, 0x83, 0x20, 0x73, 0x0f, 0x46, 0x8f, 0xce, 0x6a,
	0xe6, 0x7d, 0xa9, 0x99, 0x77, 0x51, 0x33, 0xf8, 0x55, 0x33, 0x78, 0x67, 0x19, 0x7c, 0xb0, 0x0c,
	0x3e, 0x5a, 0x06, 0x9f, 0x2c, 0x83, 0x53, 0xcb, 0xe0, 0xcc, 0x32, 0xf8, 0x6e, 0x19, 0xfc, 0xb4,
	0xcc, 0xbb, 0xb0, 0x0c, 0x4e, 0x7e, 0x30, 0xef, 0x79, 0xbf, 0xfd, 0x49, 0xe7, 0x7d, 0x77, 0xad,
	0x7b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0xaa, 0x57, 0x75, 0xb5, 0x02, 0x00, 0x00,
}
