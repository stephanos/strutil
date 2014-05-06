package strutil

import (
	"strings"
	"unicode"
)

// Tail returns a substring with all chars except the first.
func Tail(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[1:]
}

// Init returns a substring with all chars except the last.
func Init(s string) string {
	lenStr := len(s)
	if lenStr == 0 {
		return ""
	}
	return s[:lenStr-1]
}

// Head returns a substring containing the first char.
func Head(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(rune(s[0]))
}

// Last returns a substring containing the last char.
func Last(s string) string {
	strLen := len(s)
	if strLen == 0 {
		return ""
	}
	return string(rune(s[strLen-1]))
}

// Take returns a substring of the first n elements.
// If n is greater than the length of th string, the whole string is returned.
func Take(s string, n int) string {
	lenStr := len(s)
	if n > lenStr {
		n = lenStr
	}
	return s[:n]
}

// TakeRight returns a substring of the last n elements.
// If n is greater than the length of the string, the whole string is returned.
func TakeRight(s string, n int) string {
	lenStr := len(s)
	if n > lenStr {
		n = lenStr
	}
	return s[lenStr-n:]
}

// TakeBefore searches a string from left to right for a pattern and returns a
// substring consisting of the characters in the string that are to the left
// of the pattern or all string if no match found.
func TakeBefore(s, pattern string) string {
	l, _ := takeBeforeAfter(s, pattern)
	return l
}

// TakeRightBefore searches a string from left to right for a pattern and
// returns a substring consisting of the characters in the string that are to
// the right of the pattern or all string if no match found.
func TakeRightBefore(s, pattern string) string {
	l, _ := takeRightBeforeAfter(s, pattern)
	return l
}

// TakeAfter searches a string from right to left for a pattern and returns a
// substring consisting of the characters in the string that are to the left
// of the pattern or all string if no match found.
func TakeAfter(s, pattern string) string {
	_, r := takeBeforeAfter(s, pattern)
	return r
}

// TakeRightAfter searches a string from right to left for a pattern and
// returns a substring consisting of the characters in the string that are to
// the right of the pattern or all string if no match found.
func TakeRightAfter(s, pattern string) string {
	_, r := takeRightBeforeAfter(s, pattern)
	return r
}

func takeBeforeAfter(s, pattern string) (l, r string) {
	idx := strings.Index(s, pattern)
	return splitAround(s, idx, len(pattern))
}

func takeRightBeforeAfter(s, pattern string) (l, r string) {
	idx := strings.LastIndex(s, pattern)
	return splitAround(s, idx, len(pattern))
}

func splitAround(s string, idx, len int) (l, r string) {
	if idx != -1 {
		l = s[:idx]
		r = s[(idx + len):]
	} else {
		l = s
		r = s
	}
	return
}

// Truncate returns a substring of the first n elements and
// appends a suffix ("..." by default) to the result.
//
// If the suffix is longer than the truncated part,
// the original string is returned.
func Truncate(s string, n int, suffix ...string) string {
	if n < 1 {
		return ""
	}

	suf := "..."
	if len(suffix) > 0 {
		suf = suffix[0]
	}
	sufLen := len(suf)

	strLen := len(s)
	if n >= strLen || strLen-n <= sufLen {
		return s
	}
	return strings.TrimFunc(Take(s, n), unicode.IsSpace) + suf
}
