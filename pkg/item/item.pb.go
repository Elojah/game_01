// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: item.proto

package item

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import geometry "github.com/elojah/game_01/pkg/geometry"
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

// Orb represents a skill learner object. Consumer learns skill at item consumption.
type Orb struct {
	Skill *github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=Skill,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Skill,omitempty"`
}

func (m *Orb) Reset()      { *m = Orb{} }
func (*Orb) ProtoMessage() {}
func (*Orb) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_73898ef6c8cf16da, []int{0}
}
func (m *Orb) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Orb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Orb.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Orb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orb.Merge(dst, src)
}
func (m *Orb) XXX_Size() int {
	return m.Size()
}
func (m *Orb) XXX_DiscardUnknown() {
	xxx_messageInfo_Orb.DiscardUnknown(m)
}

var xxx_messageInfo_Orb proto.InternalMessageInfo

// Component represents a common object used for craft only.
type Component struct {
	Type *github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=Type,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Type,omitempty"`
}

func (m *Component) Reset()      { *m = Component{} }
func (*Component) ProtoMessage() {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_73898ef6c8cf16da, []int{1}
}
func (m *Component) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Component.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(dst, src)
}
func (m *Component) XXX_Size() int {
	return m.Size()
}
func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

type Type struct {
	Orb       *Orb       `protobuf:"bytes,1,opt,name=Orb" json:"Orb,omitempty"`
	Component *Component `protobuf:"bytes,2,opt,name=Component" json:"Component,omitempty"`
}

func (m *Type) Reset()      { *m = Type{} }
func (*Type) ProtoMessage() {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_73898ef6c8cf16da, []int{2}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Type.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(dst, src)
}
func (m *Type) XXX_Size() int {
	return m.Size()
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetOrb() *Orb {
	if m != nil {
		return m.Orb
	}
	return nil
}

func (m *Type) GetComponent() *Component {
	if m != nil {
		return m.Component
	}
	return nil
}

type I struct {
	ID       github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"ID"`
	Name     string                                `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Icon     github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,3,opt,name=Icon,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Icon"`
	Mesh     github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,4,opt,name=Mesh,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Mesh"`
	Type     Type                                  `protobuf:"bytes,5,opt,name=Type" json:"Type"`
	Position *geometry.Position                    `protobuf:"bytes,6,opt,name=Position" json:"Position,omitempty"`
}

func (m *I) Reset()      { *m = I{} }
func (*I) ProtoMessage() {}
func (*I) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_73898ef6c8cf16da, []int{3}
}
func (m *I) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *I) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_I.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *I) XXX_Merge(src proto.Message) {
	xxx_messageInfo_I.Merge(dst, src)
}
func (m *I) XXX_Size() int {
	return m.Size()
}
func (m *I) XXX_DiscardUnknown() {
	xxx_messageInfo_I.DiscardUnknown(m)
}

var xxx_messageInfo_I proto.InternalMessageInfo

func (m *I) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *I) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type{}
}

func (m *I) GetPosition() *geometry.Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func init() {
	proto.RegisterType((*Orb)(nil), "item.Orb")
	proto.RegisterType((*Component)(nil), "item.Component")
	proto.RegisterType((*Type)(nil), "item.Type")
	proto.RegisterType((*I)(nil), "item.I")
}
func (this *Orb) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Orb)
	if !ok {
		that2, ok := that.(Orb)
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
	if that1.Skill == nil {
		if this.Skill != nil {
			return false
		}
	} else if !this.Skill.Equal(*that1.Skill) {
		return false
	}
	return true
}
func (this *Component) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Component)
	if !ok {
		that2, ok := that.(Component)
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
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if !this.Type.Equal(*that1.Type) {
		return false
	}
	return true
}
func (this *Type) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Type)
	if !ok {
		that2, ok := that.(Type)
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
	if !this.Orb.Equal(that1.Orb) {
		return false
	}
	if !this.Component.Equal(that1.Component) {
		return false
	}
	return true
}
func (this *I) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*I)
	if !ok {
		that2, ok := that.(I)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.Icon.Equal(that1.Icon) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if !this.Type.Equal(&that1.Type) {
		return false
	}
	if !this.Position.Equal(that1.Position) {
		return false
	}
	return true
}
func (this *Orb) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&item.Orb{")
	s = append(s, "Skill: "+fmt.Sprintf("%#v", this.Skill)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Component) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&item.Component{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&item.Type{")
	if this.Orb != nil {
		s = append(s, "Orb: "+fmt.Sprintf("%#v", this.Orb)+",\n")
	}
	if this.Component != nil {
		s = append(s, "Component: "+fmt.Sprintf("%#v", this.Component)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *I) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&item.I{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Icon: "+fmt.Sprintf("%#v", this.Icon)+",\n")
	s = append(s, "Mesh: "+fmt.Sprintf("%#v", this.Mesh)+",\n")
	s = append(s, "Type: "+strings.Replace(this.Type.GoString(), `&`, ``, 1)+",\n")
	if this.Position != nil {
		s = append(s, "Position: "+fmt.Sprintf("%#v", this.Position)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringItem(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Orb) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Orb) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Skill != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Skill.Size()))
		n1, err := m.Skill.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Component) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Component) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Type.Size()))
		n2, err := m.Type.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Orb != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Orb.Size()))
		n3, err := m.Orb.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Component != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Component.Size()))
		n4, err := m.Component.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *I) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *I) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintItem(dAtA, i, uint64(m.ID.Size()))
	n5, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintItem(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintItem(dAtA, i, uint64(m.Icon.Size()))
	n6, err := m.Icon.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x22
	i++
	i = encodeVarintItem(dAtA, i, uint64(m.Mesh.Size()))
	n7, err := m.Mesh.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x2a
	i++
	i = encodeVarintItem(dAtA, i, uint64(m.Type.Size()))
	n8, err := m.Type.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	if m.Position != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Position.Size()))
		n9, err := m.Position.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}

func encodeVarintItem(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedOrb(r randyItem, easy bool) *Orb {
	this := &Orb{}
	this.Skill = github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedComponent(r randyItem, easy bool) *Component {
	this := &Component{}
	this.Type = github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedType(r randyItem, easy bool) *Type {
	this := &Type{}
	fieldNum := r.Intn(2)
	switch fieldNum {
	case 0:
		this.Orb = NewPopulatedOrb(r, easy)
	case 1:
		this.Component = NewPopulatedComponent(r, easy)
	}
	return this
}

func NewPopulatedI(r randyItem, easy bool) *I {
	this := &I{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	this.Name = string(randStringItem(r))
	v2 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Icon = *v2
	v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Mesh = *v3
	v4 := NewPopulatedType(r, easy)
	this.Type = *v4
	if r.Intn(10) != 0 {
		this.Position = geometry.NewPopulatedPosition(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyItem interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneItem(r randyItem) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringItem(r randyItem) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneItem(r)
	}
	return string(tmps)
}
func randUnrecognizedItem(r randyItem, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldItem(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldItem(dAtA []byte, r randyItem, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateItem(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateItem(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateItem(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateItem(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateItem(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateItem(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateItem(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Orb) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Skill != nil {
		l = m.Skill.Size()
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func (m *Component) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		l = m.Type.Size()
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Orb != nil {
		l = m.Orb.Size()
		n += 1 + l + sovItem(uint64(l))
	}
	if m.Component != nil {
		l = m.Component.Size()
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func (m *I) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovItem(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = m.Icon.Size()
	n += 1 + l + sovItem(uint64(l))
	l = m.Mesh.Size()
	n += 1 + l + sovItem(uint64(l))
	l = m.Type.Size()
	n += 1 + l + sovItem(uint64(l))
	if m.Position != nil {
		l = m.Position.Size()
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func sovItem(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozItem(x uint64) (n int) {
	return sovItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Orb) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Orb{`,
		`Skill:` + fmt.Sprintf("%v", this.Skill) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Component) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Component{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`Orb:` + strings.Replace(fmt.Sprintf("%v", this.Orb), "Orb", "Orb", 1) + `,`,
		`Component:` + strings.Replace(fmt.Sprintf("%v", this.Component), "Component", "Component", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *I) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&I{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Icon:` + fmt.Sprintf("%v", this.Icon) + `,`,
		`Mesh:` + fmt.Sprintf("%v", this.Mesh) + `,`,
		`Type:` + strings.Replace(strings.Replace(this.Type.String(), "Type", "Type", 1), `&`, ``, 1) + `,`,
		`Position:` + strings.Replace(fmt.Sprintf("%v", this.Position), "Position", "geometry.Position", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringItem(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (this *Type) GetValue() interface{} {
	if this.Orb != nil {
		return this.Orb
	}
	if this.Component != nil {
		return this.Component
	}
	return nil
}

func (this *Type) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *Orb:
		this.Orb = vt
	case *Component:
		this.Component = vt
	default:
		return false
	}
	return true
}
func (m *Orb) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: Orb: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Orb: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skill", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_01_pkg_ulid.ID
			m.Skill = &v
			if err := m.Skill.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *Component) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: Component: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Component: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_01_pkg_ulid.ID
			m.Type = &v
			if err := m.Type.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orb", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Orb == nil {
				m.Orb = &Orb{}
			}
			if err := m.Orb.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Component == nil {
				m.Component = &Component{}
			}
			if err := m.Component.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *I) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: I: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: I: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Icon.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mesh", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Mesh.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Type.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Position == nil {
				m.Position = &geometry.Position{}
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func skipItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
				return 0, ErrInvalidLengthItem
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowItem
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
				next, err := skipItem(dAtA[start:])
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
	ErrInvalidLengthItem = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItem   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("item.proto", fileDescriptor_item_73898ef6c8cf16da) }

var fileDescriptor_item_73898ef6c8cf16da = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0xcf, 0xd2, 0x50,
	0x14, 0xc6, 0x7b, 0x4a, 0x21, 0x72, 0x35, 0x31, 0xb9, 0x53, 0x83, 0xc9, 0xc1, 0x34, 0x9a, 0xb8,
	0xd0, 0xfa, 0x77, 0x31, 0x31, 0x51, 0x60, 0xa9, 0x89, 0x60, 0xaa, 0x83, 0x9b, 0xa1, 0x78, 0x2d,
	0x95, 0xb6, 0xb7, 0x29, 0x65, 0x60, 0xf3, 0x23, 0xf8, 0x11, 0x1c, 0xfd, 0x08, 0x8c, 0x8e, 0x8c,
	0x8c, 0xc4, 0x81, 0xd8, 0xcb, 0xe2, 0xc8, 0xe8, 0x68, 0x7a, 0x0a, 0xc8, 0xf4, 0xe6, 0x0d, 0xdb,
	0xb9, 0xf7, 0x3e, 0xcf, 0xef, 0x9c, 0xfb, 0xe4, 0x30, 0x16, 0xe6, 0x22, 0xb6, 0xd3, 0x4c, 0xe6,
	0x92, 0x1b, 0x65, 0xdd, 0xea, 0x04, 0x61, 0x3e, 0x99, 0xfb, 0xf6, 0x58, 0xc6, 0x4e, 0x20, 0x03,
	0xe9, 0xd0, 0xa3, 0x3f, 0xff, 0x4c, 0x27, 0x3a, 0x50, 0x55, 0x99, 0x5a, 0xcf, 0xce, 0xe4, 0x22,
	0x92, 0x5f, 0x46, 0x13, 0x27, 0x18, 0xc5, 0xe2, 0xe3, 0xc3, 0x47, 0x4e, 0x3a, 0x0d, 0x9c, 0x40,
	0xc8, 0x58, 0xe4, 0xd9, 0xc2, 0x49, 0xe5, 0x2c, 0xcc, 0x43, 0x99, 0x54, 0x36, 0xeb, 0x35, 0xab,
	0x0d, 0x33, 0x9f, 0xf7, 0x58, 0xfd, 0xdd, 0x34, 0x8c, 0x22, 0x13, 0xee, 0xc2, 0x83, 0x5b, 0xdd,
	0xce, 0x6a, 0xdb, 0x86, 0x5f, 0xdb, 0xf6, 0xfd, 0xab, 0xa1, 0xf3, 0x28, 0xfc, 0x64, 0xbb, 0x7d,
	0xaf, 0xf2, 0x5a, 0x03, 0xd6, 0xec, 0xc9, 0x38, 0x95, 0x89, 0x48, 0x72, 0xfe, 0x8a, 0x19, 0xef,
	0x17, 0xa9, 0xb8, 0x0c, 0x48, 0x56, 0xeb, 0x43, 0x85, 0xe0, 0x77, 0x68, 0x46, 0x22, 0xdd, 0x7c,
	0xdc, 0xb4, 0x29, 0xa9, 0x61, 0xe6, 0x7b, 0x34, 0x79, 0xe7, 0xac, 0xa9, 0xa9, 0x93, 0xe4, 0x76,
	0x25, 0x39, 0x5d, 0x7b, 0xff, 0x15, 0xcf, 0x8d, 0xd5, 0xf7, 0x36, 0x58, 0x4b, 0x9d, 0x81, 0xcb,
	0x5f, 0x30, 0xdd, 0xed, 0x9f, 0x0d, 0xa8, 0x5d, 0x7f, 0x40, 0xdd, 0xed, 0x73, 0xce, 0x8c, 0xc1,
	0x28, 0x16, 0xd4, 0xb4, 0xe9, 0x51, 0x5d, 0xfe, 0xda, 0x1d, 0xcb, 0xc4, 0xac, 0x5d, 0x02, 0x25,
	0x6b, 0x89, 0x78, 0x23, 0x66, 0x13, 0xd3, 0xb8, 0x08, 0x51, 0x5a, 0xf9, 0xbd, 0x43, 0xf6, 0x75,
	0x8a, 0x83, 0x55, 0x71, 0x94, 0x37, 0x5d, 0xa3, 0xc4, 0x55, 0xf1, 0xf2, 0xa7, 0xec, 0xc6, 0xdb,
	0xc3, 0x32, 0x98, 0x0d, 0x52, 0x72, 0xfb, 0xb8, 0x26, 0xf6, 0xf1, 0x85, 0x1c, 0xe0, 0x9d, 0x94,
	0xdd, 0x97, 0xeb, 0x02, 0xb5, 0x4d, 0x81, 0xda, 0xbe, 0x40, 0xf8, 0x5b, 0x20, 0x7c, 0x55, 0x08,
	0x3f, 0x14, 0xc2, 0x52, 0x21, 0xfc, 0x54, 0x08, 0x2b, 0x85, 0xb0, 0x56, 0x08, 0xbf, 0x15, 0xc2,
	0x1f, 0x85, 0xda, 0x5e, 0x21, 0x7c, 0xdb, 0xa1, 0xb6, 0xde, 0xa1, 0xb6, 0xd9, 0xa1, 0xe6, 0x37,
	0x68, 0xf3, 0x9e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xff, 0x56, 0x6c, 0x9d, 0xf3, 0x02, 0x00,
	0x00,
}
