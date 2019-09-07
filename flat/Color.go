// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Color struct {
	_tab flatbuffers.Table
}

func GetRootAsColor(buf []byte, offset flatbuffers.UOffsetT) *Color {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Color{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Color) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Color) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Color) A() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateA(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *Color) R() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateR(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Color) G() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateG(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *Color) B() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateB(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func ColorStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ColorAddA(builder *flatbuffers.Builder, a byte) {
	builder.PrependByteSlot(0, a, 0)
}
func ColorAddR(builder *flatbuffers.Builder, r byte) {
	builder.PrependByteSlot(1, r, 0)
}
func ColorAddG(builder *flatbuffers.Builder, g byte) {
	builder.PrependByteSlot(2, g, 0)
}
func ColorAddB(builder *flatbuffers.Builder, b byte) {
	builder.PrependByteSlot(3, b, 0)
}
func ColorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
