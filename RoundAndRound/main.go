package main

import "fmt"

func main() {
	square := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	size := len(square)
	fmt.Printf("size: %d\n", size)
	for x := 0; x < size; x++ {
		fmt.Println(square[x])
	}

	for x := 0; x < size-1; x++ {
		if x == 0 {
			for y := 0; y < size; y++ {
				fmt.Printf("%d ", square[0][y])
			}
		} else {
			fmt.Printf("%d ", square[x][size-1])
		}
	}

	for x := size - 1; x > 0; x-- {
		if x == (size - 1) {
			for y := size - 1; y >= 0; y-- {
				fmt.Printf("%d ", square[size-1][y])
			}
		} else {
			fmt.Printf("%d ", square[x][0])
		}
	}
	fmt.Println("")
}
