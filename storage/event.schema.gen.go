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
	ID [16]byte
}

func (d *Listener) Size() (s uint64) {

	{
		s += 16
	}
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
	return buf[:i+0], nil
}

func (d *Listener) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}

type Damage struct {
	Source [16]byte
	Amount int64
}

func (d *Damage) Size() (s uint64) {

	{
		s += 16
	}
	s += 8
	return
}
func (d *Damage) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{

		buf[i+0+0] = byte(d.Amount >> 0)

		buf[i+1+0] = byte(d.Amount >> 8)

		buf[i+2+0] = byte(d.Amount >> 16)

		buf[i+3+0] = byte(d.Amount >> 24)

		buf[i+4+0] = byte(d.Amount >> 32)

		buf[i+5+0] = byte(d.Amount >> 40)

		buf[i+6+0] = byte(d.Amount >> 48)

		buf[i+7+0] = byte(d.Amount >> 56)

	}
	return buf[:i+8], nil
}

func (d *Damage) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type DamageInflict struct {
	Target [16]byte
	Amount int64
}

func (d *DamageInflict) Size() (s uint64) {

	{
		s += 16
	}
	s += 8
	return
}
func (d *DamageInflict) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	{

		buf[i+0+0] = byte(d.Amount >> 0)

		buf[i+1+0] = byte(d.Amount >> 8)

		buf[i+2+0] = byte(d.Amount >> 16)

		buf[i+3+0] = byte(d.Amount >> 24)

		buf[i+4+0] = byte(d.Amount >> 32)

		buf[i+5+0] = byte(d.Amount >> 40)

		buf[i+6+0] = byte(d.Amount >> 48)

		buf[i+7+0] = byte(d.Amount >> 56)

	}
	return buf[:i+8], nil
}

func (d *DamageInflict) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type Heal struct {
	Source [16]byte
	Amount int64
}

func (d *Heal) Size() (s uint64) {

	{
		s += 16
	}
	s += 8
	return
}
func (d *Heal) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{

		buf[i+0+0] = byte(d.Amount >> 0)

		buf[i+1+0] = byte(d.Amount >> 8)

		buf[i+2+0] = byte(d.Amount >> 16)

		buf[i+3+0] = byte(d.Amount >> 24)

		buf[i+4+0] = byte(d.Amount >> 32)

		buf[i+5+0] = byte(d.Amount >> 40)

		buf[i+6+0] = byte(d.Amount >> 48)

		buf[i+7+0] = byte(d.Amount >> 56)

	}
	return buf[:i+8], nil
}

func (d *Heal) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type HealInflict struct {
	Target [16]byte
	Amount int64
}

func (d *HealInflict) Size() (s uint64) {

	{
		s += 16
	}
	s += 8
	return
}
func (d *HealInflict) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	{

		buf[i+0+0] = byte(d.Amount >> 0)

		buf[i+1+0] = byte(d.Amount >> 8)

		buf[i+2+0] = byte(d.Amount >> 16)

		buf[i+3+0] = byte(d.Amount >> 24)

		buf[i+4+0] = byte(d.Amount >> 32)

		buf[i+5+0] = byte(d.Amount >> 40)

		buf[i+6+0] = byte(d.Amount >> 48)

		buf[i+7+0] = byte(d.Amount >> 56)

	}
	return buf[:i+8], nil
}

func (d *HealInflict) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type Event struct {
	ID     [16]byte
	TS     int64
	Action interface{}
}

func (d *Event) Size() (s uint64) {

	{
		s += 16
	}
	{
		var v uint64
		switch d.Action.(type) {

		case Listener:
			v = 0 + 1

		case Damage:
			v = 1 + 1

		case DamageInflict:
			v = 2 + 1

		case Heal:
			v = 3 + 1

		case HealInflict:
			v = 4 + 1

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

		case Listener:

			{
				s += tt.Size()
			}

		case Damage:

			{
				s += tt.Size()
			}

		case DamageInflict:

			{
				s += tt.Size()
			}

		case Heal:

			{
				s += tt.Size()
			}

		case HealInflict:

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
		copy(buf[i+0:], d.ID[:])
		i += 16
	}
	{

		buf[i+0+0] = byte(d.TS >> 0)

		buf[i+1+0] = byte(d.TS >> 8)

		buf[i+2+0] = byte(d.TS >> 16)

		buf[i+3+0] = byte(d.TS >> 24)

		buf[i+4+0] = byte(d.TS >> 32)

		buf[i+5+0] = byte(d.TS >> 40)

		buf[i+6+0] = byte(d.TS >> 48)

		buf[i+7+0] = byte(d.TS >> 56)

	}
	{
		var v uint64
		switch d.Action.(type) {

		case Listener:
			v = 0 + 1

		case Damage:
			v = 1 + 1

		case DamageInflict:
			v = 2 + 1

		case Heal:
			v = 3 + 1

		case HealInflict:
			v = 4 + 1

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

		case Listener:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case Damage:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case DamageInflict:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case Heal:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case HealInflict:

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
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{

		d.TS = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

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
			var tt Listener

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 1 + 1:
			var tt Damage

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 2 + 1:
			var tt DamageInflict

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 3 + 1:
			var tt Heal

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 4 + 1:
			var tt HealInflict

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
