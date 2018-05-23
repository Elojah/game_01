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

type HealDirectFeedback struct {
	Amount int64
}

func (d *HealDirectFeedback) Size() (s uint64) {

	s += 8
	return
}
func (d *HealDirectFeedback) Marshal(buf []byte) ([]byte, error) {
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
	return buf[:i+8], nil
}

func (d *HealDirectFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	return i + 8, nil
}

type DamageDirectFeedback struct {
	Amount int64
}

func (d *DamageDirectFeedback) Size() (s uint64) {

	s += 8
	return
}
func (d *DamageDirectFeedback) Marshal(buf []byte) ([]byte, error) {
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
	return buf[:i+8], nil
}

func (d *DamageDirectFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Amount = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	return i + 8, nil
}

type HealOverTimeFeedback struct {
}

func (d *HealOverTimeFeedback) Size() (s uint64) {

	return
}
func (d *HealOverTimeFeedback) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	return buf[:i+0], nil
}

func (d *HealOverTimeFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	return i + 0, nil
}

type DamageOverTimeFeedback struct {
}

func (d *DamageOverTimeFeedback) Size() (s uint64) {

	return
}
func (d *DamageOverTimeFeedback) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	return buf[:i+0], nil
}

func (d *DamageOverTimeFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	return i + 0, nil
}

type AbilityFeedback struct {
	ID         [16]byte
	AbilityID  [16]byte
	Components []interface{}
}

func (d *AbilityFeedback) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += 16
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

				case HealDirectFeedback:
					v = 0 + 1

				case DamageDirectFeedback:
					v = 1 + 1

				case HealOverTimeFeedback:
					v = 2 + 1

				case DamageOverTimeFeedback:
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

				case HealDirectFeedback:

					{
						s += tt.Size()
					}

				case DamageDirectFeedback:

					{
						s += tt.Size()
					}

				case HealOverTimeFeedback:

					{
						s += tt.Size()
					}

				case DamageOverTimeFeedback:

					{
						s += tt.Size()
					}

				}
			}

		}

	}
	return
}
func (d *AbilityFeedback) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.AbilityID[:])
		i += 16
	}
	{
		l := uint64(len(d.Components))

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
		for k0 := range d.Components {

			{
				var v uint64
				switch d.Components[k0].(type) {

				case HealDirectFeedback:
					v = 0 + 1

				case DamageDirectFeedback:
					v = 1 + 1

				case HealOverTimeFeedback:
					v = 2 + 1

				case DamageOverTimeFeedback:
					v = 3 + 1

				}

				{

					t := uint64(v)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				switch tt := d.Components[k0].(type) {

				case HealDirectFeedback:

					{
						nbuf, err := tt.Marshal(buf[i+0:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				case DamageDirectFeedback:

					{
						nbuf, err := tt.Marshal(buf[i+0:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				case HealOverTimeFeedback:

					{
						nbuf, err := tt.Marshal(buf[i+0:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				case DamageOverTimeFeedback:

					{
						nbuf, err := tt.Marshal(buf[i+0:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}

				}
			}

		}
	}
	return buf[:i+0], nil
}

func (d *AbilityFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		copy(d.AbilityID[:], buf[i+0:])
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
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					v = t

				}
				switch v {

				case 0 + 1:
					var tt HealDirectFeedback

					{
						ni, err := tt.Unmarshal(buf[i+0:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 1 + 1:
					var tt DamageDirectFeedback

					{
						ni, err := tt.Unmarshal(buf[i+0:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 2 + 1:
					var tt HealOverTimeFeedback

					{
						ni, err := tt.Unmarshal(buf[i+0:])
						if err != nil {
							return 0, err
						}
						i += ni
					}

					d.Components[k0] = tt

				case 3 + 1:
					var tt DamageOverTimeFeedback

					{
						ni, err := tt.Unmarshal(buf[i+0:])
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
	return i + 0, nil
}
