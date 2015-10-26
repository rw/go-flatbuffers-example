// automatically generated, do not modify

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type User struct {
	_tab flatbuffers.Table
}

func GetRootAsUser(buf []byte, offset flatbuffers.UOffsetT) *User {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &User{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *User) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *User) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *User) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func UserStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func UserAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0) }
func UserAddId(builder *flatbuffers.Builder, id uint64) { builder.PrependUint64Slot(1, id, 0) }
func UserEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
