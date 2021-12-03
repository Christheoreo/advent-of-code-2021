package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var report [][]int

func init() {
	lines, err := filereader.ReadFileToStringArray("charlie.txt")

	if err != nil {
		panic(err)
	}

	report = make([][]int, len(lines))

	for index, line := range lines {
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			report[index] = append(report[index], num)
		}
	}

}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Charlie")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Charlie - Problem One")

	columns := len(report[0])
	gammaString := ""
	epsilonString := ""

	for i := 0; i < columns; i++ {
		col := getColumn(i, report)

		mostCommon := getMostCommonNumber(col)
		leastCommon := getLeastCommonNumber(col)

		gammaString += strconv.Itoa(mostCommon)
		epsilonString += strconv.Itoa(leastCommon)
	}

	// convert from binary to ints
	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)

	answer := gamma * epsilon

	fmt.Printf("Answer = %d\n", answer)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Charlie - Problem two")

	oxygenRatingString := ""
	scubaRatingString := ""

	col := getColumn(0, report)

	mostCommon := getMostCommonNumber(col)
	leastCommon := getLeastCommonNumber(col)

	oxygenRows := returnRowsThatContainXInCertainPoisition(0, mostCommon, report)
	scubaRatingRows := returnRowsThatContainXInCertainPoisition(0, leastCommon, report)

	columnIndex := 1
	for len(oxygenRows) > 1 {
		col := getColumn(columnIndex, oxygenRows)
		mostCommon := getMostCommonNumber(col)
		oxygenRows = returnRowsThatContainXInCertainPoisition(columnIndex, mostCommon, oxygenRows)
		columnIndex++
	}

	columnIndex = 1
	for len(scubaRatingRows) > 1 {
		col := getColumn(columnIndex, scubaRatingRows)
		leastCommon := getLeastCommonNumber(col)
		scubaRatingRows = returnRowsThatContainXInCertainPoisition(columnIndex, leastCommon, scubaRatingRows)
		columnIndex++
	}

	for _, num := range oxygenRows[0] {
		str := strconv.Itoa(num)
		oxygenRatingString += str
	}

	for _, num := range scubaRatingRows[0] {
		str := strconv.Itoa(num)
		scubaRatingString += str
	}

	// convert from binary to ints
	oxygen, _ := strconv.ParseInt(oxygenRatingString, 2, 64)
	scuba, _ := strconv.ParseInt(scubaRatingString, 2, 64)

	answer := oxygen * scuba

	fmt.Printf("Answer = %d\n", answer)

}

func getMostCommonNumber(numbers []int) (mostCommon int) {
	frequency := make(map[int]int)

	for _, num := range numbers {
		val, ok := frequency[num]
		if ok {
			frequency[num] = val + 1
			continue
		}
		frequency[num] = 1
	}

	mostCommon = 0

	for key, val := range frequency {
		if val == frequency[mostCommon] && key == 1 {
			mostCommon = key
			continue
		}
		if val > frequency[mostCommon] {
			mostCommon = key
			continue
		}
	}

	return
}

func getLeastCommonNumber(numbers []int) (leastCommon int) {
	frequency := make(map[int]int)

	for _, num := range numbers {
		val, ok := frequency[num]
		if ok {
			frequency[num] = val + 1
			continue
		}
		frequency[num] = 1
	}

	leastCommon = 0

	for key, val := range frequency {
		if val == frequency[leastCommon] && key == 0 {
			leastCommon = key
			continue
		}
		if val < frequency[leastCommon] {
			leastCommon = key
			continue
		}
	}

	return
}

func getColumn(column int, arr [][]int) []int {
	numbers := make([]int, len(arr))
	for rowIndex, row := range arr {
		numbers[rowIndex] = row[column]
	}
	return numbers
}

func returnRowsThatContainXInCertainPoisition(position int, number int, rows [][]int) [][]int {
	goodRows := make([][]int, 0)

	for _, row := range rows {
		if row[position] == number {
			goodRows = append(goodRows, row)
		}
	}

	return goodRows
}
