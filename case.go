package strutil

import (
	"unicode"
	"unicode/utf8"
)

// LowerFirst returns str with the first char lower-cased.
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// UpperFirst returns str with the first char upper-cased.
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
