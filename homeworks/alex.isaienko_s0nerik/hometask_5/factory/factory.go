package factory

import (
	"math/rand"
	"github.com/MastersAcademy/go-course-2017/homeworks/alex.isaienko_s0nerik/hometask_5/figures"
)

const (
	Cube = iota
	Sphere = iota
)

var figureTypes = [...]uint{ Cube, Sphere }

func randomFigureOfType(rnd *rand.Rand, figureType uint, maxSideLength float64) *figures.Figure {
	var figure figures.Figure = nil
	switch figureType {
	case Cube:
		figure = &figures.Cube{ Side: rnd.Float64() * maxSideLength }
	case Sphere:
		figure = &figures.Sphere{ Radius: rnd.Float64() * maxSideLength / 2 }
	}
	return &figure
}

func RandomFigure(rnd *rand.Rand, maxSideLength float64) *figures.Figure {
	figureTypeIndex := rnd.Intn(len(figureTypes))
	return randomFigureOfType(rnd, figureTypes[figureTypeIndex], maxSideLength)
}