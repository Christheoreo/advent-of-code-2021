package main

import (
	"fmt"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/days"
)

type IDay interface {
	Setup()
	Run()
	RunOne()
	RunTwo()
	GetName() string
}

func main() {
	day := days.Bravo{}
	solveProblem(day)
}

func solveProblem(day IDay) {
	defer timeTrack(time.Now(), day.GetName())
	day.Setup()
	// Run problem one for now
	// day.RunOne()
	// day.RunTwo()
	day.Run()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n-----\n%s took %s\n-----\n", name, elapsed)
}
