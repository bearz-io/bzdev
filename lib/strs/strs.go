package strs

import (
	"strings"
	"unicode"
)

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsEmptySpace(s string) bool {
	if len(s) == 0 {
		return true
	}

	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func PadLeft(s string, length int, pad string) string {
	diff := length - len(s)
	if diff <= 0 {
		return s
	}

	return strings.Repeat(pad, diff) + s
}

func PadRight(s string, length int, pad string) string {
	diff := length - len(s)
	if diff <= 0 {
		return s
	}

	return s + strings.Repeat(pad, diff)
}
