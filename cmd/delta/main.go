package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/converter"
	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

type (
	BoardWinOrder struct {
		Board   Board
		Callout int
	}
	Board struct {
		Rows []Row
	}
	Row struct {
		Numbers []Num
	}
	Num struct {
		Number int
		Marked bool
	}
)

var callouts []int
var boards []Board

func init() {
	calloutsStringArr, err := filereader.ReadFileToStringArray("delta-callouts.txt")

	if err != nil {
		panic(err)
	}

	calloutsString := calloutsStringArr[0]

	charArray := strings.Split(calloutsString, ",")

	callouts, err = converter.ConvertStringArrayToIntArray(charArray)

	if err != nil {
		panic(err)
	}

	boardsStringArr, err := filereader.ReadFileToStringArray("delta.txt")

	if err != nil {
		panic(err)
	}
	boards = make([]Board, 1)

	for _, line := range boardsStringArr {
		parts := strings.Split(line, " ")

		if line == "" {
			// make a new baord
			boards = append(boards, Board{
				Rows: make([]Row, 0),
			})
			continue
		}

		newLine := make([]int, 0)

		for _, part := range parts {

			if part == "" {
				continue
			}

			num, _ := strconv.Atoi(part)

			newLine = append(newLine, num)

		}
		if len(newLine) == 5 {
			row := Row{
				Numbers: make([]Num, len(newLine)),
			}
			for index, num := range newLine {

				row.Numbers[index] = Num{
					Number: num,
					Marked: false,
				}

			}
			boards[len(boards)-1].Rows = append(boards[len(boards)-1].Rows, row)
		}
	}
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Delta")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Delta - Problem One")

	completed := false

	answer := 0
	for calloutIndex, callout := range callouts {
		if completed {
			break
		}
		for _, board := range boards {
			board = findAndMarkNumberInBoard(board, callout)

			if calloutIndex < 4 {
				continue
			}

			completed = isBoardCompleted(board)

			if completed {
				answer = callout * findSumOfUnmarkedNumbers(board)
				break
			}

		}
	}
	fmt.Printf("Answer = %d", answer)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Delta - Problem two")

	boardWinOrders := make([]BoardWinOrder, len(boards))

	completed := false

	answer := 0
	for calloutIndex, callout := range callouts {
		for _, board := range boards {
			if isBoardCompleted(board) {
				continue
			}
			board = findAndMarkNumberInBoard(board, callout)

			if calloutIndex < 4 {
				continue
			}

			completed = isBoardCompleted(board)

			if completed {
				boardWinOrders = append(boardWinOrders, BoardWinOrder{
					Callout: callout,
					Board:   board,
				})
			}

		}
	}

	lastBoard := boardWinOrders[len(boardWinOrders)-1]

	answer = lastBoard.Callout * findSumOfUnmarkedNumbers(lastBoard.Board)
	fmt.Printf("Answer = %d", answer)
}

func findAndMarkNumberInBoard(board Board, x int) Board {
	newBoard := board
	for rowIndex, row := range board.Rows {
		for numberIndex, number := range row.Numbers {
			if number.Number == x {
				number.Marked = true
				newBoard.Rows[rowIndex].Numbers[numberIndex].Marked = true
			}
		}
	}
	return newBoard
}

func isBoardCompleted(board Board) bool {
	numbers := make([][]Num, 0)

	for _, row := range board.Rows {
		numbers = append(numbers, row.Numbers)
	}

	// check if the rows are done

	for _, row := range numbers {
		allMarked := true
		for _, num := range row {
			if !num.Marked {
				allMarked = false
				break
			}
		}

		if allMarked {
			return true
		}
	}

	// check for columns

	for colIndex := 0; colIndex < len(numbers[0]); colIndex++ {
		allMarked := true
		for _, row := range numbers {
			if !row[colIndex].Marked {
				allMarked = false
				break
			}
		}
		if allMarked {
			return true
		}
	}
	return false
}

func findSumOfUnmarkedNumbers(board Board) int {
	total := 0

	for _, row := range board.Rows {
		for _, num := range row.Numbers {

			if !num.Marked {
				total += num.Number
			}
		}
	}

	return total
}
