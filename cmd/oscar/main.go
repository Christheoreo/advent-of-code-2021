package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var origonalNumbers [][]int

func init() {
	lines, _ := filereader.ReadFileToStringArray("oscar-test.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar-test11.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar-test12.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar-test13.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar-test14.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar-test15.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar-test20.txt")
	// lines, _ := filereader.ReadFileToStringArray("oscar.txt")

	origonalNumbers = make([][]int, 0)
	for lineIndex, line := range lines {
		origonalNumbers = append(origonalNumbers, make([]int, len(line)))
		parts := strings.Split(line, "")

		for index, char := range parts {
			num, _ := strconv.Atoi(char)
			origonalNumbers[lineIndex][index] = num
		}

	}
}

// I was defeated!
func main() {
	defer timetrack.TimeTrack(time.Now(), "Oscar")
	problemOne()
	problemTwo()
}

var riskLevel int = 0

var interations int = 0

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Oscar - Problem One")

	m1 := make(map[string]int)

	recursive(len(origonalNumbers)-1, len(origonalNumbers[0])-1, m1)

	fmt.Printf("Answer = %d\n", riskLevel)
	fmt.Printf("total iterations =  = %d\n", interations)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Oscar - Problem two")
}

func recursive(row int, col int, trail map[string]int) {
	interations++

	newTrail := make(map[string]int)

	for key, val := range trail {
		newTrail[key] = val
	}

	key := fmt.Sprintf("%d%d", row, col)

	_, ok := newTrail[key]

	if ok {
		// DOnt continue, we have already been here.
		return
	}

	if row != 0 || col != 0 {
		newTrail[key] = origonalNumbers[row][col]
	}

	if row == 0 && col == 0 {
		total := getCount(newTrail)

		if total < riskLevel || riskLevel == 0 {
			riskLevel = total
		}
		return
	}

	if riskLevel > 0 {
		total := getCount(newTrail)

		if total > riskLevel {
			// too big.
			return
		}
	}

	if riskLevel != 0 {
		total := getCount(newTrail)
		if row+col+total > riskLevel {
			return
		}
	}

	canGoUp := row > 0
	canGoLeft := col > 0
	canGoRight := col < len(origonalNumbers[0])-1
	canGoDown := row < len(origonalNumbers)-1

	if canGoLeft {
		recursive(row, col-1, newTrail)
	}

	if canGoUp {
		recursive(row-1, col, newTrail)
	}

	if canGoDown {
		recursive(row+1, col, newTrail)
	}

	if canGoRight {
		recursive(row, col+1, newTrail)
	}
}

func getCount(m map[string]int) int {
	total := 0

	for _, val := range m {
		total += val
	}

	return total
}
