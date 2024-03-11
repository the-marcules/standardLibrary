package main

import (
	"fmt"
	"strings"
)

func EncryptThis(text string) string {
	words := strings.Split(text, " ")
	for i, word := range words {
		tmp := ""
		for li, letter := range word {
			switch li {
			case 0:
				tmp += fmt.Sprintf("%v", rune(word[li]))
			case 1:
				tmp += string(word[len(word)-1])
			case len(word) - 1:
				tmp += string(word[1])
			default:
				tmp += string(letter)
			}
		}
		words[i] = tmp
	}
	return strings.Join(words, " ")
}
