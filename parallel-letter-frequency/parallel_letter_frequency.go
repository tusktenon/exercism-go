// Benchmark results from an Apple M1 MacBook Air (arm64/darwin):
//
//	Sequential:                207,562 ns/op    17,533 B/op    12 allocs/op
//	Concurrent (simple merge): 113,630 ns/op    12,392 B/op    63 allocs/op
//	Concurrent (clever merge): 112,388 ns/op    10,357 B/op    51 allocs/op
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	// allow unused implementation
	_ = concurrentFrequencySimpleMerge
	_ = concurrentFrequencyCleverMerge

	// select an implementation
	return concurrentFrequencyCleverMerge(texts)
}

// Option 1: The merged map is initialized as an empty map.
func concurrentFrequencySimpleMerge(texts []string) FreqMap {
	maps := make(chan FreqMap)
	for _, t := range texts {
		go func() { maps <- Frequency(t) }()
	}
	merged := FreqMap{}
	for range texts {
		for r, f := range <-maps {
			merged[r] += f
		}
	}
	return merged
}

// Option 2: The merged map starts as the first map to be returned from one of
// the concurrent goroutines. In principle, this should offer better
// performance: one less merge is required, and the merged map starts at more
// or less the correct size. Benchmarking shows that, while this approach does
// indeed perform fewer allocations, the actual speed-up is quite small.
func concurrentFrequencyCleverMerge(texts []string) FreqMap {
	maps := make(chan FreqMap)
	for _, t := range texts {
		go func() { maps <- Frequency(t) }()
	}
	merged := <-maps
	for range len(texts) - 1 {
		for r, f := range <-maps {
			merged[r] += f
		}
	}
	return merged
}
