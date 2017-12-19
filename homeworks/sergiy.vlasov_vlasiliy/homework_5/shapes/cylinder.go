package shapes

import "math"

type Cylinder struct {
	radius float64
	height float64
}

func (cylinder *Cylinder) Volume() float64 {
	return math.Pi * math.Pow(cylinder.radius, 2) * cylinder.height
}

func NewCylinder(widthSide float64) Shape {
	return &Cylinder{radius: widthSide / 2, height:widthSide}
}