package flatten

func Flatten(nested interface{}) []interface{} {
	// allow unused implementation
	_, _ = flatten1, flatten2

	// select an implementation
	return flatten1(nested)
}

// Option 1: A "pure" functional approach (no mutation).
// Runs slower and performs more memory allocations.
func flatten1(nested interface{}) []interface{} {
	if nested == nil {
		return []interface{}{}
	}
	if s, ok := nested.([]interface{}); ok {
		if len(s) == 0 {
			return []interface{}{}
		}
		return append(Flatten(s[0]), Flatten(s[1:])...)
	}
	return []interface{}{nested}
}

// Option 2: Combine recursion with iteration.
func flatten2(nested interface{}) []interface{} {
    flattened := []interface{}{}
	switch nested := nested.(type) {
	case nil: // omit nil entries
	case []interface{}:
		for _, e := range nested {
			flattened = append(flattened, flatten2(e)...)
		}
	default:
		flattened = append(flattened, nested)
	}
	return flattened
}
