package container

type node struct {
	Key  int
	Val  int
	Next *node
	Prev *node
}

type DoublyLinkedList struct {
	size int
	head *node
	tail *node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) AddFirst(x *node) {
	if l.size == 0 {
		l.head = x
		l.tail = x
		l.size++
		return
	}

	head := l.head

	l.head = x
	l.head.Next = head
	head.Prev = x
	l.size++
}

func (l *DoublyLinkedList) AddLast(x *node) {
	if l.size == 0 {
		l.head = x
		l.tail = x
		l.size++
		return
	}

	tail := l.tail

	l.tail = x
	l.tail.Prev = tail
	tail.Next = x
	l.size++
}

func (l *DoublyLinkedList) Remove(x node) {
	if l.size == 0 {
		return
	}

	if l.head.Key == x.Key {
		l.head = l.head.Next
		l.head.Prev = nil
		l.size--
		return
	}

	if l.tail.Key == x.Key {
		prevTailNode := l.tail.Prev
		l.tail = prevTailNode
		prevTailNode.Next = nil
		l.size--
		return
	}

	var prevNode *node
	var nextNode *node
	cur := l.head

	for cur.Key != x.Key {
		if cur == nil {
			return
		}

		prevNode = cur
		cur = cur.Next
		nextNode = cur.Next
	}

	prevNode.Next = cur.Next
	nextNode.Prev = prevNode

	l.size--
}

func (l *DoublyLinkedList) RemoveLast() *node {
	if l.size == 0 {
		return nil
	}

	if l.size == 1 {
		l.size--
		tail := l.tail
		l.head = nil
		l.tail = nil
		return tail
	}

	l.size--
	tail := l.tail
	prev := tail.Prev
	prev.Next = nil
	l.tail = prev
	return tail
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}
