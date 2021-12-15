package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var origonalTemplate string
var insertionRules [][]string

type X struct {
	ToInsert string
	Index    int
}

func init() {
	setup()
}

func setup() {
	lines, _ := filereader.ReadFileToStringArray("november.txt")

	origonalTemplate = lines[0]
	lines = lines[2:]
	insertionRules = make([][]string, len(lines))

	for index, line := range lines {
		parts := strings.Split(line, "->")

		a := strings.TrimSpace(parts[0])
		b := strings.TrimSpace(parts[1])

		insertionRules[index] = []string{a, b}

	}
}
func main() {
	defer timetrack.TimeTrack(time.Now(), "November")
	problemOne()
	setup()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "November - Problem One")
	t := strings.Split(origonalTemplate, "")
	template := make(map[string]int)

	rules := make(map[string]string)

	for i := 0; i <= len(t)-2; i++ {
		key := fmt.Sprintf("%s%s", t[i], t[i+1])
		_, ok := template[key]

		if !ok {
			template[key] = 1
		} else {
			template[key]++
		}
	}

	for _, r := range insertionRules {
		rules[r[0]] = r[1]
	}
	for i := 0; i < 10; i++ {
		template = completeStep(template, rules)
	}

	largest, smallest := getCounts(template)

	answer := largest - smallest

	fmt.Printf("Tht largest (%d) - the smallest (%d) is %d\n", largest, smallest, answer)

}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "November - Problem two")
	t := strings.Split(origonalTemplate, "")
	template := make(map[string]int)

	rules := make(map[string]string)

	for i := 0; i <= len(t)-2; i++ {
		key := fmt.Sprintf("%s%s", t[i], t[i+1])
		_, ok := template[key]

		if !ok {
			template[key] = 1
		} else {
			template[key]++
		}
	}
	for _, r := range insertionRules {
		rules[r[0]] = r[1]
	}
	for i := 0; i < 40; i++ {
		template = completeStep(template, rules)
	}

	largest, smallest := getCounts(template)

	answer := largest - smallest

	fmt.Printf("Tht largest (%d) - the smallest (%d) is %d\n", largest, smallest, answer)
}

func completeStep(template map[string]int, rules map[string]string) map[string]int {
	newTemplate := make(map[string]int)

	for key, val := range template {
		char, ok := rules[key]

		if !ok {
			continue
		}

		patternA := fmt.Sprintf("%s%s", string(key[0]), char)
		patternB := fmt.Sprintf("%s%s", char, string(key[1]))

		_, aok := newTemplate[patternA]
		_, bok := newTemplate[patternB]

		if aok {
			newTemplate[patternA] += val
		} else {
			newTemplate[patternA] = val
		}

		if bok {
			newTemplate[patternB] += val
		} else {
			newTemplate[patternB] = val
		}

	}

	return newTemplate
}

func getCounts(data map[string]int) (int, int) {

	counts := make(map[string]int)

	largest := 0
	smallest := 0

	for key, val := range data {

		_, ok := counts[string(key[0])]

		if !ok {
			counts[string(key[0])] = val
		} else {
			counts[string(key[0])] += val
		}
	}

	lastKey := string(origonalTemplate[len(origonalTemplate)-1])

	_, ok := counts[lastKey]

	if ok {
		counts[lastKey] += 1
	} else {
		counts[lastKey] = 1
	}

	for _, val := range counts {

		if largest == 0 {
			largest = val
			continue
		}

		if smallest == 0 {
			smallest = val
			continue
		}

		if val > largest {
			largest = val
			continue
		}

		if val < smallest {
			smallest = val
			continue
		}
	}

	return largest, smallest
}
