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

// MakeLogin makes Login
func MakeLogin(b []byte) *fbs.Login {
	return fbs.GetRootAsLogin(b, 0)
}

// MakeLoginMessage makes Login
func MakeLoginMessage(id int64) (bytes []byte) {
	builder := flatbuffers.NewBuilder(0)

	fbs.LoginStart(builder)

	fbs.LoginAddId(builder, id)

	builder.Finish(fbs.LoginEnd(builder))

	bytes = builder.Bytes[builder.Head():]
	return bytes
}

// MakeLoginResult makes login result.
func MakeLoginResult(b []byte) *fbs.LoginResult {
	return fbs.GetRootAsLoginResult(b, 0)
}

// MakeLoginResultMessage makes login result message.
func MakeLoginResultMessage(id int64) (bytes []byte) {
	builder := flatbuffers.NewBuilder(0)

	fbs.LoginResultStart(builder)

	fbs.LoginResultAddId(builder, id)

	builder.Finish(fbs.LoginResultEnd(builder))

	bytes = builder.Bytes[builder.Head():]
	return bytes
}

// MakeNode makes node message.
func MakeNode(id int64, position vector.Vector) []byte {
	builder := flatbuffers.NewBuilder(0)

	fbs.NodeStart(builder)

	fbs.NodeAddId(builder, id)
	fbs.NodeAddPosition(builder, fbs.CreateVector(builder, position.X, position.Y))

	builder.Finish(fbs.NodeEnd(builder))

	return builder.Bytes[builder.Head():]
}
