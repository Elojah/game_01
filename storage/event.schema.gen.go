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

type Move struct {
	X      float64
	Y      float64
	Z      float64
	Target [16]byte
}

func (d *Move) Size() (s uint64) {

	{
		s += 16
	}
	s += 24
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

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[0+0] = byte(v >> 0)

		buf[1+0] = byte(v >> 8)

		buf[2+0] = byte(v >> 16)

		buf[3+0] = byte(v >> 24)

		buf[4+0] = byte(v >> 32)

		buf[5+0] = byte(v >> 40)

		buf[6+0] = byte(v >> 48)

		buf[7+0] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[0+8] = byte(v >> 0)

		buf[1+8] = byte(v >> 8)

		buf[2+8] = byte(v >> 16)

		buf[3+8] = byte(v >> 24)

		buf[4+8] = byte(v >> 32)

		buf[5+8] = byte(v >> 40)

		buf[6+8] = byte(v >> 48)

		buf[7+8] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[0+16] = byte(v >> 0)

		buf[1+16] = byte(v >> 8)

		buf[2+16] = byte(v >> 16)

		buf[3+16] = byte(v >> 24)

		buf[4+16] = byte(v >> 32)

		buf[5+16] = byte(v >> 40)

		buf[6+16] = byte(v >> 48)

		buf[7+16] = byte(v >> 56)

	}
	{
		copy(buf[i+24:], d.Target[:])
		i += 16
	}
	return buf[:i+24], nil
}

func (d *Move) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

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
		copy(d.Target[:], buf[i+24:])
		i += 16
	}
	return i + 24, nil
}

type DamageReceived struct {
	Source [16]byte
	Target [16]byte
	Amount int64
}

func (d *DamageReceived) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	s += 8
	return
}
func (d *DamageReceived) Marshal(buf []byte) ([]byte, error) {
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

func (d *DamageReceived) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type DamageDone struct {
	Source [16]byte
	Target [16]byte
	Amount int64
}

func (d *DamageDone) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	s += 8
	return
}
func (d *DamageDone) Marshal(buf []byte) ([]byte, error) {
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

func (d *DamageDone) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type HealReceived struct {
	Source [16]byte
	Target [16]byte
	Amount int64
}

func (d *HealReceived) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	s += 8
	return
}
func (d *HealReceived) Marshal(buf []byte) ([]byte, error) {
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

func (d *HealReceived) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	{

		d.Amount = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type HealDone struct {
	Source [16]byte
	Target [16]byte
	Amount int64
}

func (d *HealDone) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
	}
	s += 8
	return
}
func (d *HealDone) Marshal(buf []byte) ([]byte, error) {
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

func (d *HealDone) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
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
	Source [16]byte
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

		case DamageReceived:
			v = 1 + 1

		case DamageDone:
			v = 2 + 1

		case HealReceived:
			v = 3 + 1

		case HealDone:
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

		case Move:

			{
				s += tt.Size()
			}

		case DamageReceived:

			{
				s += tt.Size()
			}

		case DamageDone:

			{
				s += tt.Size()
			}

		case HealReceived:

			{
				s += tt.Size()
			}

		case HealDone:

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
		copy(buf[i+0:], d.Source[:])
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

		case Move:
			v = 0 + 1

		case DamageReceived:
			v = 1 + 1

		case DamageDone:
			v = 2 + 1

		case HealReceived:
			v = 3 + 1

		case HealDone:
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

		case Move:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case DamageReceived:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case DamageDone:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case HealReceived:

			{
				nbuf, err := tt.Marshal(buf[i+8:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case HealDone:

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
		copy(d.Source[:], buf[i+0:])
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
			var tt DamageReceived

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 2 + 1:
			var tt DamageDone

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 3 + 1:
			var tt HealReceived

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 4 + 1:
			var tt HealDone

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
