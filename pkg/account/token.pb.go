// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token.proto

package account

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

type Token struct {
	ID       github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"ID"`
	IP       string                                `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Account  github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,3,opt,name=Account,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Account"`
	Ping     uint64                                `protobuf:"varint,4,opt,name=Ping,proto3" json:"Ping,omitempty"`
	CorePool github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,5,opt,name=CorePool,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"CorePool"`
	SyncPool github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,6,opt,name=SyncPool,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"SyncPool"`
	PC       github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,7,opt,name=PC,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"PC"`
	Entity   github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,8,opt,name=Entity,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"Entity"`
}

func (m *Token) Reset()      { *m = Token{} }
func (*Token) ProtoMessage() {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_4da764ff1e8a81cd, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Token) GetPing() uint64 {
	if m != nil {
		return m.Ping
	}
	return 0
}

func init() {
	proto.RegisterType((*Token)(nil), "account.Token")
}
func (this *Token) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Token)
	if !ok {
		that2, ok := that.(Token)
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
	if this.IP != that1.IP {
		return false
	}
	if !this.Account.Equal(that1.Account) {
		return false
	}
	if this.Ping != that1.Ping {
		return false
	}
	if !this.CorePool.Equal(that1.CorePool) {
		return false
	}
	if !this.SyncPool.Equal(that1.SyncPool) {
		return false
	}
	if !this.PC.Equal(that1.PC) {
		return false
	}
	if !this.Entity.Equal(that1.Entity) {
		return false
	}
	return true
}
func (this *Token) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&account.Token{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "IP: "+fmt.Sprintf("%#v", this.IP)+",\n")
	s = append(s, "Account: "+fmt.Sprintf("%#v", this.Account)+",\n")
	s = append(s, "Ping: "+fmt.Sprintf("%#v", this.Ping)+",\n")
	s = append(s, "CorePool: "+fmt.Sprintf("%#v", this.CorePool)+",\n")
	s = append(s, "SyncPool: "+fmt.Sprintf("%#v", this.SyncPool)+",\n")
	s = append(s, "PC: "+fmt.Sprintf("%#v", this.PC)+",\n")
	s = append(s, "Entity: "+fmt.Sprintf("%#v", this.Entity)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringToken(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintToken(dAtA, i, uint64(m.ID.Size()))
	n1, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.IP) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintToken(dAtA, i, uint64(len(m.IP)))
		i += copy(dAtA[i:], m.IP)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintToken(dAtA, i, uint64(m.Account.Size()))
	n2, err := m.Account.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Ping != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintToken(dAtA, i, uint64(m.Ping))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintToken(dAtA, i, uint64(m.CorePool.Size()))
	n3, err := m.CorePool.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintToken(dAtA, i, uint64(m.SyncPool.Size()))
	n4, err := m.SyncPool.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x3a
	i++
	i = encodeVarintToken(dAtA, i, uint64(m.PC.Size()))
	n5, err := m.PC.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x42
	i++
	i = encodeVarintToken(dAtA, i, uint64(m.Entity.Size()))
	n6, err := m.Entity.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedToken(r randyToken, easy bool) *Token {
	this := &Token{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	this.IP = string(randStringToken(r))
	v2 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Account = *v2
	this.Ping = uint64(uint64(r.Uint32()))
	v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.CorePool = *v3
	v4 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.SyncPool = *v4
	v5 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.PC = *v5
	v6 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.Entity = *v6
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyToken interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneToken(r randyToken) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringToken(r randyToken) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneToken(r)
	}
	return string(tmps)
}
func randUnrecognizedToken(r randyToken, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldToken(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldToken(dAtA []byte, r randyToken, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateToken(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateToken(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateToken(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovToken(uint64(l))
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = m.Account.Size()
	n += 1 + l + sovToken(uint64(l))
	if m.Ping != 0 {
		n += 1 + sovToken(uint64(m.Ping))
	}
	l = m.CorePool.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.SyncPool.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.PC.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.Entity.Size()
	n += 1 + l + sovToken(uint64(l))
	return n
}

func sovToken(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Token) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Token{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`IP:` + fmt.Sprintf("%v", this.IP) + `,`,
		`Account:` + fmt.Sprintf("%v", this.Account) + `,`,
		`Ping:` + fmt.Sprintf("%v", this.Ping) + `,`,
		`CorePool:` + fmt.Sprintf("%v", this.CorePool) + `,`,
		`SyncPool:` + fmt.Sprintf("%v", this.SyncPool) + `,`,
		`PC:` + fmt.Sprintf("%v", this.PC) + `,`,
		`Entity:` + fmt.Sprintf("%v", this.Entity) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringToken(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
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
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			m.Ping = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ping |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorePool", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CorePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncPool", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SyncPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PC", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PC.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowToken
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
				next, err := skipToken(dAtA[start:])
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
	ErrInvalidLengthToken = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("token.proto", fileDescriptor_token_4da764ff1e8a81cd) }

var fileDescriptor_token_4da764ff1e8a81cd = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0x02, 0x41,
	0x14, 0x85, 0xe7, 0x2e, 0xbf, 0x8e, 0xc6, 0x62, 0xaa, 0x89, 0xc5, 0x85, 0x98, 0x98, 0xd0, 0xc0,
	0x6a, 0xac, 0x4d, 0x94, 0x9f, 0x98, 0xed, 0x26, 0x68, 0x6f, 0x60, 0x5d, 0x87, 0x15, 0xd8, 0x21,
	0x64, 0xb6, 0xa0, 0xf3, 0x11, 0x7c, 0x0c, 0x1f, 0xc1, 0xd2, 0x92, 0x4e, 0x4a, 0x62, 0x41, 0xdc,
	0xa1, 0xb1, 0xa4, 0xb4, 0x34, 0x3b, 0x88, 0xb1, 0xde, 0xee, 0x9c, 0x4c, 0xbe, 0x2f, 0x93, 0x93,
	0x4b, 0xf7, 0xb5, 0x1a, 0x06, 0x51, 0x63, 0x32, 0x55, 0x5a, 0xb1, 0x52, 0xcf, 0xf7, 0x55, 0x1c,
	0xe9, 0xa3, 0xba, 0x0c, 0xf5, 0x20, 0xee, 0x37, 0x7c, 0x35, 0x76, 0xa5, 0x92, 0xca, 0xb5, 0xef,
	0xfd, 0xf8, 0xc1, 0x36, 0x5b, 0x6c, 0xda, 0x72, 0xc7, 0xef, 0x39, 0x5a, 0xb8, 0x4d, 0x3d, 0xec,
	0x82, 0x3a, 0x5e, 0x9b, 0x43, 0x15, 0x6a, 0x07, 0xcd, 0xfa, 0x7c, 0x55, 0x21, 0x1f, 0xab, 0xca,
	0xc9, 0x3f, 0x59, 0x30, 0x52, 0x8f, 0xbd, 0x81, 0x2b, 0x7b, 0xe3, 0xe0, 0xee, 0xf4, 0xcc, 0x9d,
	0x0c, 0xa5, 0x1b, 0x8f, 0xc2, 0xfb, 0x86, 0xd7, 0xee, 0x3a, 0x5e, 0x9b, 0x1d, 0x52, 0xc7, 0x13,
	0xdc, 0xa9, 0x42, 0x6d, 0xaf, 0xeb, 0x78, 0x82, 0x5d, 0xd3, 0xd2, 0xd5, 0xf6, 0x4b, 0x3c, 0x97,
	0xc5, 0xb9, 0xa3, 0x19, 0xa3, 0x79, 0x11, 0x46, 0x92, 0xe7, 0xab, 0x50, 0xcb, 0x77, 0x6d, 0x66,
	0x1e, 0x2d, 0xb7, 0xd4, 0x34, 0x10, 0x4a, 0x8d, 0x78, 0x21, 0x8b, 0xfd, 0x0f, 0x4f, 0x55, 0x37,
	0xb3, 0xc8, 0xb7, 0xaa, 0x62, 0x26, 0xd5, 0x0e, 0x4f, 0x17, 0x14, 0x2d, 0x5e, 0xca, 0xb4, 0xa0,
	0x68, 0xb1, 0x0e, 0x2d, 0x76, 0x22, 0x1d, 0xea, 0x19, 0x2f, 0x67, 0x51, 0xfc, 0xc2, 0xcd, 0xcb,
	0x45, 0x82, 0x64, 0x99, 0x20, 0xd9, 0x24, 0x08, 0xdf, 0x09, 0xc2, 0x93, 0x41, 0x78, 0x31, 0x08,
	0xaf, 0x06, 0xe1, 0xcd, 0x20, 0xcc, 0x0d, 0xc2, 0xc2, 0x20, 0x7c, 0x1a, 0x84, 0x2f, 0x83, 0x64,
	0x63, 0x10, 0x9e, 0xd7, 0x48, 0x16, 0x6b, 0x24, 0xcb, 0x35, 0x92, 0x7e, 0xd1, 0x9e, 0xc6, 0xf9,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xd7, 0x06, 0xa7, 0x61, 0x02, 0x00, 0x00,
}
