package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

type (
	Fold struct {
		Direction rune
		Value     int
	}
)

var paper [][]rune
var folds []Fold

func resetPaper() {
	lines, _ := filereader.ReadFileToStringArray("mike.txt")
	paper = make([][]rune, 0)
	folds = make([]Fold, 0)
	numbers := make([][]int, 0)

	largestX := 0
	largestY := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, ",") {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			numbers = append(numbers, []int{x, y})
			if x > largestX {
				largestX = x
			}
			if y > largestY {
				largestY = y
			}

			continue
		}

		s := line[11:]

		val, _ := strconv.Atoi(string(s[2:]))

		fold := Fold{
			Direction: rune(s[0]),
			Value:     val,
		}

		folds = append(folds, fold)

	}

	for row := 0; row <= largestY; row++ {
		paper = append(paper, make([]rune, largestX+1))
		for col := 0; col <= largestX; col++ {
			paper[row][col] = '.'
		}
	}

	for _, number := range numbers {
		paper[number[1]][number[0]] = '#'
	}

}
func init() {
	resetPaper()
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Mike")
	problemOne()
	resetPaper()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Mike - Problem One")
	paper = fold(paper, folds[0])
	holes := 0

	for _, row := range paper {
		for _, val := range row {
			if val == '#' {
				holes++
			}
		}
	}

	fmt.Printf("Totoal dots after the first fold = %d\n", holes)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Mike - Problem two")
	for _, f := range folds {
		paper = fold(paper, f)
	}

	for _, row := range paper {
		fmt.Printf("%v\n", string(row))
	}
}

func fold(paper [][]rune, fold Fold) [][]rune {

	if fold.Direction == 'y' {
		topHalf := paper[:fold.Value]
		bottomHalf := paper[fold.Value+1:]

		for topRow, bottomRow := len(topHalf)-1, 0; bottomRow < len(bottomHalf); topRow, bottomRow = topRow-1, bottomRow+1 {
			for col := 0; col < len(paper[0]); col++ {
				if topHalf[topRow][col] == '#' {
					continue
				}

				if bottomHalf[bottomRow][col] == '.' {
					continue
				}
				topHalf[topRow][col] = '#'
			}
		}
		return topHalf
	}
	// else its X..

	leftHalf := make([][]rune, len(paper))
	rightHalf := make([][]rune, len(paper))

	for rowIndex, row := range paper {
		for index, col := range row {
			if index < fold.Value {
				leftHalf[rowIndex] = append(leftHalf[rowIndex], col)
			} else if index > fold.Value {
				rightHalf[rowIndex] = append(rightHalf[rowIndex], col)
			}
		}
	}

	for row := 0; row < len(paper); row++ {
		for colLeft, colRight := len(leftHalf[0])-1, 0; colRight < len(rightHalf) && colLeft >= 0; colLeft, colRight = colLeft-1, colRight+1 {
			if leftHalf[row][colLeft] == '#' {
				continue
			}
			if rightHalf[row][colRight] == '.' {
				continue
			}
			leftHalf[row][colLeft] = '#'
		}
	}

	return leftHalf

}
