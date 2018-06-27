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

type Listener struct {
	ID     [16]byte
	Action uint8
}

func (d *Listener) Size() (s uint64) {

	{
		s += 16
	}
	s += 1
	return
}
func (d *Listener) Marshal(buf []byte) ([]byte, error) {
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

		buf[i+0+0] = byte(d.Action >> 0)

	}
	return buf[:i+1], nil
}

func (d *Listener) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{

		d.Action = 0 | (uint8(buf[i+0+0]) << 0)

	}
	return i + 1, nil
}
