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

type Token struct {
	ID       [16]byte
	IP       string
	Account  [16]byte
	Ping     uint64
	CorePool [16]byte
	SyncPool [16]byte
	PC       [16]byte
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
	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	s += 8
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
	{

		buf[i+0+0] = byte(d.Ping >> 0)

		buf[i+1+0] = byte(d.Ping >> 8)

		buf[i+2+0] = byte(d.Ping >> 16)

		buf[i+3+0] = byte(d.Ping >> 24)

		buf[i+4+0] = byte(d.Ping >> 32)

		buf[i+5+0] = byte(d.Ping >> 40)

		buf[i+6+0] = byte(d.Ping >> 48)

		buf[i+7+0] = byte(d.Ping >> 56)

	}
	{
		copy(buf[i+8:], d.CorePool[:])
		i += 16
	}
	{
		copy(buf[i+8:], d.SyncPool[:])
		i += 16
	}
	{
		copy(buf[i+8:], d.PC[:])
		i += 16
	}
	return buf[:i+8], nil
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
	{

		d.Ping = 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)

	}
	{
		copy(d.CorePool[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.SyncPool[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.PC[:], buf[i+8:])
		i += 16
	}
	return i + 8, nil
}
