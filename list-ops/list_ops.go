package listops

type IntList []int

func (s IntList) Foldl(f func(int, int) int, initial int) int {
	acc := initial
	for _, x := range s {
		acc = f(acc, x)
	}
	return acc
}

func (s IntList) Foldr(f func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = f(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(f func(int) bool) IntList {
	filtered := IntList{}
	for _, x := range s {
		if f(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(f func(int) int) IntList {
	mapped := make(IntList, len(s))
	for i, x := range s {
		mapped[i] = f(x)
	}
	return mapped
}

func (s IntList) Reverse() IntList {
	r := make(IntList, len(s))
	for i := range r {
		r[i] = s[len(s)-i-1]
	}
	return r
}

func (s IntList) Append(list IntList) IntList {
	var a IntList
	aLen := len(s) + len(list)
	if aLen <= cap(s) {
		// s has sufficient capacity to append list; just extend the slice.
		a = s[:aLen]
	} else {
		// s has insufficient space. Allocate a new backing array and copy the
		// contents of s.
		a = make(IntList, aLen)
		copy(a, s)
	}
    // copy the contents of list onto the end of a
	copy(a[len(s):], list)
	return a
}

func (s IntList) Concat(lists []IntList) IntList {
	var c IntList
	// calculate the length of the concatenated list
	cLen := len(s)
	for _, list := range lists {
		cLen += len(list)
	}
	if cLen <= cap(s) {
		// s has sufficient capacity to append all the elements in lists;
		// simply extend the slice.
		c = s[:cLen]
	} else {
		// s has insufficient space. Allocate a new backing array and copy the
		// contents of s.
		c = make(IntList, cLen)
		copy(c, s)
	}
    // copy the contents of each list in lists onto the end of c
	start := len(s)
	for _, list := range lists {
		copy(c[start:], list)
		start += len(list)
	}
	return c
}
