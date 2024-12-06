package romannumerals

import (
	"fmt"
	"strings"
)

var (
	arabic = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman  = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3999 {
		return "", fmt.Errorf("out of range: %d", n)
	}
	var b strings.Builder
	for n > 0 {
		for i, a := range arabic {
			if n >= a {
				b.WriteString(roman[i])
				n -= a
				break
			}
		}
	}
	return b.String(), nil
}
