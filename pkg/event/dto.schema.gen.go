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

func (d *DTO) Size() (s uint64) {

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

func (d *DTO) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Token[:])
		i += 16
	}
	{

		*(*int64)(unsafe.Pointer(&buf[i+0])) = d.TS

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

func (d *DTO) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	{

		d.TS = *(*int64)(unsafe.Pointer(&buf[i+0]))

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

func (d *DTO) UnmarshalSafe(buf []byte) (uint64, error) {
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
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	{

		d.TS = *(*int64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		v := uint64(0)

		{

			if i+8 >= lb {
				return 0, io.EOF
			}
			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for i < lb && buf[i+8]&0x80 == 0x80 {
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
				adjust := i + 8
				if adjust >= lb {
					return 0, io.EOF
				}
				ni, err := tt.UnmarshalSafe(buf[adjust:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Action = tt

		case 1 + 1:
			var tt Cast

			{
				adjust := i + 8
				if adjust >= lb {
					return 0, io.EOF
				}
				ni, err := tt.UnmarshalSafe(buf[adjust:])
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
