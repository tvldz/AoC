package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func output(out int) {
	strNum := strconv.Itoa(out)
	for _, digit := range strNum {
		fmt.Printf("%s,", string(digit))
	}
}

func combo(operand int, registers map[string]int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return registers["A"]
	} else if operand == 5 {
		return registers["B"]
	} else if operand == 6 {
		return registers["C"]
	}
	return 0
}

func main() {
	registers := map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
	}

	str := "0,0,0,0,0" //instructions
	var instructions []int
	nums := strings.Split(str, ",")
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		instructions = append(instructions, num)
	}

	i := 0

	for i < len(instructions) {
		if instructions[i] == 0 {
			operand := instructions[i+1]
			c := combo(operand, registers)
			denominator := int(math.Pow(2, float64(c)))
			registers["A"] = registers["A"] / denominator
			i += 2
			continue
		} else if instructions[i] == 1 {
			registers["B"] = registers["B"] ^ instructions[i+1]
			i += 2
			continue
		} else if instructions[i] == 2 {
			operand := instructions[i+1]
			c := combo(operand, registers)
			registers["B"] = c % 8
			i += 2
			continue
		} else if instructions[i] == 3 {
			if registers["A"] == 0 {
				i += 2
				continue
			}
			i = instructions[i+1]
			continue
		} else if instructions[i] == 4 {
			registers["B"] = registers["B"] ^ registers["C"]
			i += 2
			continue
		} else if instructions[i] == 5 {
			operand := instructions[i+1]
			c := combo(operand, registers)
			output(c % 8)
			i += 2
			continue
		} else if instructions[i] == 6 {
			operand := instructions[i+1]
			c := combo(operand, registers)
			denominator := int(math.Pow(2, float64(c)))
			registers["B"] = registers["A"] / denominator
			i += 2
			continue
		} else if instructions[i] == 7 {
			operand := instructions[i+1]
			c := combo(operand, registers)
			denominator := int(math.Pow(2, float64(c)))
			registers["C"] = registers["A"] / denominator
			i += 2
			continue
		}
	}
}
