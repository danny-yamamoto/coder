package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func parseLog(output string) map[string]string {
	ret := make(map[string]string)
	if output == "" {
		return ret
	}
	lastIndex := strings.Index(output, ":")
	if lastIndex == -1 {
		return ret
	}
	rows := strings.Split(output, "\n")
	for _, line := range rows {
		if line == "" {
			continue
		}
		dataLog := strings.Split(line, ":")
		if len(dataLog) != 2 {
			continue
		}
		logDate := dataLog[0]
		logData := strings.Trim(dataLog[1], " ")
		ret[logDate] = logData
	}
	return ret
}

func main() {
	logStr := "2023-01-01: New Year's Day\n2023-02-14: Valentine's Day\n2023-03-17: St. Patrick's Day\n2023-03-20\n"
	logMap := parseLog(logStr)
	logJson, err := json.MarshalIndent(logMap, "", "    ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(logJson))
}
