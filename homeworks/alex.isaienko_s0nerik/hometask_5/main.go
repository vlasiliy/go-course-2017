package main

import (
	"fmt"
	"os"
	"strconv"
	"sort"
	"github.com/MastersAcademy/go-course-2017/homeworks/alex.isaienko_s0nerik/hometask_5/sortable"
)

func main() {
	var N uint64 = 0
	var X float64 = 0
	if len(os.Args) == 3 {
		if parsedN, err := strconv.ParseUint(os.Args[1], 10, 0); err == nil {
			N = parsedN
		} else {
			fmt.Errorf("error while parsing arguments: %v", err)
			return
		}
		if parsedX, err := strconv.ParseFloat(os.Args[2], 0); err == nil {
			X = parsedX
		} else {
			fmt.Errorf("error while parsing arguments: %v", err)
			return
		}
	} else {
		fmt.Print("You must enter two arguments: N, X")
		return
	}

	figuresSlice := sortable.RandomFiguresSlice(N, X)

	// Sort figures by their volumes
	sort.Sort(figuresSlice)

	// Print info about each figure
	fmt.Println("Generated figures, sorted by their volume:")
	fmt.Println("Volume    Figure")
	for i := range figuresSlice {
		fmt.Printf("%-10.2f%v\n", (*figuresSlice[i]).Volume(), *figuresSlice[i])
	}
	fmt.Print("\n")

	// Arrange sorted figures diagonally
	arrangedDiagonally := figuresSlice.ArrangedDiagonally()

	// Print results of "shaking" using every possible vertex
	fmt.Println("Rotated to match a \"shaking\" operation using 1st vertex:")
	arrangedDiagonally.Rotated(1).PrintVolumes()

	fmt.Println("Rotated to match a \"shaking\" operation using 2nd vertex:")
	arrangedDiagonally.Rotated(2).PrintVolumes()

	fmt.Println("Rotated to match a \"shaking\" operation using 3rd vertex:")
	arrangedDiagonally.Rotated(3).PrintVolumes()

	fmt.Println("Rotated to match a \"shaking\" operation using 4th vertex:")
	arrangedDiagonally.Rotated(4).PrintVolumes()

}
