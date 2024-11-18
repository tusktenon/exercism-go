package lasagna

// PreparationTime returns an estimate for the total preparation time, based on
// the number of layers and the average preparation time per layer.
func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}

// Quantities returns the amounts of noodles and sauce needed.
func Quantities(layers []string) (noodles int, sauce float64) {
	const noodlesPerLayer = 50
	const saucePerLayer = 0.2
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += noodlesPerLayer
		case "sauce":
			sauce += saucePerLayer
		}
	}
	return
}

// AddSecretIngredient replaces the final ingredient in the destination list
// with the final ingredient in the source list.
func AddSecretIngredient(src, dst []string) {
	dst[len(dst)-1] = src[len(src)-1]
}

// ScaleRecipe scales the amounts in a recipe for the desired number of
// portions, assuming the original recipe produces 2 portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	factor := float64(portions) / 2
	scaled := make([]float64, 0, len(amounts))
	for _, a := range amounts {
		scaled = append(scaled, factor*a)
	}
	return scaled
}
