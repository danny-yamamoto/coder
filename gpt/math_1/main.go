package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func calc(output string) ([]float64, error) {
	var ret []float64
	if output == "" {
		return ret, nil
	}
	var medianFloat float64 = 0
	elements := strings.Split(output, ",")
	var total float64 = 0
	var numbers []float64
	for _, v := range elements {
		if v == "" {
			continue
		}
		row := strings.Trim(v, " ")
		numValue, err := strconv.ParseFloat(row, 64)
		if err != nil {
			continue
		}
		total = total + numValue
		numbers = append(numbers, numValue)
	}
	sort.Float64s(numbers)
	numLen := len(numbers)
	quotient := numLen / 2
	remainder := numLen % 2
	numAvg := total / float64(numLen)
	ret = append(ret, numAvg)
	if remainder == 0 {
		first := float64(numbers[quotient-1])
		second := float64(numbers[quotient])
		medianFloat = (first + second) / 2
	} else {
		medianFloat = float64(numbers[quotient])
	}
	ret = append(ret, medianFloat)
	return ret, nil
}

func main() {
	output := "1, 3, 3, 6, 7, 8, 9"
	//output := "1, 3, 3, 6, 7, 8"
	calcAry, err := calc(output)
	if err != nil {
		log.Fatal("error")
	}
	averageValue := calcAry[0]
	medianValue := calcAry[1]
	fmt.Printf("average = %f, median = %f\n", averageValue, medianValue)
}
