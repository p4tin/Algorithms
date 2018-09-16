package main

import "fmt"

// FindNthOccurance = Returns the position of the nth occurance or -1
func FindNthOccurance(haystack, needle string, occurance int) int {
	for i := 0; i < len(haystack)-(len(needle)-1); i++ {
		if haystack[i:i+len(needle)] == needle {
			occurance--
			if occurance == 0 {
				return i
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(FindNthOccurance("the cat jumped out of the", "the", 2))
}
