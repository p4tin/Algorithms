package BinarySearch

func Bsearch(key int, ints []int) int {
	lo := 0
	hi := len(ints) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < ints[mid] {
			hi = mid - 1
		} else if key > ints[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
