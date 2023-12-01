package main

import "fmt"

func charCnt(input string, searchText string) int {
	ret := 0
	if input == "" {
		return ret
	}
	for _, char := range input {
		if string(char) == searchText {
			ret++
		}
	}
	return ret
}

func main() {
	input := "英語ではHello, world!"
	searchText := "l"
	cnt := charCnt(input, searchText)
	fmt.Printf("%d\n", cnt)
}
