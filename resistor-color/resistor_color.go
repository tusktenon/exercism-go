package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	_, _ = colorCodeArraySearch, colorCodeMapLookup
    // select an implementation:
	return colorCodeArraySearch(color)
}

// Finds the resistance value of the given color by searching the slice
// returned by the Colors function. 
func colorCodeArraySearch(color string) int {
	for i, c := range Colors() {
		if c == color {
			return i
		}
	}
	return -1
}

// A global variable required for the colorCodeMapLookup function.
var colorMap map[string]int = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Finds the resistance value of the given color with a map lookup. This
// approach slightly outperforms colorCodeArraySearch, but only if the map is
// defined as a global variable.
func colorCodeMapLookup(color string) int {
	return colorMap[color]
}
