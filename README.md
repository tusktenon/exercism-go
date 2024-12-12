# Go on Exercism: Solutions to Exercises

These are my solutions to the exercises of the [Go track](https://exercism.org/tracks/go) on [Exercism](https://exercism.org).


## Completing Exercises

Remember to do an initial commit after downloading the exercise. Once you think you have a working solution, verify with `go test`. Most of the exercises in the Go track include some benchmark tests; you can run these with `go test -bench .` (if you're just interested in execution speed) or `go test -bench . --benchmem` (to add memory use and allocation stats).

The `go.mod` file that comes with the exercises is set to Go 1.18, but I often like to use features from Go 1.21 and 1.22 (in particular, the built-in `min`, `max` and `clear` functions; the [fix for the "loop variable capture" issue](https://go.dev/blog/loopvar-preview); and `for` loops that can `range` over an `int`). Exercism supports Go 1.22, but you must submit your modified `go.mod` along with your main `.go` file.

### Handling Multiple Solutions

I often like to write several solutions to an exercise, either to explore specific language features or to compare performance characteristics. Go is *very* strict about unused unexported ("private") variables, including functions: if it starts with a lowercase letter and you don't use it, it's an error. In a real-world program, this is great, but it makes alternate implementations a little trickier.

If the exercise only has one or two exported functions, I just write all the alternates as unexported functions in the same file. The exported function is then just used as a wrapper to select one of the options, with blank assignments to suppress the "unused" errors:

```go
func TheFunction(x int, y string) bool {
	// allow unused implementations
	_, _ = theFunctionA, theFunctionB
	_ = theFunctionC
	_ = theFunctionD

	// select an implementation
	return theFunctionA(x, y)
}
```

If the exercise has more than a few exported types and functions (e.g., you're writing a custom type with a number of methods), I place each implementation in its own subfolder and just symlink one of them to the parent folder. If you expect to switch between implementations later (perhaps to compare benchmark results), it's also worth adding a `.gitignore` containing the symlink target. See the [Circular Buffer](circular-buffer) exercise for an example.

Admittedly, this all feels a little suspect to me. If someone knows of a better way to handle these situations, please let me know!


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

### Bank Account
Go's `defer` statement makes working with mutex locks much simpler.

### Binary Search

The classic binary search algorithm, implemented in both the closed and half-open styles.

### Binary Search Tree

A simple binary search tree, with no deletion or rebalancing.

### Circular Buffer

The classic circular buffer, implemented three ways:

- the classic fixed-sized slice with modular indexing;
- the "sliding slice", using the subslice operator to remove the element at the front and the `append` function to add an element at the back;
- using a buffered channel.

Compared to the fixed slice approach, the sliding slice has more concise and readable code, but the backing slice has to be reallocated as it "slides to the right" with repeated write and read calls. Surprisingly, it actually outperforms the fixed-slice solution in the benchmarks, despite the additional allocations.

Using a buffered channel to implement a simple queue is tempting in Go (which has no built-in queue type, but does have beautiful syntax for sending and receiving on channels), but unwise (see *The Go Programming Language* p. 233). This implementation did far worse in the benchmarks than either of the slice-based ones, presumably because of all the thread safety that's built into channels but not needed here.

### Collatz Conjecture

When the divisor is 2, there is a small but noticeable performance advantage in using bit operations instead of arithmetic ones (e.g., `n&1 == 0` instead of `n%2 == 0`, `n >>= 1` instead of `n /= 2`), at least for Go 1.23 on darwin/arm64.

### Custom Set

This exercise appears in many language tracks. Since most languages already provide a built-in set type, people typically implement their custom set with a resizable array (`vector` in C++, `ArrayList` in Java, `Vec` in Rust, etc.); since Go lacks a built-in set type, it's completely reasonable to implement the custom set as a simple wrapper around `map[string]struct{}`. I did it both ways.

I also have to mention a great solution to the common task of pretty-printing a custom collection type. I had originally implemented the `String` method with the classic reset-the-prefix approach:
```go
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

### Flatten Array

Recall that type switches provide a `nil` case. Very useful!

### Isogram

Four different approaches with interesting performance differences. A good reminder that big-O analysis can be misleading when data sizes are small: a quadratic array-based solution can significantly outperform a linear hash-map-based one. Also, one of these solutions looks like it shouldn't work for non-ASCII strings, but actually does.

### Leap

Even the simplest of problems can have some room for optimization. Also, the README for this exercise explains how Exercism generates test data from a cross-language repository.

### Linked List
The classic doubly linked list, suitable for implementing a deque.

### List Ops

The classic functional, higher-order list functions (`map`, `filter`, `foldl`, `foldr`, `concat`, etc.), implemented with imperative techniques (for efficiency).

### Parallel Letter Frequency

Go's built-in concurrency features really shine here: goroutines and channels make this exercise embarrassingly easy. I experimented with two approaches to merging the maps into a final result.

Looking through the community solutions, I found that many users used a *buffered* channel, with a capacity equal to the number of strings to be processed: `ch := make(chan FreqMap, len(texts))`. There's really no reason to use a buffered channel here:

- on the sending side, the very last thing each worker goroutine does is send the frequency map for their text, so it's fine if they block here while waiting for the channel to clear.
- on the receiving side, the main goroutine can only process (merge) one frequency map at a time.

Using a buffered channel still produces the correct result, but does impose a small time and memory cost (verified with benchmarking), so best not to use one where it isn't needed.

Also note that most of the community solutions were written before the "loop variable capture" issue [was fixed in Go 1.22](https://go.dev/blog/loopvar-preview), so the function passed to each goroutine takes a (seemingly unnecessary) string argument.

### Reverse String
Using two loop variables instead of one makes the solution more expressive (and ever so slightly faster).

### Roman Numerals
We can store the Arabic-to-Roman conversions in a pair of arrays (of types `int` and `string`), or a single array of `struct` type. The latter feels a little safer and more expressive, but the former runs slightly faster.

### Sieve

As in the Sieve of Eratosthenes. The canonical structure for keeping track of the marked/unmarked numbers is a slice of bools, but just for fun, I also wrote a solution using a custom bit-array type to save memory.

### Simple Linked List
The classic singly linked list (with no tail pointer), suitable for implementing a stack.

### Strain

Generic functions in Go.
