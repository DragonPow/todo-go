package helper

import "strconv"

func IntToString(input []int32) (result []string) {
	for _, v := range input {
		result = append(result, strconv.Itoa(int(v)))
	}
	return result
}
