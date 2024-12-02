package sieve

func Sieve(limit int) []int {
	marked := make([]bool, limit+1)
	marked[0], marked[1] = true, true
	for i := 2; i*i <= limit; i++ {
		if !marked[i] {
			for j := i * i; j <= limit; j += i {
				marked[j] = true
			}
		}
	}
	primes := []int{}
	for i, m := range marked {
		if !m {
			primes = append(primes, i)
		}
	}
	return primes
}
