// This implementation uses the slice operator to remove the element at the
// front of the buffer and the built-in append function to add an element at
// the end. It's a popular community solution, and the code is more succinct
// and readable than the "standard" implementation (a fixed data slice with
// modular indexing), but it suffers from a major performance disadvantage: as
// the data slice "shifts to the right" with multiple writes and reads, it must
// be repeatedly reallocated.
package circular

import "errors"

type Buffer struct {
	data []byte
	cap  int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: []byte{}, cap: size}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.data) == 0 {
		return 0, errors.New("buffer is empty")
	}
	val := b.data[0]
	b.data = b.data[1:]
	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.data) == b.cap {
		return errors.New("buffer is full")
	}
	b.data = append(b.data, c)
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.data) == b.cap {
		b.data = b.data[1:]
	}
	b.data = append(b.data, c)
}

func (b *Buffer) Reset() {
	b.data = []byte{}
}
