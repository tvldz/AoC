package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	var left []int
	var right []int
	var total int = 0

	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		left_int, _ := strconv.Atoi(strings.Fields(line)[0])
		right_int, _ := strconv.Atoi(strings.Fields(line)[1])
		left = append(left, left_int)
		right = append(right, right_int)
	}
	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		total += absDiffInt(left[i], right[i])
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
