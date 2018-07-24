package account

import (
	"errors"
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
		copy(buf[i:], d.ID[:])
		i += 16
	}
	{
		l := uint64(len(d.Username))
		{
			t := uint64(l)
			for t >= 0x80 {
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
			i++
		}
		copy(buf[i:], d.Username)
		i += l
	}
	{
		l := uint64(len(d.Password))
		{
			t := uint64(l)
			for t >= 0x80 {
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
			i++
		}
		copy(buf[i:], d.Password)
		i += l
	}
	{
		copy(buf[i:], d.Token[:])
		i += 16
	}
	return buf[:i], nil
}

func (d *A) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		d.Username = string(buf[i : i+l])
		i += l
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		d.Password = string(buf[i : i+l])
		i += l
	}
	{
		copy(d.Token[:], buf[i:])
		i += 16
	}
	return i + 0, nil
}

func (d *A) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < 32 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for i < lb && buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		if i+l >= lb {
			return 0, errors.New("invalid buffer")
		}
		d.Username = string(buf[i : i+l])
		i += l
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for i < lb && buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		if i+l >= lb {
			return 0, errors.New("invalid buffer")
		}
		d.Password = string(buf[i : i+l])
		i += l
	}
	{
		copy(d.Token[:], buf[i:])
		i += 16
	}
	return i + 0, nil
}
