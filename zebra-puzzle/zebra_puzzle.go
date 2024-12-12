package zebra

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	// allow unused implementation
	_, _ = solve, solveBrute

	// select an implementation
	return solve()
}

// In this implementation, we iterate through all possible combinations of
// house properties, ruling out candidates as early as possible to minimize the
// total number of iterations required.
func solve() Solution {
	pFive := permuteFive()
	for _, colors := range pFive() {
		// 6. The green house is immediately to the right of the ivory house.
		if !(colors[Green]-colors[Ivory] == 1) {
			continue
		}
		for _, nations := range pFive() {
			// 10. The Norwegian lives in the first house.
			ok := nations[Norwegian] == 0 &&
				// 2. The Englishman lives in the red house.
				nations[Englishman] == colors[Red] &&
				// 15. The Norwegian lives next to the blue house.
				nextTo(nations[Norwegian], colors[Blue])
			if !ok {
				continue
			}
			for _, drinks := range pFive() {
				// 9. Milk is drunk in the middle house.
				ok := drinks[Milk] == 2 &&
					// 4. Coffee is drunk in the green house.
					drinks[Coffee] == colors[Green] &&
					// 5. The Ukrainian drinks tea.
					nations[Ukrainian] == drinks[Tea]
				if !ok {
					continue
				}
				for _, smokes := range pFive() {
					// 8. Kools are smoked in the yellow house.
					ok := smokes[Kools] == colors[Yellow] &&
						// 13. The Lucky Strike smoker drinks orange juice.
						smokes[LuckyStrike] == drinks[OrangeJuice] &&
						// 14. The Japanese smokes Parliaments.
						nations[Japanese] == smokes[Parliaments]
					if !ok {
						continue
					}
					for _, pets := range pFive() {
						// 3. The Spaniard owns the dog.
						ok := nations[Spaniard] == pets[Dog] &&
							// 7. The Old Gold smoker owns snails.
							smokes[OldGold] == pets[Snails] &&
							// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
							nextTo(smokes[Chesterfields], pets[Fox]) &&
							// 12. Kools are smoked in the house next to the house where the horse is kept.
							nextTo(smokes[Kools], pets[Horse])
						if ok {
							return getSolution(nations, drinks, pets)
						}
					}
				}
			}
		}
	}
	panic("no solution found")
}

// In this implementation, we test all the rules at once, in the innermost
// loop. This is extremely inefficient (we might have to consider all 5!^5 ~
// 24.9 billion possibilities), but a modern PC can still arrive at the correct
// solution in under 2 minutes.
func solveBrute() Solution {
	pFive := permuteFive()
	for _, colors := range pFive() {
		for _, nations := range pFive() {
			for _, drinks := range pFive() {
				for _, smokes := range pFive() {
					for _, pets := range pFive() {
						// 1. There are five houses. (Always true.)
						// 2. The Englishman lives in the red house.
						ok := nations[Englishman] == colors[Red] &&
							// 3. The Spaniard owns the dog.
							nations[Spaniard] == pets[Dog] &&
							// 4. Coffee is drunk in the green house.
							drinks[Coffee] == colors[Green] &&
							// 5. The Ukrainian drinks tea.
							nations[Ukrainian] == drinks[Tea] &&
							// 6. The green house is immediately to the right of the ivory house.
							colors[Green]-colors[Ivory] == 1 &&
							// 7. The Old Gold smoker owns snails.
							smokes[OldGold] == pets[Snails] &&
							// 8. Kools are smoked in the yellow house.
							smokes[Kools] == colors[Yellow] &&
							// 9. Milk is drunk in the middle house.
							drinks[Milk] == 2 &&
							// 10. The Norwegian lives in the first house.
							nations[Norwegian] == 0 &&
							// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
							nextTo(smokes[Chesterfields], pets[Fox]) &&
							// 12. Kools are smoked in the house next to the house where the horse is kept.
							nextTo(smokes[Kools], pets[Horse]) &&
							// 13. The Lucky Strike smoker drinks orange juice.
							smokes[LuckyStrike] == drinks[OrangeJuice] &&
							// 14. The Japanese smokes Parliaments.
							nations[Japanese] == smokes[Parliaments] &&
							// 15. The Norwegian lives next to the blue house.
							nextTo(nations[Norwegian], colors[Blue])
						if ok {
							return getSolution(nations, drinks, pets)
						}
					}
				}
			}
		}
	}
	panic("no solution found")
}

// getSolution extracts a Solution object from slices containing the correct
// house assignments for nationalities, drinks and pets.
func getSolution(nations, drinks, pets []uint8) (s Solution) {
	nationStrings := []string{"Englishman", "Japanese", "Norwegian", "Spaniard", "Ukrainian"}
	for i, house := range nations {
		if house == drinks[Water] {
			s.DrinksWater = nationStrings[i]
		}
		if house == pets[Zebra] {
			s.OwnsZebra = nationStrings[i]
		}
	}
	return
}

// Colors
const (
	Blue int = iota
	Green
	Ivory
	Red
	Yellow
)

// Nationalities
const (
	Englishman int = iota
	Japanese
	Norwegian
	Spaniard
	Ukrainian
)

// Pets
const (
	Dog int = iota
	Fox
	Horse
	Snails
	Zebra
)

// Drinks
const (
	Coffee int = iota
	Milk
	OrangeJuice
	Tea
	Water
)

// Cigarettes
const (
	Chesterfields int = iota
	Kools
	LuckyStrike
	OldGold
	Parliaments
)

// nextTo returns true if x and y differ by one.
func nextTo(x, y uint8) bool {
	// This is slighter faster than
	// return math.Abs(float64(x) - float64(y)) == 1
	return (x == y+1) || (y == x+1)
}

// permuteFive returns a function that returns a slice containing all 120 (5!)
// permutations of (0, 1, 2, 3, 4). For speed, the slice of permutations is
// generated only once (via the permutations function) and then acts as a
// template: each call to the returned function makes and returns a (deep)
// copy.
func permuteFive() func() [][]uint8 {
	template := permutations([]uint8{0, 1, 2, 3, 4})
	return func() [][]uint8 {
		tCopy := make([][]uint8, 0, 120)
		for _, p := range template {
			pCopy := make([]uint8, 5)
			copy(pCopy, p)
			tCopy = append(tCopy, pCopy)
		}
		return tCopy
	}
}

// permutations takes a slice `list` of length n and returns a slice containing
// all n! permutations of `list`. This implementation uses Heap's algorithm; see
// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations[T any](list []T) [][]T {
	result := make([][]T, 0, factorial(len(list)))

	// Recall: anonymous recursive functions must be explicitly declared with var
	var helper func([]T, int)
	helper = func(s []T, n int) {
		if n == 1 {
			cp := make([]T, len(s))
			copy(cp, s)
			result = append(result, cp)
		} else {
			for i := range n {
				helper(s, n-1)
				if n&1 != 0 {
					s[i], s[n-1] = s[n-1], s[i]
				} else {
					s[0], s[n-1] = s[n-1], s[0]
				}
			}
		}
	}

	helper(list, len(list))
	return result
}

// factorial is a simple implementation of the factorial function
// NOTE: A robust implementation would check that n is neither negative nor so
// large that n! will overflow.
func factorial(n int) int {
	fact := 1
	for ; n > 1; n-- {
		fact *= n
	}
	return fact
}
