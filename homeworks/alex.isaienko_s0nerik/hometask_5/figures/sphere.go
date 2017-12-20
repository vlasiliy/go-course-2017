package figures

import (
	"math"
	"fmt"
)

type Sphere struct {
	Radius float64
}

func (it *Sphere) Volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(it.Radius, 3)
}

func (it *Sphere) String() string {
	return fmt.Sprintf("Sphere(radius = %v)", it.Radius)
}