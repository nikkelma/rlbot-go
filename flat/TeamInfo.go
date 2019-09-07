// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TeamInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsTeamInfo(buf []byte, offset flatbuffers.UOffsetT) *TeamInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TeamInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TeamInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TeamInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TeamInfo) TeamIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TeamInfo) MutateTeamIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

/// number of goals scored.
func (rcv *TeamInfo) Score() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// number of goals scored.
func (rcv *TeamInfo) MutateScore(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func TeamInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TeamInfoAddTeamIndex(builder *flatbuffers.Builder, teamIndex int32) {
	builder.PrependInt32Slot(0, teamIndex, 0)
}
func TeamInfoAddScore(builder *flatbuffers.Builder, score int32) {
	builder.PrependInt32Slot(1, score, 0)
}
func TeamInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}