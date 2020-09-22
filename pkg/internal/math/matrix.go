package math

type Matrix struct {
	M [2][2]float64
}

func (m Matrix) TransformVector(v Vector) (o Vector) {
	o.X = m.M[0][0]*v.X + m.M[0][1]*v.Y
	o.Y = m.M[1][0]*v.X + m.M[1][1]*v.Y
	return o
}
