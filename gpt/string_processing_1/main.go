package main

import "fmt"

func charCnt(input string, searchText string) int {
	count := 0
	searchRune := rune(searchText[0])
	for _, char := range input {
		if char == searchRune {
			count++
		}
	}
	return count
}

func main() {
	input := "英語ではHello, world!"
	searchText := "l"
	cnt := charCnt(input, searchText)
	fmt.Printf("%d\n", cnt)
}
