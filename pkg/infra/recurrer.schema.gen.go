package infra

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

func (d *Recurrer) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	s += 1
	return
}

func (d *Recurrer) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.EntityID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.TokenID[:])
		i += 16
	}
	{

		*(*uint8)(unsafe.Pointer(&buf[i+0])) = d.Action

	}
	{
		copy(buf[i+1:], d.Pool[:])
		i += 16
	}
	return buf[:i+1], nil
}

func (d *Recurrer) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.EntityID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.TokenID[:], buf[i+0:])
		i += 16
	}
	{

		d.Action = *(*uint8)(unsafe.Pointer(&buf[i+0]))

	}
	{
		copy(d.Pool[:], buf[i+1:])
		i += 16
	}
	return i + 1, nil
}

func (d *Recurrer) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.EntityID[:], buf[i+0:])
		i += 16
	}
	{
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.TokenID[:], buf[i+0:])
		i += 16
	}
	{

		d.Action = *(*uint8)(unsafe.Pointer(&buf[i+0]))

	}
	{
		if i+1 >= lb {
			return 0, io.EOF
		}
		copy(d.Pool[:], buf[i+1:])
		i += 16
	}
	return i + 1, nil
}
