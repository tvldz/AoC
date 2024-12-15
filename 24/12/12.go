package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var NORTH = []int{-1, 0}
var SOUTH = []int{1, 0}
var EAST = []int{0, 1}
var WEST = []int{0, -1}

var plotCount = 0
var fenceCount = 0
var grandTotal = 0

func flood(land [][]rune, plant rune, location []int, landSize []int, landCopy [][]rune) {
	if location[0] < 0 || location[0] >= landSize[0] {
		return
	} else if location[1] < 0 || location[1] >= landSize[1] {
		return
	} else if land[location[0]][location[1]] != plant {
		return
	} else {
		getFenceCount(landCopy, plant, location, landSize)
		land[location[0]][location[1]] = '0'
		plotCount++
		north := []int{location[0] + NORTH[0], location[1] + NORTH[1]}
		flood(land, plant, north, landSize, landCopy)
		south := []int{location[0] + SOUTH[0], location[1] + SOUTH[1]}
		flood(land, plant, south, landSize, landCopy)
		east := []int{location[0] + EAST[0], location[1] + EAST[1]}
		flood(land, plant, east, landSize, landCopy)
		west := []int{location[0] + WEST[0], location[1] + WEST[1]}
		flood(land, plant, west, landSize, landCopy)
	}
	return
}

func getFenceCount(landCopy [][]rune, plant rune, location []int, landSize []int) {
	count := 0
	north := []int{location[0] + NORTH[0], location[1] + NORTH[1]}
	south := []int{location[0] + SOUTH[0], location[1] + SOUTH[1]}
	west := []int{location[0] + WEST[0], location[1] + WEST[1]}
	east := []int{location[0] + EAST[0], location[1] + EAST[1]}
	if north[0] < 0 {
		count++
	} else {
		if landCopy[north[0]][north[1]] != plant {
			count++
		}
	}
	if south[0] >= landSize[0] {
		count++
	} else {
		if landCopy[south[0]][south[1]] != plant {
			count++
		}
	}
	if west[1] < 0 {
		count++
	} else {
		if landCopy[west[0]][west[1]] != plant {
			count++
		}
	}
	if east[1] >= landSize[1] {
		count++
	} else {
		if landCopy[east[0]][east[1]] != plant {
			count++
		}
	}
	fenceCount += count
	return
}

func main() {
	var land [][]rune

	file, err := os.Open("12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		land = append(land, runes)
	}
	landSize := []int{len(land[0]), len(land)}

	landCopy := make([][]rune, len(land))
	for i := range land {
		landCopy[i] = make([]rune, len(land[i]))
		copy(landCopy[i], land[i])
	}

	for rowIndex, row := range land {
		for colIndex, column := range row {
			if column != '0' {
				plotCount = 0
				fenceCount = 0
				node := []int{rowIndex, colIndex}
				flood(land, column, node, landSize, landCopy)
				grandTotal += plotCount * fenceCount
			}
		}
	}
	fmt.Println(grandTotal)
}
