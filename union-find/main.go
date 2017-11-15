package main

import "fmt"

type DataSet struct {
	Size int
	DirectPaths map[int]int
}

func CreateDataSet(size int) *DataSet {
		return &DataSet{
			Size: size,
			DirectPaths: make(map[int]int),
		}
}

func (ds *DataSet)Connect(x, y int) {
	if _, ok := ds.DirectPaths[x]; ok {
		return
	} else if _, ok := ds.DirectPaths[y]; ok {
		return
	} else {
		ds.DirectPaths[x] = y
	}
}

func main() {
	ds := CreateDataSet(10)
	fmt.Printf("%+v", ds)
}
