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
		copy(buf[i:], d.EntityID[:])
		i += 16
	}
	{
		copy(buf[i:], d.TokenID[:])
		i += 16
	}
	{
		buf[i] = byte(d.Action >> 0)
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
		copy(d.EntityID[:], buf[i:])
		i += 16
	}
	{
		copy(d.TokenID[:], buf[i:])
		i += 16
	}
	{
		d.Action = QAction(0 | (uint8(buf[i]) << 0))
	}
	{
		copy(d.Pool[:], buf[i+1:])
		i += 16
	}
	return i + 1, nil
}
