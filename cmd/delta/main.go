package main

import (
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

func init() {
	//
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Delta")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Delta - Problem One")
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Delta - Problem two")
}
