// Package leap provides the IsLeapYear function, which reports whether the
// given year is a leap year.
package leap

// IsLeapYear reports whether the year is a leap year.
func IsLeapYear(year int) bool {
	// allow unused implementations
	_ = isLeapYearBooleanChain
	_ = isLeapYearMaxTwo

	// select an implementation
	return isLeapYearMaxTwo(year)
}

// This approach is concise and expressive (it's essentially a one-to-one
// translation of the provided definition of a leap year into Go code), but it
// might perform up to three checks.
func isLeapYearBooleanChain(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// This approach is only slightly more verbose, and never performs more than
// two checks.
func isLeapYearMaxTwo(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}
