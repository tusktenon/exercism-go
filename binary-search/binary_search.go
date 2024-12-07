package binarysearch

func SearchInts(list []int, key int) int {
	// allow unused implementation
	_, _ = searchClosed, searchHalfOpen

	// select an implementation
	return searchHalfOpen(list, key)
}

// In this approach, we search for the key in the closed interval between left
// and right (including both endpoints).
//
// NOTE: While mid could also be defined as (left + right)/2, using
// left + (right - left)/2 guards against overflow.
func searchClosed(list []int, key int) int {
	for left, right := 0, len(list)-1; left <= right; {
		switch mid := left + (right-left)/2; {
		case list[mid] < key:
			left = mid + 1
		case list[mid] == key:
			return mid
		case list[mid] > key:
			right = mid - 1
		}
	}
	return -1
}

// In this approach, right is a "one past the end" index, and we search for the
// key in the half-open interval from left (inclusive) to right (exclusive).
func searchHalfOpen(list []int, key int) int {
	for left, right := 0, len(list); left != right; {
		switch mid := left + (right-left)/2; {
		case list[mid] < key:
			left = mid + 1
		case list[mid] == key:
			return mid
		case list[mid] > key:
			right = mid
		}
	}
	return -1
}
