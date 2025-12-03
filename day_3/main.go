package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := parseInstructions("day_3/input.txt")
	var biggestNumberCollection []int
	for _, line := range instructions {
		biggestNumberInLine := largestTwoDigit(line)
		biggestNumberCollection = append(biggestNumberCollection, biggestNumberInLine)
	}
	result := sumArray(biggestNumberCollection)
	fmt.Println(result)
}

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func largestTwoDigit(s string) int {
	maxNum := -1

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			twoDigits := string(s[i]) + string(s[j])

			num, err := strconv.Atoi(twoDigits)
			if err != nil {
				continue
			}
			// Update maximum
			if num > maxNum {
				maxNum = num
			}
		}
	}
	return maxNum
}

func parseInstructions(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		instructions = append(instructions, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file")
	}
	return instructions
}
