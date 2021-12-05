package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

type Cooridinate struct {
	X int
	Y int
}

type CooridinateWithOverlap struct {
	X            int
	Y            int
	OverlapCount int
}

var coordinates [][]Cooridinate

var grid [][]int

func init() {
	coordinates = make([][]Cooridinate, 0)
	lines, err := filereader.ReadFileToStringArray("echo.txt")

	if err != nil {
		panic(err)
	}

	grid = make([][]int, 0)

	for _, line := range lines {
		parts := strings.Split(line, " -> ")

		pair := make([]Cooridinate, 2)

		for index, part := range parts {
			points := strings.Split(part, ",")

			x, _ := strconv.Atoi(points[0])
			y, _ := strconv.Atoi(points[1])

			co := Cooridinate{
				X: x, Y: y,
			}

			pair[index] = co
		}

		coordinates = append(coordinates, pair)
	}

	// get the smallest coodinate and the largest so we can plot a grid
	bottomRight := Cooridinate{
		X: 1,
		Y: 1,
	}
	for _, pair := range coordinates {

		for _, co := range pair {

			if co.X > bottomRight.X {
				bottomRight.X = co.X
			}
			if co.Y > bottomRight.Y {
				bottomRight.Y = co.Y
			}
		}

	}

	for row := 0; row <= bottomRight.Y; row++ {
		grid = append(grid, make([]int, 0))

		for col := 0; col <= bottomRight.X; col++ {
			grid[row] = append(grid[row], 0)
		}
	}
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Echo")
	// Make sure to only run at a time
	// problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Echo - Problem One")
	lines := returnLines("both", coordinates)

	points := drawOutLines(lines)

	// fmt.Println(points)

	for _, point := range points {
		// gridTwo[point.X][point.Y] += 1
		grid[point.Y][point.X] += 1
	}

	answer := returnPointsWithXOrMoreOverlaps(2, grid)
	fmt.Printf("Answer = %d", answer)

}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Echo - Problem two")

	lines := returnLines("all", coordinates)

	points := drawOutLines(lines)

	for _, point := range points {
		grid[point.Y][point.X]++
	}
	answer := returnPointsWithXOrMoreOverlaps(2, grid)
	fmt.Printf("Answer = %d", answer)
}

func returnLines(lineType string, coordinates [][]Cooridinate) [][]Cooridinate {

	cos := make([][]Cooridinate, 0)

	for _, pair := range coordinates {
		start := pair[0]
		end := pair[1]

		if lineType == "horizontal" {
			if start.X == end.X {
				cos = append(cos, pair)
			}
		} else if lineType == "vertical" {
			if start.Y == end.Y {
				cos = append(cos, pair)
			}
		} else if lineType == "both" {
			if start.Y == end.Y || start.X == end.X {
				cos = append(cos, pair)
			}
		} else if lineType == "all" {
			if start.Y == end.Y || start.X == end.X {
				cos = append(cos, pair)
				continue
			}

			xDiff := 0
			yDiff := 0
			if start.X > end.X {
				xDiff = start.X - end.X
			} else {
				xDiff = end.X - start.X
			}
			if start.Y > end.Y {
				yDiff = start.Y - end.Y
			} else {
				yDiff = end.Y - start.Y
			}

			// as long as the differenes are the same number (ignore the minus) then thats all good

			if math.Abs(float64(xDiff)) == math.Abs(float64(yDiff)) {
				cos = append(cos, pair)
			}

		}

	}

	return cos
}

// returns all the coordinates that will be croseed in the lines.
func drawOutLines(coordinates [][]Cooridinate) []Cooridinate {
	lines := make([]Cooridinate, 0)
	for _, pair := range coordinates {
		start := pair[0]
		end := pair[1]

		lines = append(lines, start, end)

		if start.X == end.X {
			startingPoint := start.Y + 1
			endingPoint := end.Y
			if start.Y > end.Y {
				startingPoint = end.Y + 1
				endingPoint = start.Y
			}
			for i := startingPoint; i < endingPoint; i++ {
				lines = append(lines, Cooridinate{
					X: start.X,
					Y: i,
				})
			}
		} else if start.Y == end.Y {
			startingPoint := start.X + 1
			endingPoint := end.X
			if start.X > end.X {
				startingPoint = end.X + 1
				endingPoint = start.X
			}
			for i := startingPoint; i < endingPoint; i++ {
				lines = append(lines, Cooridinate{
					X: i,
					Y: start.Y,
				})
			}
		} else {
			// diagonal

			direction := "top left"

			if start.X > end.X {
				// we are on the right of end end point
				if start.Y < end.Y {
					direction = "top right"
				} else {
					direction = "bottom right"
				}
			} else {
				if start.Y > end.Y {
					direction = "bottom left"
				}
			}

			if direction == "top left" {
				//
				// start is on the top left of end
				// we need to add to the x and add to the  y
				xDistance := (end.X - start.X) - 1

				for i := 1; i <= xDistance; i++ {
					x := start.X + i
					y := start.Y + i
					lines = append(lines, Cooridinate{
						X: x,
						Y: y,
					})
				}
			} else if direction == "top right" {
				//
				// start is on the top right of end
				// we need to minus the x and add to the  y
				xDistance := (start.X - end.X) - 1
				for i := 1; i <= xDistance; i++ {
					x := start.X - i
					y := start.Y + i
					lines = append(lines, Cooridinate{
						X: x,
						Y: y,
					})
				}
			} else if direction == "bottom left" {
				//
				// start is on the bottom left of end
				// we need to add to the x and minus the  y
				xDistance := (end.X - start.X) - 1
				for i := 1; i <= xDistance; i++ {
					x := start.X + i
					y := start.Y - i
					lines = append(lines, Cooridinate{
						X: x,
						Y: y,
					})
				}
			} else {
				//
				// start is on the bottom right of end
				// we need to minus the x and minus the  y
				xDistance := (start.X - end.X) - 1
				for i := 1; i <= xDistance; i++ {
					x := start.X - i
					y := start.Y - i
					lines = append(lines, Cooridinate{
						X: x,
						Y: y,
					})
				}
			}

		}
	}
	return lines
}

func returnPointsWithXOrMoreOverlaps(overlapCount int, grid [][]int) int {
	count := 0

	for _, row := range grid {
		for _, col := range row {
			if col >= overlapCount {
				count++
			}
		}
	}
	return count
}
