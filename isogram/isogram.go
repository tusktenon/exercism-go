package isogram

func IsIsogram(word string) bool {
	var seen, flag int32
	for _, b := range []byte(word) {
		flag = 0
		if 'a' <= b && b <= 'z' {
			flag = 1 << (b - 'a')
		} else if 'A' <= b && b <= 'Z' {
			flag = 1 << (b - 'A')
		}
		if flag != 0 {
			if seen&flag != 0 {
				return false
			} else {
				seen |= flag
			}
		}
	}
	return true
}
