package flatten

func Flatten(nested interface{}) []interface{} {
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
