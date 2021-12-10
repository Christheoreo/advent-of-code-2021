package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var lines []string

type CorruptLine struct {
	Char rune
	Line string
}

func init() {
	lines, _ = filereader.ReadFileToStringArray("juliet.txt")
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Juliet")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Juliet - Problem One")

	corruptedLines := make([]CorruptLine, 0)
	mapping := make([]rune, 0)

	for _, line := range lines {
		corrupted := false
		var r rune
		for index, char := range line {

			if index == 0 {
				if !isOpeningChar(char) {
					// corrupted
					r = char
					corrupted = true
					break
				}
				mapping = append(mapping, char)
				continue
			}

			// we are either looking for an opening bracket or a closing bracket for the last opened bracket
			if isClosingChar(char) {
				lastChar := getLastOpenedChar(mapping)

				pair := isPair(lastChar, char)

				if !pair {
					r = char
					corrupted = true
					break
				}
				// remove the last opened chaar from the array
				mapping = removeLastOpenedChar(mapping)
				continue
			}

			// its an open char
			mapping = append(mapping, char)
			continue
		}

		if corrupted {
			corruptedLines = append(corruptedLines, CorruptLine{
				Char: r,
				Line: line,
			})
		}
	}
	total := 0
	for _, line := range corruptedLines {
		val := getScore(line.Char, true)
		subTotal := val
		total += subTotal
	}
	fmt.Printf("The total syntax error score = %d\n", total)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Juliet - Problem two")
	incompleteLines := make([][]rune, 0)

	for _, line := range lines {
		corrupted := false
		mapping := make([]rune, 0)
		for index, char := range line {

			if index == 0 {
				if !isOpeningChar(char) {
					// corrupted
					corrupted = true
					break
				}
				mapping = append(mapping, char)
				continue
			}

			// we are either looking for an opening bracket or a closing bracket for the last opened bracket

			if isClosingChar(char) {
				lastChar := getLastOpenedChar(mapping)

				pair := isPair(lastChar, char)

				if !pair {
					corrupted = true
					break
				}
				// remove the last opened chaar from the array
				mapping = removeLastOpenedChar(mapping)
				continue
			}

			// its an open char
			mapping = append(mapping, char)
			continue
		}

		if !corrupted {
			l := make([]rune, 0)
			for i := len(mapping) - 1; i >= 0; i-- {
				val := mapping[i]
				p := getPair(val)
				l = append(l, p)
			}
			incompleteLines = append(incompleteLines, l)
		}
	}
	totals := make([]int, 0)
	for _, line := range incompleteLines {
		lineTotal := 0
		for _, char := range line {
			lineTotal *= 5
			val := getScore(char, false)
			lineTotal += val
		}
		totals = append(totals, lineTotal)
	}

	sort.Ints(totals)
	length := len(totals) - 1

	answer := totals[(length / 2)]
	fmt.Printf("The middle score is = %d\n", answer)
}

func isOpeningChar(r rune) bool {
	openingChars := "<([{"
	return strings.ContainsRune(openingChars, r)
}

func isClosingChar(r rune) bool {
	closingChars := "}])>"
	return strings.ContainsRune(closingChars, r)
}

func getLastOpenedChar(arr []rune) rune {
	var r rune

	for i := len(arr) - 1; i >= 0; i-- {
		if isOpeningChar(arr[i]) {
			r = arr[i]
			break
		}
	}

	return r
}

func removeLastOpenedChar(arr []rune) []rune {
	index := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if isOpeningChar(arr[i]) {
			index = i
			break
		}
	}
	newArr := make([]rune, 0)
	for i, v := range arr {
		if i != index {
			newArr = append(newArr, v)
		}
	}

	return newArr
}

func isPair(char rune, charTwo rune) bool {
	if char == '{' && charTwo == '}' ||
		char == '[' && charTwo == ']' ||
		char == '(' && charTwo == ')' ||
		char == '<' && charTwo == '>' {
		return true
	}
	return false
}

func getPair(char rune) rune {
	if char == '{' {
		return '}'
	} else if char == '}' {
		return '{'
	} else if char == '[' {
		return ']'
	} else if char == ']' {
		return '['
	} else if char == '(' {
		return ')'
	} else if char == ')' {
		return '('
	} else if char == '<' {
		return '>'
	} else {
		return '<'
	}
}

func getScore(char rune, problemOne bool) int {
	if char == ')' {
		if problemOne {
			return 3
		}
		return 1

	} else if char == ']' {
		if problemOne {
			return 57
		} else {
			return 2
		}
	} else if char == '}' {
		if problemOne {
			return 1197
		}
		return 3
	} else if char == '>' {
		if problemOne {
			return 25137
		}
		return 4
	}
	return 0
}
