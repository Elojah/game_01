// automatically generated by the FlatBuffers compiler, do not modify

package game

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Position struct {
	_tab flatbuffers.Struct
}

func (rcv *Position) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Position) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Position) X() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Position) MutateX(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Position) Y() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Position) MutateY(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func (rcv *Position) Z() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(16))
}
func (rcv *Position) MutateZ(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(16), n)
}

func CreatePosition(builder *flatbuffers.Builder, x uint64, y uint64, z uint64) flatbuffers.UOffsetT {
	builder.Prep(8, 24)
	builder.PrependUint64(z)
	builder.PrependUint64(y)
	builder.PrependUint64(x)
	return builder.Offset()
}
