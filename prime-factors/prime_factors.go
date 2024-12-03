package prime

func Factors(n int64) (factors []int64) {
    for d := int64(2); n > 1; {
		if n%d == 0 {
			factors = append(factors, d)
			n /= d
		} else {
            d++
        }
	}
	return
}
