package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findCommonElements(input1 string, input2 string) []string {
	var commonElements []string
	numbers1 := strings.Split(input1, ",")
	numbers2 := strings.Split(input2, ",")
	for _, val1 := range numbers1 {
		val1Trim := strings.TrimSpace(val1)
		if val1Trim == "" {
			continue
		}
		num1, _ := strconv.Atoi(val1Trim)
		for _, val2 := range numbers2 {
			val2Trim := strings.TrimSpace(val2)
			if val2Trim == "" {
				continue
			}
			num2, _ := strconv.Atoi(val2Trim)
			if num1 == num2 {
				commonElements = append(commonElements, val1Trim)
				continue
			}
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
