package binarysearch

func SearchInts(list []int, key int) int {
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
