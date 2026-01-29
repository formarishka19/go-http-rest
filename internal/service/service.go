package service

import (
	morse "go-http-rest/pkg/morse"
	"strings"
)

func ConvertData(data string) (string, error) {
	f := func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	}
	if strings.ContainsFunc(data, f) {
		return morse.ToMorse(data), nil
	}
	return morse.ToText(data), nil

}
