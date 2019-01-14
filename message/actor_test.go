package message

import (
	"testing"

	"github.com/hueypark/marsettler/message/fbs"
)

func TestActor(t *testing.T) {
	var id int64 = 1
	var posX float32 = 100.0
	var posY float32 = 200.0
	actorBytes := MakeActor(id, posX, posY)

	actor := fbs.GetRootAsActor(actorBytes, 0)
	pos := actor.Position(nil)
	if actor.Id() != id || pos.X() != posX || pos.Y() != posY {
		t.Error("Value not equal.")
	}
}
