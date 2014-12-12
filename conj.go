package strutil

import "strings"

// And creates a conjunction with "and" from the passed-in words.
func And(s ...string) string {
	return conj(s, " and ")
}

// Or creates a conjunction with "or" from the passed-in words.
func Or(s ...string) string {
	return conj(s, " or ")
}

func conj(s []string, conjWith string) string {
	items := len(s)
	if items == 0 {
		return ""
	} else if items == 1 {
		return s[0]
	}

	return strings.Join(s[:items-1], ", ") + conjWith + s[items-1]
}
