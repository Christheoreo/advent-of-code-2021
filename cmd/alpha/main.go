package main

import (
	"fmt"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/converter"
	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var depths []int

func init() {
	depthsStringArr, err := filereader.ReadFileToStringArray("alpha.txt")

	if err != nil {
		panic(err)
	}

	depths, err = converter.ConvertStringArrayToIntArray(depthsStringArr)

	if err != nil {
		panic(err)
	}
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Alpha")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Alpha - Problem one")

	bar := depths[0]
	incremenets := 0

	for _, depth := range depths {
		if depth > bar {
			incremenets++
		}
		bar = depth
	}

	fmt.Printf("Total incremenets = %d", incremenets)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Alpha - Problem two")

	depthsOfThree := make([]int, 0)
	total := 0

	for i := 0; i < len(depths)-2; i++ {
		total = depths[i] + depths[i+1] + depths[i+2]
		depthsOfThree = append(depthsOfThree, total)
	}

	bar := depthsOfThree[0]
	incremenets := 0

	for _, depth := range depthsOfThree {
		if depth > bar {
			incremenets++
		}
		bar = depth
	}

	fmt.Printf("Total incremenets = %d", incremenets)
}
