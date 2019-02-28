package message

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// MakeActor makes actor
func MakeActor(id, nodID int64) []byte {
	builder := flatbuffers.NewBuilder(0)

	fbs.ActorStart(builder)
	fbs.ActorAddId(builder, id)
	fbs.ActorAddNodeId(builder, nodID)
	builder.Finish(fbs.ActorEnd(builder))

	return builder.Bytes[builder.Head():]
}

// NewLogin makes Login
func NewLogin(b []byte) *fbs.Login {
	return fbs.GetRootAsLogin(b, 0)
}

// MakeLogin makes Login
func MakeLogin(id int64) (bytes []byte) {
	builder := flatbuffers.NewBuilder(0)

	fbs.LoginStart(builder)

	fbs.LoginAddId(builder, id)

	builder.Finish(fbs.LoginEnd(builder))

	bytes = builder.Bytes[builder.Head():]
	return bytes
}

// NewLoginResult creates login result.
func NewLoginResult(b []byte) *fbs.LoginResult {
	return fbs.GetRootAsLoginResult(b, 0)
}

// MakeLoginResult makes login result message.
func MakeLoginResult(id int64) (bytes []byte) {
	builder := flatbuffers.NewBuilder(0)

	fbs.LoginResultStart(builder)

	fbs.LoginResultAddId(builder, id)

	builder.Finish(fbs.LoginResultEnd(builder))

	bytes = builder.Bytes[builder.Head():]
	return bytes
}

// NewNode creates node.
func NewNode(b []byte) *fbs.Node {
	return fbs.GetRootAsNode(b, 0)
}

// MakeNode makes node message.
func MakeNode(id int64, position vector.Vector) []byte {
	builder := flatbuffers.NewBuilder(0)

	fbs.NodeStart(builder)

	fbs.NodeAddID(builder, id)
	fbs.NodeAddPosition(builder, fbs.CreateVector(builder, position.X, position.Y))

	builder.Finish(fbs.NodeEnd(builder))

	return builder.Bytes[builder.Head():]
}
