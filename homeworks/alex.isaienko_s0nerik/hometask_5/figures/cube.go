package figures

import (
	"math"
	"fmt"
)

type Cube struct {
	Side float64
}

func (it *Cube) Volume() float64 {
	return math.Pow(it.Side, 3)
}

func (it *Cube) String() string {
	return fmt.Sprintf("Cube(side = %v)", it.Side)
}