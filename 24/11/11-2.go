package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

var BLINKS int = 75

func blink(input map[int]int) map[int]int {
    var updatedStones = make(map[int]int)

    for stone_value, count := range input {
        if stone_value == 0 {
            updatedStones[1] += count
        } else if len(strconv.Itoa(stone_value))%2 == 0 {
            inputString := strconv.Itoa(stone_value)
            inputStringLength := len(inputString)
            leftDigit, _ := strconv.Atoi(inputString[:inputStringLength/2])
            rightDigit, _ := strconv.Atoi(inputString[inputStringLength/2:])
            updatedStones[leftDigit] += count
            updatedStones[rightDigit] += count
        } else {
            updatedStones[stone_value*2024] += count
        }
    }
    return updatedStones
}

func main() {
    var stones = make(map[int]int)
    total := 0

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
            stones[digit] += 1
        }
    }
    for range BLINKS {
        stones = blink(stones)
    }

    for _, count := range stones {
        total += count
    }
    fmt.Println(total)
}
