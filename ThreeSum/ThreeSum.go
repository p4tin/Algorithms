package ThreeSum

//Given N distinct integers, how many triplesums exactly to zero
var ints []int

func Calculate3Sum_BruteForce() int {
	counter := 0
	for x := 0; x < len(ints)-2; x++ {
		for y := x + 1; y < len(ints)-1; y++ {
			for z := y + 1; z < len(ints); z++ {
				if ints[x]+ints[y]+ints[z] == 0 {
					counter++
				}
			}
		}
	}
	return counter
}
