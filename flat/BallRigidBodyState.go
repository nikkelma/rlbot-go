// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Rigid body state for the ball.
type BallRigidBodyState struct {
	_tab flatbuffers.Table
}

func GetRootAsBallRigidBodyState(buf []byte, offset flatbuffers.UOffsetT) *BallRigidBodyState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BallRigidBodyState{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BallRigidBodyState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BallRigidBodyState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BallRigidBodyState) State(obj *RigidBodyState) *RigidBodyState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(RigidBodyState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func BallRigidBodyStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BallRigidBodyStateAddState(builder *flatbuffers.Builder, state flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(state), 0)
}
func BallRigidBodyStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
