// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Login struct {
	_tab flatbuffers.Table
}

func GetRootAsLogin(buf []byte, offset flatbuffers.UOffsetT) *Login {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Login{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Login) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Login) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Login) Id() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Login) MutateId(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func LoginStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func LoginAddId(builder *flatbuffers.Builder, Id int64) {
	builder.PrependInt64Slot(0, Id, 0)
}
func LoginEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}