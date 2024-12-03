package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getOrder(x, y int) int {
	if x-y < 0 {
		return 0
	} else {
		return 1
	}
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func getNumSlice(line []string) []int {
	var return_slice []int
	for _, str_num := range line {
		int_num, _ := strconv.Atoi(str_num)
		return_slice = append(return_slice, int_num)
	}
	return return_slice
}

func main() {
	var matrix [][]int
	var last int
	direction := -1
	check := 1
	total := 0

	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, getNumSlice(strings.Fields(line)))
	}

	for _, array := range matrix {
		last = -1
		direction = -1
		check = 1
		for _, value := range array {
			if last == -1 {
				last = value
				continue
			}
			if direction == -1 {
				direction = getOrder(value, last)
			}
			if getOrder(value, last) != direction {
				check = 0
				last = value
				break
			}
			diff := absDiffInt(last, value)
			if diff < 1 || diff > 3 {
				check = 0
				last = value
				break
			}
			last = value
		}
		total += check
	}
	fmt.Println(total)
}
