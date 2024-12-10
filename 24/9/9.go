package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func buildFile(input []int) []int {
	var output []int
	tick := 0
	for index, value := range input {
		if index%2 == 0 {
			for range value {
				output = append(output, tick)
			}
			tick += 1
		} else {
			for range value {
				output = append(output, -1)
			}
		}
	}
	return output
}

func main() {
	var input []int
	file, err := os.Open("9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			fmt.Println("Error reading character:", err)
			return
		}
		digit, err := strconv.Atoi(string(char))
		input = append(input, digit)
	}

	inputFile := buildFile(input)
	//very inefficient
	for i := len(inputFile) - 1; i >= 0; i-- {
		if inputFile[i] != -1 {
			for j := 0; j < len(inputFile); j++ {
				if inputFile[j] == -1 {
					inputFile[j], inputFile[i] = inputFile[i], inputFile[j]
					break
				}
			}
		}
	}
	checksum := 0
	for k := 0; inputFile[1:][k] != -1; k++ {
		checksum += k * inputFile[1:][k]
	}
	fmt.Println(checksum)
}
