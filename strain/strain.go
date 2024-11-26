package strain

func Keep[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range collection {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
	return Keep(collection, func(t T) bool { return !predicate(t) })
}
