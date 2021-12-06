package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var timers []int

func init() {
	resetData()
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "foxtrot")
	problemOne()
	problemTwo()
}

func resetData() {
	timers = make([]int, 0)
	input, err := filereader.ReadFileToStringArray("foxtrot.txt")

	if err != nil {
		panic(err)
	}

	numberString := input[0]

	numberStringArr := strings.Split(numberString, ",")

	for _, numString := range numberStringArr {
		num, _ := strconv.Atoi(numString)
		timers = append(timers, num)
	}
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "foxtrot - Problem One")
	days := 80
	answer := calculate(days)
	fmt.Printf("Total fish at %d days = %d", days, answer)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "foxtrot - Problem two")
	days := 256
	answer := calculate(days)
	fmt.Printf("Total fish at %d days = %d", days, answer)
}

func calculate(days int) int {
	buckets := make(map[int]int)
	for _, timer := range timers {
		buckets[timer] += 1
	}

	for day := 1; day <= days; day++ {
		keyOrder := []int{8, 7, 6, 5, 4, 3, 2, 1, 0}
		newBuckets := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
			8: 0,
		}

		for _, key := range keyOrder {
			_, ok := buckets[key]
			if !ok {
				buckets[key] = 0
			}

			if key != 8 {
				newBuckets[key] = buckets[key+1]
			}

			if key == 0 {
				newBuckets[6] += buckets[key]
				newBuckets[8] += buckets[key]
			}
		}

		buckets = newBuckets

	}

	answer := 0

	for key := range buckets {
		answer += buckets[key]
	}
	return answer
}
