package sector

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

func (d *Connection) Size() (s uint64) {

	{
		s += d.Coord.Size()
	}
	{
		s += d.External.Size()
	}
	return
}

func (d *Connection) Marshal(buf []byte) ([]byte, error) {
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
		nbuf, err := d.Coord.Marshal(buf[0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		nbuf, err := d.External.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *Connection) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		ni, err := d.Coord.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		ni, err := d.External.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

func (d *Connection) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{
		adjust := i + 0
		if adjust >= lb {
			return 0, io.EOF
		}
		ni, err := d.Coord.UnmarshalSafe(buf[adjust:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		adjust := i + 0
		if adjust >= lb {
			return 0, io.EOF
		}
		ni, err := d.External.UnmarshalSafe(buf[adjust:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

func (d *S) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += d.Dim.Size()
	}
	{
		l := uint64(len(d.Connections))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Connections {
			_ = k0 // make compiler happy in case k is unused

			{
				s += d.Connections[k0].Size()
			}

		}

	}
	return
}

func (d *S) Marshal(buf []byte) ([]byte, error) {
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
		nbuf, err := d.Dim.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		l := uint64(len(d.Connections))

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
		for k0 := range d.Connections {

			{
				nbuf, err := d.Connections[k0].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *S) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{
		ni, err := d.Dim.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
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
		if uint64(cap(d.Connections)) >= l {
			d.Connections = d.Connections[:l]
		} else {
			d.Connections = make([]Connection, l)
		}
		for k0 := range d.Connections {

			{
				ni, err := d.Connections[k0].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

func (d *S) UnmarshalSafe(buf []byte) (uint64, error) {
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
		adjust := i + 0
		if adjust >= lb {
			return 0, io.EOF
		}
		ni, err := d.Dim.UnmarshalSafe(buf[adjust:])
		if err != nil {
			return 0, err
		}
		i += ni
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
		if uint64(cap(d.Connections)) >= l {
			d.Connections = d.Connections[:l]
		} else {
			d.Connections = make([]Connection, l)
		}
		for k0 := range d.Connections {

			{
				adjust := i + 0
				if adjust >= lb {
					return 0, io.EOF
				}
				ni, err := d.Connections[k0].UnmarshalSafe(buf[adjust:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}
