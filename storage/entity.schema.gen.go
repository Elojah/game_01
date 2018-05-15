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
	ID [16]byte
	HP uint8
	MP uint8
	X  float64
	Y  float64
	Z  float64
}

func (d *Entity) Size() (s uint64) {

	{
		s += 16
	}
	s += 26
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

		buf[i+0+0] = byte(d.HP >> 0)

	}
	{

		buf[i+0+1] = byte(d.MP >> 0)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[i+0+2] = byte(v >> 0)

		buf[i+1+2] = byte(v >> 8)

		buf[i+2+2] = byte(v >> 16)

		buf[i+3+2] = byte(v >> 24)

		buf[i+4+2] = byte(v >> 32)

		buf[i+5+2] = byte(v >> 40)

		buf[i+6+2] = byte(v >> 48)

		buf[i+7+2] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[i+0+10] = byte(v >> 0)

		buf[i+1+10] = byte(v >> 8)

		buf[i+2+10] = byte(v >> 16)

		buf[i+3+10] = byte(v >> 24)

		buf[i+4+10] = byte(v >> 32)

		buf[i+5+10] = byte(v >> 40)

		buf[i+6+10] = byte(v >> 48)

		buf[i+7+10] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[i+0+18] = byte(v >> 0)

		buf[i+1+18] = byte(v >> 8)

		buf[i+2+18] = byte(v >> 16)

		buf[i+3+18] = byte(v >> 24)

		buf[i+4+18] = byte(v >> 32)

		buf[i+5+18] = byte(v >> 40)

		buf[i+6+18] = byte(v >> 48)

		buf[i+7+18] = byte(v >> 56)

	}
	return buf[:i+26], nil
}

func (d *Entity) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{

		d.HP = 0 | (uint8(buf[i+0+0]) << 0)

	}
	{

		d.MP = 0 | (uint8(buf[i+0+1]) << 0)

	}
	{

		v := 0 | (uint64(buf[i+0+2]) << 0) | (uint64(buf[i+1+2]) << 8) | (uint64(buf[i+2+2]) << 16) | (uint64(buf[i+3+2]) << 24) | (uint64(buf[i+4+2]) << 32) | (uint64(buf[i+5+2]) << 40) | (uint64(buf[i+6+2]) << 48) | (uint64(buf[i+7+2]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+10]) << 0) | (uint64(buf[i+1+10]) << 8) | (uint64(buf[i+2+10]) << 16) | (uint64(buf[i+3+10]) << 24) | (uint64(buf[i+4+10]) << 32) | (uint64(buf[i+5+10]) << 40) | (uint64(buf[i+6+10]) << 48) | (uint64(buf[i+7+10]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+18]) << 0) | (uint64(buf[i+1+18]) << 8) | (uint64(buf[i+2+18]) << 16) | (uint64(buf[i+3+18]) << 24) | (uint64(buf[i+4+18]) << 32) | (uint64(buf[i+5+18]) << 40) | (uint64(buf[i+6+18]) << 48) | (uint64(buf[i+7+18]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 26, nil
}
