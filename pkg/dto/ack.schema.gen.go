package dto

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

func (d *ACK) Size() (s uint64) {
	{
		s += 16
	}
	return
}

func (d *ACK) Marshal(buf []byte) ([]byte, error) {
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
	return buf[:i], nil
}

func (d *ACK) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	return i, nil
}

func (d *ACK) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 16 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	return i, nil
}
