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

### Isogram

Four different approaches with interesting performance differences. A good reminder that big-O analysis can be misleading when data sizes are small: a quadratic array-based solution can significantly outperform a linear hash-map-based one. Also, one of these solutions looks like it shouldn't work for non-ASCII strings, but actually does.

### Leap

Even the simplest of problems can have some room for optimization. Also, the README for this exercise explains how Exercism generates test data from a cross-language repository.

### Strain

Generic functions in Go.
