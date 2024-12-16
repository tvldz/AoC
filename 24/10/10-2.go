package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var total = 0
var grandTotal = 0

func dfs(visited [][]bool, matrix [][]int, searchValue int, rowIndex int, colIndex int, matrixHeight int, matrixWidth int) {
	if searchValue == 9 && matrix[rowIndex][colIndex] == 9 {
		total += 1
		return
	}
	if matrix[rowIndex][colIndex] == searchValue {
		visited[rowIndex][colIndex] = true
		if rowIndex+1 < matrixHeight {
			dfs(visited, matrix, searchValue+1, rowIndex+1, colIndex, matrixHeight, matrixWidth)
		}
		if rowIndex-1 >= 0 {
			dfs(visited, matrix, searchValue+1, rowIndex-1, colIndex, matrixHeight, matrixWidth)
		}
		if colIndex+1 < matrixWidth {
			dfs(visited, matrix, searchValue+1, rowIndex, colIndex+1, matrixHeight, matrixWidth)
		}
		if colIndex-1 >= 0 {
			dfs(visited, matrix, searchValue+1, rowIndex, colIndex-1, matrixHeight, matrixWidth)
		}
	}
	return
}

func main() {
	var matrix [][]int

	file, err := os.Open("10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersStr := strings.Fields(line)
		var numbers []int
		for _, numberStr := range numbersStr {
			for _, char := range numberStr {
				numInt, _ := strconv.Atoi(string(char))
				numbers = append(numbers, numInt)
			}
		}
		matrix = append(matrix, numbers)
	}

	matrixHeight := len(matrix[0])
	matrixWidth := len(matrix)
	visited := make([][]bool, matrixHeight)
	for i := range visited {
		visited[i] = make([]bool, matrixWidth)
	}

	for rowIndex, row := range matrix {
		for colIndex, _ := range row {
			if matrix[rowIndex][colIndex] == 0 {
				total = 0
				for i := 0; i < matrixHeight; i++ {
					for j := 0; j < matrixWidth; j++ {
						visited[i][j] = false
					}
				}
				dfs(visited, matrix, 0, rowIndex, colIndex, matrixHeight, matrixWidth)
				grandTotal += total
			}
		}
	}
	fmt.Println(grandTotal)
}
