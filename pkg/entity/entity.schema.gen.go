package entity

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

func (d *Position) Size() (s uint64) {
	{
		s += d.Coord.Size()
	}
	{
		s += 16
	}
	return
}

func (d *Position) Marshal(buf []byte) ([]byte, error) {
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
		nbuf, err := d.Coord.Marshal(buf[0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		copy(buf[i:], d.SectorID[:])
		i += 16
	}
	return buf[:i], nil
}

func (d *Position) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		ni, err := d.Coord.Unmarshal(buf[i:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		copy(d.SectorID[:], buf[i:])
		i += 16
	}
	return i + 0, nil
}

func (d *Position) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 24+16 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		ni, err := d.Coord.Unmarshal(buf[i:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		copy(d.SectorID[:], buf[i:])
		i += 16
	}
	return i + 0, nil
}

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
		copy(buf[i:], d.ID[:])
		i += 16
	}
	{
		copy(buf[i:], d.Type[:])
		i += 16
	}
	{
		l := uint64(len(d.Name))
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
		copy(buf[i:], d.Name)
		i += l
	}
	{
		buf[i] = byte(d.HP >> 0)
		buf[i+1] = byte(d.HP >> 8)
		buf[i+2] = byte(d.HP >> 16)
		buf[i+3] = byte(d.HP >> 24)
		buf[i+4] = byte(d.HP >> 32)
		buf[i+5] = byte(d.HP >> 40)
		buf[i+6] = byte(d.HP >> 48)
		buf[i+7] = byte(d.HP >> 56)
	}
	{
		buf[i+8] = byte(d.MP >> 0)
		buf[i+1+8] = byte(d.MP >> 8)
		buf[i+2+8] = byte(d.MP >> 16)
		buf[i+3+8] = byte(d.MP >> 24)
		buf[i+4+8] = byte(d.MP >> 32)
		buf[i+5+8] = byte(d.MP >> 40)
		buf[i+6+8] = byte(d.MP >> 48)
		buf[i+7+8] = byte(d.MP >> 56)
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
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Type[:], buf[i:])
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
		d.Name = string(buf[i : i+l])
		i += l
	}
	{
		d.HP = 0 | (uint64(buf[i]) << 0) | (uint64(buf[i+1]) << 8) | (uint64(buf[i+2]) << 16) | (uint64(buf[i+3]) << 24) | (uint64(buf[i+4]) << 32) | (uint64(buf[i+5]) << 40) | (uint64(buf[i+6]) << 48) | (uint64(buf[i+7]) << 56)
	}
	{
		d.MP = 0 | (uint64(buf[i+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)
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
	if lb < 32+16+24+16+1 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Type[:], buf[i:])
		i += 16
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 && i < lb {
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
		d.Name = string(buf[i : i+l])
		i += l
	}
	if lb-i < 17 {
		return 0, errors.New("invalid buffer")
	}
	{
		d.HP = 0 | (uint64(buf[i]) << 0) | (uint64(buf[i+1]) << 8) | (uint64(buf[i+2]) << 16) | (uint64(buf[i+3]) << 24) | (uint64(buf[i+4]) << 32) | (uint64(buf[i+5]) << 40) | (uint64(buf[i+6]) << 48) | (uint64(buf[i+7]) << 56)
	}
	{
		d.MP = 0 | (uint64(buf[i+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)
	}
	{
		ni, err := d.Position.UnmarshalSafe(buf[i+16:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 16, nil
}
