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
	if len(str) == 0 {
		return "", nil
	} else if unicode.IsDigit(rune(str[0])) {
		return "", ErrInvalidString
	} else if len(str) == 1 {
		return str, nil
	}

	var newStr strings.Builder

	for i := 1; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) && unicode.IsDigit(rune(str[i-1])) {
			return "", ErrInvalidString
		} else if unicode.IsDigit(rune(str[i])) && !unicode.IsDigit(rune(str[i-1])) {
			count, err := strconv.Atoi(string(str[i]))
			if err != nil {
				fmt.Println(err)
			}
			if count == 0 {
				continue
			}
			char := strings.Repeat(string(str[i-1]), count)
			newStr.WriteString(char)
		} else if !unicode.IsDigit(rune(str[i-1])) {
			char := strings.Repeat(string(str[i-1]), 1)
			newStr.WriteString(char)
		}
	}
	if !unicode.IsDigit(rune(str[len(str)-1])) {
		newStr.WriteString(string(str[len(str)-1]))
	}
	return newStr.String(), nil
}
