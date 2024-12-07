package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getStart(matrix [][]string) []int {
	var startIndex []int
	for row_index, row := range matrix {
		for column_index, column := range row {
			if column == "^" {
				startIndex = []int{row_index, column_index}
			}
		}
	}
	return startIndex
}

func main() {
	type location struct {
		row    int
		column int
	}

	var matrix [][]string

	file, err := os.Open("6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		charSlice := strings.Split(line, "")
		matrix = append(matrix, charSlice)
	}
	matrixSize := []int{len(matrix), len(matrix[0])}
	currentLocation := getStart(matrix)

	movements := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	currentMovement := 0
	var visitedLocations = make(map[location]int)

	for {
		visitedLocation := location{row: currentLocation[0], column: currentLocation[1]}
		visitedLocations[visitedLocation] = 1
		if currentLocation[0]+movements[currentMovement%4][0] >= matrixSize[0] || currentLocation[1]+movements[currentMovement%4][1] >= matrixSize[1] {
			break
		}
		if currentLocation[0]+movements[currentMovement%4][0] < 0 || currentLocation[1]+movements[currentMovement%4][1] < 0 {
			break
		}
		if matrix[currentLocation[0]+movements[currentMovement%4][0]][currentLocation[1]+movements[currentMovement%4][1]] == "#" {
			currentMovement++
		} else {
			currentLocation[0] = currentLocation[0] + movements[currentMovement%4][0]
			currentLocation[1] = currentLocation[1] + movements[currentMovement%4][1]
		}
	}
	fmt.Println(len(visitedLocations))
}
