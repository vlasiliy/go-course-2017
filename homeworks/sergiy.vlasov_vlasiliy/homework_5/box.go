package main

import (
	"github.com/MastersAcademy/go-course-2017/homeworks/sergiy.vlasov_vlasiliy/homework_5/shapes"
	"math/rand"
	"fmt"
)

type Box [][]float64

func createBox(countCeil int, height float64) Box {
	var box Box

	for i := 0; i < countCeil; i++ {
		row := []float64{}
		for j := 0; j < countCeil; j++ {
			randomSide := rand.Float64() * height
			row = append(row, shapes.CreateShape(shapes.RandomShapeName(), randomSide).Volume())
		}
		box = append(box, row);
	}

	return box
}

func shakeBox(box *Box, angle int) {
	fillAfterDefaultShake(box)

	for i := 0; i < (angle - 1); i++ {
		turn90(box)
	}
}

func turn90(box *Box) {
	b := *box
	countCeil := len(b)

	for i := 0; i < int(countCeil / 2); i++ {
		for j := i; j < countCeil - 1 - i; j++ {
			tmp := b[i][j];
			b[i][j] = b[countCeil - j -1][i];
			b[countCeil - j - 1][i] = b[countCeil - i - 1][countCeil - j - 1];
			b[countCeil - i - 1][countCeil - j - 1] = b[j][countCeil-i-1];
			b[j][countCeil - i - 1] = tmp;
		}
	}
}

func fillAfterDefaultShake(box *Box) {
	var lineKey int

	b := *box
	countCeil := len(b)

	for sumKeys := 0; sumKeys < countCeil * 2; sumKeys++ {
		for i := 0; i < countCeil; i++ {
			for j := 0; j < countCeil; j++ {
				if i + j == sumKeys {
					b[i][j] = line[lineKey]
					lineKey++
				}
				if i + j > sumKeys {
					break
				}
			}
		}
	}
}

func showBox (box *Box) {
	b := *box
	countCeil := len(b)

	for i := 0; i < countCeil; i++ {
		for j := 0; j < countCeil; j++ {
			fmt.Printf("%-7.1f" ,b[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func createSortLine(box *Box) Line {
	var line Line
	var temp float64

	b := *box;
	countCeil := len(b)

	for i := 0; i < countCeil; i++ {
		for j := 0; j < countCeil; j++ {
			line = append(line, b[i][j])
		}
	}

	for i := 0; i < countCeil * countCeil; i++ {
		for j := i; j < countCeil * countCeil; j++ {
			if line[i] > line[j] {
				temp = line[i]
				line[i] = line[j]
				line[j] = temp
			}
		}
	}

	return line
}