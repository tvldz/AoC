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

func dfs(visited [][]bool, matrix [][]int, searchValue int, rowIndex int, colIndex int, matrixHeight int, matrixWidth int, destinations map[[2]int]bool) {
	if searchValue == 9 && matrix[rowIndex][colIndex] == 9 && destinations[[2]int{rowIndex, colIndex}] != true {
		destinations[[2]int{rowIndex, colIndex}] = true
		total += 1
		return
	}
	if matrix[rowIndex][colIndex] == searchValue {
		visited[rowIndex][colIndex] = true
		if rowIndex+1 < matrixHeight {
			if visited[rowIndex+1][colIndex] != true {
				dfs(visited, matrix, searchValue+1, rowIndex+1, colIndex, matrixHeight, matrixWidth, destinations)
			}
		}
		if rowIndex-1 >= 0 {
			if visited[rowIndex-1][colIndex] != true {
				dfs(visited, matrix, searchValue+1, rowIndex-1, colIndex, matrixHeight, matrixWidth, destinations)
			}
		}
		if colIndex+1 < matrixWidth {
			if visited[rowIndex][colIndex+1] != true {
				dfs(visited, matrix, searchValue+1, rowIndex, colIndex+1, matrixHeight, matrixWidth, destinations)
			}
		}
		if colIndex-1 >= 0 {
			if visited[rowIndex][colIndex-1] != true {
				dfs(visited, matrix, searchValue+1, rowIndex, colIndex-1, matrixHeight, matrixWidth, destinations)
			}
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
				destinations := make(map[[2]int]bool)
				dfs(visited, matrix, 0, rowIndex, colIndex, matrixHeight, matrixWidth, destinations)
				grandTotal += total
			}
		}
	}
	fmt.Println(grandTotal)
}