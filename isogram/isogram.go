// Package isogram contains various implementations of the `IsIsogram` function,
// which returns true if the given string is an isogram (contains no repeating
// letter character).
//
// Benchmarking results are from an Apple M1 MacBook Air (darwin/arm64).
package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	// allow unused implementations
	_ = isIsogramBitSet
	_ = isIsogramMap
	_ = isIsogramCount
	_ = isIsogramTriangle

	// select an implementation
	return isIsogramTriangle(word)
}

// Use a bit set to track which letters have already been encountered.
// This is by far the fastest approach, but only works for ASCII input.
// Benchmark: 105.7 ns/op    0 B/op    0 allocs/op
func isIsogramBitSet(word string) bool {
	var seen, flag int32
	for _, b := range []byte(word) {
		flag = 0
		if 'a' <= b && b <= 'z' {
			flag = 1 << (b - 'a')
		} else if 'A' <= b && b <= 'Z' {
			flag = 1 << (b - 'A')
		}
		if seen&flag != 0 {
			return false
		}
		seen |= flag
	}
	return true
}

// Use a map with ignored values as a set to track the letters.
// This approach supports non-ASCII characters and is the most obvious and
// expressive (at least to my eye), but also the slowest.
// Benchmark: 3638 ns/op    1231 B/op    16 allocs/op
func isIsogramMap(word string) bool {
	seen := map[rune]struct{}{}
	for _, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) {
			if _, ok := seen[r]; ok {
				return false
			}
			seen[r] = struct{}{}
		}
	}
	return true
}

// Use the strings.Count function.
// Despite the obvious inefficiency of reexamining the entire string for every
// letter, this approach is significantly faster than the one using a map
// (while also supporting Unicode).
// Benchmark: 1377 ns/op    40 B/op    3 allocs/op
func isIsogramCount(word string) bool {
	word = strings.ToLower(word)
	for _, r := range word {
		if unicode.IsLetter(r) && strings.Count(word, string(r)) > 1 {
			return false
		}
	}
	return true
}

// We can further improve the performance of the strings.Count approach by only
// examining the portion of the string after each letter. This is the favourite
// community solution.
// Perhaps surprisingly, this implementation is Unicode-safe! In the for loop,
// i is iterating through the bytes of word but r is iterating through the
// runes, so i+1 may not be the start of the next character after r, but rather
// part of r itself. But thanks to the magic of UTF-8 encoding, Go can still
// differentiate between any invalid bytes and the start of the next (valid)
// rune.
// Benchmark: 842.5 ns/op    40 B/op    3 allocs/op
func isIsogramTriangle(word string) bool {
	word = strings.ToLower(word)
	for i, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) && strings.ContainsRune(word[i+1:], r) {
			return false
		}
	}
	return true
}
