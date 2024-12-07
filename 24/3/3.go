package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	total := 0
	last := 0
	that := 0
	data, err := os.ReadFile("3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(data)
	re := regexp.MustCompile(`mul\(\d*,\d*\)`)
	re2 := regexp.MustCompile(`\d*`)
	matches := re.FindAllStringSubmatch(text, -1)

	counter := 0
	for _, match := range matches {
		matches2 := re2.FindAllStringSubmatch(match[0], -1)
		for _, match2 := range matches2 {
			if len(match2[0]) > 0 {
				if counter%2 == 0 {
					last, _ = strconv.Atoi(match2[0])
				}
				if counter%2 == 1 {
					that, _ = strconv.Atoi(match2[0])
					total += last * that
				}
				counter += 1
			}
		}
	}
	fmt.Println(total)
}