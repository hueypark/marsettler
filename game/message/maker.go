package message

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// MakeActor makes actor
func MakeActor(id int64, posX, posY float32) []byte {
	builder := flatbuffers.NewBuilder(0)

	fbs.ActorStart(builder)
	fbs.ActorAddId(builder, id)
	fbs.ActorAddPosition(builder, fbs.CreateVector(builder, posX, posY))
	builder.Finish(fbs.ActorEnd(builder))

	return builder.Bytes[builder.Head():]
}

// MakeLogin makes Login
func MakeLogin(b []byte) *fbs.Login {
	return fbs.GetRootAsLogin(b, 0)
}

// MakeLoginMessage makes Login
func MakeLoginMessage(id int64) (bytes []byte, size int) {
	builder := flatbuffers.NewBuilder(0)

	fbs.LoginStart(builder)

	fbs.LoginAddId(builder, id)

	builder.Finish(fbs.LoginEnd(builder))

	bytes = builder.Bytes[builder.Head():]
	return bytes, len(bytes)
}

// MakeLoginResult makes login result.
func MakeLoginResult(b []byte) *fbs.LoginResult {
	return fbs.GetRootAsLoginResult(b, 0)
}

// MakeLoginResultMessage makes login result message.
func MakeLoginResultMessage(id int64) (bytes []byte, size int) {
	builder := flatbuffers.NewBuilder(0)

	fbs.LoginResultStart(builder)

	fbs.LoginResultAddId(builder, id)

	builder.Finish(fbs.LoginResultEnd(builder))

	bytes = builder.Bytes[builder.Head():]
	return bytes, len(bytes)
}
