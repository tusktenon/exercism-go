package linkedlist

import "errors"

// Define List and Node types here.
type List struct {
	head, tail *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value      interface{}
	prev, next *Node
}

// NewList returns a new linked list, preserving the order of the values.
func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

// Next returns a pointer to the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns a pointer to the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift inserts a value at the front of the list.
func (l *List) Unshift(v interface{}) {
	newhead := Node{Value: v, next: l.head}
	if l.head == nil {
		// list is currently empty
		l.tail = &newhead
	} else {
		l.head.prev = &newhead
	}
	l.head = &newhead
}

// Push inserts a value at the back of the list.
func (l *List) Push(v interface{}) {
	newtail := Node{Value: v, prev: l.tail}
	if l.tail == nil {
		// list is currently empty
		l.head = &newtail
	} else {
		l.tail.next = &newtail
	}
	l.tail = &newtail
}

// Shift removes the value at the front of the list, returning an error if the
// list is empty.
func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("empty list")
	}
	v := l.head.Value
	l.head = l.head.next
	if l.head == nil {
		// list is now empty
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return v, nil
}

// Pop removes the value at the back of the list, returning an error if the
// list is empty.
func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("empty list")
	}
	v := l.tail.Value
	l.tail = l.tail.prev
	if l.tail == nil {
		// list is now empty
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return v, nil
}

func (l *List) Reverse() {
	for curr := l.head; curr != nil; {
		next := curr.next
		curr.prev, curr.next = curr.next, curr.prev
		curr = next
	}
	l.head, l.tail = l.tail, l.head
}

// First returns a pointer to the first node in the list (the head).
func (l *List) First() *Node {
	return l.head
}

// Last returns a pointer to the last node in the list (the tail).
func (l *List) Last() *Node {
	return l.tail
}
