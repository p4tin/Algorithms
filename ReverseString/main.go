package main

import "fmt"

func ReverseString(input string) string {
	inputBytes := []rune(input)
	outputBytes := make([]rune, len(inputBytes))

	ind := 0
	for i := len(inputBytes) - 1; i >= 0; i-- {
		outputBytes[ind] = inputBytes[i]
		ind++
	}
	return string(outputBytes)
}

func main() {
	Str := "This is an awesome function"
	fmt.Println(ReverseString(Str))
}
