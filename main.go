package main

import (
	"fmt"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/days"
)

type IDay interface {
	Run()
	RunOne()
	RunTwo()
	GetName() string
}

func main() {
	day := days.Alpha{}
	solveProblem(day)
}

func solveProblem(day IDay) {
	defer timeTrack(time.Now(), day.GetName())
	// Run problem one for now
	// day.RunOne()
	day.RunTwo()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n-----\n%s took %s\n-----\n", name, elapsed)
}
