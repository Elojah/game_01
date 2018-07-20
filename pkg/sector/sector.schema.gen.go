package sector

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

func (d *BondPoint) Size() (s uint64) {
	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += d.Position.Size()
	}
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
		copy(buf[i:], d.ID[:])
		i += 16
	}
	{
		copy(buf[i:], d.SectorID[:])
		i += 16
	}
	{
		nbuf, err := d.Position.Marshal(buf[i:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i], nil
}
func (d *BondPoint) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.SectorID[:], buf[i:])
		i += 16
	}
	{
		ni, err := d.Position.Unmarshal(buf[i:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

func (d *S) Size() (s uint64) {
	{
		s += 16
	}
	{
		s += d.Dim.Size()
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
	return
}
func (d *S) Marshal(buf []byte) ([]byte, error) {
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
		nbuf, err := d.Dim.Marshal(buf[i:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		l := uint64(len(d.BondPoints))
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
		for k0 := range d.BondPoints {
			{
				nbuf, err := d.BondPoints[k0].Marshal(buf[i:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
		}
	}
	return buf[:i], nil
}
func (d *S) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		ni, err := d.Dim.Unmarshal(buf[i:])
		if err != nil {
			return 0, err
		}
		i += ni
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
		if uint64(cap(d.BondPoints)) >= l {
			d.BondPoints = d.BondPoints[:l]
		} else {
			d.BondPoints = make([]BondPoint, l)
		}
		for k0 := range d.BondPoints {
			{
				ni, err := d.BondPoints[k0].Unmarshal(buf[i:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
		}
	}
	return i + 0, nil
}
