package util

import "unicode/utf8"

func Len(s string) int {
	return utf8.RuneCountInString(s)
}
