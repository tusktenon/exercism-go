package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
// NOTE: We could also just use the strings.ReplaceAll function.
func Replace(log string, oldRune, newRune rune) string {
	rs := make([]rune, 0, utf8.RuneCountInString(log))
	for _, r := range log {
		if r == oldRune {
			rs = append(rs, newRune)
		} else {
			rs = append(rs, r)
		}
	}
	return string(rs)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
