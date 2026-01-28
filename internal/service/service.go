package service

import (
	morse "go-http-rest/pkg/morse"
	"slices"
)

var morseSymbols = []rune{'.', '-', ' '}

func ConvertData(data string) (string, error) {
	input := []rune(data)
	for _, sybmol := range input {
		if !slices.Contains(morseSymbols, sybmol) {
			return morse.ToMorse(data), nil
		}
	}
	return morse.ToText(data), nil

}
