package linkedlist

import "errors"

type List struct {
	head *Node
	size int
}

type Node struct {
	data int
	next *Node
}

func New(elements []int) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (l *List) Size() (size int) {
	return l.size
}

func (l *List) Push(element int) {
	l.head = &Node{data: element, next: l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("empty list")
	}
	popped := l.head.data
	l.head = l.head.next
	l.size--
	return popped, nil
}

func (l *List) Array() []int {
	s := make([]int, l.size)
	for i, node := l.size-1, l.head; node != nil; i, node = i-1, node.next {
		s[i] = node.data
	}
	return s
}

func (l *List) Reverse() *List {
	r := &List{}
	for node := l.head; node != nil; node = node.next {
		r.Push(node.data)
	}
	return r
}
