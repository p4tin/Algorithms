package quickunionweighted

import "testing"

func TestConnections(t *testing.T) {
	maxSize := 10
	ds := CreateDataSet(maxSize)
	tests := []struct {
		Name string
		X    int
		Y    int
	}{
		{"Test1", 1, 2},
		{"Test2", 4, 5},
		{"Test3", 7, 9},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ds.Union(test.X, test.Y)
			if !ds.Connected(1, 2) {
				t.Errorf("Expected %d and %d to be connected but they are not\n", test.X, test.Y)
			}
		})
	}
}

func TestComplexPath(t *testing.T) {
	maxSize := 10
	ds := CreateDataSet(maxSize)
	ds.Union(1, 2)
	ds.Union(5, 6)
	ds.Union(3, 4)
	ds.Union(7, 9)
	ds.Union(2, 3)
	if !ds.Connected(1, 4) {
		t.Errorf("Expected %d and %d to be connected but they are not\n", 1, 4)
	}
	if ds.Connected(2, 9) {
		t.Errorf("Expected %d and %d not to be connected but they are\n", 2, 9)
	}

}
