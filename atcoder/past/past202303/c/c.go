package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var p string
	var err error
	var num int
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if scanner.Scan() {
		p = scanner.Text()
	}
	split := strings.Split(p, " ")
	for i := 1; i <= n; i++ {
		for index, part := range split {
			num, err = strconv.Atoi(part)
			if err != nil {
				fmt.Println(err)
			}
			if i == num {
				numbers = append(numbers, index+1)
				break
			}
		}
	}
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), " "), "[]"))
}
