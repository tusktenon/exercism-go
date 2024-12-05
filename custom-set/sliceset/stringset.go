// Package stringset implements a Set as a collection of unique string values.
// Internally, Set elements are stored in a sorted string slice.
//
// Keeping the slice in sorted order allows most operations to be more
// efficient. Specifically, if Sets s, s1 and s2 contain n, n1 and n2 elements,
// respectively, we have the following bounds:
//   - New(): O(1)
//   - NewFromSlice(src): O(m log(m)), where m = len(src)
//   - s.String(): O(n)
//   - s.IsEmpty(): O(1)
//   - s.Has: O(log(n))
//   - s.Add(e): O(n)
//   - Equal(s1, s2): O(min(n1, n2))
//   - Disjoint(s1, s2) and Subset(s1, s2): O(n1 + n2)
//   - Difference(s1, s2), Intersection(s1, s2) and Union(s1, s2): O(n1 + n2)
//
// NOTE: Since Go does not have a built-in set type, a completely reasonable
// approach to this exercise is to implement Set as a thin wrapper around
// map[string]struct{}. However, the sorted string slice implementation
//   - is a more interesting challenge;
//   - makes for a better comparison with the same exercise in other language
//     tracks (most languages have a built-in set type);
//   - might outperform the map-based approach in certain situations (e.g, when
//     sets are small and calls to Add are uncommon).
package stringset

import (
	"reflect"
	"slices"
	"sort"
	"strings"
)

type Set struct{ elements []string }

func New() Set { return Set{elements: []string{}} }

func NewFromSlice(l []string) Set {
	// work from a copy to avoid mutating the source slice
	c := make([]string, len(l))
	copy(c, l)
	if len(c) < 2 {
		return Set{elements: c}
	}
	sort.Strings(c)
	prev := c[0]
	dedup := []string{prev}
	for _, curr := range c[1:] {
		if curr != prev {
			dedup = append(dedup, curr)
			prev = curr
		}
	}
	return Set{elements: dedup}
}

func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}
	return `{"` + strings.Join(s.elements, `", "`) + `"}`
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(elem string) bool {
	_, found := slices.BinarySearch(s.elements, elem)
	return found
}

// NOTE: Add must take a pointer receiver to be able to modify s.elements. In a
// realistic Go program, convention dictates that if any method has a pointer
// receiver, then they all should. Since none of the other Set methods modify
// their receiver, a better solution would be to have Add take its receiver by
// value and return the modified Set (in analogy with the built-in append
// function, slices.Insert, etc.).
func (s *Set) Add(elem string) {
	if i, found := slices.BinarySearch(s.elements, elem); !found {
		s.elements = slices.Insert(s.elements, i, elem)
	}
}

func Subset(s1, s2 Set) bool {
	if len(s1.elements) > len(s2.elements) {
		return false
	}
	var i, j int
	for i < len(s1.elements) && j < len(s2.elements) {
		e1, e2 := s1.elements[i], s2.elements[j]
		switch {
		case e1 < e2:
			return false
		case e1 == e2:
			i++
			j++
		case e1 > e2:
			j++
		}
	}
	return i == len(s1.elements)
}

func Disjoint(s1, s2 Set) bool {
	for i, j := 0, 0; i < len(s1.elements) && j < len(s2.elements); {
		e1, e2 := s1.elements[i], s2.elements[j]
		switch {
		case e1 < e2:
			i++
		case e1 == e2:
			return false
		case e1 > e2:
			j++
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1.elements, s2.elements)
}

func Intersection(s1, s2 Set) Set {
	common := []string{}
	for i, j := 0, 0; i < len(s1.elements) && j < len(s2.elements); {
		e1, e2 := s1.elements[i], s2.elements[j]
		switch {
		case e1 < e2:
			i++
		case e1 == e2:
			common = append(common, e1)
			i++
			j++
		case e1 > e2:
			j++
		}
	}
	return Set{elements: common}
}

func Difference(s1, s2 Set) Set {
	diff := []string{}
	var i, j int
	for i < len(s1.elements) && j < len(s2.elements) {
		e1, e2 := s1.elements[i], s2.elements[j]
		switch {
		case e1 < e2:
			diff = append(diff, e1)
			i++
		case e1 == e2:
			i++
			j++
		case e1 > e2:
			j++
		}
	}
	diff = append(diff, s1.elements[i:]...)
	return Set{elements: diff}
}

func Union(s1, s2 Set) Set {
	either := make([]string, 0, max(len(s1.elements), len(s2.elements)))
	var i, j int
	for i < len(s1.elements) && j < len(s2.elements) {
		e1, e2 := s1.elements[i], s2.elements[j]
		switch {
		case e1 < e2:
			either = append(either, e1)
			i++
		case e1 == e2:
			either = append(either, e1)
			i++
			j++
		case e1 > e2:
			either = append(either, e2)
			j++
		}
	}
	either = append(either, s1.elements[i:]...)
	either = append(either, s2.elements[j:]...)
	return Set{elements: either}
}
