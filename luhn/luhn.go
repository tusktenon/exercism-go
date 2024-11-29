package luhn

func Valid(id string) bool {
	var luhn, count int
	for i := len(id) - 1; i >= 0; i-- {
		switch c := id[i]; {
		case '0' <= c && c <= '9':
			count++
			d := int(c - '0')
			if count&1 == 0 {
				d <<= 1
				if d > 9 {
					d -= 9
				}
			}
			luhn += d
		case c == ' ':
			continue
		default:
			return false
		}
	}
	return count > 1 && luhn%10 == 0
}
