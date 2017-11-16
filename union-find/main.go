package main

import (
	"Algorithms/union-find/quickunionweighted"
	"fmt"
)

func main() {
	maxSize := 10
	ds := quickunionweighted.CreateDataSet(maxSize)
	ds.Union(1, 2)
	ds.Union(5, 6)
	ds.Union(3, 4)
	ds.Union(7, 9)
	ds.Union(2, 3)
	fmt.Printf("%+v", ds)
}
