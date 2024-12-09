package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	type location struct {
		row    int
		column int
	}

	var matrix [][]rune

	file, err := os.Open("8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		matrix = append(matrix, runes)
	}
	height := len(matrix[0]) - 1
	width := len(matrix) - 1

	antennas := make(map[rune][]location)
	antinodeLocations := make(map[location]int)
	for rowIndex, row := range matrix {
		for colIndex, column := range row {
			if column != '.' {
				newLocation := location{row: rowIndex, column: colIndex}
				antennas[column] = append(antennas[column], newLocation)
			}
		}
	}
	for frequency, locations := range antennas {
		for location1Index, location1 := range locations {
			for location2Index, location2 := range locations {
				if location1Index == location2Index {
					continue
				} else {

					antinodeX := 2*location2.column - location1.column
					antinodeY := 2*location2.row - location1.row
					antinode := location{column: antinodeX, row: antinodeY}
					prevAntinode := location2
					antinodeLocations[prevAntinode] = 1
					for antinode.row <= height && antinode.column <= width && antinode.row >= 0 && antinode.column >= 0 {
						antinodeLocations[antinode] = 1
						antinodeX = 2*antinode.column - prevAntinode.column
						antinodeY = 2*antinode.row - prevAntinode.row
						prevAntinode = antinode
						antinode = location{column: antinodeX, row: antinodeY}
					}
				}
			}
		}
	}
	fmt.Println(len(antinodeLocations))
}
