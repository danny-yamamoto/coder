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
	ary := strings.Split(output, ",")
	var total float64 = 0
	var numbers []int
	for _, v := range ary {
		if v == "" {
			continue
		}
		row := strings.Trim(v, " ")
		numValue, err := strconv.ParseFloat(row, 64)
		numInt, _ := strconv.Atoi(row)
		if err != nil {
			continue
		}
		total = total + numValue
		numbers = append(numbers, numInt)
	}
	sort.Ints(numbers)
	numLen := len(ary)
	quotient := numLen / 2
	remainder := numLen % 2
	numAvg := total / float64(numLen)
	ret = append(ret, numAvg)
	if remainder == 0 {
		first := float64(numbers[quotient-1])
		second := float64(numbers[quotient])
		medianFloat = (first + second) / 2
		//fmt.Printf("medianFloat: %f\n", medianFloat)
	} else {
		medianFloat = float64(numbers[quotient])
		//fmt.Printf("medianFloat: %f\n", medianFloat)
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
