// Package word do things with letters.
package word

import "strings"

// UseCount counts the words of a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the words of a string.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
