package prime

func Factors(n int64) (factors []int64) {
	// We can double the function's speed (verified with benchmarking) by only
	// considering odd numbers as possible factors after first considering 2.
	var d, step int64 = 2, 1
	for n > 1 {
		if n%d == 0 {
			factors = append(factors, d)
			n /= d
		} else {
			d += step
			step = 2
		}
	}
	return
}
