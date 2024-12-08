// Package circular provides a circular buffer of bytes supporting both
// overflow-checked writes and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader and
// io.ByteWriter and can be used (size permitting) as a drop in replacement for
// anything using that interface.
package circular

import "errors"

type Buffer struct {
	data        []byte
	start, used int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.used == 0 {
		return 0, errors.New("buffer is empty")
	}
	val := b.data[b.start]
	b.start = (b.start + 1) % len(b.data)
	b.used--
	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.used == len(b.data) {
		return errors.New("buffer is full")
	}
	b.data[(b.start+b.used)%len(b.data)] = c
	b.used++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.used == len(b.data) {
		b.data[b.start] = c
		b.start = (b.start + 1) % len(b.data)
	} else {
		b.data[(b.start+b.used)%len(b.data)] = c
		b.used++
	}
}

func (b *Buffer) Reset() {
	b.start, b.used = 0, 0
}
