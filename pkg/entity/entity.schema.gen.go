package entity

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

func (d *E) Size() (s uint64) {

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
		s += d.Position.Size()
	}
	s += 16
	return
}

func (d *E) Marshal(buf []byte) ([]byte, error) {
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

		*(*uint64)(unsafe.Pointer(&buf[i+0])) = d.HP

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[i+8])) = d.MP

	}
	{
		nbuf, err := d.Position.Marshal(buf[i+16:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+16], nil
}

func (d *E) Unmarshal(buf []byte) (uint64, error) {
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

		d.HP = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.MP = *(*uint64)(unsafe.Pointer(&buf[i+8]))

	}
	{
		ni, err := d.Position.Unmarshal(buf[i+16:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 16, nil
}

func (d *E) UnmarshalSafe(buf []byte) (uint64, error) {
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
		copy(d.Type[:], buf[i+0:])
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
		d.Name = string(buf[i+0 : i+0+l])
		i += l
		if d.Size() > lb {
			return 0, io.EOF
		}
	}
	{

		d.HP = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.MP = *(*uint64)(unsafe.Pointer(&buf[i+8]))

	}
	{
		adjust := i + 16
		if adjust >= lb {
			return 0, io.EOF
		}
		ni, err := d.Position.UnmarshalSafe(buf[adjust:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 16, nil
}
