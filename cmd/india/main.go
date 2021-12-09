package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var numbers [][]int

func init() {
	numbers = make([][]int, 0)
	lines, _ := filereader.ReadFileToStringArray("india.txt")

	for lineIndex, line := range lines {
		parts := strings.Split(line, "")
		numbers = append(numbers, make([]int, 0))

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers[lineIndex] = append(numbers[lineIndex], num)
		}
	}
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "india")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "india - Problem One")

	lowestPoints := getLowestPoints(numbers)

	riskLevel := 0
	for _, point := range lowestPoints {
		riskLevel += (point[2] + 1)
	}

	fmt.Printf("Risk level = %d\n", riskLevel)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "india - Problem two")

	positionsOfLowestPoints := getLowestPoints(numbers)

	largestBasins := make([]int, 0)

	for _, pos := range positionsOfLowestPoints {
		key := fmt.Sprintf("%d,%d", pos[1], pos[1])
		m := map[string]int{
			key: pos[2],
		}
		positions := getBasinCount(pos[0], pos[1], pos[2], numbers, m)

		// get the unique poistions
		largestBasins = append(largestBasins, len(positions))
	}

	sort.Ints(largestBasins)

	largest := largestBasins[len(largestBasins)-3:]

	fmt.Printf("The product of the largest 3 basins (%d,%d,%d) = %d\n", largest[0], largest[1], largest[2], largest[0]*largest[1]*largest[2])
}

func getLowestPoints(numbers [][]int) [][]int {
	positionsOfLowestPoints := make([][]int, 0)

	for rowIndex, row := range numbers {
		for colIndex, number := range row {
			if colIndex != 0 {
				if number >= row[colIndex-1] {
					continue
				}
			}

			if rowIndex != 0 {
				if number >= numbers[rowIndex-1][colIndex] {
					continue
				}
			}

			if colIndex != len(row)-1 {
				if number >= row[colIndex+1] {
					continue
				}
			}

			if rowIndex != len(numbers)-1 {
				if number >= numbers[rowIndex+1][colIndex] {
					continue
				}
			}

			positionsOfLowestPoints = append(positionsOfLowestPoints, []int{colIndex, rowIndex, number})
		}
	}

	return positionsOfLowestPoints
}

func helper(num int, number int, rowIndex int, colIndex int, positions map[string]int) map[string]int {
	key := fmt.Sprintf("%d,%d", rowIndex, colIndex)

	_, ok := positions[key]

	if ok {
		return positions
	}
	positions[key] = num
	return getBasinCount(colIndex, rowIndex, num, numbers, positions)
}
func getBasinCount(colIndex int, rowIndex int, number int, numbers [][]int, positions map[string]int) map[string]int {
	row := numbers[rowIndex]

	if colIndex != 0 {
		num := row[colIndex-1]
		if num-number >= 0 && num != 9 {
			positions = helper(num, number, rowIndex, colIndex-1, positions)
		}
	}

	if rowIndex != 0 {
		num := numbers[rowIndex-1][colIndex]
		if num-number >= 0 && num != 9 {
			positions = helper(num, number, rowIndex-1, colIndex, positions)
		}
	}

	if colIndex != len(row)-1 {
		num := row[colIndex+1]
		if num-number >= 0 && num != 9 {
			positions = helper(num, number, rowIndex, colIndex+1, positions)
		}
	}

	if rowIndex != len(numbers)-1 {
		num := numbers[rowIndex+1][colIndex]
		if num-number >= 0 && num != 9 {
			positions = helper(num, number, rowIndex+1, colIndex, positions)
		}

	}
	return positions
}
