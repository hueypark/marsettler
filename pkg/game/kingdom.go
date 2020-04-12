package game

import (
	"image/color"

	"github.com/hueypark/marsettler/core/id_generator"
	"golang.org/x/image/colornames"
)

// Kingdoms has all the kingdom in the game.
var Kingdoms map[int64]*Kingdom

func init() {
	Kingdoms = make(map[int64]*Kingdom)
}

// NeutralKingdomID represents neutral kingdom id.
const NeutralKingdomID = 0

// Kingdom represents the kingdom.
type Kingdom struct {
	// id is kingdom unique id.
	id int64

	// color represents the symbolic color of the kingdom.
	color color.RGBA

	// score is the score for the victory.
	score int
}

// NewKingdom creates new kingdom.
func NewKingdom() *Kingdom {
	kingdom := &Kingdom{
		id:    id_generator.Generate(),
		color: colornames.Blue,
		score: 0,
	}

	Kingdoms[kingdom.id] = kingdom

	return kingdom
}

func (k *Kingdom) ID() int64 {
	return k.id
}
