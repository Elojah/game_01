package dto

import (
	"errors"
	"io"
	"time"
	"unsafe"

	"github.com/elojah/game_01/pkg/geometry"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Move struct {
	Source   [16]byte
	Target   [16]byte
	Position geometry.Vec3
}

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
	return i, nil
}

func (d *Move) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := len(buf)
	if lb < 32 {
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
	return i, nil
}

type Cast struct {
	AbilityID [16]byte
	Source    [16]byte
	Targets   [][16]byte
	Position  geometry.Vec3
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
	return i, nil
}

type Event struct {
	ID     [16]byte
	Token  [16]byte
	TS     int64
	Action interface{}
}

func (d *Event) Size() (s uint64) {
	{
		s += 16
	}
	{
		s += 16
	}
	{
		var v uint64
		switch d.Action.(type) {
		case Move:
			v = 0 + 1
		case Cast:
			v = 1 + 1
		}
		{
			t := v
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++
		}
		switch tt := d.Action.(type) {
		case Move:
			{
				s += tt.Size()
			}
		case Cast:
			{
				s += tt.Size()
			}
		}
	}
	s += 8
	return
}

func (d *Event) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i:], d.Token[:])
		i += 16
	}
	{
		buf[i] = byte(d.TS >> 0)
		buf[i+1] = byte(d.TS >> 8)
		buf[i+2] = byte(d.TS >> 16)
		buf[i+3] = byte(d.TS >> 24)
		buf[i+4] = byte(d.TS >> 32)
		buf[i+5] = byte(d.TS >> 40)
		buf[i+6] = byte(d.TS >> 48)
		buf[i+7] = byte(d.TS >> 56)
	}
	{
		var v uint64
		switch d.Action.(type) {
		case Move:
			v = 0 + 1
		case Cast:
			v = 1 + 1
		}
		{
			t := uint64(v)
			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++
		}
		switch tt := d.Action.(type) {
		case Move:
			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
		case Cast:
			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
		}
	}
	return buf[:i+8], nil
}

func (d *Event) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Token[:], buf[i:])
		i += 16
	}
	{
		d.TS = 0 | (int64(buf[i]) << 0) | (int64(buf[i+1]) << 8) | (int64(buf[i+2]) << 16) | (int64(buf[i+3]) << 24) | (int64(buf[i+4]) << 32) | (int64(buf[i+5]) << 40) | (int64(buf[i+6]) << 48) | (int64(buf[i+7]) << 56)
	}
	{
		v := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++
			v = t
		}
		switch v {
		case 0 + 1:
			var tt Move
			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			d.Action = tt
		case 1 + 1:
			var tt Cast
			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			d.Action = tt
		default:
			d.Action = nil
		}
	}
	return i + 8, nil
}
