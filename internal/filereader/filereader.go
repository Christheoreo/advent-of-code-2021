package filereader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToStringArray(filepath string) ([]string, error) {

	mydir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fmt.Sprintf("%s/inputs/%s", mydir, filepath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
