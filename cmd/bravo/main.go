package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

type instruction struct {
	Command string
	Value   int
}

var lines []string
var instructions []instruction

func init() {
	l, err := filereader.ReadFileToStringArray("bravo.txt")
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

func main() {
	defer timetrack.TimeTrack(time.Now(), "Bravo")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Bravo - Problem One")
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

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Bravo - Problem two")
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
