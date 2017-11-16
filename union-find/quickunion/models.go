package quickunion

type DataSet struct {
	Size       int
	QuickUnion []int
}

func CreateDataSet(size int) *DataSet {
	ds := &DataSet{
		Size:       size,
		QuickUnion: make([]int, size),
	}
	for x := 0; x < size; x++ {
		ds.QuickUnion[x] = x
	}
	return ds
}

func (ds *DataSet) Union(x, y int) {
	root := ds.root(x)
	ds.QuickUnion[y] = root
}

func (ds *DataSet) Connected(x, y int) bool {
	return ds.root(ds.QuickUnion[x]) == ds.root(ds.QuickUnion[y])
}

func (ds *DataSet) root(x int) int {
	testRoot := x
	for ds.QuickUnion[testRoot] != testRoot {
		testRoot = ds.QuickUnion[testRoot]
	}
	return testRoot
}
