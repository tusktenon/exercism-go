package strand

import "strings"

func ToRNA(dna string) string {
	_, _, _, _ = toRNA1, toRNA2, toRNA3, toRNA4
	return toRNA1(dna)
}

// Option 1: Use a switch statement.
func toRNA1(dna string) string {
	strand := []byte(dna)
	for i, n := range strand {
		switch n {
		case 'A':
			strand[i] = 'U'
		case 'C':
			strand[i] = 'G'
		case 'G':
			strand[i] = 'C'
		case 'T':
			strand[i] = 'A'
		default:
			panic("invalid nucleotide symbol")
		}
	}
	return string(strand)
}

// Option 2: Use string search and indexing.
// A little more compact than Option 1, but also slower.
func toRNA2(dna string) string {
	strand := []byte(dna)
	for i, n := range strand {
		strand[i] = "UGCA"[strings.IndexByte("ACGT", n)]
	}
	return string(strand)
}

// Option 3: Use the strings.Map function.
// Elegant, but also slower than the previous options.
func toRNA3(dna string) string {
	translate := func(r rune) rune {
		return rune("UGCA"[strings.IndexRune("ACGT", r)])
	}
	return strings.Map(translate, dna)
}

// Option 4: Use a map.
// As expected, this is by far the slowest option.
func toRNA4(dna string) string {
	translate := map[byte]byte{
		'A': 'U',
		'C': 'G',
		'G': 'C',
		'T': 'A',
	}
	strand := []byte(dna)
	for i, n := range strand {
		strand[i] = translate[n]
	}
	return string(strand)
}
