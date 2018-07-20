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
		buf[0] = byte(d.Amount >> 0)
		buf[1] = byte(d.Amount >> 8)
		buf[2] = byte(d.Amount >> 16)
		buf[3] = byte(d.Amount >> 24)
		buf[4] = byte(d.Amount >> 32)
		buf[5] = byte(d.Amount >> 40)
		buf[6] = byte(d.Amount >> 48)
		buf[7] = byte(d.Amount >> 56)
	}
	return buf[:i+8], nil
}
func (d *HealDirectFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		d.Amount = 0 | (int64(buf[0]) << 0) | (int64(buf[1]) << 8) | (int64(buf[2]) << 16) | (int64(buf[3]) << 24) | (int64(buf[4]) << 32) | (int64(buf[5]) << 40) | (int64(buf[6]) << 48) | (int64(buf[7]) << 56)
	}
	return i + 8, nil
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
		buf[0] = byte(d.Amount >> 0)
		buf[1] = byte(d.Amount >> 8)
		buf[2] = byte(d.Amount >> 16)
		buf[3] = byte(d.Amount >> 24)
		buf[4] = byte(d.Amount >> 32)
		buf[5] = byte(d.Amount >> 40)
		buf[6] = byte(d.Amount >> 48)
		buf[7] = byte(d.Amount >> 56)
	}
	return buf[:i+8], nil
}
func (d *DamageDirectFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		d.Amount = 0 | (int64(buf[0]) << 0) | (int64(buf[1]) << 8) | (int64(buf[2]) << 16) | (int64(buf[3]) << 24) | (int64(buf[4]) << 32) | (int64(buf[5]) << 40) | (int64(buf[6]) << 48) | (int64(buf[7]) << 56)
	}
	return i + 8, nil
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
	return buf[:i], nil
}
func (d *HealOverTimeFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	return i + 0, nil
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
	return buf[:i], nil
}
func (d *DamageOverTimeFeedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
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
		copy(buf[i:], d.ID[:])
		i += 16
	}
	{
		copy(buf[i:], d.AbilityID[:])
		i += 16
	}
	{
		l := uint64(len(d.Components))
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
						buf[i] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i] = byte(t)
					i++
				}
				switch tt := d.Components[k0].(type) {
				case HealDirectFeedback:
					{
						nbuf, err := tt.Marshal(buf[i:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
				case DamageDirectFeedback:
					{
						nbuf, err := tt.Marshal(buf[i:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
				case HealOverTimeFeedback:
					{
						nbuf, err := tt.Marshal(buf[i:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
				case DamageOverTimeFeedback:
					{
						nbuf, err := tt.Marshal(buf[i:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
				}
			}
		}
	}
	return buf[:i], nil
}
func (d *Feedback) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		copy(d.AbilityID[:], buf[i:])
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
		if uint64(cap(d.Components)) >= l {
			d.Components = d.Components[:l]
		} else {
			d.Components = make([]FeedbackComponent, l)
		}
		for k0 := range d.Components {
			{
				v := uint64(0)
				{
					bs := uint8(7)
					t := uint64(buf[i] & 0x7F)
					for buf[i]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i]&0x7F) << bs
						bs += 7
					}
					i++
					v = t
				}
				switch v {
				case 0 + 1:
					var tt HealDirectFeedback
					{
						ni, err := tt.Unmarshal(buf[i:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					d.Components[k0] = tt
				case 1 + 1:
					var tt DamageDirectFeedback
					{
						ni, err := tt.Unmarshal(buf[i:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					d.Components[k0] = tt
				case 2 + 1:
					var tt HealOverTimeFeedback
					{
						ni, err := tt.Unmarshal(buf[i:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					d.Components[k0] = tt
				case 3 + 1:
					var tt DamageOverTimeFeedback
					{
						ni, err := tt.Unmarshal(buf[i:])
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
