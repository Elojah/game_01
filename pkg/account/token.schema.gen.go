package account

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

		*(*uint64)(unsafe.Pointer(&buf[i+0])) = d.Ping

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
	{
		copy(buf[i+8:], d.Entity[:])
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

		d.Ping = *(*uint64)(unsafe.Pointer(&buf[i+0]))

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
	{
		copy(d.Entity[:], buf[i+8:])
		i += 16
	}
	return i + 8, nil
}

func (d *Token) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		l := uint64(0)

		{

			if i+0 >= lb {
				return 0, io.EOF
			}
			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for i < lb && buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.IP = string(buf[i+0 : i+0+l])
		i += l
		if d.Size() > lb {
			return 0, io.EOF
		}
	}
	{
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.Account[:], buf[i+0:])
		i += 16
	}
	{

		d.Ping = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		if i+8 >= lb {
			return 0, io.EOF
		}
		copy(d.CorePool[:], buf[i+8:])
		i += 16
	}
	{
		if i+8 >= lb {
			return 0, io.EOF
		}
		copy(d.SyncPool[:], buf[i+8:])
		i += 16
	}
	{
		if i+8 >= lb {
			return 0, io.EOF
		}
		copy(d.PC[:], buf[i+8:])
		i += 16
	}
	{
		if i+8 >= lb {
			return 0, io.EOF
		}
		copy(d.Entity[:], buf[i+8:])
		i += 16
	}
	return i + 8, nil
}
