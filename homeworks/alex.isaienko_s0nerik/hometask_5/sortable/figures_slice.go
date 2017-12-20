package sortable

import (
	"math"
	"fmt"
	"time"
	"math/rand"
	"github.com/MastersAcademy/go-course-2017/homeworks/alex.isaienko_s0nerik/hometask_5/figures"
	"github.com/MastersAcademy/go-course-2017/homeworks/alex.isaienko_s0nerik/hometask_5/factory"
)

type FiguresSlice []*figures.Figure

func (it FiguresSlice) Len() int {
	return len(it)
}
func (it FiguresSlice) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
func (it FiguresSlice) Less(i, j int) bool {
	return (*it[i]).Volume() < (*it[j]).Volume()
}

func (it FiguresSlice) N() int64 {
	return int64(math.Pow(float64(len(it)), 1/2.0))
}

// Return a FiguresSlice where figures are arranged diagonally
func (it FiguresSlice) ArrangedDiagonally() FiguresSlice {
	arrangedFigures := make(FiguresSlice, len(it))
	N := arrangedFigures.N()

	var i int64 = 0
	var x int64
	var y int64
	var subMatrixSize int64
	for subMatrixSize = 0; subMatrixSize < N; subMatrixSize++ {
		x = 0
		for y = subMatrixSize; y >= 0; y-- {
			arrangedFigures[y*N + x] = it[i]
			x++
			i++
		}
	}

	lastIndex := N - 1
	i = N*N - 1
	for subMatrixSize = 0; subMatrixSize < lastIndex; subMatrixSize++ {
		x = lastIndex - subMatrixSize
		for y = lastIndex; y >= lastIndex - subMatrixSize; y-- {
			arrangedFigures[x*N + y] = it[i]
			x++
			i--
		}
	}

	return arrangedFigures
}

// Returns a rotated FiguresSlice where the first element lies at the
// vertex that corresponds to the side number
// Vertices:
// 1 2
// 3 4
func (it FiguresSlice) Rotated(vertex uint64) FiguresSlice {
	newSlice := make(FiguresSlice, len(it))
	copy(newSlice, it)

	N := newSlice.N()

	if vertex == 2 || vertex == 3 {
		for i := int64(0); i < N; i++ {
			for j := int64(0); j < N; j++ {
				newSlice[i*N+j] = it[(N-j-1)*N+i]
			}
		}
	}

	if vertex == 3 || vertex == 4 {
		for i := len(newSlice)/2-1; i >= 0; i-- {
			opp := len(newSlice)-1-i
			newSlice[i], newSlice[opp] = newSlice[opp], newSlice[i]
		}
	}

	return newSlice
}

func (it FiguresSlice) PrintVolumes() {
	N := it.N()

	// Print
	for y := int64(0); y < N; y++ {
		for x := int64(0); x < N; x++ {
			fmt.Printf("%-10.2f ", (*it[y*N+x]).Volume())
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func RandomFiguresSlice(N uint64, max float64) FiguresSlice {
	var figuresSlice FiguresSlice

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := uint64(0); i < N*N; i++ {
		fig := factory.RandomFigure(rnd, max)
		figuresSlice = append(figuresSlice, fig)
	}

	return figuresSlice
}