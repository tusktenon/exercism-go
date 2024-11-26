package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	// allow unused implementations
	_, _, _, _ = isIsogramBitSet, isIsogramMap, isIsogramCount, isIsogramTriangle

	// select an implementation
	return isIsogramMap(word)
}

// Use a bit set to track which letters have already been encountered.
// This is by far the fastest approach, but only works for ASCII input.
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
// letter, this approach is significantly faster than the one using a map (while
// also supporting Unicode).
func isIsogramCount(word string) bool {
	word = strings.ToLower(word)
	for _, r := range word {
		if unicode.IsLetter(r) && strings.Count(word, string(r)) > 1 {
			return false
		}
	}
	return true
}

// Given the surprisingly strong performance of the strings.Count approach, we
// might hope to do even better by only examining the portion of the string
// after each letter. Sure enough, this approach offers a noticeable
// improvement.
func isIsogramTriangle(word string) bool {
	letters := []rune(strings.ToLower(word))
	for i, left := range letters {
		if !unicode.IsLetter(left) {
			continue
		}
		for _, right := range letters[i+1:] {
			if right == left {
				return false
			}
		}
	}
	return true
}
