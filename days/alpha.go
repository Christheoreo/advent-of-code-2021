package days

import (
	"fmt"

	"github.com/Christheoreo/advent-of-code-2021/utils"
)

// implements IDay
type Alpha struct{}

func (a Alpha) Setup() {
	//
}

func (a Alpha) Run() {
	a.RunOne()
	a.RunTwo()
}

// Problem one
func (a Alpha) RunOne() {

	depthsStringArr, err := utils.ReadFileToStringArray("alpha.txt")

	if err != nil {
		panic(err)
	}

	depths, err := utils.ConvertStringArrayToIntArray(depthsStringArr)

	if err != nil {
		panic(err)
	}

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

func (a Alpha) RunTwo() {
	depthsStringArr, err := utils.ReadFileToStringArray("alpha.txt")

	if err != nil {
		panic(err)
	}

	depths, err := utils.ConvertStringArrayToIntArray(depthsStringArr)

	if err != nil {
		panic(err)
	}

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

func (a Alpha) GetName() string {
	return "Day 01"
}
