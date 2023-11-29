package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sorting(output string) []int {
	var ret []int
	if output == "" {
		return ret
	}
	lastIndex := strings.Index(output, ",")
	if lastIndex == -1 {
		return ret
	}
	splitOutput := strings.Split(output, ",")
	var members []int
	for _, v := range splitOutput {
		if v == "" {
			continue
		}
		fmt.Printf("value: %s\n", v)
		num, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			fmt.Printf("error: %d\n", err)
			continue
		}
		members = append(members, num)
	}
	membersLen := len(members)
	if membersLen == 0 {
		return ret
	}
	sort.Ints(members)
	// min
	ret = append(ret, members[0])
	// max
	ret = append(ret, members[membersLen-1])
	return ret
}

func main() {
	ary := "3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5"
	sortAry := sorting(ary)
	for _, v := range sortAry {
		fmt.Println(v)
	}
}
