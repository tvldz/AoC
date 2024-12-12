package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var BLINKS int = 25

func main() {
	var input []int
	file, err := os.Open("11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			digit, _ := strconv.Atoi(word)
			input = append(input, digit)
		}
	}

	for i := 0; i < BLINKS; i++ {
		for j := 0; j < len(input); j++ {

			if input[j] == 0 {
				input[j] = 1
				continue
			} else if len(strconv.Itoa(input[j]))%2 == 0 {
				inputString := strconv.Itoa(input[j])
				inputStringLength := len(inputString)
				leftDigit, _ := strconv.Atoi(inputString[:inputStringLength/2])
				rightDigit, _ := strconv.Atoi(inputString[inputStringLength/2:])
				input[j] = leftDigit
				input = slices.Insert(input, j+1, rightDigit)
				j++
				continue
			} else {
				input[j] = input[j] * 2024
			}
		}
	}
	fmt.Println(len(input))
}
