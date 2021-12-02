package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Christheoreo/advent-of-code-2021/utils"
)

// implements IDay
type Bravo struct{}

type instruction struct {
	Command string
	Value   int
}

var lines []string
var instructions []instruction

func (b Bravo) Setup() {
	l, err := utils.ReadFileToStringArray("bravo.txt")
	if err != nil {
		panic(err)
	}
	lines = l
	instructions = make([]instruction, len(lines))
	for index, line := range lines {

		split := strings.Split(line, " ")

		command := split[0]
		valueString := split[1]

		value, _ := strconv.Atoi(valueString)
		i := instruction{
			Command: command,
			Value:   value,
		}

		instructions[index] = i
	}
}

func (b Bravo) Run() {
	b.RunOne()
	b.RunTwo()
}

// Problem one
func (b Bravo) RunOne() {

	horizontalPos := 0
	depth := 0

	for _, instruction := range instructions {
		if instruction.Command == "forward" {
			horizontalPos += instruction.Value
		}

		if instruction.Command == "down" {
			depth += instruction.Value
		}
		if instruction.Command == "up" {
			depth -= instruction.Value
		}
	}

	answer := depth * horizontalPos

	fmt.Printf("Depth * horizontal posistion = %d\n", answer)
}

func (b Bravo) RunTwo() {
	horizontalPos := 0
	depth := 0
	aim := 0

	for _, instruction := range instructions {
		if instruction.Command == "forward" {
			horizontalPos += instruction.Value
			depth += instruction.Value * aim
		}

		if instruction.Command == "down" {
			aim += instruction.Value
		}
		if instruction.Command == "up" {
			aim -= instruction.Value
		}
	}

	answer := depth * horizontalPos

	fmt.Printf("Depth * horizontal posistion = %d\n", answer)
}

func (b Bravo) GetName() string {
	return "Day 02"
}
