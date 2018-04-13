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

type Attack struct {
}

func (d *Attack) Size() (s uint64) {

	return
}
func (d *Attack) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	return buf[:i+0], nil
}

func (d *Attack) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	return i + 0, nil
}

type Message struct {
	Token    [16]byte
	Actor    [16]byte
	Position *Vec3
	Action   interface{}
	ACK      *int64
}

func (d *Message) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	{
		if d.Position != nil {

			{
				s += (*d.Position).Size()
			}
			s += 0
		}
	}
	{
		var v uint64
		switch d.Action.(type) {

		case Attack:
			v = 0 + 1

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

		case Attack:

			{
				s += tt.Size()
			}

		}
	}
	{
		if d.ACK != nil {

			s += 8
		}
	}
	s += 2
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
		copy(buf[i+0:], d.Actor[:])
		i += 16
	}
	{
		if d.Position == nil {
			buf[i+0] = 0
		} else {
			buf[i+0] = 1

			{
				nbuf, err := (*d.Position).Marshal(buf[i+1:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
			i += 0
		}
	}
	{
		var v uint64
		switch d.Action.(type) {

		case Attack:
			v = 0 + 1

		}

		{

			t := uint64(v)

			for t >= 0x80 {
				buf[i+1] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+1] = byte(t)
			i++

		}
		switch tt := d.Action.(type) {

		case Attack:

			{
				nbuf, err := tt.Marshal(buf[i+1:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	{
		if d.ACK == nil {
			buf[i+1] = 0
		} else {
			buf[i+1] = 1

			{

				buf[i+0+2] = byte((*d.ACK) >> 0)

				buf[i+1+2] = byte((*d.ACK) >> 8)

				buf[i+2+2] = byte((*d.ACK) >> 16)

				buf[i+3+2] = byte((*d.ACK) >> 24)

				buf[i+4+2] = byte((*d.ACK) >> 32)

				buf[i+5+2] = byte((*d.ACK) >> 40)

				buf[i+6+2] = byte((*d.ACK) >> 48)

				buf[i+7+2] = byte((*d.ACK) >> 56)

			}
			i += 8
		}
	}
	return buf[:i+2], nil
}

func (d *Message) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Actor[:], buf[i+0:])
		i += 16
	}
	{
		if buf[i+0] == 1 {
			if d.Position == nil {
				d.Position = new(Vec3)
			}

			{
				ni, err := (*d.Position).Unmarshal(buf[i+1:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			i += 0
		} else {
			d.Position = nil
		}
	}
	{
		v := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+1] & 0x7F)
			for buf[i+1]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+1]&0x7F) << bs
				bs += 7
			}
			i++

			v = t

		}
		switch v {

		case 0 + 1:
			var tt Attack

			{
				ni, err := tt.Unmarshal(buf[i+1:])
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
	{
		if buf[i+1] == 1 {
			if d.ACK == nil {
				d.ACK = new(int64)
			}

			{

				(*d.ACK) = 0 | (int64(buf[i+0+2]) << 0) | (int64(buf[i+1+2]) << 8) | (int64(buf[i+2+2]) << 16) | (int64(buf[i+3+2]) << 24) | (int64(buf[i+4+2]) << 32) | (int64(buf[i+5+2]) << 40) | (int64(buf[i+6+2]) << 48) | (int64(buf[i+7+2]) << 56)

			}
			i += 8
		} else {
			d.ACK = nil
		}
	}
	return i + 2, nil
}
