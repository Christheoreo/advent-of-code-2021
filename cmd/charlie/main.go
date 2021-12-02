package main

import (
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

func init() {
	//
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Charlie")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Charlie - Problem One")
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Charlie - Problem two")
}
