// Package romannumerals converts Arabic numbers (as ints) to Roman numerals
// (as strings).
//
// This version stores the conversion data in a pair of arrays of types int and
// string. Compared to using a single array of struct type, this paired-arrays
// approach feels slightly less elegant, but also runs slightly faster.
package romannumerals

import (
	"fmt"
	"strings"
)

var (
	arabic = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman  = [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3999 {
		return "", fmt.Errorf("out of range: %d", n)
	}
	var b strings.Builder
	for i, a := range arabic {
		for ; n >= a; n -= a {
			b.WriteString(roman[i])
		}
	}
	return b.String(), nil
}
