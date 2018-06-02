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

type SectorEntities struct {
	SectorID  [16]byte
	EntityIDs [][16]byte
}

func (d *SectorEntities) Size() (s uint64) {

	{
		s += 16
	}
	{
		l := uint64(len(d.EntityIDs))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for _ = range d.EntityIDs {

			{
				s += 16
			}

		}

	}
	return
}
func (d *SectorEntities) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.SectorID[:])
		i += 16
	}
	{
		l := uint64(len(d.EntityIDs))

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
		for k0 := range d.EntityIDs {

			{
				copy(buf[i+0:], d.EntityIDs[k0][:])
				i += 16
			}

		}
	}
	return buf[:i+0], nil
}

func (d *SectorEntities) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.SectorID[:], buf[i+0:])
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
		if uint64(cap(d.EntityIDs)) >= l {
			d.EntityIDs = d.EntityIDs[:l]
		} else {
			d.EntityIDs = make([][16]byte, l)
		}
		for k0 := range d.EntityIDs {

			{
				copy(d.EntityIDs[k0][:], buf[i+0:])
				i += 16
			}

		}
	}
	return i + 0, nil
}
