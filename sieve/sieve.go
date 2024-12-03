package sieve

func Sieve(limit int) []int {
	// allow unused implementation
	_, _ = sieveBitArray, sieveBoolSlice

	// select an implementation
	return sieveBoolSlice(limit)
}

// Approach 1: Use a slice of bools to track the marked/unmarked numbers.
// This is expressive and convenient, but not very space efficient:
// conceptually, we only need a single bit to store a boolean value, but each
// bool actually occupies a full byte, so we are using 8 times more space than
// necessary.
func sieveBoolSlice(limit int) []int {
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

// Approach 2: Save memory by using a custom bit-array type to track the
// marked/unmarked numbers.
func sieveBitArray(limit int) []int {
	primes := []int{}
	s := newSieve(limit)
	for i := 2; i <= limit; i++ {
		if !s.isMarked(i) {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				s.mark(j)
			}
		}
	}
	return primes
}

// The bit array is a slice of uints.
type sieve []uint

// A uint is either 32 or 64 bits, depending on the platform.
const wordSize = 32 << (^uint(0) >> 63)

// newSieve returns a slice with at least n + 1 bits.
func newSieve(n int) sieve {
	return make([]uint, (n/wordSize)+1)
}

func (s sieve) mark(i int) {
	q, r := i/wordSize, i%wordSize
	s[q] |= (1 << r)
}

func (s sieve) isMarked(i int) bool {
	q, r := i/wordSize, i%wordSize
	return s[q]&(1<<r) != 0
}
