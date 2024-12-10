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
	maps := make(chan FreqMap)
	for _, t := range texts {
		go func() { maps <- Frequency(t) }()
	}
	merged := <-maps
	for range len(texts) - 1 {
		curr := <-maps
		for r, f := range curr {
			merged[r] += f
		}
	}
	return merged
}
