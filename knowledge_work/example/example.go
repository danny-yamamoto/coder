package main

import (
	"fmt"
	"strings"
)

func parse(output string) map[string]string {
	ret := make(map[string]string)
	if output == "" {
		return ret
	}
	splitData := strings.Split(output, "\n")
	for _, row := range splitData {
		if len(row) == 0 {
			continue
		}
		lastIndex := strings.Index(row, ":")
		if lastIndex == -1 {
			continue
		}
		trimRow := strings.Trim(row, " ")
		if trimRow == "" {
			continue
		}
		rowAry := strings.Split(row, ":")
		fileName := strings.Trim(rowAry[0], " ")
		mineType := strings.Trim(rowAry[1], " ")
		ret[fileName] = mineType
	}
	return ret
}

func main() {
	// sample data
	fileOutput := "go.mod: text/plain\ngo.sum: text/plain\nmain: application/x-match-binary"
	fileData := parse(fileOutput)
	for k, v := range fileData {
		fmt.Println(k + ":" + v)
	}
}
