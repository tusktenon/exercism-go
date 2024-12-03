package sieve

func Sieve(limit int) []int {
	primes := []int{}
	marked := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if !marked[i] {
            primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				marked[j] = true
			}
		}
	}
	return primes
}
