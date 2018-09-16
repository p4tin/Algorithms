package ThreeSum

import "testing"

func TestCalculate3Sum_BruteForce(t *testing.T) {
	ints = []int{1, 2, 3, 4, 5}
	answer := Calculate3Sum_BruteForce()
	if answer != 0 {
		t.Errorf("Expected zero but got %d", answer)
	}

	ints = []int{30, -40, -20, -10, 40, 0, 10, 5}
	answer = Calculate3Sum_BruteForce()
	if answer != 4 {
		t.Errorf("Expected 4 but got %d", answer)
	}
}

func TestCalculate3Sum_BinarySearch(t *testing.T) {
	bints = []int{1, 2, 3, 4, 5}
	answer := Calculate3Sum_BinarySearch()
	if answer != 0 {
		t.Errorf("Expected zero but got %d", answer)
	}

	bints = []int{30, -40, -20, -10, 40, 0, 10, 5}
	answer = Calculate3Sum_BinarySearch()
	if answer != 4 {
		t.Errorf("Expected 4 but got %d", answer)
	}
}
