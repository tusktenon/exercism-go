package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("non-positive input %d", n)
	}
	s := 0
	for ; n > 1; s++ {
        // benchmarking shows a 10% performance improvement in using bit operations
        // (n&1 and n >>= 1) over arithmetic ones (n%2 and n /= 2)
		if n&1 == 0 {
			n >>= 1
		} else {
			n = 3*n + 1
		}
	}
	return s, nil
}
