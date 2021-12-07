package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var posistions []int
var largest int

func init() {
	lines, err := filereader.ReadFileToStringArray("golf.txt")

	if err != nil {
		panic(err)
	}

	line := lines[0]

	parts := strings.Split(line, ",")

	numbers := make([]int, len(parts))

	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers[i] = num
	}
	posistions = numbers

	largest = 0
	for _, pos := range posistions {
		if pos > largest {
			largest = pos
		}
	}
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Golf")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Golf - Problem One")

	minimalMovementCost := 0

	for i := 0; i < largest; i++ {
		total := calculateMovementCost(i, false, posistions)
		if total < minimalMovementCost || minimalMovementCost == 0 {
			minimalMovementCost = total
		}
	}

	fmt.Printf("Minimal movement cost = %d\n", minimalMovementCost)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Golf - Problem two")

	minimalMovementCost := 0

	for i := 0; i < largest; i++ {
		total := calculateMovementCost(i, true, posistions)
		if total < minimalMovementCost || minimalMovementCost == 0 {
			minimalMovementCost = total
		}
	}

	fmt.Printf("Minimal movement cost = %d\n", minimalMovementCost)
}

func calculateMovementCost(x int, incrementalFuelCost bool, positions []int) int {
	total := 0
	for _, pos := range positions {
		diff := 0
		if x > pos {
			diff = (x) - pos
		} else {
			diff = pos - x
		}

		if incrementalFuelCost {
			toAdd := 0
			for i := 1; i <= diff; i++ {
				toAdd += i
			}
			total += toAdd

		} else {
			total += diff
		}
	}
	return total
}
