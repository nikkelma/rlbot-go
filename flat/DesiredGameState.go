// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DesiredGameState struct {
	_tab flatbuffers.Table
}

func GetRootAsDesiredGameState(buf []byte, offset flatbuffers.UOffsetT) *DesiredGameState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DesiredGameState{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DesiredGameState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DesiredGameState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DesiredGameState) BallState(obj *DesiredBallState) *DesiredBallState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DesiredBallState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredGameState) CarStates(obj *DesiredCarState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DesiredGameState) CarStatesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DesiredGameState) BoostStates(obj *DesiredBoostState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DesiredGameState) BoostStatesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DesiredGameState) GameInfoState(obj *DesiredGameInfoState) *DesiredGameInfoState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DesiredGameInfoState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredGameState) ConsoleCommands(obj *ConsoleCommand, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DesiredGameState) ConsoleCommandsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DesiredGameStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DesiredGameStateAddBallState(builder *flatbuffers.Builder, ballState flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ballState), 0)
}
func DesiredGameStateAddCarStates(builder *flatbuffers.Builder, carStates flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(carStates), 0)
}
func DesiredGameStateStartCarStatesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DesiredGameStateAddBoostStates(builder *flatbuffers.Builder, boostStates flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(boostStates), 0)
}
func DesiredGameStateStartBoostStatesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DesiredGameStateAddGameInfoState(builder *flatbuffers.Builder, gameInfoState flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(gameInfoState), 0)
}
func DesiredGameStateAddConsoleCommands(builder *flatbuffers.Builder, consoleCommands flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(consoleCommands), 0)
}
func DesiredGameStateStartConsoleCommandsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DesiredGameStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
