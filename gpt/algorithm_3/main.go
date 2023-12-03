package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findCommonElements(input1 string, input2 string) []string {
	var commonElements []string
	set := make(map[int]struct{})
	numbers1 := strings.Split(input1, ",")
	for _, val := range numbers1 {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		set[num] = struct{}{}
	}
	numbers2 := strings.Split(input2, ",")
	for _, val2 := range numbers2 {
		num2, _ := strconv.Atoi(strings.TrimSpace(val2))
		if _, exists := set[num2]; exists {
			commonElements = append(commonElements, strconv.Itoa(num2))
			delete(set, num2)
		}
	}

	return commonElements
}

func main() {
	input1 := "1, 3, 4, 6, 7, 9"
	input2 := "1, 2, 4, 5, 9, 10"
	matchNumbers := findCommonElements(input1, input2)
	result := strings.Join(matchNumbers, ",")
	fmt.Println(result)
}
