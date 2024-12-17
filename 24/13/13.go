package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const aCost = 3
const bCost = 1

type Machine struct {
	aX      int
	aY      int
	bX      int
	bY      int
	targetX int
	targetY int
}

func main() {
	var machines []Machine
	var newMachine Machine
	re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	prizeRe := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	grandTotal := 0

	file, err := os.Open("13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	section := 0
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Text() != "" {
			if section%3 == 0 {
				matches := re.FindStringSubmatch(line)
				newMachine.aX, _ = strconv.Atoi(matches[1])
				newMachine.aY, _ = strconv.Atoi(matches[2])
			}
			if section%3 == 1 {
				matches := re.FindStringSubmatch(line)
				newMachine.bX, _ = strconv.Atoi(matches[1])
				newMachine.bY, _ = strconv.Atoi(matches[2])
			}
			if section%3 == 2 {
				matches := prizeRe.FindStringSubmatch(line)
				newMachine.targetX, _ = strconv.Atoi(matches[1])
				newMachine.targetY, _ = strconv.Atoi(matches[2])
				machines = append(machines, newMachine)
			}
			section++
		}
	}
	for _, machine := range machines {
		buttonA := []int{machine.aX, machine.aY}
		buttonB := []int{machine.bX, machine.bY}
		total := 999999
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if buttonA[0]*i+buttonB[0]*j == machine.targetX && buttonA[1]*i+buttonB[1]*j == machine.targetY {
					if i*3+j < total {
						total = i*3 + j
					}
				}
			}
		}
		if total != 999999 {
			grandTotal += total
		}
	}
	fmt.Println(grandTotal)
}
