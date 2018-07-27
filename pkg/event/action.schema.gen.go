package event

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

func (d *Move) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
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
		adjust := i + 0
		if adjust >= lb {
			return 0, io.EOF
		}
		ni, err := d.Position.UnmarshalSafe(buf[adjust:])
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

		for k0 := range d.Targets {

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

func (d *Cast) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
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

			if i+0 >= lb {
				return 0, io.EOF
			}
			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for i < lb && buf[i+0]&0x80 == 0x80 {
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
		adjust := i + 0
		if adjust >= lb {
			return 0, io.EOF
		}
		ni, err := d.Position.UnmarshalSafe(buf[adjust:])
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
		copy(buf[i+0:], d.ID[:])
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
		copy(d.ID[:], buf[i+0:])
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

func (d *Feedback) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
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
		copy(d.Target[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}
