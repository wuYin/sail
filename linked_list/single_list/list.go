package single_list

import (
	"bytes"
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
}

func NewList() *List {
	return new(List)
}

func (l *List) Append(v int) {
	node := &Node{val: v}
	if l.IsEmpty() {
		l.head = node
		return
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
}

func (l *List) Prepend(v int) {
	head := &Node{
		val:  v,
		next: l.head,
	}
	l.head = head
}

func (l *List) Remove(v int) bool {
	if l.IsEmpty() {
		return false
	}

	// 删除头结点
	if l.head.val == v {
		l.head = l.head.next
		return true
	}

	cur := l.head
	for cur.next != nil {
		if cur.next.val == v {
			cur.next = cur.next.next
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) String() string {
	if l.IsEmpty() {
		return ""
	}

	buf := bytes.NewBuffer(nil)
	cur := l.head
	for cur != nil {
		buf.WriteString(fmt.Sprintf("%d->", cur.val))
		cur = cur.next
	}
	buf.WriteString("Nil")
	return buf.String()
}
