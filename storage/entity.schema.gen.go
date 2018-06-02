package storage

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

type Entity struct {
	ID       [16]byte
	Type     [16]byte
	Name     string
	HP       uint64
	MP       uint64
	SectorID [16]byte
	X        float64
	Y        float64
	Z        float64
}

func (d *Entity) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		s += 16
	}
	s += 40
	return
}
func (d *Entity) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.ID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Type[:])
		i += 16
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	{

		buf[i+0+0] = byte(d.HP >> 0)

		buf[i+1+0] = byte(d.HP >> 8)

		buf[i+2+0] = byte(d.HP >> 16)

		buf[i+3+0] = byte(d.HP >> 24)

		buf[i+4+0] = byte(d.HP >> 32)

		buf[i+5+0] = byte(d.HP >> 40)

		buf[i+6+0] = byte(d.HP >> 48)

		buf[i+7+0] = byte(d.HP >> 56)

	}
	{

		buf[i+0+8] = byte(d.MP >> 0)

		buf[i+1+8] = byte(d.MP >> 8)

		buf[i+2+8] = byte(d.MP >> 16)

		buf[i+3+8] = byte(d.MP >> 24)

		buf[i+4+8] = byte(d.MP >> 32)

		buf[i+5+8] = byte(d.MP >> 40)

		buf[i+6+8] = byte(d.MP >> 48)

		buf[i+7+8] = byte(d.MP >> 56)

	}
	{
		copy(buf[i+16:], d.SectorID[:])
		i += 16
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[i+0+16] = byte(v >> 0)

		buf[i+1+16] = byte(v >> 8)

		buf[i+2+16] = byte(v >> 16)

		buf[i+3+16] = byte(v >> 24)

		buf[i+4+16] = byte(v >> 32)

		buf[i+5+16] = byte(v >> 40)

		buf[i+6+16] = byte(v >> 48)

		buf[i+7+16] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[i+0+24] = byte(v >> 0)

		buf[i+1+24] = byte(v >> 8)

		buf[i+2+24] = byte(v >> 16)

		buf[i+3+24] = byte(v >> 24)

		buf[i+4+24] = byte(v >> 32)

		buf[i+5+24] = byte(v >> 40)

		buf[i+6+24] = byte(v >> 48)

		buf[i+7+24] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[i+0+32] = byte(v >> 0)

		buf[i+1+32] = byte(v >> 8)

		buf[i+2+32] = byte(v >> 16)

		buf[i+3+32] = byte(v >> 24)

		buf[i+4+32] = byte(v >> 32)

		buf[i+5+32] = byte(v >> 40)

		buf[i+6+32] = byte(v >> 48)

		buf[i+7+32] = byte(v >> 56)

	}
	return buf[:i+40], nil
}

func (d *Entity) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Type[:], buf[i+0:])
		i += 16
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{

		d.HP = 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)

	}
	{

		d.MP = 0 | (uint64(buf[i+0+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)

	}
	{
		copy(d.SectorID[:], buf[i+16:])
		i += 16
	}
	{

		v := 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+24]) << 0) | (uint64(buf[i+1+24]) << 8) | (uint64(buf[i+2+24]) << 16) | (uint64(buf[i+3+24]) << 24) | (uint64(buf[i+4+24]) << 32) | (uint64(buf[i+5+24]) << 40) | (uint64(buf[i+6+24]) << 48) | (uint64(buf[i+7+24]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+32]) << 0) | (uint64(buf[i+1+32]) << 8) | (uint64(buf[i+2+32]) << 16) | (uint64(buf[i+3+32]) << 24) | (uint64(buf[i+4+32]) << 32) | (uint64(buf[i+5+32]) << 40) | (uint64(buf[i+6+32]) << 48) | (uint64(buf[i+7+32]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 40, nil
}
