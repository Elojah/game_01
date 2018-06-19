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

type Recurrer struct {
	ID       [16]byte
	EntityID [16]byte
	TokenID  [16]byte
	Action   uint8
}

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
		copy(buf[i+0:], d.ID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.EntityID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.TokenID[:])
		i += 16
	}
	{

		buf[i+0+0] = byte(d.Action >> 0)

	}
	return buf[:i+1], nil
}

func (d *Recurrer) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.EntityID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.TokenID[:], buf[i+0:])
		i += 16
	}
	{

		d.Action = 0 | (uint8(buf[i+0+0]) << 0)

	}
	return i + 1, nil
}
