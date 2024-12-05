// Package stringset implements a Set as a collection of unique string values.
package stringset

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strings"
)

// Implement the Set type using a sorted string slice.
type Set struct{ elements []string }

func New() Set {
	return Set{elements: []string{}}
}

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
		}
		prev = curr
	}
	return Set{elements: dedup}
}

// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be
// formatted as {"a", "b"}. Format the empty set as {}.
func (s Set) String() string {
	var b strings.Builder
	b.WriteByte('{')
	prefix := ""
	for _, e := range s.elements {
		fmt.Fprintf(&b, "%s\"%s\"", prefix, e)
		prefix = ", "
	}
	b.WriteByte('}')
	return b.String()
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
