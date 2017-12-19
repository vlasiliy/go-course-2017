package shapes

var Shapes = map[string]func(widthSide float64) Shape {
	"Cube":NewCube,
	"Cylinder":NewCylinder,
	"Cone":NewCone,
}

func CreateShape(name string, widthSide float64) Shape {
	if _, ok := Shapes[name]; ok {
		return Shapes[name](widthSide)
	} else {
		return nil
	}
}

func RandomShapeName() string {
	for key, _ := range Shapes {
		return key
	}

	return ""
}