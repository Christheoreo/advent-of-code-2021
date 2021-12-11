package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var numbers [][]int

func init() {
	resetNumbers()
}

func resetNumbers() {
	lines, _ := filereader.ReadFileToStringArray("kilo.txt")
	numbers = make([][]int, 0)
	for index, line := range lines {
		parts := strings.Split(line, "")
		numbers = append(numbers, make([]int, 0))

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers[index] = append(numbers[index], num)
		}
	}
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Kilo")
	problemOne()
	resetNumbers()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Kilo - Problem One")
	flashCount := 0

	for i := 1; i <= 100; i++ {
		for rowIndex := 0; rowIndex < len(numbers); rowIndex++ {
			row := numbers[rowIndex]
			for colIndex := 0; colIndex < len(row); colIndex++ {
				numbers[rowIndex][colIndex] += 1
			}
		}

		for rowIndex := 0; rowIndex < len(numbers); rowIndex++ {
			for colIndex := 0; colIndex < len(numbers[rowIndex]); colIndex++ {
				numbers = recursivelyAddOne(rowIndex, colIndex, numbers)
			}
		}

		for rowIndex := 0; rowIndex < len(numbers); rowIndex++ {
			row := numbers[rowIndex]
			for colIndex := 0; colIndex < len(row); colIndex++ {

				if numbers[rowIndex][colIndex] == -1 {
					numbers[rowIndex][colIndex] = 0
					flashCount++
				}
			}
		}
	}
	fmt.Printf("Total flashes after 100 goes is %d\n", flashCount)

}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Kilo - Problem two")

	answer := 0
	i := 1
	for answer == 0 {

		for rowIndex := 0; rowIndex < len(numbers); rowIndex++ {
			row := numbers[rowIndex]
			for colIndex := 0; colIndex < len(row); colIndex++ {
				numbers[rowIndex][colIndex] += 1
			}
		}
		for rowIndex := 0; rowIndex < len(numbers); rowIndex++ {
			for colIndex := 0; colIndex < len(numbers[rowIndex]); colIndex++ {
				numbers = recursivelyAddOne(rowIndex, colIndex, numbers)
			}
		}
		allFlashed := true
		for rowIndex := 0; rowIndex < len(numbers); rowIndex++ {
			row := numbers[rowIndex]
			for colIndex := 0; colIndex < len(row); colIndex++ {
				if numbers[rowIndex][colIndex] == -1 {
					numbers[rowIndex][colIndex] = 0
				} else {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			answer = i
			break
		}
		i++
	}
	fmt.Printf("Totak cycles = %d\n", answer)
}

func recursivelyAddOne(rowIndex int, colIndex int, numbers [][]int) [][]int {

	if numbers[rowIndex][colIndex] < 10 {
		return numbers
	}
	newNumbers := numbers
	possibilities := [][]int{
		// top left
		{rowIndex - 1, colIndex - 1},
		//top
		{rowIndex - 1, colIndex},
		// top right
		{rowIndex - 1, colIndex + 1},
		// Left
		{rowIndex, colIndex - 1},
		// Right
		{rowIndex, colIndex + 1},
		// Bottom left
		{rowIndex + 1, colIndex - 1},
		// Botton
		{rowIndex + 1, colIndex},
		// Bottom Right
		{rowIndex + 1, colIndex + 1},
	}
	// Add 1 to the surrounding numbers

	newNumbers[rowIndex][colIndex] = -1

	for _, p := range possibilities {
		if p[0] < 0 || p[1] < 0 || p[0] >= len(numbers) || p[1] >= len(numbers[rowIndex]) {
			continue
		}
		// We only want to add 1 to those greater than minus 1
		if newNumbers[p[0]][p[1]] > -1 {
			newNumbers[p[0]][p[1]] += 1
		}
		if newNumbers[p[0]][p[1]] > 9 {
			newNumbers = recursivelyAddOne(p[0], p[1], newNumbers)
		}
	}

	return newNumbers
}
