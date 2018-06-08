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

type BondPoint struct {
	ID       [16]byte
	SectorID [16]byte
	X        float64
	Y        float64
	Z        float64
}

func (d *BondPoint) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	s += 24
	return
}
func (d *BondPoint) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.SectorID[:])
		i += 16
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[i+0+0] = byte(v >> 0)

		buf[i+1+0] = byte(v >> 8)

		buf[i+2+0] = byte(v >> 16)

		buf[i+3+0] = byte(v >> 24)

		buf[i+4+0] = byte(v >> 32)

		buf[i+5+0] = byte(v >> 40)

		buf[i+6+0] = byte(v >> 48)

		buf[i+7+0] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[i+0+8] = byte(v >> 0)

		buf[i+1+8] = byte(v >> 8)

		buf[i+2+8] = byte(v >> 16)

		buf[i+3+8] = byte(v >> 24)

		buf[i+4+8] = byte(v >> 32)

		buf[i+5+8] = byte(v >> 40)

		buf[i+6+8] = byte(v >> 48)

		buf[i+7+8] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[i+0+16] = byte(v >> 0)

		buf[i+1+16] = byte(v >> 8)

		buf[i+2+16] = byte(v >> 16)

		buf[i+3+16] = byte(v >> 24)

		buf[i+4+16] = byte(v >> 32)

		buf[i+5+16] = byte(v >> 40)

		buf[i+6+16] = byte(v >> 48)

		buf[i+7+16] = byte(v >> 56)

	}
	return buf[:i+24], nil
}

func (d *BondPoint) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.SectorID[:], buf[i+0:])
		i += 16
	}
	{

		v := 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 24, nil
}

type Sector struct {
	ID         [16]byte
	X          float64
	Y          float64
	Z          float64
	BondPoints []BondPoint
}

func (d *Sector) Size() (s uint64) {

	{
		s += 16
	}
	{
		l := uint64(len(d.BondPoints))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.BondPoints {

			{
				s += d.BondPoints[k0].Size()
			}

		}

	}
	s += 24
	return
}
func (d *Sector) Marshal(buf []byte) ([]byte, error) {
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

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[i+0+0] = byte(v >> 0)

		buf[i+1+0] = byte(v >> 8)

		buf[i+2+0] = byte(v >> 16)

		buf[i+3+0] = byte(v >> 24)

		buf[i+4+0] = byte(v >> 32)

		buf[i+5+0] = byte(v >> 40)

		buf[i+6+0] = byte(v >> 48)

		buf[i+7+0] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[i+0+8] = byte(v >> 0)

		buf[i+1+8] = byte(v >> 8)

		buf[i+2+8] = byte(v >> 16)

		buf[i+3+8] = byte(v >> 24)

		buf[i+4+8] = byte(v >> 32)

		buf[i+5+8] = byte(v >> 40)

		buf[i+6+8] = byte(v >> 48)

		buf[i+7+8] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[i+0+16] = byte(v >> 0)

		buf[i+1+16] = byte(v >> 8)

		buf[i+2+16] = byte(v >> 16)

		buf[i+3+16] = byte(v >> 24)

		buf[i+4+16] = byte(v >> 32)

		buf[i+5+16] = byte(v >> 40)

		buf[i+6+16] = byte(v >> 48)

		buf[i+7+16] = byte(v >> 56)

	}
	{
		l := uint64(len(d.BondPoints))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+24] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+24] = byte(t)
			i++

		}
		for k0 := range d.BondPoints {

			{
				nbuf, err := d.BondPoints[k0].Marshal(buf[i+24:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+24], nil
}

func (d *Sector) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{

		v := 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+24] & 0x7F)
			for buf[i+24]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+24]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.BondPoints)) >= l {
			d.BondPoints = d.BondPoints[:l]
		} else {
			d.BondPoints = make([]BondPoint, l)
		}
		for k0 := range d.BondPoints {

			{
				ni, err := d.BondPoints[k0].Unmarshal(buf[i+24:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 24, nil
}
