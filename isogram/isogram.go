package isogram

import "unicode"

func IsIsogram(word string) bool {
	// allow unused implementations
	_, _ = isIsogramBitFlag, isIsogramSet

	// select an implementation
	return isIsogramBitFlag(word)
}

func isIsogramBitFlag(word string) bool {
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

func isIsogramSet(word string) bool {
	seen := map[rune]struct{}{}
	for _, r := range word {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			if _, ok := seen[r]; ok {
				return false
			}
			seen[r] = struct{}{}
		}
	}
	return true
}
