package strutil

import "strconv"

// ToInt returns the corresponding int (base 10) for a passed-in string.
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// FromInt returns the string representation of i (with base 10).
func FromInt(i int) string {
	return strconv.Itoa(i)
}
