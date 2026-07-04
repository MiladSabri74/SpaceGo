package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value   any
	NextPtr *Node
	PrevPtr *Node
}

type List struct {
	size int
	Head *Node
	Tail *Node
}

func NewList(elements ...any) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func NewNode(element any) *Node {
	return &Node{
		Value:   element,
		NextPtr: nil,
		PrevPtr: nil,
	}
}

func (n *Node) Next() *Node {
	return n.NextPtr
}

func (n *Node) Prev() *Node {
	return n.PrevPtr
}

func (l *List) Unshift(v any) {
	n := NewNode(v)
	if l.Head != nil {
		n.NextPtr = l.Head
		l.Head.PrevPtr = n
	} else {
		l.Tail = n
	}
	l.Head = n
	l.size++
}

func (l *List) Push(v any) {
	n := NewNode(v)
	if l.Tail != nil {
		n.PrevPtr = l.Tail
		l.Tail.NextPtr = n
	} else {
		l.Head = n
	}
	l.Tail = n
	l.size++
}

func (l *List) Shift() (any, error) {
	if l.Head == nil {
		return nil, errors.New("list is empty")
	}

	value := l.Head.Value
	newHead := l.Head.NextPtr

	l.Head = newHead
	if newHead != nil {
		newHead.PrevPtr = nil
	} else {
		l.Tail = nil
	}

	l.size--
	return value, nil
}

func (l *List) Pop() (any, error) {
	if l.Tail == nil {
		return nil, errors.New("list is empty")
	}

	value := l.Tail.Value
	newTail := l.Tail.PrevPtr

	l.Tail = newTail
	if newTail != nil {
		newTail.NextPtr = nil
	} else {
		l.Head = nil
	}

	l.size--
	return value, nil
}

func (l *List) Reverse() {
	current := l.Head
	var temp *Node

	for current != nil {
		temp = current.PrevPtr
		current.PrevPtr = current.NextPtr
		current.NextPtr = temp
		current = current.PrevPtr
	}
	temp = l.Head
	l.Head = l.Tail
	l.Tail = temp
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}

func (l *List) Count() int {
	return l.size
}

func (l *List) Delete(v any) bool {
	current := l.Head
	for current != nil {
		if current.Value == v {
			if current.PrevPtr != nil {
				current.PrevPtr.NextPtr = current.NextPtr
			} else {
				l.Head = current.NextPtr
			}

			if current.NextPtr != nil {
				current.NextPtr.PrevPtr = current.PrevPtr
			} else {
				l.Tail = current.PrevPtr
			}
			l.size--
			return true
		}
		current = current.NextPtr
	}
	return false
}
