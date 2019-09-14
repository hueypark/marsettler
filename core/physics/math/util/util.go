package util

import "math"

func DegToRad(angle float64) float64 {
	return angle * math.Pi / 180
}

func RadToDeg(angle float64) float64 {
	return angle * 180 / math.Pi
}
