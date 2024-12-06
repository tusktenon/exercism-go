// Package romannumerals converts Arabic numbers (as ints) to Roman numerals
// (as strings).
//
// This version uses a single conversion array whose elements are instances of
// an anonymous struct type. This feels safer and more expressive than the
// paired-arrays approach, but runs slightly slower.
package romannumerals

import (
	"fmt"
	"strings"
)

var conversions = [...]struct {
	arabic int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3999 {
		return "", fmt.Errorf("out of range: %d", n)
	}
	var b strings.Builder
	for _, c := range conversions {
		for ; n >= c.arabic; n -= c.arabic {
			b.WriteString(c.roman)
		}
	}
	return b.String(), nil
}
