package BinarySearch

import "testing"

func Test_ReturnOK(t *testing.T) {
	ints := []int{1, 3, 5, 7, 9, 11, 13}
	key := 3
	resp := Bsearch(key, ints)
	if resp != 1 {
		t.Errorf("Expected Key 1 but got %d", resp)
	}
}

func Test_ReturnMinusOne(t *testing.T) {
	ints := []int{1, 3, 5, 7, 9, 11, 13}
	key := 4
	resp := Bsearch(key, ints)
	if resp != -1 {
		t.Errorf("Expected Key -1 but got %d", resp)
	}
}
