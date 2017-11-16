package quickunionweighted

type DataSet struct {
	Size       int
	QuickUnion []int
	TreeSize   []int
}

func CreateDataSet(size int) *DataSet {
	ds := &DataSet{
		Size:       size,
		QuickUnion: make([]int, size),
		TreeSize:   make([]int, size),
	}
	for x := 0; x < size; x++ {
		ds.QuickUnion[x] = x
		ds.TreeSize[x] = 1
	}
	return ds
}

func (ds *DataSet) Union(x, y int) {
	rootx := ds.root(x)
	rooty := ds.root(y)
	if ds.TreeSize[y] < ds.TreeSize[x] {
		ds.QuickUnion[rooty] = rootx
		ds.TreeSize[rooty] = ds.TreeSize[rooty] + ds.TreeSize[rootx]
	} else {
		ds.QuickUnion[rootx] = rooty
		ds.TreeSize[rootx] = ds.TreeSize[rootx] + ds.TreeSize[rooty]
	}
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
