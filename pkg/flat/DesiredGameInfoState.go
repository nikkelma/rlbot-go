// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DesiredGameInfoState struct {
	_tab flatbuffers.Table
}

func GetRootAsDesiredGameInfoState(buf []byte, offset flatbuffers.UOffsetT) *DesiredGameInfoState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DesiredGameInfoState{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DesiredGameInfoState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DesiredGameInfoState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DesiredGameInfoState) WorldGravityZ(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredGameInfoState) GameSpeed(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DesiredGameInfoStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DesiredGameInfoStateAddWorldGravityZ(builder *flatbuffers.Builder, worldGravityZ flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(worldGravityZ), 0)
}
func DesiredGameInfoStateAddGameSpeed(builder *flatbuffers.Builder, gameSpeed flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(gameSpeed), 0)
}
func DesiredGameInfoStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
