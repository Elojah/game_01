// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: items.proto

package sector

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

type Items struct {
	SectorID github_com_elojah_game_01_pkg_ulid.ID   `protobuf:"bytes,1,opt,name=SectorID,proto3,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"SectorID"`
	ItemIDs  []github_com_elojah_game_01_pkg_ulid.ID `protobuf:"bytes,2,rep,name=ItemIDs,customtype=github.com/elojah/game_01/pkg/ulid.ID" json:"ItemIDs"`
}

func (m *Items) Reset()      { *m = Items{} }
func (*Items) ProtoMessage() {}
func (*Items) Descriptor() ([]byte, []int) {
	return fileDescriptor_items_ae69c5adc6372288, []int{0}
}
func (m *Items) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Items) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Items.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Items) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Items.Merge(dst, src)
}
func (m *Items) XXX_Size() int {
	return m.Size()
}
func (m *Items) XXX_DiscardUnknown() {
	xxx_messageInfo_Items.DiscardUnknown(m)
}

var xxx_messageInfo_Items proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Items)(nil), "sector.Items")
}
func (this *Items) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Items)
	if !ok {
		that2, ok := that.(Items)
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
	if !this.SectorID.Equal(that1.SectorID) {
		return false
	}
	if len(this.ItemIDs) != len(that1.ItemIDs) {
		return false
	}
	for i := range this.ItemIDs {
		if !this.ItemIDs[i].Equal(that1.ItemIDs[i]) {
			return false
		}
	}
	return true
}
func (this *Items) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&sector.Items{")
	s = append(s, "SectorID: "+fmt.Sprintf("%#v", this.SectorID)+",\n")
	s = append(s, "ItemIDs: "+fmt.Sprintf("%#v", this.ItemIDs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringItems(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Items) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Items) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintItems(dAtA, i, uint64(m.SectorID.Size()))
	n1, err := m.SectorID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.ItemIDs) > 0 {
		for _, msg := range m.ItemIDs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintItems(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintItems(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedItems(r randyItems, easy bool) *Items {
	this := &Items{}
	v1 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
	this.SectorID = *v1
	v2 := r.Intn(10)
	this.ItemIDs = make([]github_com_elojah_game_01_pkg_ulid.ID, v2)
	for i := 0; i < v2; i++ {
		v3 := github_com_elojah_game_01_pkg_ulid.NewPopulatedID(r)
		this.ItemIDs[i] = *v3
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyItems interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneItems(r randyItems) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringItems(r randyItems) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneItems(r)
	}
	return string(tmps)
}
func randUnrecognizedItems(r randyItems, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldItems(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldItems(dAtA []byte, r randyItems, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateItems(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateItems(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateItems(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateItems(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateItems(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateItems(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateItems(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Items) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SectorID.Size()
	n += 1 + l + sovItems(uint64(l))
	if len(m.ItemIDs) > 0 {
		for _, e := range m.ItemIDs {
			l = e.Size()
			n += 1 + l + sovItems(uint64(l))
		}
	}
	return n
}

func sovItems(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozItems(x uint64) (n int) {
	return sovItems(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Items) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Items{`,
		`SectorID:` + fmt.Sprintf("%v", this.SectorID) + `,`,
		`ItemIDs:` + fmt.Sprintf("%v", this.ItemIDs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringItems(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Items) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItems
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
			return fmt.Errorf("proto: Items: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Items: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SectorID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItems
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
				return ErrInvalidLengthItems
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SectorID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemIDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItems
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
				return ErrInvalidLengthItems
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_01_pkg_ulid.ID
			m.ItemIDs = append(m.ItemIDs, v)
			if err := m.ItemIDs[len(m.ItemIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItems(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItems
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
func skipItems(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItems
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
					return 0, ErrIntOverflowItems
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
					return 0, ErrIntOverflowItems
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
				return 0, ErrInvalidLengthItems
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowItems
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
				next, err := skipItems(dAtA[start:])
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
	ErrInvalidLengthItems = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItems   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("items.proto", fileDescriptor_items_ae69c5adc6372288) }

var fileDescriptor_items_ae69c5adc6372288 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2c, 0x49, 0xcd,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4e, 0x4d, 0x2e, 0xc9, 0x2f, 0x92,
	0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf,
	0xd7, 0x07, 0x4b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa6, 0x34, 0x9b,
	0x91, 0x8b, 0xd5, 0x13, 0x64, 0x8c, 0x90, 0x27, 0x17, 0x47, 0x30, 0xd8, 0x08, 0x4f, 0x17, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x1e, 0x27, 0xdd, 0x13, 0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93, 0x57, 0x45,
	0x32, 0x32, 0x35, 0x27, 0x3f, 0x2b, 0x31, 0x43, 0x3f, 0x3d, 0x31, 0x37, 0x35, 0xde, 0xc0, 0x50,
	0xbf, 0x20, 0x3b, 0x5d, 0xbf, 0x34, 0x27, 0x33, 0x45, 0xcf, 0xd3, 0x25, 0x08, 0xae, 0x5d, 0xc8,
	0x9d, 0x8b, 0x1d, 0x64, 0xa6, 0xa7, 0x4b, 0xb1, 0x04, 0x93, 0x02, 0x33, 0xe9, 0x26, 0xc1, 0x74,
	0x3b, 0x39, 0x5c, 0x78, 0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x3f,
	0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0x71, 0xc7, 0x23, 0x39, 0xc6,
	0x03, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0xde, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x17, 0x55, 0x72, 0x1c, 0x2c, 0x01, 0x00, 0x00,
}