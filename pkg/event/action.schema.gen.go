package event

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

func (d *Move) Size() (s uint64) {
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

func (d *Move) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i:], d.Source[:])
		i += 16
	}
	{
		copy(buf[i:], d.Target[:])
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

func (d *Move) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.Source[:], buf[i:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i:])
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

func (d *Move) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 32+1 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.Source[:], buf[i:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i:])
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

func (d *Cast) Size() (s uint64) {
	{
		s += 16
	}
	{
		s += 16
	}
	{
		l := uint64(len(d.Targets))
		{
			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++
		}
		for _ = range d.Targets {
			{
				s += 16
			}
		}
	}
	{
		s += d.Position.Size()
	}
	return
}

func (d *Cast) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i:], d.AbilityID[:])
		i += 16
	}
	{
		copy(buf[i:], d.Source[:])
		i += 16
	}
	{
		l := uint64(len(d.Targets))
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
		for k0 := range d.Targets {
			{
				copy(buf[i:], d.Targets[k0][:])
				i += 16
			}
		}
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

func (d *Cast) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.AbilityID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i:])
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
		if uint64(cap(d.Targets)) >= l {
			d.Targets = d.Targets[:l]
		} else {
			d.Targets = make([][16]byte, l)
		}
		for k0 := range d.Targets {
			{
				copy(d.Targets[k0][:], buf[i:])
				i += 16
			}
		}
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

func (d *Cast) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < 32+24+16 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.AbilityID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i:])
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
		if uint64(cap(d.Targets)) >= l {
			d.Targets = d.Targets[:l]
		} else {
			d.Targets = make([][16]byte, l)
		}
		for k0 := range d.Targets {
			{
				if i >= lb {
					return 0, errors.New("invalid buffer")
				}
				copy(d.Targets[k0][:], buf[i:])
				i += 16
			}
		}
	}
	if i >= lb {
		return 0, errors.New("invalid buffer")
	}
	{
		ni, err := d.Position.UnmarshalSafe(buf[i:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

func (d *Feedback) Size() (s uint64) {
	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	return
}

func (d *Feedback) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i:], d.Source[:])
		i += 16
	}
	{
		copy(buf[i:], d.Target[:])
		i += 16
	}
	return buf[:i], nil
}

func (d *Feedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i:])
		i += 16
	}
	return i + 0, nil
}

func (d *Feedback) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 16+16+16 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i:])
		i += 16
	}
	return i + 0, nil
}
