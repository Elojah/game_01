package ability

import (
	"errors"
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

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
		buf[0] = byte(d.Amount >> 0)
		buf[1] = byte(d.Amount >> 8)
		buf[2] = byte(d.Amount >> 16)
		buf[3] = byte(d.Amount >> 24)
		buf[4] = byte(d.Amount >> 32)
		buf[5] = byte(d.Amount >> 40)
		buf[6] = byte(d.Amount >> 48)
		buf[7] = byte(d.Amount >> 56)
	}
	{
		buf[8] = byte(d.Type >> 0)
	}
	return buf[:i+9], nil
}

func (d *HealDirect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		d.Amount = 0 | (uint64(buf[0]) << 0) | (uint64(buf[1]) << 8) | (uint64(buf[2]) << 16) | (uint64(buf[3]) << 24) | (uint64(buf[4]) << 32) | (uint64(buf[5]) << 40) | (uint64(buf[6]) << 48) | (uint64(buf[7]) << 56)
	}
	{
		d.Type = 0 | (uint8(buf[8]) << 0)
	}
	return i + 9, nil
}

func (d *HealDirect) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 9 {
		return 0, errors.New("invalid buffer")
	}
	return d.Unmarshal(buf)
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
		buf[0] = byte(d.Amount >> 0)
		buf[1] = byte(d.Amount >> 8)
		buf[2] = byte(d.Amount >> 16)
		buf[3] = byte(d.Amount >> 24)
		buf[4] = byte(d.Amount >> 32)
		buf[5] = byte(d.Amount >> 40)
		buf[6] = byte(d.Amount >> 48)
		buf[7] = byte(d.Amount >> 56)
	}
	{
		buf[8] = byte(d.Type >> 0)
	}
	return buf[:i+9], nil
}

func (d *DamageDirect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		d.Amount = 0 | (uint64(buf[0]) << 0) | (uint64(buf[1]) << 8) | (uint64(buf[2]) << 16) | (uint64(buf[3]) << 24) | (uint64(buf[4]) << 32) | (uint64(buf[5]) << 40) | (uint64(buf[6]) << 48) | (uint64(buf[7]) << 56)
	}
	{
		d.Type = 0 | (uint8(buf[8]) << 0)
	}
	return i + 9, nil
}

func (d *DamageDirect) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 9 {
		return 0, errors.New("invalid buffer")
	}
	return d.Unmarshal(buf)
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
		buf[0] = byte(d.Amount >> 0)
		buf[1] = byte(d.Amount >> 8)
		buf[2] = byte(d.Amount >> 16)
		buf[3] = byte(d.Amount >> 24)
		buf[4] = byte(d.Amount >> 32)
		buf[5] = byte(d.Amount >> 40)
		buf[6] = byte(d.Amount >> 48)
		buf[7] = byte(d.Amount >> 56)
	}
	{
		buf[8] = byte(d.Type >> 0)
	}
	{
		buf[9] = byte(d.Frequency >> 0)
		buf[1+9] = byte(d.Frequency >> 8)
		buf[2+9] = byte(d.Frequency >> 16)
		buf[3+9] = byte(d.Frequency >> 24)
		buf[4+9] = byte(d.Frequency >> 32)
		buf[5+9] = byte(d.Frequency >> 40)
		buf[6+9] = byte(d.Frequency >> 48)
		buf[7+9] = byte(d.Frequency >> 56)
	}
	{
		buf[17] = byte(d.Duration >> 0)
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
		d.Amount = 0 | (uint64(buf[0]) << 0) | (uint64(buf[1]) << 8) | (uint64(buf[2]) << 16) | (uint64(buf[3]) << 24) | (uint64(buf[4]) << 32) | (uint64(buf[5]) << 40) | (uint64(buf[6]) << 48) | (uint64(buf[7]) << 56)
	}
	{
		d.Type = 0 | (uint8(buf[8]) << 0)
	}
	{
		d.Frequency = 0 | (uint64(buf[9]) << 0) | (uint64(buf[1+9]) << 8) | (uint64(buf[2+9]) << 16) | (uint64(buf[3+9]) << 24) | (uint64(buf[4+9]) << 32) | (uint64(buf[5+9]) << 40) | (uint64(buf[6+9]) << 48) | (uint64(buf[7+9]) << 56)
	}
	{
		d.Duration = 0 | (uint64(buf[17]) << 0) | (uint64(buf[1+17]) << 8) | (uint64(buf[2+17]) << 16) | (uint64(buf[3+17]) << 24) | (uint64(buf[4+17]) << 32) | (uint64(buf[5+17]) << 40) | (uint64(buf[6+17]) << 48) | (uint64(buf[7+17]) << 56)
	}
	return i + 25, nil
}

func (d *HealOverTime) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 25 {
		return 0, errors.New("invalid buffer")
	}
	return d.Unmarshal(buf)
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
		buf[0] = byte(d.Amount >> 0)
		buf[1] = byte(d.Amount >> 8)
		buf[2] = byte(d.Amount >> 16)
		buf[3] = byte(d.Amount >> 24)
		buf[4] = byte(d.Amount >> 32)
		buf[5] = byte(d.Amount >> 40)
		buf[6] = byte(d.Amount >> 48)
		buf[7] = byte(d.Amount >> 56)
	}
	{
		buf[8] = byte(d.Type >> 0)
	}
	{
		buf[9] = byte(d.Frequency >> 0)
		buf[1+9] = byte(d.Frequency >> 8)
		buf[2+9] = byte(d.Frequency >> 16)
		buf[3+9] = byte(d.Frequency >> 24)
		buf[4+9] = byte(d.Frequency >> 32)
		buf[5+9] = byte(d.Frequency >> 40)
		buf[6+9] = byte(d.Frequency >> 48)
		buf[7+9] = byte(d.Frequency >> 56)
	}
	{
		buf[17] = byte(d.Duration >> 0)
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
		d.Amount = 0 | (uint64(buf[0]) << 0) | (uint64(buf[1]) << 8) | (uint64(buf[2]) << 16) | (uint64(buf[3]) << 24) | (uint64(buf[4]) << 32) | (uint64(buf[5]) << 40) | (uint64(buf[6]) << 48) | (uint64(buf[7]) << 56)
	}
	{
		d.Type = 0 | (uint8(buf[8]) << 0)
	}
	{
		d.Frequency = 0 | (uint64(buf[9]) << 0) | (uint64(buf[1+9]) << 8) | (uint64(buf[2+9]) << 16) | (uint64(buf[3+9]) << 24) | (uint64(buf[4+9]) << 32) | (uint64(buf[5+9]) << 40) | (uint64(buf[6+9]) << 48) | (uint64(buf[7+9]) << 56)
	}
	{
		d.Duration = 0 | (uint64(buf[17]) << 0) | (uint64(buf[1+17]) << 8) | (uint64(buf[2+17]) << 16) | (uint64(buf[3+17]) << 24) | (uint64(buf[4+17]) << 32) | (uint64(buf[5+17]) << 40) | (uint64(buf[6+17]) << 48) | (uint64(buf[7+17]) << 56)
	}
	return i + 25, nil
}

func (d *DamageOverTime) UnmarshalSafe(buf []byte) (uint64, error) {
	if len(buf) < 25 {
		return 0, errors.New("invalid buffer")
	}
	return d.Unmarshal(buf)
}

func (d *A) Size() (s uint64) {
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

func (d *A) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i:], d.Type[:])
		i += 16
	}
	{
		l := uint64(len(d.Name))
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
		copy(buf[i:], d.Name)
		i += l
	}
	{
		buf[i] = byte(d.MPConsumption >> 0)
		buf[i+1] = byte(d.MPConsumption >> 8)
		buf[i+2] = byte(d.MPConsumption >> 16)
		buf[i+3] = byte(d.MPConsumption >> 24)
		buf[i+4] = byte(d.MPConsumption >> 32)
		buf[i+5] = byte(d.MPConsumption >> 40)
		buf[i+6] = byte(d.MPConsumption >> 48)
		buf[i+7] = byte(d.MPConsumption >> 56)
	}
	{
		buf[i+8] = byte(d.CD >> 0)
		buf[i+1+8] = byte(d.CD >> 8)
		buf[i+2+8] = byte(d.CD >> 16)
		buf[i+3+8] = byte(d.CD >> 24)
	}
	{
		buf[i+12] = byte(d.CurrentCD >> 0)
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

func (d *A) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Type[:], buf[i:])
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
		d.Name = string(buf[i : i+l])
		i += l
	}
	{
		d.MPConsumption = 0 | (uint64(buf[i]) << 0) | (uint64(buf[i+1]) << 8) | (uint64(buf[i+2]) << 16) | (uint64(buf[i+3]) << 24) | (uint64(buf[i+4]) << 32) | (uint64(buf[i+5]) << 40) | (uint64(buf[i+6]) << 48) | (uint64(buf[i+7]) << 56)
	}
	{
		d.CD = 0 | (uint32(buf[i+8]) << 0) | (uint32(buf[i+1+8]) << 8) | (uint32(buf[i+2+8]) << 16) | (uint32(buf[i+3+8]) << 24)
	}
	{
		d.CurrentCD = 0 | (uint32(buf[i+12]) << 0) | (uint32(buf[i+1+12]) << 8) | (uint32(buf[i+2+12]) << 16) | (uint32(buf[i+3+12]) << 24)
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
			d.Components = make([]Component, l)
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

func (d *A) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < 48 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.Type[:], buf[i:])
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
		d.Name = string(buf[i : i+l])
		i += l
	}
	if i > lb || lb-i < 17 {
		return 0, errors.New("invalid buffer")
	}
	{
		d.MPConsumption = 0 | (uint64(buf[i]) << 0) | (uint64(buf[i+1]) << 8) | (uint64(buf[i+2]) << 16) | (uint64(buf[i+3]) << 24) | (uint64(buf[i+4]) << 32) | (uint64(buf[i+5]) << 40) | (uint64(buf[i+6]) << 48) | (uint64(buf[i+7]) << 56)
	}
	{
		d.CD = 0 | (uint32(buf[i+8]) << 0) | (uint32(buf[i+1+8]) << 8) | (uint32(buf[i+2+8]) << 16) | (uint32(buf[i+3+8]) << 24)
	}
	{
		d.CurrentCD = 0 | (uint32(buf[i+12]) << 0) | (uint32(buf[i+1+12]) << 8) | (uint32(buf[i+2+12]) << 16) | (uint32(buf[i+3+12]) << 24)
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
			d.Components = make([]Component, l)
		}
		for k0 := range d.Components {
			{
				if i > lb || lb-i < 17 {
					return 0, errors.New("invalid buffer")
				}
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
						ni, err := tt.UnmarshalSafe(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					d.Components[k0] = tt
				case 1 + 1:
					var tt DamageDirect
					{
						ni, err := tt.UnmarshalSafe(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					d.Components[k0] = tt
				case 2 + 1:
					var tt HealOverTime
					{
						ni, err := tt.UnmarshalSafe(buf[i+16:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					d.Components[k0] = tt
				case 3 + 1:
					var tt DamageOverTime
					{
						ni, err := tt.UnmarshalSafe(buf[i+16:])
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
