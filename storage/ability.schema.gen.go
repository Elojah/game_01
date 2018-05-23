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

type HealDirect struct {
	Amount uint64
	Type   uint8
}

func (d *HealDirect) Size() (s uint64) {

	s += 9
	return
}
func (d *HealDirect) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.Amount >> 0)

		buf[1+0] = byte(d.Amount >> 8)

		buf[2+0] = byte(d.Amount >> 16)

		buf[3+0] = byte(d.Amount >> 24)

		buf[4+0] = byte(d.Amount >> 32)

		buf[5+0] = byte(d.Amount >> 40)

		buf[6+0] = byte(d.Amount >> 48)

		buf[7+0] = byte(d.Amount >> 56)

	}
	{

		buf[0+8] = byte(d.Type >> 0)

	}
	return buf[:i+9], nil
}

func (d *HealDirect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)

	}
	{

		d.Type = 0 | (uint8(buf[0+8]) << 0)

	}
	return i + 9, nil
}

type DamageDirect struct {
	Amount uint64
	Type   uint8
}

func (d *DamageDirect) Size() (s uint64) {

	s += 9
	return
}
func (d *DamageDirect) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.Amount >> 0)

		buf[1+0] = byte(d.Amount >> 8)

		buf[2+0] = byte(d.Amount >> 16)

		buf[3+0] = byte(d.Amount >> 24)

		buf[4+0] = byte(d.Amount >> 32)

		buf[5+0] = byte(d.Amount >> 40)

		buf[6+0] = byte(d.Amount >> 48)

		buf[7+0] = byte(d.Amount >> 56)

	}
	{

		buf[0+8] = byte(d.Type >> 0)

	}
	return buf[:i+9], nil
}

func (d *DamageDirect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)

	}
	{

		d.Type = 0 | (uint8(buf[0+8]) << 0)

	}
	return i + 9, nil
}

type HealOverTime struct {
	Amount    uint64
	Type      uint8
	Frequency uint64
	Duration  uint64
}

func (d *HealOverTime) Size() (s uint64) {

	s += 25
	return
}
func (d *HealOverTime) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.Amount >> 0)

		buf[1+0] = byte(d.Amount >> 8)

		buf[2+0] = byte(d.Amount >> 16)

		buf[3+0] = byte(d.Amount >> 24)

		buf[4+0] = byte(d.Amount >> 32)

		buf[5+0] = byte(d.Amount >> 40)

		buf[6+0] = byte(d.Amount >> 48)

		buf[7+0] = byte(d.Amount >> 56)

	}
	{

		buf[0+8] = byte(d.Type >> 0)

	}
	{

		buf[0+9] = byte(d.Frequency >> 0)

		buf[1+9] = byte(d.Frequency >> 8)

		buf[2+9] = byte(d.Frequency >> 16)

		buf[3+9] = byte(d.Frequency >> 24)

		buf[4+9] = byte(d.Frequency >> 32)

		buf[5+9] = byte(d.Frequency >> 40)

		buf[6+9] = byte(d.Frequency >> 48)

		buf[7+9] = byte(d.Frequency >> 56)

	}
	{

		buf[0+17] = byte(d.Duration >> 0)

		buf[1+17] = byte(d.Duration >> 8)

		buf[2+17] = byte(d.Duration >> 16)

		buf[3+17] = byte(d.Duration >> 24)

		buf[4+17] = byte(d.Duration >> 32)

		buf[5+17] = byte(d.Duration >> 40)

		buf[6+17] = byte(d.Duration >> 48)

		buf[7+17] = byte(d.Duration >> 56)

	}
	return buf[:i+25], nil
}

func (d *HealOverTime) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)

	}
	{

		d.Type = 0 | (uint8(buf[0+8]) << 0)

	}
	{

		d.Frequency = 0 | (uint64(buf[0+9]) << 0) | (uint64(buf[1+9]) << 8) | (uint64(buf[2+9]) << 16) | (uint64(buf[3+9]) << 24) | (uint64(buf[4+9]) << 32) | (uint64(buf[5+9]) << 40) | (uint64(buf[6+9]) << 48) | (uint64(buf[7+9]) << 56)

	}
	{

		d.Duration = 0 | (uint64(buf[0+17]) << 0) | (uint64(buf[1+17]) << 8) | (uint64(buf[2+17]) << 16) | (uint64(buf[3+17]) << 24) | (uint64(buf[4+17]) << 32) | (uint64(buf[5+17]) << 40) | (uint64(buf[6+17]) << 48) | (uint64(buf[7+17]) << 56)

	}
	return i + 25, nil
}

type DamageOverTime struct {
	Amount    uint64
	Type      uint8
	Frequency uint64
	Duration  uint64
}

func (d *DamageOverTime) Size() (s uint64) {

	s += 25
	return
}
func (d *DamageOverTime) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.Amount >> 0)

		buf[1+0] = byte(d.Amount >> 8)

		buf[2+0] = byte(d.Amount >> 16)

		buf[3+0] = byte(d.Amount >> 24)

		buf[4+0] = byte(d.Amount >> 32)

		buf[5+0] = byte(d.Amount >> 40)

		buf[6+0] = byte(d.Amount >> 48)

		buf[7+0] = byte(d.Amount >> 56)

	}
	{

		buf[0+8] = byte(d.Type >> 0)

	}
	{

		buf[0+9] = byte(d.Frequency >> 0)

		buf[1+9] = byte(d.Frequency >> 8)

		buf[2+9] = byte(d.Frequency >> 16)

		buf[3+9] = byte(d.Frequency >> 24)

		buf[4+9] = byte(d.Frequency >> 32)

		buf[5+9] = byte(d.Frequency >> 40)

		buf[6+9] = byte(d.Frequency >> 48)

		buf[7+9] = byte(d.Frequency >> 56)

	}
	{

		buf[0+17] = byte(d.Duration >> 0)

		buf[1+17] = byte(d.Duration >> 8)

		buf[2+17] = byte(d.Duration >> 16)

		buf[3+17] = byte(d.Duration >> 24)

		buf[4+17] = byte(d.Duration >> 32)

		buf[5+17] = byte(d.Duration >> 40)

		buf[6+17] = byte(d.Duration >> 48)

		buf[7+17] = byte(d.Duration >> 56)

	}
	return buf[:i+25], nil
}

func (d *DamageOverTime) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)

	}
	{

		d.Type = 0 | (uint8(buf[0+8]) << 0)

	}
	{

		d.Frequency = 0 | (uint64(buf[0+9]) << 0) | (uint64(buf[1+9]) << 8) | (uint64(buf[2+9]) << 16) | (uint64(buf[3+9]) << 24) | (uint64(buf[4+9]) << 32) | (uint64(buf[5+9]) << 40) | (uint64(buf[6+9]) << 48) | (uint64(buf[7+9]) << 56)

	}
	{

		d.Duration = 0 | (uint64(buf[0+17]) << 0) | (uint64(buf[1+17]) << 8) | (uint64(buf[2+17]) << 16) | (uint64(buf[3+17]) << 24) | (uint64(buf[4+17]) << 32) | (uint64(buf[5+17]) << 40) | (uint64(buf[6+17]) << 48) | (uint64(buf[7+17]) << 56)

	}
	return i + 25, nil
}

type Ability struct {
	ID            [16]byte
	Type          [16]byte
	Name          string
	MPConsumption uint64
	CD            uint32
	CurrentCD     uint32
	Components    []interface{}
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
	{
		l := uint64(len(d.Components))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Components {

			{
				var v uint64
				switch d.Components[k0].(type) {

				case HealDirect:
					v = 0 + 1

				case DamageDirect:
					v = 1 + 1

				case HealOverTime:
					v = 2 + 1

				case DamageOverTime:
					v = 3 + 1

				}

				{

					t := v
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				switch tt := d.Components[k0].(type) {

				case HealDirect:

					{
						s += tt.Size()
					}

				case DamageDirect:

					{
						s += tt.Size()
					}

				case HealOverTime:

					{
						s += tt.Size()
					}

				case DamageOverTime:

					{
						s += tt.Size()
					}

				}
			}

		}

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
	{
		l := uint64(len(d.Components))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+16] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+16] = byte(t)
			i++

		}
		for k0 := range d.Components {

			{
				var v uint64
				switch d.Components[k0].(type) {

				case HealDirect:
					v = 0 + 1

				case DamageDirect:
					v = 1 + 1

				case HealOverTime:
					v = 2 + 1

				case DamageOverTime:
					v = 3 + 1

				}

				{

					t := uint64(v)

					for t >= 0x80 {
						buf[i+16] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+16] = byte(t)
					i++

				}
				switch tt := d.Components[k0].(type) {

				case HealDirect:

					{
						nbuf, err := tt.Marshal(buf[i+16:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				case DamageDirect:

					{
						nbuf, err := tt.Marshal(buf[i+16:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				case HealOverTime:

					{
						nbuf, err := tt.Marshal(buf[i+16:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				case DamageOverTime:

					{
						nbuf, err := tt.Marshal(buf[i+16:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				}
			}

		}
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
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+16] & 0x7F)
			for buf[i+16]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+16]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Components)) >= l {
			d.Components = d.Components[:l]
		} else {
			d.Components = make([]interface{}, l)
		}
		for k0 := range d.Components {

			{
				v := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+16] & 0x7F)
					for buf[i+16]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+16]&0x7F) << bs
						bs += 7
					}
					i++

					v = t

				}
				switch v {

				case 0 + 1:
					var tt HealDirect

					{
						ni, err := tt.Unmarshal(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 1 + 1:
					var tt DamageDirect

					{
						ni, err := tt.Unmarshal(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 2 + 1:
					var tt HealOverTime

					{
						ni, err := tt.Unmarshal(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 3 + 1:
					var tt DamageOverTime

					{
						ni, err := tt.Unmarshal(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				default:
					d.Components[k0] = nil
				}
			}

		}
	}
	return i + 16, nil
}
