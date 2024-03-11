package main

import (
	"strconv"
	"strings"
)

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	sum := 0

	for i, value := range isbn {
		num, err := parseInt(value, i)
		if err != nil {
			return false
		}
		sum += num * (i + 1)
	}
	return sum%11 == 0
}

func parseInt(char rune, pos int) (int, error) {
	toStr := string(char)

	if strings.ToUpper(toStr) == "X" && pos == 9 {
		return 10, nil
	}
	toInt, err := strconv.Atoi(toStr)
	return toInt, err
}
