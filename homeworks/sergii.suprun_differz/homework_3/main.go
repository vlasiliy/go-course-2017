package main

import "fmt"

func main() {
	height := 4

	if height < 2 {
		fmt.Println("Wrong height")
		return
	}

	for i := 1; i <= height; i++ {
		first := true
		for j := i; j < height; j++ {
			fmt.Print(" ")
		}
		for {
			for j := 0; j < i; j++ {
				fmt.Print("#")
			}
			if first {
				fmt.Print("  ")
				first = false
				continue
			}
			break
		}
		fmt.Println()
	}

	for i := 0; i < height; i++ {
		for j := 0; j < height*2+2; j++ {
			s := " "
			if (j < height || j >= height+2) && (j-i <= height+2) && (i+j > height-2) {
				s = "#"
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}
