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

func (d *E) Size() (s uint64) {
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
		copy(buf[i:], d.Source[:])
		i += 16
	}
	{
		t := d.TS.UnixNano()
		buf[i] = byte(t >> 0)
		buf[i+1] = byte(t >> 8)
		buf[i+2] = byte(t >> 16)
		buf[i+3] = byte(t >> 24)
		buf[i+4] = byte(t >> 32)
		buf[i+5] = byte(t >> 40)
		buf[i+6] = byte(t >> 48)
		buf[i+7] = byte(t >> 56)
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
func (d *E) Unmarshal(buf []byte) (uint64, error) {
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
		d.TS = time.Unix(0, 0|(int64(buf[i])<<0)|(int64(buf[i+1])<<8)|(int64(buf[i+2])<<16)|(int64(buf[i+3])<<24)|(int64(buf[i+4])<<32)|(int64(buf[i+5])<<40)|(int64(buf[i+6])<<48)|(int64(buf[i+7])<<56))
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
