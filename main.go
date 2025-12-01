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
	instructions := parseInstructions("input.txt")

	dial := NewDial(50)
	for _, instr := range instructions {
		if instr.direction == "L" {
			dial.Left(instr.steps)
		} else {
			dial.Right(instr.steps)
		}
		fmt.Printf("%s%d -> %d\n", instr.direction, instr.steps, dial.Value())
	}
	fmt.Println(dial.ZeroCrossings())
}

func parseInstructions(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := string(line[0])
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Print("invalid instruction")
			continue
		}

		instructions = append(instructions, Instruction{
			direction: direction,
			steps:     steps,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file")
	}

	return instructions
}

// Dial represents a rotating dial object.
type Dial struct {
	value         int // current position (0-99)
	zeroCrossings int
}

// Instruction represents how to move the dial.
type Instruction struct {
	direction string
	steps     int
}

// NewDial instantiates a dial.
func NewDial(start int) *Dial {
	return &Dial{value: ((start % 100) + 100) % 100}
}

// Left moves the dial counterclockwise: 0 → 99, 1 → 0, etc.
func (d *Dial) Left(steps int) {
	for i := 0; i < steps; i++ {
		d.value = (d.value - 1 + 100) % 100
		if d.value == 0 {
			d.zeroCrossings++
		}
	}
}

// Right moves the dial clockwise: 99 → 0, 98 → 99, etc.
func (d *Dial) Right(steps int) {
	for i := 0; i < steps; i++ {
		d.value = (d.value + 1) % 100
		if d.value == 0 {
			d.zeroCrossings++
		}
	}
}

// Value returns the value of the dial.
func (d *Dial) Value() int {
	return d.value
}

// ZeroCrossings returns the number of times the dial has landed on, or crossed 0.
func (d *Dial) ZeroCrossings() int {
	return d.zeroCrossings
}
