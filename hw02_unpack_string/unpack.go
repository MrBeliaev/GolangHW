package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	switch len(str) {
	case 0:
		return "", nil
	case 1:
		if !unicode.IsDigit(rune(str[0])) {
			return str, nil
		}
		return "", ErrInvalidString
	default:
		if unicode.IsDigit(rune(str[0])) {
			return "", ErrInvalidString
		}
	}

	var newStr strings.Builder

	for i := 1; i < len(str); i++ {
		switch unicode.IsDigit(rune(str[i])) {
		case true:
			switch unicode.IsDigit(rune(str[i-1])) {
			case true:
				return "", ErrInvalidString
			case false:
				count, err := strconv.Atoi(string(str[i]))
				if err != nil {
					fmt.Println(err)
				}
				if count == 0 {
					continue
				}
				char := strings.Repeat(string(str[i-1]), count)
				newStr.WriteString(char)
			}
		case false:
			if !unicode.IsDigit(rune(str[i-1])) {
				char := strings.Repeat(string(str[i-1]), 1)
				newStr.WriteString(char)
			}
		}
	}
	if !unicode.IsDigit(rune(str[len(str)-1])) {
		newStr.WriteString(string(str[len(str)-1]))
	}
	return newStr.String(), nil
}
