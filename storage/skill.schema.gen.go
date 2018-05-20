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

type Skill struct {
	ID            [16]byte
	Type          [16]byte
	Name          string
	MPConsumption uint64
	DirectDamage  uint64
	DirectHeal    uint64
	CD            uint32
	CurrentCD     uint32
}

func (d *Skill) Size() (s uint64) {

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
	s += 32
	return
}
func (d *Skill) Marshal(buf []byte) ([]byte, error) {
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

		buf[i+0+8] = byte(d.DirectDamage >> 0)

		buf[i+1+8] = byte(d.DirectDamage >> 8)

		buf[i+2+8] = byte(d.DirectDamage >> 16)

		buf[i+3+8] = byte(d.DirectDamage >> 24)

		buf[i+4+8] = byte(d.DirectDamage >> 32)

		buf[i+5+8] = byte(d.DirectDamage >> 40)

		buf[i+6+8] = byte(d.DirectDamage >> 48)

		buf[i+7+8] = byte(d.DirectDamage >> 56)

	}
	{

		buf[i+0+16] = byte(d.DirectHeal >> 0)

		buf[i+1+16] = byte(d.DirectHeal >> 8)

		buf[i+2+16] = byte(d.DirectHeal >> 16)

		buf[i+3+16] = byte(d.DirectHeal >> 24)

		buf[i+4+16] = byte(d.DirectHeal >> 32)

		buf[i+5+16] = byte(d.DirectHeal >> 40)

		buf[i+6+16] = byte(d.DirectHeal >> 48)

		buf[i+7+16] = byte(d.DirectHeal >> 56)

	}
	{

		buf[i+0+24] = byte(d.CD >> 0)

		buf[i+1+24] = byte(d.CD >> 8)

		buf[i+2+24] = byte(d.CD >> 16)

		buf[i+3+24] = byte(d.CD >> 24)

	}
	{

		buf[i+0+28] = byte(d.CurrentCD >> 0)

		buf[i+1+28] = byte(d.CurrentCD >> 8)

		buf[i+2+28] = byte(d.CurrentCD >> 16)

		buf[i+3+28] = byte(d.CurrentCD >> 24)

	}
	return buf[:i+32], nil
}

func (d *Skill) Unmarshal(buf []byte) (uint64, error) {
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

		d.DirectDamage = 0 | (uint64(buf[i+0+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)

	}
	{

		d.DirectHeal = 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)

	}
	{

		d.CD = 0 | (uint32(buf[i+0+24]) << 0) | (uint32(buf[i+1+24]) << 8) | (uint32(buf[i+2+24]) << 16) | (uint32(buf[i+3+24]) << 24)

	}
	{

		d.CurrentCD = 0 | (uint32(buf[i+0+28]) << 0) | (uint32(buf[i+1+28]) << 8) | (uint32(buf[i+2+28]) << 16) | (uint32(buf[i+3+28]) << 24)

	}
	return i + 32, nil
}
