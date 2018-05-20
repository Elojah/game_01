package dto

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (d *Vec3) Size() (s uint64) {

	s += 24
	return
}
func (d *Vec3) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[0+0] = byte(v >> 0)

		buf[1+0] = byte(v >> 8)

		buf[2+0] = byte(v >> 16)

		buf[3+0] = byte(v >> 24)

		buf[4+0] = byte(v >> 32)

		buf[5+0] = byte(v >> 40)

		buf[6+0] = byte(v >> 48)

		buf[7+0] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[0+8] = byte(v >> 0)

		buf[1+8] = byte(v >> 8)

		buf[2+8] = byte(v >> 16)

		buf[3+8] = byte(v >> 24)

		buf[4+8] = byte(v >> 32)

		buf[5+8] = byte(v >> 40)

		buf[6+8] = byte(v >> 48)

		buf[7+8] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[0+16] = byte(v >> 0)

		buf[1+16] = byte(v >> 8)

		buf[2+16] = byte(v >> 16)

		buf[3+16] = byte(v >> 24)

		buf[4+16] = byte(v >> 32)

		buf[5+16] = byte(v >> 40)

		buf[6+16] = byte(v >> 48)

		buf[7+16] = byte(v >> 56)

	}
	return buf[:i+24], nil
}

func (d *Vec3) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		v := 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+8]) << 0) | (uint64(buf[1+8]) << 8) | (uint64(buf[2+8]) << 16) | (uint64(buf[3+8]) << 24) | (uint64(buf[4+8]) << 32) | (uint64(buf[5+8]) << 40) | (uint64(buf[6+8]) << 48) | (uint64(buf[7+8]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+16]) << 0) | (uint64(buf[1+16]) << 8) | (uint64(buf[2+16]) << 16) | (uint64(buf[3+16]) << 24) | (uint64(buf[4+16]) << 32) | (uint64(buf[5+16]) << 40) | (uint64(buf[6+16]) << 48) | (uint64(buf[7+16]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 24, nil
}

type Move struct {
	Source   [16]byte
	Target   [16]byte
	Position Vec3
}

func (d *Move) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += d.Position.Size()
	}
	return
}
func (d *Move) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	{
		nbuf, err := d.Position.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *Move) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{
		ni, err := d.Position.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type Cast struct {
	SkillID  [16]byte
	Source   [16]byte
	Target   [16]byte
	Position Vec3
}

func (d *Cast) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += d.Position.Size()
	}
	return
}
func (d *Cast) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		copy(buf[i+0:], d.SkillID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	{
		nbuf, err := d.Position.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *Cast) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.SkillID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{
		ni, err := d.Position.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type SetPC struct {
	Type [16]byte
}

func (d *SetPC) Size() (s uint64) {

	{
		s += 16
	}
	return
}
func (d *SetPC) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		copy(buf[i+0:], d.Type[:])
		i += 16
	}
	return buf[:i+0], nil
}

func (d *SetPC) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Type[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}

type ConnectPC struct {
	Target [16]byte
}

func (d *ConnectPC) Size() (s uint64) {

	{
		s += 16
	}
	return
}
func (d *ConnectPC) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	return buf[:i+0], nil
}

func (d *ConnectPC) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}

type Message struct {
	Token  [16]byte
	ACK    *[16]byte
	TS     int64
	Action interface{}
}

func (d *Message) Size() (s uint64) {

	{
		s += 16
	}
	{
		if d.ACK != nil {

			{
				s += 16
			}
			s += 0
		}
	}
	{
		var v uint64
		switch d.Action.(type) {

		case Move:
			v = 0 + 1

		case Cast:
			v = 1 + 1

		case ConnectPC:
			v = 2 + 1

		case SetPC:
			v = 3 + 1

		}

		{

			t := v
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		switch tt := d.Action.(type) {

		case Move:

			{
				s += tt.Size()
			}

		case Cast:

			{
				s += tt.Size()
			}

		case ConnectPC:

			{
				s += tt.Size()
			}

		case SetPC:

			{
				s += tt.Size()
			}

		}
	}
	s += 9
	return
}
func (d *Message) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		copy(buf[i+0:], d.Token[:])
		i += 16
	}
	{
		if d.ACK == nil {
			buf[i+0] = 0
		} else {
			buf[i+0] = 1

			{
				copy(buf[i+1:], (*d.ACK)[:])
				i += 16
			}
			i += 0
		}
	}
	{

		buf[i+0+1] = byte(d.TS >> 0)

		buf[i+1+1] = byte(d.TS >> 8)

		buf[i+2+1] = byte(d.TS >> 16)

		buf[i+3+1] = byte(d.TS >> 24)

		buf[i+4+1] = byte(d.TS >> 32)

		buf[i+5+1] = byte(d.TS >> 40)

		buf[i+6+1] = byte(d.TS >> 48)

		buf[i+7+1] = byte(d.TS >> 56)

	}
	{
		var v uint64
		switch d.Action.(type) {

		case Move:
			v = 0 + 1

		case Cast:
			v = 1 + 1

		case ConnectPC:
			v = 2 + 1

		case SetPC:
			v = 3 + 1

		}

		{

			t := uint64(v)

			for t >= 0x80 {
				buf[i+9] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+9] = byte(t)
			i++

		}
		switch tt := d.Action.(type) {

		case Move:

			{
				nbuf, err := tt.Marshal(buf[i+9:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case Cast:

			{
				nbuf, err := tt.Marshal(buf[i+9:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case ConnectPC:

			{
				nbuf, err := tt.Marshal(buf[i+9:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case SetPC:

			{
				nbuf, err := tt.Marshal(buf[i+9:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+9], nil
}

func (d *Message) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	{
		if buf[i+0] == 1 {
			if d.ACK == nil {
				d.ACK = new([16]byte)
			}

			{
				copy((*d.ACK)[:], buf[i+1:])
				i += 16
			}
			i += 0
		} else {
			d.ACK = nil
		}
	}
	{

		d.TS = 0 | (int64(buf[i+0+1]) << 0) | (int64(buf[i+1+1]) << 8) | (int64(buf[i+2+1]) << 16) | (int64(buf[i+3+1]) << 24) | (int64(buf[i+4+1]) << 32) | (int64(buf[i+5+1]) << 40) | (int64(buf[i+6+1]) << 48) | (int64(buf[i+7+1]) << 56)

	}
	{
		v := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+9] & 0x7F)
			for buf[i+9]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+9]&0x7F) << bs
				bs += 7
			}
			i++

			v = t

		}
		switch v {

		case 0 + 1:
			var tt Move

			{
				ni, err := tt.Unmarshal(buf[i+9:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 1 + 1:
			var tt Cast

			{
				ni, err := tt.Unmarshal(buf[i+9:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 2 + 1:
			var tt ConnectPC

			{
				ni, err := tt.Unmarshal(buf[i+9:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 3 + 1:
			var tt SetPC

			{
				ni, err := tt.Unmarshal(buf[i+9:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		default:
			d.Action = nil
		}
	}
	return i + 9, nil
}
