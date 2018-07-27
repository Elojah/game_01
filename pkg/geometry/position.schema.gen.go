package geometry

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

func (d *Vec3) Size() (s uint64) {

	s += 24
	return
}

func (d *Vec3) Marshal(buf []byte) ([]byte, error) {
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

		*(*float64)(unsafe.Pointer(&buf[0])) = d.X

	}
	{

		*(*float64)(unsafe.Pointer(&buf[8])) = d.Y

	}
	{

		*(*float64)(unsafe.Pointer(&buf[16])) = d.Z

	}
	return buf[:i+24], nil
}

func (d *Vec3) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.X = *(*float64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Y = *(*float64)(unsafe.Pointer(&buf[8]))

	}
	{

		d.Z = *(*float64)(unsafe.Pointer(&buf[16]))

	}
	return i + 24, nil
}

func (d *Vec3) UnmarshalSafe(buf []byte) (uint64, error) {
	lb := uint64(len(buf))
	if lb < d.Size() {
		return 0, io.EOF
	}
	i := uint64(0)

	{

		d.X = *(*float64)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Y = *(*float64)(unsafe.Pointer(&buf[8]))

	}
	{

		d.Z = *(*float64)(unsafe.Pointer(&buf[16]))

	}
	return i + 24, nil
}

func (d *Position) Size() (s uint64) {

	{
		s += d.Coord.Size()
	}
	{
		s += 16
	}
	return
}

func (d *Position) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.SectorID[:])
		i += 16
	}
	return buf[:i+0], nil
}

func (d *Position) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		ni, err := d.Coord.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		copy(d.SectorID[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}

func (d *Position) UnmarshalSafe(buf []byte) (uint64, error) {
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
		if i+0 >= lb {
			return 0, io.EOF
		}
		copy(d.SectorID[:], buf[i+0:])
		i += 16
	}
	return i + 0, nil
}
