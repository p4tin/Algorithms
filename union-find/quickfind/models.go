package quickfind

type DataSet struct {
	Size      int
	QuickFind []int
}

func CreateDataSet(size int) *DataSet {
	ds := &DataSet{
		Size:      size,
		QuickFind: make([]int, size),
	}
	for x := 0; x < size; x++ {
		ds.QuickFind[x] = x
	}
	return ds
}

func (ds *DataSet) Union(x, y int) {
	old := ds.QuickFind[x]
	new := ds.QuickFind[y]
	ds.QuickFind[x] = new
	for x := 0; x < ds.Size; x++ {
		if ds.QuickFind[x] == old {
			ds.QuickFind[x] = new
		}
	}
}

func (ds *DataSet) Connected(x, y int) bool {
	return ds.QuickFind[x] == ds.QuickFind[y]
}
