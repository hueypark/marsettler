package game

import (
	"testing"

	"github.com/hueypark/marsettler/data"
)

func BenchmarkNewActor(b *testing.B) {
	kingdom := NewKingdom()
	w := NewWorld()
	startNodeID, err := w.StartNodeID()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.NewActor(kingdom.ID(), startNodeID, data.Hero)
	}
}
