package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getAllKLength(set []rune, k int) [][]rune {
	n := len(set)
	result := make([][]rune, 0)
	getAllKLengthRec(set, []rune{}, n, k, &result)
	return result
}

func getAllKLengthRec(set []rune, prefix []rune, n, k int, result *[][]rune) {
	if k == 0 {
		combination := make([]rune, len(prefix))
		copy(combination, prefix)
		*result = append(*result, combination)
		return
	}

	for i := 0; i < n; i++ {
		newPrefix := append(prefix, set[i])
		getAllKLengthRec(set, newPrefix, n, k-1, result)
	}
}

func main() {
	type equation struct {
		targetValue int
		values      []int
	}

	operands := []rune{'+', '*'}
	grandTotal := 0

	file, err := os.Open("7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var valueSliceInt []int
		line := scanner.Text()
		lineSlice := strings.Split(line, ":")
		targetInt, _ := strconv.Atoi(lineSlice[0])
		trimmed := strings.TrimLeftFunc(lineSlice[1], unicode.IsSpace)
		valueSlice := strings.Split(trimmed, " ")
		for _, value := range valueSlice {
			valueInt, _ := strconv.Atoi(value)
			valueSliceInt = append(valueSliceInt, valueInt)
		}
		newEquation := equation{targetValue: targetInt, values: valueSliceInt}
		operands := getAllKLength(operands, len(newEquation.values)-1)
		for i := 0; i < len(operands); i++ {
			operand := operands[i]
			total := newEquation.values[0]
			for j := 1; j < len(newEquation.values); j++ {
				if operand[j-1] == '*' {
					total *= newEquation.values[j]
				} else if operand[j-1] == '+' {
					total += newEquation.values[j]
				}
			}
			if total == newEquation.targetValue {
				grandTotal += total
				fmt.Printf("Found: %d\n", newEquation)
				break
			}
		}
	}
	fmt.Println(grandTotal)
}
