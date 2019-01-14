package message

import "github.com/hueypark/marsettler/message/fbs"

func MakeActor(id int64, posX, posY float32) []byte {
	builder.Reset()

	fbs.ActorStart(builder)
	fbs.ActorAddId(builder, id)
	fbs.ActorAddPosition(builder, fbs.CreateVector(builder, posX, posY))
	builder.Finish(fbs.ActorEnd(builder))

	return builder.Bytes[builder.Head():]
}
