package shapes

import "math"

type Cone struct {
	radius float64
	height float64
}

func (cone *Cone) Volume() float64 {
	return math.Pi * math.Pow(cone.radius, 2) * cone.height / 3
}

func NewCone(widthSide float64) Shape {
	return &Cone{radius: widthSide / 2, height:widthSide}
}