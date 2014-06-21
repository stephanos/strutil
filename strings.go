package strutil

import (
	"bytes"
	"go/doc"
	"regexp"
	"strings"
	"unicode"
)

var (
	whitespaceRegex = regexp.MustCompile("\\s+")
)

// DeleteWhitespace returns a version of the passed-in string with all
// whitespaces (as by unicode.IsSpace) removed.
func DeleteWhitespace(s string) string {
	if len(s) == 0 {
		return s
	}

	var hasSpace bool
	var buf bytes.Buffer
	for _, r := range s {
		if unicode.IsSpace(r) {
			hasSpace = true
		} else {
			buf.WriteRune(r)
		}
	}

	if !hasSpace {
		return s
	}
	return buf.String()
}

// ShrinkWhitespace returns a version of the passed-in string with whitespace
// removed from the end and beginning as well as all whitespace in-between
// shrunken to a single space.
func ShrinkWhitespace(s string) string {
	return whitespaceRegex.ReplaceAllString(strings.TrimSpace(s), " ")
}

// IndexAfter returns the index directly after the first instance of sep in s,
// or -1 if sep is not present in s.
func IndexAfter(s, sep string) int {
	pos := strings.Index(s, sep)
	if pos == -1 {
		return -1
	}
	return pos + len(sep)
}

// Reverse returns a reversed version of s.
func Reverse(s string) (rev string) {
	for _, r := range s {
		rev = string(r) + rev
	}
	return rev
}

// Lines returns sequence of lines of passed-in str.
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// SplitCamelCase splits the string s at each run of upper case runes and
// returns and array of slices of s.
//func SplitCamelCase(s string) (words []string) {
//	fieldStart := 0
//	for i, r := range s {
//		if i != 0 && unicode.IsUpper(r) {
//			words = append(words, s[fieldStart:i])
//			fieldStart = i
//		}
//	}
//	if fieldStart != len(s) {
//		words = append(words, s[fieldStart:])
//	}
//	return words
//}

// IsBlank returns whether a string is whitespace or empty.
func IsBlank(s string) bool {
	strLen := len(s)
	if strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(s[i])) == false {
			return false
		}
	}
	return true
}

// IsEmpty returns whether a string is empty ("").
func IsEmpty(s string) bool {
	return len(s) == 0
}

// WordWrap wraps paragraphs of text to width.
func WordWrap(s string, width int) string {
	buf := new(bytes.Buffer)
	doc.ToText(buf, s, "", "", width)
	return buf.String()
}

// Surround appends and prepends a string to another.
func Surround(s, around string) string {
	return around + s + around
}
