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

type Permission struct {
	ID    [16]byte
	Right uint8
}

func (d *Permission) Size() (s uint64) {

	{
		s += 16
	}
	s += 1
	return
}
func (d *Permission) Marshal(buf []byte) ([]byte, error) {
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

		buf[i+0+0] = byte(d.Right >> 0)

	}
	return buf[:i+1], nil
}

func (d *Permission) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{

		d.Right = 0 | (uint8(buf[i+0+0]) << 0)

	}
	return i + 1, nil
}

type Token struct {
	ID      [16]byte
	IP      string
	Account [16]byte
}

func (d *Token) Size() (s uint64) {

	{
		s += 16
	}
	{
		l := uint64(len(d.IP))

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
	return
}
func (d *Token) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.IP))

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
		copy(buf[i+0:], d.IP)
		i += l
	}
	{
		copy(buf[i+0:], d.Account[:])
		i += 16
	}
	return buf[:i+0], nil
}

func (d *Token) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
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
		d.IP = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		copy(d.Account[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}
