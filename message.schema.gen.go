package game

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

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

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

		v := *(*uint64)(unsafe.Pointer(&(d.X)))

		buf[0+0] = byte(v >> 0)

		buf[1+0] = byte(v >> 8)

		buf[2+0] = byte(v >> 16)

		buf[3+0] = byte(v >> 24)

		buf[4+0] = byte(v >> 32)

		buf[5+0] = byte(v >> 40)

		buf[6+0] = byte(v >> 48)

		buf[7+0] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Y)))

		buf[0+8] = byte(v >> 0)

		buf[1+8] = byte(v >> 8)

		buf[2+8] = byte(v >> 16)

		buf[3+8] = byte(v >> 24)

		buf[4+8] = byte(v >> 32)

		buf[5+8] = byte(v >> 40)

		buf[6+8] = byte(v >> 48)

		buf[7+8] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Z)))

		buf[0+16] = byte(v >> 0)

		buf[1+16] = byte(v >> 8)

		buf[2+16] = byte(v >> 16)

		buf[3+16] = byte(v >> 24)

		buf[4+16] = byte(v >> 32)

		buf[5+16] = byte(v >> 40)

		buf[6+16] = byte(v >> 48)

		buf[7+16] = byte(v >> 56)

	}
	return buf[:i+24], nil
}

func (d *Vec3) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		v := 0 | (uint64(buf[0+0]) << 0) | (uint64(buf[1+0]) << 8) | (uint64(buf[2+0]) << 16) | (uint64(buf[3+0]) << 24) | (uint64(buf[4+0]) << 32) | (uint64(buf[5+0]) << 40) | (uint64(buf[6+0]) << 48) | (uint64(buf[7+0]) << 56)
		d.X = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+8]) << 0) | (uint64(buf[1+8]) << 8) | (uint64(buf[2+8]) << 16) | (uint64(buf[3+8]) << 24) | (uint64(buf[4+8]) << 32) | (uint64(buf[5+8]) << 40) | (uint64(buf[6+8]) << 48) | (uint64(buf[7+8]) << 56)
		d.Y = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+16]) << 0) | (uint64(buf[1+16]) << 8) | (uint64(buf[2+16]) << 16) | (uint64(buf[3+16]) << 24) | (uint64(buf[4+16]) << 32) | (uint64(buf[5+16]) << 40) | (uint64(buf[6+16]) << 48) | (uint64(buf[7+16]) << 56)
		d.Z = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 24, nil
}

type Actor struct {
	ID       [16]byte
	HP       uint8
	MP       uint8
	Position Vec3
}

func (d *Actor) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += d.Position.Size()
	}
	s += 2
	return
}
func (d *Actor) Marshal(buf []byte) ([]byte, error) {
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

		buf[i+0+0] = byte(d.HP >> 0)

	}
	{

		buf[i+0+1] = byte(d.MP >> 0)

	}
	{
		nbuf, err := d.Position.Marshal(buf[i+2:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+2], nil
}

func (d *Actor) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.ID[:], buf[i+0:])
		i += 16
	}
	{

		d.HP = 0 | (uint8(buf[i+0+0]) << 0)

	}
	{

		d.MP = 0 | (uint8(buf[i+0+1]) << 0)

	}
	{
		ni, err := d.Position.Unmarshal(buf[i+2:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 2, nil
}

type ActorPatch struct {
	HP       *uint8
	MP       *uint8
	Position *Vec3
}

func (d *ActorPatch) Size() (s uint64) {

	{
		if d.HP != nil {

			s += 1
		}
	}
	{
		if d.MP != nil {

			s += 1
		}
	}
	{
		if d.Position != nil {

			{
				s += (*d.Position).Size()
			}
			s += 0
		}
	}
	s += 3
	return
}
func (d *ActorPatch) Marshal(buf []byte) ([]byte, error) {
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
		if d.HP == nil {
			buf[0] = 0
		} else {
			buf[0] = 1

			{

				buf[0+1] = byte((*d.HP) >> 0)

			}
			i += 1
		}
	}
	{
		if d.MP == nil {
			buf[i+1] = 0
		} else {
			buf[i+1] = 1

			{

				buf[i+0+2] = byte((*d.MP) >> 0)

			}
			i += 1
		}
	}
	{
		if d.Position == nil {
			buf[i+2] = 0
		} else {
			buf[i+2] = 1

			{
				nbuf, err := (*d.Position).Marshal(buf[i+3:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
			i += 0
		}
	}
	return buf[:i+3], nil
}

func (d *ActorPatch) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		if buf[i+0] == 1 {
			if d.HP == nil {
				d.HP = new(uint8)
			}

			{

				(*d.HP) = 0 | (uint8(buf[i+0+1]) << 0)

			}
			i += 1
		} else {
			d.HP = nil
		}
	}
	{
		if buf[i+1] == 1 {
			if d.MP == nil {
				d.MP = new(uint8)
			}

			{

				(*d.MP) = 0 | (uint8(buf[i+0+2]) << 0)

			}
			i += 1
		} else {
			d.MP = nil
		}
	}
	{
		if buf[i+2] == 1 {
			if d.Position == nil {
				d.Position = new(Vec3)
			}

			{
				ni, err := (*d.Position).Unmarshal(buf[i+3:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			i += 0
		} else {
			d.Position = nil
		}
	}
	return i + 3, nil
}

type ActorSubset struct {
	IDs [][16]byte
}

func (d *ActorSubset) Size() (s uint64) {

	{
		l := uint64(len(d.IDs))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for _ = range d.IDs {

			{
				s += 16
			}

		}

	}
	return
}
func (d *ActorSubset) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.IDs))

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
		for k0 := range d.IDs {

			{
				copy(buf[i+0:], d.IDs[k0][:])
				i += 16
			}

		}
	}
	return buf[:i+0], nil
}

func (d *ActorSubset) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

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
		if uint64(cap(d.IDs)) >= l {
			d.IDs = d.IDs[:l]
		} else {
			d.IDs = make([][16]byte, l)
		}
		for k0 := range d.IDs {

			{
				copy(d.IDs[k0][:], buf[i+0:])
				i += 16
			}

		}
	}
	return i + 0, nil
}

type ActorCreate struct {
	Token  [16]byte
	Actors []Actor
}

func (d *ActorCreate) Size() (s uint64) {

	{
		s += 16
	}
	{
		l := uint64(len(d.Actors))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Actors {

			{
				s += d.Actors[k0].Size()
			}

		}

	}
	return
}
func (d *ActorCreate) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Token[:])
		i += 16
	}
	{
		l := uint64(len(d.Actors))

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
		for k0 := range d.Actors {

			{
				nbuf, err := d.Actors[k0].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *ActorCreate) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Token[:], buf[i+0:])
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
		if uint64(cap(d.Actors)) >= l {
			d.Actors = d.Actors[:l]
		} else {
			d.Actors = make([]Actor, l)
		}
		for k0 := range d.Actors {

			{
				ni, err := d.Actors[k0].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

type ActorUpdate struct {
	Token  [16]byte
	Subset ActorSubset
	Patch  ActorPatch
}

func (d *ActorUpdate) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += d.Subset.Size()
	}
	{
		s += d.Patch.Size()
	}
	return
}
func (d *ActorUpdate) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Token[:])
		i += 16
	}
	{
		nbuf, err := d.Subset.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		nbuf, err := d.Patch.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *ActorUpdate) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	{
		ni, err := d.Subset.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		ni, err := d.Patch.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type ActorDelete struct {
	Token  [16]byte
	Subset ActorSubset
}

func (d *ActorDelete) Size() (s uint64) {

	{
		s += 16
	}
	{
		s += d.Subset.Size()
	}
	return
}
func (d *ActorDelete) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i+0:], d.Token[:])
		i += 16
	}
	{
		nbuf, err := d.Subset.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *ActorDelete) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Token[:], buf[i+0:])
		i += 16
	}
	{
		ni, err := d.Subset.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type Message struct {
	Val interface{}
}

func (d *Message) Size() (s uint64) {

	{
		var v uint64
		switch d.Val.(type) {

		case ActorCreate:
			v = 0 + 1

		case ActorUpdate:
			v = 1 + 1

		case ActorDelete:
			v = 2 + 1

		}

		{

			t := v
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		switch tt := d.Val.(type) {

		case ActorCreate:

			{
				s += tt.Size()
			}

		case ActorUpdate:

			{
				s += tt.Size()
			}

		case ActorDelete:

			{
				s += tt.Size()
			}

		}
	}
	return
}
func (d *Message) Marshal(buf []byte) ([]byte, error) {
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
		var v uint64
		switch d.Val.(type) {

		case ActorCreate:
			v = 0 + 1

		case ActorUpdate:
			v = 1 + 1

		case ActorDelete:
			v = 2 + 1

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
		switch tt := d.Val.(type) {

		case ActorCreate:

			{
				nbuf, err := tt.Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case ActorUpdate:

			{
				nbuf, err := tt.Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		case ActorDelete:

			{
				nbuf, err := tt.Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *Message) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

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
			var tt ActorCreate

			{
				ni, err := tt.Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Val = tt

		case 1 + 1:
			var tt ActorUpdate

			{
				ni, err := tt.Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Val = tt

		case 2 + 1:
			var tt ActorDelete

			{
				ni, err := tt.Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

			d.Val = tt

		default:
			d.Val = nil
		}
	}
	return i + 0, nil
}
