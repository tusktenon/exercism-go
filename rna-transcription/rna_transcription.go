package strand

func ToRNA(dna string) string {
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
			panic("Invalid nucleotide symbol")
		}
	}
	return string(strand)
}
