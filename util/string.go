// The inCommon/util package provides very few utility helpers
package util

import "unicode"

// UpperCaseFirst converts the first character of a string to upper case
func UpperCaseFirst(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
