package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := []string{
		"a",
		"e",
		"i",
		"o",
		"u",
	}
	for _, letter := range strings.ToLower(str) {
		for _, vowel := range vowels {
			if fmt.Sprintf("%c", letter) == vowel {
				count++
			}
		}
	}
	return count
}
