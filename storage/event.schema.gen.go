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

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (d *Vec3) Size() (s uint64) {

	s += 24
	return
}
func (d *Vec3) Marshal(buf []byte) ([]byte, error) {
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
	return buf[:i+24], nil
}

func (d *Vec3) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		v := 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+8]) << 0) | (uint64(buf[1+8]) << 8) | (uint64(buf[2+8]) << 16) | (uint64(buf[3+8]) << 24) | (uint64(buf[4+8]) << 32) | (uint64(buf[5+8]) << 40) | (uint64(buf[6+8]) << 48) | (uint64(buf[7+8]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+16]) << 0) | (uint64(buf[1+16]) << 8) | (uint64(buf[2+16]) << 16) | (uint64(buf[3+16]) << 24) | (uint64(buf[4+16]) << 32) | (uint64(buf[5+16]) << 40) | (uint64(buf[6+16]) << 48) | (uint64(buf[7+16]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 24, nil
}

type Move struct {
	Source   [16]byte
	Target   [16]byte
	Position Vec3
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
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	{
		nbuf, err := d.Position.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *Move) Unmarshal(buf []byte) (uint64, error) {
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
		ni, err := d.Position.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type Cast struct {
	AbilityID [16]byte
	Source    [16]byte
	Targets   [][16]byte
	Position  Vec3
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
		copy(buf[i+0:], d.AbilityID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{
		l := uint64(len(d.Targets))

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
		for k0 := range d.Targets {

			{
				copy(buf[i+0:], d.Targets[k0][:])
				i += 16
			}

		}
	}
	{
		nbuf, err := d.Position.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *Cast) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.AbilityID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i+0:])
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
		if uint64(cap(d.Targets)) >= l {
			d.Targets = d.Targets[:l]
		} else {
			d.Targets = make([][16]byte, l)
		}
		for k0 := range d.Targets {

			{
				copy(d.Targets[k0][:], buf[i+0:])
				i += 16
			}

		}
	}
	{
		ni, err := d.Position.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type Feedback struct {
	AfbID  [16]byte
	Source [16]byte
	Target [16]byte
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
		copy(buf[i+0:], d.AfbID[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Source[:])
		i += 16
	}
	{
		copy(buf[i+0:], d.Target[:])
		i += 16
	}
	return buf[:i+0], nil
}

func (d *Feedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.AfbID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Source[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
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

		case Cast:
			v = 1 + 1

		case Feedback:
			v = 2 + 1

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

		case Feedback:

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

		case Cast:
			v = 1 + 1

		case Feedback:
			v = 2 + 1

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

		case Feedback:

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
			var tt Cast

			{
				ni, err := tt.Unmarshal(buf[i+8:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 2 + 1:
			var tt Feedback

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
