package account

import (
	"errors"
	"io"
	"net"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

func (d *Token) Size() (s uint64) {
	{
		s += 16
	}
	{
		l := uint64(len(d.IP.String()))
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
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	{
		s += 16
	}
	s += 8
	return
}

func (d *Token) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.IP.String()))
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
		d.IP, _ = net.ResolveUDPAddr("udp", string(buf[i:]))
		i += l
	}
	{
		copy(buf[i:], d.Account[:])
		i += 16
	}
	{
		buf[i] = byte(d.Ping >> 0)
		buf[i+1] = byte(d.Ping >> 8)
		buf[i+2] = byte(d.Ping >> 16)
		buf[i+3] = byte(d.Ping >> 24)
		buf[i+4] = byte(d.Ping >> 32)
		buf[i+5] = byte(d.Ping >> 40)
		buf[i+6] = byte(d.Ping >> 48)
		buf[i+7] = byte(d.Ping >> 56)
	}
	{
		copy(buf[i+8:], d.CorePool[:])
		i += 16
	}
	{
		copy(buf[i+8:], d.SyncPool[:])
		i += 16
	}
	{
		copy(buf[i+8:], d.PC[:])
		i += 16
	}
	{
		copy(buf[i+8:], d.Entity[:])
		i += 16
	}
	return buf[:i+8], nil
}

func (d *Token) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
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
		d.IP, _ = net.ResolveUDPAddr("udp", string(buf[i:i+l]))
		i += l
	}
	{
		copy(d.Account[:], buf[i:])
		i += 16
	}
	{
		d.Ping = 0 | (uint64(buf[i]) << 0) | (uint64(buf[i+1]) << 8) | (uint64(buf[i+2]) << 16) | (uint64(buf[i+3]) << 24) | (uint64(buf[i+4]) << 32) | (uint64(buf[i+5]) << 40) | (uint64(buf[i+6]) << 48) | (uint64(buf[i+7]) << 56)
	}
	{
		copy(d.CorePool[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.SyncPool[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.PC[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.Entity[:], buf[i+8:])
		i += 16
	}
	return i + 8, nil
}

func (d *Token) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < 104 {
		return 0, errors.New("invalid buffer")
	}
	i := uint64(0)
	{
		copy(d.ID[:], buf[i:])
		i += 16
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for i < lb && buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		if i+l >= lb {
			return 0, errors.New("invalid buffer")
		}
		d.IP, _ = net.ResolveUDPAddr("udp", string(buf[i:i+l]))
		i += l
	}
	if lb-i < 72 {
		return 0, errors.New("invalid buffer")
	}
	{
		copy(d.Account[:], buf[i:])
		i += 16
	}
	{
		d.Ping = 0 | (uint64(buf[i]) << 0) | (uint64(buf[i+1]) << 8) | (uint64(buf[i+2]) << 16) | (uint64(buf[i+3]) << 24) | (uint64(buf[i+4]) << 32) | (uint64(buf[i+5]) << 40) | (uint64(buf[i+6]) << 48) | (uint64(buf[i+7]) << 56)
	}
	{
		copy(d.CorePool[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.SyncPool[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.PC[:], buf[i+8:])
		i += 16
	}
	{
		copy(d.Entity[:], buf[i+8:])
		i += 16
	}
	return i + 8, nil
}
