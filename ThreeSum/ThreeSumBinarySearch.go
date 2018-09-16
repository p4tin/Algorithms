package ThreeSum

import (
	"Algorithms/BinarySearch"
	"fmt"
	"sort"
)

//Given N distinct integers, how many triplesums exactly to zero
var bints []int

func Calculate3Sum_BinarySearch() int {
	sints := make([]int, len(bints))

	copy(sints, bints)
	sort.Ints(sints)

	counter := 0
	for x := 0; x < len(bints)-2; x++ {
		for y := x + 1; y < len(bints)-1; y++ {
			z := -(bints[x] + bints[y])
			if BinarySearch.Bsearch(z, sints) == -1 {
				fmt.Println(x, y, z)
				counter++
			}
		}
	}
	return counter
}
