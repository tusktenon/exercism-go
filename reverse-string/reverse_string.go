package reverse

func Reverse(input string) string {
	// allow unused implementation
	_, _ = reverse1, reverse2

	// select an implementation
	return reverse2(input)
}

// Using a single loop variable, it may not be immediately apparent that the
// loop condition (i < n/2) is correct.
func reverse1(input string) string {
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}

// Using two loop variables makes both the symmetry of the algorithm and the
// correctness of the loop condition obvious. Interestingly, benchmarking shows
// that this version also runs ever so slightly faster.
func reverse2(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
