package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := parseInstructions("day_2/input.csv")
	var invalidIDs []int
	for _, ins := range instructions {
		split := strings.Split(ins, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		invalidIDs = append(invalidIDs, checkRange(start, end)...)
	}
	result := sumArray(invalidIDs)

	fmt.Println(result)
}

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func parseInstructions(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV")
	}

	return lines[0]
}

func checkRange(start, end int) []int {
	var invalid []int

	for id := start; id <= end; id++ {
		if isInvalid(id) {
			invalid = append(invalid, id)
		}
	}
	return invalid
}

func isInvalid(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)

	// odd length strings cannot be invalid
	if length%2 != 0 {
		return false
	}

	midPoint := length / 2
	firstHalf := idStr[:midPoint]
	secondHalf := idStr[midPoint:]

	return firstHalf == secondHalf
}
