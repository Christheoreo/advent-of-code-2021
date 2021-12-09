package main

import (
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

func init() {
	//
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Juliet")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Juliet - Problem One")
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Juliet - Problem two")
}
