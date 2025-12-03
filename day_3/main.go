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
		biggestNumberInLine := largestKDigits(line, 12)
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

func largestKDigits(s string, k int) int {
	n := len(s)
	if n < k {
		return 0
	}

	var result []int

	// For each position in the result
	for i := 0; i < k; i++ {
		// Determine search range
		start := 0
		if len(result) > 0 {
			start = result[len(result)-1] + 1 // Start after last picked position
		}

		// We need k - i digits total (including this one)
		// Must leave at least k - i - 1 digits after this position
		end := n - k + i

		// Find the maximum digit in the valid range
		maxDigit := byte('0') - 1
		maxPos := -1

		for j := start; j <= end; j++ {
			if s[j] > maxDigit {
				maxDigit = s[j]
				maxPos = j
			}
		}

		result = append(result, maxPos)
	}

	// Build the result string from selected positions
	resultStr := ""
	for _, pos := range result {
		resultStr += string(s[pos])
	}

	largest, _ := strconv.Atoi(resultStr)

	return largest
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
