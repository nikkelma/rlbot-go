// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PlayerInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerInfo(buf []byte, offset flatbuffers.UOffsetT) *PlayerInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PlayerInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PlayerInfo) Physics(obj *Physics) *Physics {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Physics)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PlayerInfo) ScoreInfo(obj *ScoreInfo) *ScoreInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ScoreInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PlayerInfo) IsDemolished() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInfo) MutateIsDemolished(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

/// True if your wheels are on the ground, the wall, or the ceiling. False if you're midair or turtling.
func (rcv *PlayerInfo) HasWheelContact() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

/// True if your wheels are on the ground, the wall, or the ceiling. False if you're midair or turtling.
func (rcv *PlayerInfo) MutateHasWheelContact(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *PlayerInfo) IsSupersonic() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInfo) MutateIsSupersonic(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func (rcv *PlayerInfo) IsBot() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInfo) MutateIsBot(n byte) bool {
	return rcv._tab.MutateByteSlot(14, n)
}

/// True if the player has jumped. Falling off the ceiling / driving off the goal post does not count.
func (rcv *PlayerInfo) Jumped() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

/// True if the player has jumped. Falling off the ceiling / driving off the goal post does not count.
func (rcv *PlayerInfo) MutateJumped(n byte) bool {
	return rcv._tab.MutateByteSlot(16, n)
}

///  True if player has double jumped. False does not mean you have a jump remaining, because the
///  aerial timer can run out, and that doesn't affect this flag.
func (rcv *PlayerInfo) DoubleJumped() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///  True if player has double jumped. False does not mean you have a jump remaining, because the
///  aerial timer can run out, and that doesn't affect this flag.
func (rcv *PlayerInfo) MutateDoubleJumped(n byte) bool {
	return rcv._tab.MutateByteSlot(18, n)
}

func (rcv *PlayerInfo) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PlayerInfo) Team() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInfo) MutateTeam(n int32) bool {
	return rcv._tab.MutateInt32Slot(22, n)
}

func (rcv *PlayerInfo) Boost() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInfo) MutateBoost(n int32) bool {
	return rcv._tab.MutateInt32Slot(24, n)
}

func (rcv *PlayerInfo) Hitbox(obj *BoxShape) *BoxShape {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BoxShape)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PlayerInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func PlayerInfoAddPhysics(builder *flatbuffers.Builder, physics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(physics), 0)
}
func PlayerInfoAddScoreInfo(builder *flatbuffers.Builder, scoreInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(scoreInfo), 0)
}
func PlayerInfoAddIsDemolished(builder *flatbuffers.Builder, isDemolished byte) {
	builder.PrependByteSlot(2, isDemolished, 0)
}
func PlayerInfoAddHasWheelContact(builder *flatbuffers.Builder, hasWheelContact byte) {
	builder.PrependByteSlot(3, hasWheelContact, 0)
}
func PlayerInfoAddIsSupersonic(builder *flatbuffers.Builder, isSupersonic byte) {
	builder.PrependByteSlot(4, isSupersonic, 0)
}
func PlayerInfoAddIsBot(builder *flatbuffers.Builder, isBot byte) {
	builder.PrependByteSlot(5, isBot, 0)
}
func PlayerInfoAddJumped(builder *flatbuffers.Builder, jumped byte) {
	builder.PrependByteSlot(6, jumped, 0)
}
func PlayerInfoAddDoubleJumped(builder *flatbuffers.Builder, doubleJumped byte) {
	builder.PrependByteSlot(7, doubleJumped, 0)
}
func PlayerInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(name), 0)
}
func PlayerInfoAddTeam(builder *flatbuffers.Builder, team int32) {
	builder.PrependInt32Slot(9, team, 0)
}
func PlayerInfoAddBoost(builder *flatbuffers.Builder, boost int32) {
	builder.PrependInt32Slot(10, boost, 0)
}
func PlayerInfoAddHitbox(builder *flatbuffers.Builder, hitbox flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(hitbox), 0)
}
func PlayerInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
