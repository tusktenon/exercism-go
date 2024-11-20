package chessboard

// File records if a square is occupied by a piece.
type File []bool

// Chessboard contains a map of eight Files, accessed with keys from "A" to "H".
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (count int) {
	for _, b := range cb[file] {
		if b {
			count++
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (count int) {
	if 1 <= rank && rank <= 8 {
		for _, s := range cb {
			if s[rank-1] {
				count++
			}
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (count int) {
	for _, s := range cb {
		for range s {
			count++
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (count int) {
    for f := range cb {
        count += CountInFile(cb, f)
    }
    return
}
