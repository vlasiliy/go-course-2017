package figures

import "fmt"

type Figure interface {
	Volume() float64
	fmt.Stringer
}
