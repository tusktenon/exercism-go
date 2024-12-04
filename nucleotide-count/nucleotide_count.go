package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
    // allow unused implementation
    _, _ = DNA.countsSwitch, DNA.countsKeyTest

	// select an implementation
	return d.countsKeyTest()
}

// Approach 1: validate the nucleotides with a switch statement.
func (d DNA) countsSwitch() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range d {
		switch r {
		case 'A', 'C', 'G', 'T':
			h[r]++
		default:
			return nil, fmt.Errorf("invalid nucleotide: %c", r)
		}
	}
	return h, nil
}

// Approach 2: test that each nucleotide is already a key in the histogram.
// More idiomatic but slightly slower than the switch-statement approach.
func (d DNA) countsKeyTest() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range d {
		if _, ok := h[r]; ok {
			h[r]++
		} else {
			return nil, fmt.Errorf("invalid nucleotide: %c", r)
		}
	}
	return h, nil
}
