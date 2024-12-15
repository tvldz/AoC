package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	//simulationSeconds = 100 //part1
	simulationSeconds = 10000 //part2
	maxX              = 100
	maxY              = 102
)

type Robot struct {
	posX int
	posY int
	velX int
	velY int
}

func pause() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func printMatrix(robots []Robot) {
	matrix := make([][]string, maxY+1) // 3 rows
	for i := range matrix {
		matrix[i] = make([]string, maxX+1) // 4 columns per row
	}
	for _, robot := range robots {
		matrix[robot.posY][robot.posX] = "X"
	}

	for i, row := range matrix {
		for j, _ := range row {
			if matrix[i][j] != "X" {
				matrix[i][j] = " "
			}
		}
	}
	for _, row := range matrix {
		for _, column := range row {
			fmt.Printf("%s", column)
		}
		fmt.Println()
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	file, err := os.Open("14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var robots []Robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		pParts := strings.Split(parts[0], "=")
		vParts := strings.Split(parts[1], "=")
		pValue := pParts[1]
		vValue := vParts[1]
		pValues := strings.Split(pValue, ",")
		vValues := strings.Split(vValue, ",")
		p1, _ := strconv.Atoi(pValues[0])
		p2, _ := strconv.Atoi(pValues[1])
		v1, _ := strconv.Atoi(vValues[0])
		v2, _ := strconv.Atoi(vValues[1])
		robots = append(robots, Robot{posX: p1, posY: p2, velX: v1, velY: v2})
	}

	for i := 0; i < simulationSeconds; i++ {
		truths := 0
		for j, robot := range robots {
			robots[j].posX = mod(robot.posX+robot.velX, maxX+1)
			robots[j].posY = mod(robot.posY+robot.velY, maxY+1)
			if robots[j].posX == 50 && robots[j].posY == 50 {
				truths += 1
			}
			if robots[j].posX == 51 && robots[j].posY == 51 {
				truths += 1
			}
			if robots[j].posX == 52 && robots[j].posY == 52 {
				truths += 1
			}
			if robots[j].posX == 53 && robots[j].posY == 53 {
				truths += 1
			}
			if robots[j].posX == 54 && robots[j].posY == 54 {
				truths += 1
			}
		}
		if truths > 4 {
			printMatrix(robots)
			fmt.Println(i)
			pause()
		}
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	xMid := (maxX + 1) / 2
	yMid := (maxY + 1) / 2

	for _, robot := range robots {
		if robot.posX == xMid || robot.posY == yMid {
			continue
		}
		if robot.posX < xMid && robot.posY < yMid {
			q1 += 1
		}
		if robot.posX >= xMid && robot.posY < yMid {
			q2 += 1
		}
		if robot.posX < xMid && robot.posY >= yMid {
			q3 += 1
		}
		if robot.posX >= xMid && robot.posY >= yMid {
			q4 += 1
		}
	}
	fmt.Println(q1 * q2 * q3 * q4)
}
