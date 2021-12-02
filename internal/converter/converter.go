package converter

import "strconv"

func ConvertStringArrayToIntArray(arr []string) ([]int, error) {
	intArr := make([]int, len(arr))

	for i, str := range arr {
		val, err := strconv.Atoi(str)

		if err != nil {
			return nil, err
		}

		intArr[i] = val

	}

	return intArr, nil
}
