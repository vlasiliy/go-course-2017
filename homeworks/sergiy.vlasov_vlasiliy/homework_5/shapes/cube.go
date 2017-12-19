package shapes

import "math"

type Cube struct {
	side float64
}

func (cube *Cube) Volume() float64 {
	return math.Pow(cube.side, 3)
}

func NewCube(widthSide float64) Shape {
	return &Cube{side:widthSide}
}