package flatten

func Flatten(nested interface{}) []interface{} {
	// allow unused implementation
	_, _ = flatten1, flatten2

	// select an implementation
	return flatten2(nested)
}

// Option 1: A "pure" functional approach (no mutation).
// Runs slower, uses more memory and performs more allocations.
func flatten1(nested interface{}) []interface{} {
	switch nested := nested.(type) {
	case nil:
		return []interface{}{}
	case []interface{}:
		if len(nested) == 0 {
			return []interface{}{}
		}
		return append(Flatten(nested[0]), Flatten(nested[1:])...)
	default:
		return []interface{}{nested}
	}
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
