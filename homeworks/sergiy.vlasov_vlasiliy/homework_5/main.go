package main

import (
	"fmt"
	"strconv"
	"os"
)

type Line []float64

var line Line

func main() {
	var sizes [3]int

	args := os.Args[1:]
	//test
	//args := [3]string{"4", "4", "10"};

	if len(args) == 3 {
		for i, value := range args {
			param, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println(err)

				return
			} else {
				sizes[i] = param
			}
		}
		if (sizes[0] != sizes[1]) {
			fmt.Println("The second parameter is wrong")
		}
	} else {
		fmt.Println("Count of parameters is wrong")

		return
	}

	box := createBox(sizes[0] ,float64(sizes[2]))
	showBox(&box)

	line = createSortLine(&box)

	shakeBox(&box, 1)
	showBox(&box)

	shakeBox(&box, 2)
	showBox(&box)

	shakeBox(&box, 3)
	showBox(&box)

	shakeBox(&box, 4)
	showBox(&box)


}