package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var caves map[string][]string

var routesThatMakeItToEnd [][]string

func init() {
	lines, _ := filereader.ReadFileToStringArray("lima.txt")
	caves = make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		name := parts[0]
		subCave := parts[1]

		_, ok := caves[name]
		if ok {
			caves[name] = append(caves[name], subCave)
		} else {
			caves[name] = make([]string, 1)
			caves[name][0] = subCave
		}

		_, ok = caves[subCave]

		if ok {
			caves[subCave] = append(caves[subCave], name)
		} else {
			caves[subCave] = make([]string, 1)
			caves[subCave][0] = name
		}

	}

	caves["end"] = make([]string, 0)
}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Lima")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Lima - Problem One")
	exploreCave("start", make([]string, 0))
	fmt.Printf("Total routes = %d\n", len(routesThatMakeItToEnd))
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Lima - Problem two")
	routesThatMakeItToEnd = make([][]string, 0)
	exploreCaveTwo("start", make([]string, 0))
	fmt.Printf("Total routes = %d\n", len(routesThatMakeItToEnd))
}

func isBigCave(cave string) bool {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	contains := true
	for _, char := range cave {
		if !strings.ContainsRune(str, char) {
			contains = false
			break
		}
	}

	return contains
}

func exploreCave(cave string, currentRoute []string) {

	newRoute := make([]string, len(currentRoute)+1)
	copy(newRoute, currentRoute)

	newRoute = append(newRoute, cave)
	possCaves := caves[cave]

	hasEnd := false
	for _, possCave := range possCaves {
		if possCave == "end" {
			hasEnd = true
			continue
		}
		if isBigCave(possCave) {
			exploreCave(possCave, newRoute)
			continue
		}

		if !inStringArray(possCave, newRoute) {
			exploreCave(possCave, newRoute)
		}
	}

	if hasEnd {
		if !inStringArray("end", newRoute) {
			newRoute = append(newRoute, "end")
			routesThatMakeItToEnd = append(routesThatMakeItToEnd, newRoute)
			return
		}
	}

}

func exploreCaveTwo(cave string, currentRoute []string) {

	newRoute := make([]string, len(currentRoute)+1)
	copy(newRoute, currentRoute)

	newRoute = append(newRoute, cave)
	possCaves := caves[cave]

	hasEnd := false
	for _, possCave := range possCaves {
		if possCave == "end" {
			hasEnd = true
			continue
		}
		if isBigCave(possCave) {
			exploreCaveTwo(possCave, newRoute)
			continue
		}

		if !inStringArray(possCave, newRoute) {
			exploreCaveTwo(possCave, newRoute)
		} else if possCave != "start" && possCave != "end" {
			//
			appearsMoreThanOnce := false
			for _, pc := range newRoute {
				if !isBigCave(pc) && pc != "start" && pc != "end" {
					c := countOccurences(pc, newRoute)
					appearsMoreThanOnce = c > 1
					if appearsMoreThanOnce {
						break
					}
				}
			}
			if !appearsMoreThanOnce {
				exploreCaveTwo(possCave, newRoute)
			}
		}
	}

	if hasEnd {
		if !inStringArray("end", newRoute) {
			newRoute = append(newRoute, "end")
			routesThatMakeItToEnd = append(routesThatMakeItToEnd, newRoute)
			return
		}
	}

}

func inStringArray(str string, arr []string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

func countOccurences(char string, arr []string) int {
	count := 0

	for _, val := range arr {
		if val == char {
			count++
		}
	}
	return count
}
