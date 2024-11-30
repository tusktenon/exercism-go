package prime

import "fmt"

// Nth returns the nth prime number, or an error if n is less than 1.
func Nth(n int) (int, error) {
    // allow unused implementations
	_, _ = localList, sharedList

    // select an implementation
	return sharedList(n)
}

// localList stores the list of primes in a local variable. While this does
// allow us to pre-allocate sufficient space for the list, it means that every
// call to localList is starting over from scratch.
func localList(n int) (int, error) {
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

var shared = []int{2, 3}

// sharedList uses a shared global variable to store the list of primes, so it
// can be reused for subsequent calls.
func sharedList(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("expected a positive integer, got %v", n)
	}
outer:
	for candidate := shared[len(shared)-1] + 2; len(shared) < n; candidate += 2 {
		for _, p := range shared {
			if p*p > candidate {
				break
			}
			if candidate%p == 0 {
				continue outer
			}
		}
		shared = append(shared, candidate)
	}
	return shared[n-1], nil
}
