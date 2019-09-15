package physics

import (
	"log"
	"strconv"
	"strings"

	"github.com/hueypark/marsettler/core/math/vector"
)

type AreaOfEffect struct {
	Left   float64
	Right  float64
	Bottom float64
	Top    float64
}

func NewAreaOfEffect(param string) *AreaOfEffect {
	values := strings.Split(param, ",")
	left, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		log.Fatalln(err)
	}
	right, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		log.Fatalln(err)
	}
	bottom, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		log.Fatalln(err)
	}
	top, err := strconv.ParseFloat(values[3], 64)
	if err != nil {
		log.Fatalln(err)
	}

	return &AreaOfEffect{left, right, bottom, top}
}

func (aoe *AreaOfEffect) InArea(position vector.Vector) bool {
	if position.X <= aoe.Left ||
		aoe.Right <= position.X ||
		position.Y <= aoe.Bottom ||
		aoe.Top <= position.Y {
		return false
	}

	return true
}
