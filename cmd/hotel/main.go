package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Christheoreo/advent-of-code-2021/internal/filereader"
	"github.com/Christheoreo/advent-of-code-2021/internal/timetrack"
)

var data []map[string][]string

func init() {
	lines, _ := filereader.ReadFileToStringArray("hotel.txt")

	for _, line := range lines {
		parts := strings.Split(line, "|")

		pattern := strings.TrimSpace(parts[0])
		output := strings.TrimSpace(parts[1])

		patternParts := strings.Split(pattern, " ")
		outputParts := strings.Split(output, " ")

		arr := make(map[string][]string)

		arr["pattern"] = patternParts
		arr["output"] = outputParts

		data = append(data, arr)

	}

}

func main() {
	defer timetrack.TimeTrack(time.Now(), "Hotel")
	problemOne()
	problemTwo()
}

func problemOne() {
	defer timetrack.TimeTrack(time.Now(), "Hotel - Problem One")

	total := 0

	for _, val := range data {
		output := val["output"]
		for _, str := range output {
			if len(str) == 2 || len(str) == 4 || len(str) == 3 || len(str) == 7 {
				total++
			}
		}

	}
	fmt.Printf("Number of occurances = %d\n", total)
}

func problemTwo() {
	defer timetrack.TimeTrack(time.Now(), "Hotel - Problem two")
	total := 0
	for _, val := range data {
		pattern := val["pattern"]

		knownValues := map[int]string{
			0: "",
			1: "",
			2: "",
			3: "",
			4: "",
			5: "",
			6: "",
			7: "",
			8: "",
			9: "",
		}

		knownPositions := map[string]string{
			"tl": "",
			"bl": "",
			"b":  "",
			"br": "",
			"tr": "",
			"t":  "",
			"m":  "",
		}

		// lets try to find the easy numbers
		loop := true
		for _, str := range pattern {
			strLen := len(str)
			if strLen == 2 {
				// it is a 1
				knownValues[1] = str
				// continue
			} else if strLen == 4 {
				// it is a 4
				knownValues[4] = str
				// continue
			} else if strLen == 3 {
				// it is a 7
				knownValues[7] = str
				// continue
			} else if strLen == 7 {
				// it is an 8
				knownValues[8] = str
				// continue
			}
		}
		// we can now work out positions of top by comparing 1 & 7
		oneVal := knownValues[1]
		sevenVal := knownValues[7]

		for _, char := range sevenVal {
			if !strings.ContainsRune(oneVal, char) {
				knownPositions["t"] = string(char)
			}
		}

		for loop {

			// we have the basis

			for _, str := range pattern {

				for i := 0; i < 10; i++ {
					if knownValues[i] != "" {
						continue
					}
				}
				strLen := len(str)

				// we can work out whih one 3 is, its the only one out of 2,3,5 that has all of 1 in it.

				if knownValues[3] == "" {
					if strLen != 5 {
						continue
					}
					a := string(knownValues[1][0])
					b := string(knownValues[1][1])
					if strings.Contains(str, a) && strings.Contains(str, b) {
						knownValues[3] = str
					}
				}

				if knownValues[3] == "" {
					continue
				}

				if knownValues[6] == "" {
					if strLen != 6 {
						continue
					}
					a := string(knownValues[1][0])
					b := string(knownValues[1][1])
					if !strings.Contains(str, a) || !strings.Contains(str, b) {
						knownValues[6] = str
					}
				}

				if knownValues[6] == "" {
					continue
				}

				//so here we know 1,3,4,6,7,8
				// so we will only be deailing with 0,2,5,9

				three := knownValues[3]
				four := knownValues[4]
				one := knownValues[1]

				charsToRemove := []string{string(one[0]), string(one[1]), knownPositions["t"]}

				for _, char := range charsToRemove {
					three = strings.Replace(three, char, "", 1)
					four = strings.Replace(four, char, "", 1)
				}

				// the common letter is our middle letter

				for _, char := range three {
					if strings.ContainsRune(four, char) {
						knownPositions["m"] = string(char)
					} else {
						// this is  b
						knownPositions["b"] = string(char)
					}
				}

				for _, char := range four {
					if !strings.ContainsRune(three, char) {
						// this is  tl
						knownPositions["tl"] = string(char)
					}
				}

				if strings.Contains(knownValues[6], string(one[0])) {
					knownPositions["br"] = string(one[0])
					knownPositions["tr"] = string(one[1])
				} else {
					knownPositions["br"] = string(one[1])
					knownPositions["tr"] = string(one[0])
				}

				// now we know t, tl, tr, br, m, b

				charsToRemove = []string{knownPositions["t"], knownPositions["tr"], knownPositions["tl"], knownPositions["m"], knownPositions["br"], knownPositions["b"]}

				eight := knownValues[8]
				for _, char := range charsToRemove {
					eight = strings.Replace(eight, char, "", 1)
				}

				knownPositions["bl"] = eight
				loop = false
				break
			}

		}

		zero := fmt.Sprintf("%s%s%s%s%s%s", knownPositions["t"], knownPositions["tr"], knownPositions["tl"], knownPositions["br"], knownPositions["bl"], knownPositions["b"])
		one := fmt.Sprintf("%s%s", knownPositions["tr"], knownPositions["br"])

		two := fmt.Sprintf("%s%s%s%s%s", knownPositions["t"], knownPositions["tr"], knownPositions["m"], knownPositions["bl"], knownPositions["b"])

		three := fmt.Sprintf("%s%s%s%s%s", knownPositions["t"], knownPositions["tr"], knownPositions["m"], knownPositions["br"], knownPositions["b"])

		four := fmt.Sprintf("%s%s%s%s", knownPositions["tl"], knownPositions["m"], knownPositions["tr"], knownPositions["br"])

		five := fmt.Sprintf("%s%s%s%s%s", knownPositions["t"], knownPositions["tl"], knownPositions["m"], knownPositions["br"], knownPositions["b"])

		six := fmt.Sprintf("%s%s%s%s%s%s", knownPositions["t"], knownPositions["tl"], knownPositions["m"], knownPositions["bl"], knownPositions["b"], knownPositions["br"])

		seven := fmt.Sprintf("%s%s%s", knownPositions["t"], knownPositions["tr"], knownPositions["br"])

		eight := fmt.Sprintf("%s%s%s%s%s%s%s", knownPositions["t"], knownPositions["tr"], knownPositions["tl"], knownPositions["br"], knownPositions["bl"], knownPositions["b"], knownPositions["m"])

		nine := fmt.Sprintf("%s%s%s%s%s%s", knownPositions["t"], knownPositions["tr"], knownPositions["tl"], knownPositions["m"], knownPositions["b"], knownPositions["br"])

		output := val["output"]

		arr := []string{zero, one, two, three, four, five, six, seven, eight, nine}

		var stringTotal strings.Builder

		for _, val := range output {

			for index, arrayValue := range arr {

				if len(arrayValue) != len(val) {
					continue
				}
				match := true
				for _, char := range val {
					if !strings.ContainsRune(arrayValue, char) {
						match = false
						break
					}
				}

				if match {
					stringIndex := strconv.Itoa(index)
					stringTotal.WriteString(stringIndex)
					break
				}
			}
		}

		num, _ := strconv.Atoi(stringTotal.String())
		total += num

	}
	fmt.Printf("Answer = %d\n", total)
}
