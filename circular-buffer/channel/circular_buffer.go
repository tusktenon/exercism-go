// An implementation using a buffered channel. This makes for very concise and
// readable code, but is probably unwise all the same. Note that "The Go
// Programming Language" (p. 233) specifically warns against using a buffered
// channel as a queue, and recommends a slice-based implementation instead.
//
// Unsurprisingly, this implementation has far worse performance than either of
// the slice-based approaches (presumably because of all the thread safety
// that's built into channels but not needed here).
package circular

import "errors"

type Buffer struct {
	data chan byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make(chan byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.data) == 0 {
		return 0, errors.New("buffer is empty")
	}
	return <-b.data, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.data) == cap(b.data) {
		return errors.New("buffer is full")
	}
	b.data <- c
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.data) == cap(b.data) {
		<-b.data
	}
	b.data <- c
}

func (b *Buffer) Reset() {
	for len(b.data) != 0 {
		<-b.data
	}
}
