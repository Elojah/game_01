package geometry

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
		buf[0] = byte(v >> 0)
		buf[1] = byte(v >> 8)
		buf[2] = byte(v >> 16)
		buf[3] = byte(v >> 24)
		buf[4] = byte(v >> 32)
		buf[5] = byte(v >> 40)
		buf[6] = byte(v >> 48)
		buf[7] = byte(v >> 56)
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
		v := 0 | (uint64(buf[0]) << 0) | (uint64(buf[1]) << 8) | (uint64(buf[2]) << 16) | (uint64(buf[3]) << 24) | (uint64(buf[4]) << 32) | (uint64(buf[5]) << 40) | (uint64(buf[6]) << 48) | (uint64(buf[7]) << 56)
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
