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

type Ability struct {
	ID            [16]byte
	Type          [16]byte
	Name          string
	MPConsumption uint64
	CD            uint32
	CurrentCD     uint32
}

func (d *Ability) Size() (s uint64) {

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
	s += 16
	return
}
func (d *Ability) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Type[:])
		i += 16
	}
	{
		l := uint64(len(d.Name))

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
		copy(buf[i+0:], d.Name)
		i += l
	}
	{

		buf[i+0+0] = byte(d.MPConsumption >> 0)

		buf[i+1+0] = byte(d.MPConsumption >> 8)

		buf[i+2+0] = byte(d.MPConsumption >> 16)

		buf[i+3+0] = byte(d.MPConsumption >> 24)

		buf[i+4+0] = byte(d.MPConsumption >> 32)

		buf[i+5+0] = byte(d.MPConsumption >> 40)

		buf[i+6+0] = byte(d.MPConsumption >> 48)

		buf[i+7+0] = byte(d.MPConsumption >> 56)

	}
	{

		buf[i+0+8] = byte(d.CD >> 0)

		buf[i+1+8] = byte(d.CD >> 8)

		buf[i+2+8] = byte(d.CD >> 16)

		buf[i+3+8] = byte(d.CD >> 24)

	}
	{

		buf[i+0+12] = byte(d.CurrentCD >> 0)

		buf[i+1+12] = byte(d.CurrentCD >> 8)

		buf[i+2+12] = byte(d.CurrentCD >> 16)

		buf[i+3+12] = byte(d.CurrentCD >> 24)

	}
	return buf[:i+16], nil
}

func (d *Ability) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.Type[:], buf[i+0:])
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
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{

		d.MPConsumption = 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)

	}
	{

		d.CD = 0 | (uint32(buf[i+0+8]) << 0) | (uint32(buf[i+1+8]) << 8) | (uint32(buf[i+2+8]) << 16) | (uint32(buf[i+3+8]) << 24)

	}
	{

		d.CurrentCD = 0 | (uint32(buf[i+0+12]) << 0) | (uint32(buf[i+1+12]) << 8) | (uint32(buf[i+2+12]) << 16) | (uint32(buf[i+3+12]) << 24)

	}
	return i + 16, nil
}
