// Package stringset implements a Set as a collection of unique string values.
// Internally, Set elements are stored as the keys of a map[string]struct{}.
package stringset

import (
	"sort"
	"strings"
)

type Set struct{ elements map[string]struct{} }

func New() Set { return Set{elements: map[string]struct{}{}} }

func NewWithCapacity(c int) Set {
	return Set{elements: make(map[string]struct{}, c)}
}

func NewFromSlice(list []string) Set {
	s := New()
	for _, e := range list {
		s.Add(e)
	}
	return s
}

func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}
	elementList := make([]string, 0, len(s.elements))
	for e := range s.elements {
		elementList = append(elementList, e)
	}
	sort.Strings(elementList)
	return `{"` + strings.Join(elementList, `", "`) + `"}`
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(e string) bool {
	_, found := s.elements[e]
	return found
}

func (s *Set) Add(e string) {
	s.elements[e] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	return len(s1.elements) <= len(s2.elements) && subset(s1, s2)
}

func subset(s1, s2 Set) bool {
	for e := range s1.elements {
		if !s2.Has(e) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	// it's faster to iterate through the smaller set
	if len(s2.elements) < len(s1.elements) {
		s1, s2 = s2, s1
	}
	for e := range s1.elements {
		if s2.Has(e) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return len(s1.elements) == len(s2.elements) && subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	intersection := New()
	// it's faster to iterate through the smaller set
	if len(s2.elements) < len(s1.elements) {
		s1, s2 = s2, s1
	}
	for e := range s1.elements {
		if s2.Has(e) {
			intersection.Add(e)
		}
	}
	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := New()
	for e := range s1.elements {
		if !s2.Has(e) {
			difference.Add(e)
		}
	}
	return difference
}

func Union(s1, s2 Set) Set {
	union := NewWithCapacity(max(len(s1.elements), len(s2.elements)))
	for e := range s1.elements {
		union.Add(e)
	}
	for e := range s2.elements {
		union.Add(e)
	}
	return union
}
