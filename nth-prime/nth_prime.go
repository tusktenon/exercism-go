package prime

import "fmt"

// Nth returns the nth prime number, or an error if n is less than 1
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("expected a positive integer, got %v", n)
	}
	primes := make([]int, 0, n)
	primes = append(primes, 2)
outer:
	for candidate := 3; len(primes) < n; candidate += 2 {
		for _, p := range primes {
			if p*p > candidate {
				break
			}
			if candidate%p == 0 {
				continue outer
			}
		}
		primes = append(primes, candidate)
	}
	return primes[n-1], nil
}
