# Go on Exercism: Solutions to Exercises

These are my solutions to the exercises of the [Go track](https://exercism.org/tracks/go) on [Exercism](https://exercism.org).

## Learning Exercises

These are completed during the Go track's "Learning Mode" to illustrate important language [concepts](https://exercism.org/tracks/go/concepts). They're invariably short and simple, and there generally aren't too many reasonable ways to solve them, but they can provide nice examples of Go features and built-in functions.

### Animal Magic

The `math/rand` package provides a convenient `Shuffle` function.

### Blackjack

A very simple exercise, but also a nice showcase of Go's `switch` statement.

### Chessboard

Every function in my solution uses a bare return.

### Election Day

In Go, it's perfectly correct for a function to return a pointer to a local variable or a literal value (coming from C/C++, this is shocking).

### Gross Store

You can include a simple statement (like a variable declaration) with a tagless `switch`; just remember to add a semicolon after the statement. See the `RemoveItem` function for an example.

### Lasagna Master

The `Quantities` function provides an opportunity to use a *bare return*.

### Parsing Log Files

Regular expressions in Go.

### The Farm

Custom error types in Go.


## Practice Exercises

These are meant for students who've completed Learning Mode or otherwise acquired basic Go proficiency, and vary considerably in length and difficulty.

### Collatz Conjecture

When the divisor is 2, there is a small but noticeable performance advantage in using bit operations instead of arithmetic ones (e.g., `n&1 == 0` instead of `n%2 == 0`, `n >>= 1` instead of `n /= 2`), at least for Go 1.23 on darwin/arm64.

### Custom Set

This exercise appears in many language tracks. Since most languages already provide a built-in set type, people typically implement their custom set with a resizable array (`vector` in C++, `ArrayList` in Java, `Vec` in Rust, etc.); since Go lacks a built-in set type, it's completely reasonable to implement the custom set as a simple wrapper around `map[string]struct{}`. I did it both ways.

I also have to mention a great solution to the common task of pretty-printing a custom collection type. I had originally implemented the `String` method with the classic reset-the-prefix approach:
```go
func (s Set) Strin1() string {
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
```
But then I saw this [lovely implementation](https://exercism.org/tracks/go/exercises/custom-set/solutions/martinohmann) among the community solutions, which reminded me of the `strings.Join` function:
```go
func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}
	return `{"` + strings.Join(s.elements, `", "`) + `"}`
}
```

### Isogram

Four different approaches with interesting performance differences. A good reminder that big-O analysis can be misleading when data sizes are small: a quadratic array-based solution can significantly outperform a linear hash-map-based one. Also, one of these solutions looks like it shouldn't work for non-ASCII strings, but actually does.

### Leap

Even the simplest of problems can have some room for optimization. Also, the README for this exercise explains how Exercism generates test data from a cross-language repository.

### List Ops

The classic functional, higher-order list functions (`map`, `filter`, `foldl`, `foldr`, `concat`, etc.), implemented with imperative techniques (for efficiency).

### Reverse String
Using two loop variables instead of one makes the solution more expressive (and ever so slightly faster).

### Sieve

As in the Sieve of Eratosthenes. The canonical structure for keeping track of the marked/unmarked numbers is a slice of bools, but just for fun, I also wrote a solution using a custom bit-array type to save memory.

### Strain

Generic functions in Go.
