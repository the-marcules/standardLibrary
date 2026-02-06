package util

import "fmt"

func ListUtil(input []string) {

	for _, item := range input {
		print(item)
	}
}

func Worker[T any](data []T) []string {
	var result []string

	for i, item := range data {
		result = append(result, fmt.Sprintf("%d. item %v", i, item))
	}

	return result
}
