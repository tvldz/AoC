package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkForString(matrix [][]rune, r int, c int) int {
	total := 0
	offsets := [][][]int{{
		{0, 1}, {0, 2}, {0, 3}},
		{{0, -1}, {0, -2}, {0, -3}},
		{{-1, 0}, {-2, 0}, {-3, 0}},
		{{1, 0}, {2, 0}, {3, 0}},
		{{1, 1}, {2, 2}, {3, 3}},
		{{1, -1}, {2, -2}, {3, -3}},
		{{-1, -1}, {-2, -2}, {-3, -3}},
		{{-1, 1}, {-2, 2}, {-3, 3}},
	}

	for _, set := range offsets {
		test_string := ""
		for _, set2 := range set {
			if r+set2[0] >= 0 && r+set2[0] < len(matrix) && c+set2[1] >= 0 && c+set2[1] < len(matrix[r+set2[0]]) {
				test_string += string(matrix[r+set2[0]][c+set2[1]])
			} else {
				break
			}
		}
		if test_string == "MAS" {
			total += 1
		}
	}
	return total
}

func main() {
	var matrix [][]rune
	total := 0

	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var charSlice []rune
		for _, char := range line {
			charSlice = append(charSlice, char)
		}
		matrix = append(matrix, charSlice)
	}

	rows := len(matrix)
	columns := len(matrix[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if string(matrix[r][c]) == "X" {
				total += checkForString(matrix, r, c)
			}
		}
	}
	fmt.Println(total)
}
