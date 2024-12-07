package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLocation(inputSlice []string, rule string) int {
	index := -1
	for i, input := range inputSlice {
		if input == rule {
			index = i
			break
		}
	}
	return index
}

func isValidSlice(inputSlice []string, rulesList [][]string) bool {
	returnValue := true
	for _, rule := range rulesList {
		higher := getLocation(inputSlice, rule[0])
		lower := getLocation(inputSlice, rule[1])
		if higher == -1 || lower == -1 {
			continue
		} else if higher > lower {
			returnValue = false
			break
		}
	}
	return returnValue
}

func main() {
	var rulesSlice []string
	var inputSlice []string
	var rulesList [][]string
	var inputList [][]string
	total := 0
	file, err := os.Open("5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rulesSlice = strings.Split(line, "|")
			rulesList = append(rulesList, rulesSlice)
		}
		if strings.Contains(line, ",") {
			inputSlice = strings.Split(line, ",")
			inputList = append(inputList, inputSlice)
		}
	}
	for _, inputSlice := range inputList {
		if isValidSlice(inputSlice, rulesList) {
			i, _ := strconv.Atoi(inputSlice[len(inputSlice)/2])
			total += i
		}
	}
	fmt.Println(total)
}