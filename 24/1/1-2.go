package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countInstances(x int, int_slice []int) int {
	var total int = 0
	for i := 0; i < len(int_slice); i++ {
		if x == int_slice[i] {
			total += 1
		}
	}
	return total
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

	for i := 0; i < len(left); i++ {
		num_instances := countInstances(left[i], right)
		total += num_instances * left[i]
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
