package game

import (
	"os"
	"testing"

	"github.com/hueypark/marsettler/data"
)

func BenchmarkNewActor(b *testing.B) {
	if os.Getenv(`SKIP_DISPLAY`) != "" {
		b.Skip("Skip display test.")
	}

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
