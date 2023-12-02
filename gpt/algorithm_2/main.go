package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumOfNumbersInString(input string) int {
	total := 0
	numbers := strings.Split(input, ",")
	for _, row := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(row))
		if err != nil {
			continue
		}
		total += num
	}
	return total
}

func main() {
	input := "1, 2, 3, 4, 5"
	total := sumOfNumbersInString(input)
	fmt.Println(total)
}
