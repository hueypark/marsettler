package message

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/hueypark/marsettler/game/message/fbs"
)

var builder *flatbuffers.Builder

func init() {
	builder = flatbuffers.NewBuilder(0)
}

// MakeActor makes actor
func MakeActor(id int64, posX, posY float32) []byte {
	builder.Reset()

	fbs.ActorStart(builder)
	fbs.ActorAddId(builder, id)
	fbs.ActorAddPosition(builder, fbs.CreateVector(builder, posX, posY))
	builder.Finish(fbs.ActorEnd(builder))

	return builder.Bytes[builder.Head():]
}

// MakeLogin makes Login
func MakeLogin(b []byte) fbs.Login {
	var login fbs.Login
	login.Init(b, 0)

	return login
}

// MakeLoginResult makes LoginResult
func MakeLoginResult(id int64) (bytes []byte, size int) {
	builder.Reset()
	fbs.CreateLoginResult(builder, id)
	return builder.Bytes, len(builder.Bytes)
}
