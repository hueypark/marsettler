package message

import flatbuffers "github.com/google/flatbuffers/go"

var builder *flatbuffers.Builder

func init() {
	builder = flatbuffers.NewBuilder(0)
}
