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

func (d *A) Size() (s uint64) {

	{
		s += 16
	}
	{
		l := uint64(len(d.Username))

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
		l := uint64(len(d.Password))

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

func (d *A) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.Username))

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
		copy(buf[i+0:], d.Username)
		i += l
	}
	{
		l := uint64(len(d.Password))

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
		copy(buf[i+0:], d.Password)
		i += l
	}
	{
		copy(buf[i+0:], d.Token[:])
		i += 16
	}
	return buf[:i+0], nil
}

func (d *A) Unmarshal(buf []byte) (uint64, error) {
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
		d.Username = string(buf[i+0 : i+0+l])
		i += l
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
		d.Password = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}

func (d *A) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{
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
		d.Username = string(buf[i+0 : i+0+l])
		i += l
		if d.Size() > lb {
			return 0, io.EOF
		}
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
		d.Password = string(buf[i+0 : i+0+l])
		i += l
		if d.Size() > lb {
			return 0, io.EOF
		}
	}
	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}
