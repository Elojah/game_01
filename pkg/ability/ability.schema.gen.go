package ability

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

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.Amount

	}
	{

		*(*uint8)(unsafe.Pointer(&buf[8])) = d.Type

	}
	return buf[:i+9], nil
}

func (d *HealDirect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	return i + 9, nil
}

func (d *HealDirect) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	return i + 9, nil
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

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.Amount

	}
	{

		*(*uint8)(unsafe.Pointer(&buf[8])) = d.Type

	}
	return buf[:i+9], nil
}

func (d *DamageDirect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	return i + 9, nil
}

func (d *DamageDirect) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	return i + 9, nil
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

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.Amount

	}
	{

		*(*uint8)(unsafe.Pointer(&buf[8])) = d.Type

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[9])) = d.Frequency

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[17])) = d.Duration

	}
	return buf[:i+25], nil
}

func (d *HealOverTime) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	{

		d.Frequency = *(*uint64)(unsafe.Pointer(&buf[9]))

	}
	{

		d.Duration = *(*uint64)(unsafe.Pointer(&buf[17]))

	}
	return i + 25, nil
}

func (d *HealOverTime) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	{

		d.Frequency = *(*uint64)(unsafe.Pointer(&buf[9]))

	}
	{

		d.Duration = *(*uint64)(unsafe.Pointer(&buf[17]))

	}
	return i + 25, nil
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

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.Amount

	}
	{

		*(*uint8)(unsafe.Pointer(&buf[8])) = d.Type

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[9])) = d.Frequency

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[17])) = d.Duration

	}
	return buf[:i+25], nil
}

func (d *DamageOverTime) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	{

		d.Frequency = *(*uint64)(unsafe.Pointer(&buf[9]))

	}
	{

		d.Duration = *(*uint64)(unsafe.Pointer(&buf[17]))

	}
	return i + 25, nil
}

func (d *DamageOverTime) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{

		d.Amount = *(*uint64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Type = *(*uint8)(unsafe.Pointer(&buf[8]))

	}
	{

		d.Frequency = *(*uint64)(unsafe.Pointer(&buf[9]))

	}
	{

		d.Duration = *(*uint64)(unsafe.Pointer(&buf[17]))

	}
	return i + 25, nil
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
			_ = k0 // make compiler happy in case k is unused

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

		*(*uint64)(unsafe.Pointer(&buf[i+0])) = d.MPConsumption

	}
	{

		*(*uint32)(unsafe.Pointer(&buf[i+8])) = d.CD

	}
	{

		*(*uint32)(unsafe.Pointer(&buf[i+12])) = d.CurrentCD

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

		d.MPConsumption = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.CD = *(*uint32)(unsafe.Pointer(&buf[i+8]))

	}
	{

		d.CurrentCD = *(*uint32)(unsafe.Pointer(&buf[i+12]))

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
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.Type[:], buf[i+0:])
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
		d.Name = string(buf[i+0 : i+0+l])
		i += l
		if d.Size() > lb {
			return 0, io.EOF
		}
	}
	{

		d.MPConsumption = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.CD = *(*uint32)(unsafe.Pointer(&buf[i+8]))

	}
	{

		d.CurrentCD = *(*uint32)(unsafe.Pointer(&buf[i+12]))

	}
	{
		l := uint64(0)

		{

			if i+16 >= lb {
				return 0, io.EOF
			}
			bs := uint8(7)
			t := uint64(buf[i+16] & 0x7F)
			for i < lb && buf[i+16]&0x80 == 0x80 {
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

					if i+16 >= lb {
						return 0, io.EOF
					}
					bs := uint8(7)
					t := uint64(buf[i+16] & 0x7F)
					for i < lb && buf[i+16]&0x80 == 0x80 {
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
						adjust := i + 16
						if adjust >= lb {
							return 0, io.EOF
						}
						ni, err := tt.UnmarshalSafe(buf[adjust:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 1 + 1:
					var tt DamageDirect

					{
						adjust := i + 16
						if adjust >= lb {
							return 0, io.EOF
						}
						ni, err := tt.UnmarshalSafe(buf[adjust:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 2 + 1:
					var tt HealOverTime

					{
						adjust := i + 16
						if adjust >= lb {
							return 0, io.EOF
						}
						ni, err := tt.UnmarshalSafe(buf[adjust:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 3 + 1:
					var tt DamageOverTime

					{
						adjust := i + 16
						if adjust >= lb {
							return 0, io.EOF
						}
						ni, err := tt.UnmarshalSafe(buf[adjust:])
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
