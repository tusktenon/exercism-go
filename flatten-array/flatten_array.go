package flatten

func Flatten(nested any) []any {
	// allow unused implementation
	_, _ = flatten1, flatten2

	// select an implementation
	return flatten2(nested)
}

// Option 1: A "pure" functional approach (no mutation).
// Runs slower, uses more memory and performs more allocations.
func flatten1(nested any) []any {
	switch nested := nested.(type) {
	case nil:
		return []any{}
	case []any:
		if len(nested) == 0 {
			return []any{}
		}
		return append(Flatten(nested[0]), Flatten(nested[1:])...)
	default:
		return []any{nested}
	}
}

// Option 2: Combine recursion and iteration.
func flatten2(nested any) []any {
	switch nested := nested.(type) {
	case nil:
		return []any{}
	case []any:
		flattened := []any{}
		for _, e := range nested {
			flattened = append(flattened, flatten2(e)...)
		}
		return flattened
	default:
		return []any{nested}
	}
}
