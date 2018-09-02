// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ability.proto

package ability

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import github_com_elojah_game_01_pkg_ulid "github.com/elojah/game_01/pkg/ulid"
import time "time"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

type A struct {
	ID                   github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,json=iD,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"ID"`
	Type                 github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,2,opt,name=Type,json=type,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Type"`
	Animation            github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,3,opt,name=Animation,json=animation,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Animation"`
	Name                 string                                `protobuf:"bytes,4,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	MPConsumption        uint64                                `protobuf:"varint,5,opt,name=MPConsumption,json=mPConsumption,proto3" json:"MPConsumption,omitempty"`
	PostMPConsumption    uint64                                `protobuf:"varint,6,opt,name=PostMPConsumption,json=postMPConsumption,proto3" json:"PostMPConsumption,omitempty"`
	CD                   time.Duration                         `protobuf:"bytes,7,opt,name=CD,json=cD,stdduration" json:"CD"`
	LastUsed             time.Time                             `protobuf:"bytes,8,opt,name=LastUsed,json=lastUsed,stdtime" json:"LastUsed"`
	CastTime             time.Duration                         `protobuf:"bytes,9,opt,name=CastTime,json=castTime,stdduration" json:"CastTime"`
	Components           map[string]Component                  `protobuf:"bytes,10,rep,name=Components,json=components" json:"Components" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *A) Reset()      { *m = A{} }
func (*A) ProtoMessage() {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_ability_3bbdbc7fd4d4dca8, []int{0}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(dst, src)
}
func (m *A) XXX_Size() int {
	return m.Size()
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *A) GetMPConsumption() uint64 {
	if m != nil {
		return m.MPConsumption
	}
	return 0
}

func (m *A) GetPostMPConsumption() uint64 {
	if m != nil {
		return m.PostMPConsumption
	}
	return 0
}

func (m *A) GetCD() time.Duration {
	if m != nil {
		return m.CD
	}
	return 0
}

func (m *A) GetLastUsed() time.Time {
	if m != nil {
		return m.LastUsed
	}
	return time.Time{}
}

func (m *A) GetCastTime() time.Duration {
	if m != nil {
		return m.CastTime
	}
	return 0
}

func (m *A) GetComponents() map[string]Component {
	if m != nil {
		return m.Components
	}
	return nil
}

func init() {
	proto.RegisterType((*A)(nil), "ability.A")
	proto.RegisterMapType((map[string]Component)(nil), "ability.A.ComponentsEntry")
}
func (this *A) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*A)
	if !ok {
		that2, ok := that.(A)
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
	if !this.Animation.Equal(that1.Animation) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.MPConsumption != that1.MPConsumption {
		return false
	}
	if this.PostMPConsumption != that1.PostMPConsumption {
		return false
	}
	if this.CD != that1.CD {
		return false
	}
	if !this.LastUsed.Equal(that1.LastUsed) {
		return false
	}
	if this.CastTime != that1.CastTime {
		return false
	}
	if len(this.Components) != len(that1.Components) {
		return false
	}
	for i := range this.Components {
		a := this.Components[i]
		b := that1.Components[i]
		if !(&a).Equal(&b) {
			return false
		}
	}
	return true
}
func (this *A) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&ability.A{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Animation: "+fmt.Sprintf("%#v", this.Animation)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "MPConsumption: "+fmt.Sprintf("%#v", this.MPConsumption)+",\n")
	s = append(s, "PostMPConsumption: "+fmt.Sprintf("%#v", this.PostMPConsumption)+",\n")
	s = append(s, "CD: "+fmt.Sprintf("%#v", this.CD)+",\n")
	s = append(s, "LastUsed: "+fmt.Sprintf("%#v", this.LastUsed)+",\n")
	s = append(s, "CastTime: "+fmt.Sprintf("%#v", this.CastTime)+",\n")
	keysForComponents := make([]string, 0, len(this.Components))
	for k, _ := range this.Components {
		keysForComponents = append(keysForComponents, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForComponents)
	mapStringForComponents := "map[string]Component{"
	for _, k := range keysForComponents {
		mapStringForComponents += fmt.Sprintf("%#v: %#v,", k, this.Components[k])
	}
	mapStringForComponents += "}"
	if this.Components != nil {
		s = append(s, "Components: "+mapStringForComponents+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAbility(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *A) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAbility(dAtA, i, uint64(m.ID.Size()))
	n1, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintAbility(dAtA, i, uint64(m.Type.Size()))
	n2, err := m.Type.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintAbility(dAtA, i, uint64(m.Animation.Size()))
	n3, err := m.Animation.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Name) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAbility(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.MPConsumption != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAbility(dAtA, i, uint64(m.MPConsumption))
	}
	if m.PostMPConsumption != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAbility(dAtA, i, uint64(m.PostMPConsumption))
	}
	dAtA[i] = 0x3a
	i++
	i = encodeVarintAbility(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.CD)))
	n4, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.CD, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x42
	i++
	i = encodeVarintAbility(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.LastUsed)))
	n5, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastUsed, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x4a
	i++
	i = encodeVarintAbility(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.CastTime)))
	n6, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.CastTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.Components) > 0 {
		for k, _ := range m.Components {
			dAtA[i] = 0x52
			i++
			v := m.Components[k]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovAbility(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovAbility(uint64(len(k))) + msgSize
			i = encodeVarintAbility(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintAbility(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintAbility(dAtA, i, uint64((&v).Size()))
			n7, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n7
		}
	}
	return i, nil
}

func encodeVarintAbility(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedA(r randyAbility, easy bool) *A {
	this := &A{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	v2 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Type = *v2
	v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Animation = *v3
	this.Name = string(randStringAbility(r))
	this.MPConsumption = uint64(uint64(r.Uint32()))
	this.PostMPConsumption = uint64(uint64(r.Uint32()))
	v4 := github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.CD = *v4
	v5 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.LastUsed = *v5
	v6 := github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.CastTime = *v6
	if r.Intn(10) != 0 {
		v7 := r.Intn(10)
		this.Components = make(map[string]Component)
		for i := 0; i < v7; i++ {
			this.Components[randStringAbility(r)] = *NewPopulatedComponent(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAbility interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAbility(r randyAbility) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAbility(r randyAbility) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneAbility(r)
	}
	return string(tmps)
}
func randUnrecognizedAbility(r randyAbility, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAbility(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAbility(dAtA []byte, r randyAbility, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAbility(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *A) Size() (n int) {
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovAbility(uint64(l))
	l = m.Type.Size()
	n += 1 + l + sovAbility(uint64(l))
	l = m.Animation.Size()
	n += 1 + l + sovAbility(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAbility(uint64(l))
	}
	if m.MPConsumption != 0 {
		n += 1 + sovAbility(uint64(m.MPConsumption))
	}
	if m.PostMPConsumption != 0 {
		n += 1 + sovAbility(uint64(m.PostMPConsumption))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.CD)
	n += 1 + l + sovAbility(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastUsed)
	n += 1 + l + sovAbility(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.CastTime)
	n += 1 + l + sovAbility(uint64(l))
	if len(m.Components) > 0 {
		for k, v := range m.Components {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovAbility(uint64(len(k))) + 1 + l + sovAbility(uint64(l))
			n += mapEntrySize + 1 + sovAbility(uint64(mapEntrySize))
		}
	}
	return n
}

func sovAbility(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAbility(x uint64) (n int) {
	return sovAbility(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *A) String() string {
	if this == nil {
		return "nil"
	}
	keysForComponents := make([]string, 0, len(this.Components))
	for k, _ := range this.Components {
		keysForComponents = append(keysForComponents, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForComponents)
	mapStringForComponents := "map[string]Component{"
	for _, k := range keysForComponents {
		mapStringForComponents += fmt.Sprintf("%v: %v,", k, this.Components[k])
	}
	mapStringForComponents += "}"
	s := strings.Join([]string{`&A{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Animation:` + fmt.Sprintf("%v", this.Animation) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`MPConsumption:` + fmt.Sprintf("%v", this.MPConsumption) + `,`,
		`PostMPConsumption:` + fmt.Sprintf("%v", this.PostMPConsumption) + `,`,
		`CD:` + strings.Replace(strings.Replace(this.CD.String(), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`LastUsed:` + strings.Replace(strings.Replace(this.LastUsed.String(), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`CastTime:` + strings.Replace(strings.Replace(this.CastTime.String(), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`Components:` + mapStringForComponents + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAbility(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbility
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
			return fmt.Errorf("proto: A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
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
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
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
				return fmt.Errorf("proto: wrong wireType = %d for field Animation", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Animation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MPConsumption", wireType)
			}
			m.MPConsumption = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MPConsumption |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostMPConsumption", wireType)
			}
			m.PostMPConsumption = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostMPConsumption |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CD", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.CD, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUsed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastUsed, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CastTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.CastTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Components", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Components == nil {
				m.Components = make(map[string]Component)
			}
			var mapkey string
			mapvalue := &Component{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAbility
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAbility
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthAbility
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAbility
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthAbility
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthAbility
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Component{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipAbility(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthAbility
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Components[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAbility
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
func skipAbility(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAbility
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
					return 0, ErrIntOverflowAbility
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
					return 0, ErrIntOverflowAbility
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
				return 0, ErrInvalidLengthAbility
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAbility
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
				next, err := skipAbility(dAtA[start:])
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
	ErrInvalidLengthAbility = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAbility   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ability.proto", fileDescriptor_ability_3bbdbc7fd4d4dca8) }

var fileDescriptor_ability_3bbdbc7fd4d4dca8 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xe3, 0x34, 0xd7, 0x26, 0x2e, 0xa7, 0x52, 0x4f, 0xe1, 0x06, 0x5f, 0x84, 0x40, 0xca,
	0x40, 0x13, 0x68, 0x97, 0x0a, 0x09, 0xd1, 0xbb, 0x0b, 0x43, 0xc5, 0x0f, 0x95, 0xa8, 0xcc, 0xc8,
	0x77, 0x67, 0xd2, 0xd0, 0x38, 0x8e, 0x2e, 0x0e, 0x52, 0x36, 0x36, 0xd6, 0x8e, 0xfc, 0x09, 0xfc,
	0x09, 0x8c, 0x8c, 0x1d, 0x3b, 0x22, 0x86, 0xc2, 0x99, 0x85, 0xb1, 0x23, 0x23, 0x8a, 0x93, 0x5c,
	0xcb, 0xb1, 0xa0, 0x6e, 0x7e, 0x7e, 0xdf, 0xcf, 0x7b, 0xcf, 0xcf, 0x5f, 0xd8, 0x25, 0xe3, 0x38,
	0x89, 0x45, 0xe9, 0x65, 0x33, 0x2e, 0x38, 0x5a, 0x6b, 0xc2, 0xde, 0x56, 0x14, 0x8b, 0xa3, 0x62,
	0xec, 0x4d, 0x38, 0xf3, 0x23, 0x1e, 0x71, 0x5f, 0xe5, 0xc7, 0xc5, 0x1b, 0x15, 0xa9, 0x40, 0x9d,
	0x6a, 0xae, 0xb7, 0x31, 0xe1, 0x2c, 0xe3, 0x29, 0x4d, 0x45, 0x73, 0xd1, 0x8f, 0x38, 0x8f, 0x12,
	0x7a, 0x89, 0x89, 0x98, 0xd1, 0x5c, 0x10, 0x96, 0x35, 0x02, 0xbc, 0x2c, 0x98, 0x16, 0x33, 0x22,
	0x62, 0x9e, 0xd6, 0xf9, 0xdb, 0x1f, 0x3a, 0x10, 0x0c, 0xd0, 0x23, 0xa8, 0xef, 0x07, 0x36, 0x70,
	0x80, 0x7b, 0x63, 0xb8, 0x75, 0x7a, 0xde, 0xd7, 0xbe, 0x9d, 0xf7, 0xef, 0x5e, 0x19, 0x8d, 0x26,
	0xfc, 0x2d, 0x39, 0xf2, 0x23, 0xc2, 0xe8, 0xeb, 0xfb, 0x0f, 0xfc, 0xec, 0x38, 0xf2, 0x8b, 0x24,
	0x9e, 0x7a, 0xfb, 0x41, 0xa8, 0xc7, 0x01, 0x1a, 0x40, 0xe3, 0xb0, 0xcc, 0xa8, 0xad, 0x5f, 0xa7,
	0x80, 0x21, 0xca, 0x8c, 0xa2, 0xa7, 0xd0, 0x1a, 0xa4, 0x31, 0x53, 0xa3, 0xd9, 0x2b, 0xd7, 0xa9,
	0x63, 0x91, 0x96, 0x47, 0x08, 0x1a, 0x2f, 0x08, 0xa3, 0xb6, 0xe1, 0x00, 0xd7, 0x0a, 0x8d, 0x94,
	0x30, 0x8a, 0xee, 0xc0, 0xee, 0xf3, 0x83, 0x11, 0x4f, 0xf3, 0x82, 0x65, 0xaa, 0x49, 0xc7, 0x01,
	0xae, 0x11, 0x76, 0xd9, 0xd5, 0x4b, 0x74, 0x0f, 0x6e, 0x1e, 0xf0, 0x5c, 0xfc, 0xad, 0x5c, 0x55,
	0xca, 0xcd, 0x6c, 0x39, 0x81, 0x76, 0xa0, 0x3e, 0x0a, 0xec, 0x35, 0x07, 0xb8, 0xeb, 0xdb, 0xb7,
	0xbc, 0x7a, 0xd3, 0x5e, 0xbb, 0x69, 0x2f, 0x68, 0x36, 0x3d, 0x34, 0xab, 0x87, 0x7c, 0xfc, 0xde,
	0x07, 0xa1, 0x3e, 0x09, 0xd0, 0x1e, 0x34, 0x9f, 0x91, 0x5c, 0xbc, 0xca, 0xe9, 0xd4, 0x36, 0x15,
	0xda, 0xfb, 0x07, 0x3d, 0x6c, 0x7f, 0xb1, 0x66, 0x4f, 0x2a, 0xd6, 0x4c, 0x1a, 0x0a, 0x3d, 0x86,
	0xe6, 0x88, 0xe4, 0xa2, 0x12, 0xd9, 0xd6, 0xff, 0x37, 0x37, 0x27, 0x0d, 0x84, 0xf6, 0x20, 0x1c,
	0xb5, 0x46, 0xca, 0x6d, 0xe8, 0xac, 0xa8, 0x21, 0x5a, 0x8b, 0x0e, 0xbc, 0xcb, 0xe4, 0x93, 0x54,
	0xcc, 0xca, 0xa1, 0x51, 0xd5, 0x08, 0xe1, 0xc2, 0x7c, 0x79, 0xef, 0x25, 0xdc, 0x58, 0x12, 0xa1,
	0x9b, 0x70, 0xe5, 0x98, 0x96, 0xca, 0x44, 0x56, 0x58, 0x1d, 0x91, 0x0b, 0x3b, 0xef, 0x48, 0x52,
	0xd4, 0xbe, 0x58, 0xdf, 0x46, 0x8b, 0x0e, 0x0b, 0x34, 0xac, 0x05, 0x0f, 0xf5, 0x5d, 0x30, 0xdc,
	0x3d, 0x9b, 0x63, 0xed, 0xeb, 0x1c, 0x6b, 0x17, 0x73, 0x0c, 0x7e, 0xcf, 0x31, 0x78, 0x2f, 0x31,
	0xf8, 0x24, 0x31, 0xf8, 0x2c, 0x31, 0xf8, 0x22, 0x31, 0x38, 0x95, 0x18, 0x9c, 0x49, 0x0c, 0x7e,
	0x48, 0x0c, 0x7e, 0x49, 0xac, 0x5d, 0x48, 0x0c, 0x4e, 0x7e, 0x62, 0x6d, 0xbc, 0xaa, 0x5e, 0xbd,
	0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x2b, 0x54, 0x5c, 0x65, 0x03, 0x00, 0x00,
}
