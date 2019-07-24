// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package traceflatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StringValue struct {
	_tab flatbuffers.Table
}

func GetRootAsStringValue(buf []byte, offset flatbuffers.UOffsetT) *StringValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StringValue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *StringValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StringValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StringValue) StringValue() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StringValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StringValueAddStringValue(builder *flatbuffers.Builder, stringValue flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(stringValue), 0)
}
func StringValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
