package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	l.AddFirst(&node{
		Key:  1,
		Val:  1,
		Next: nil,
		Prev: nil,
	})
	l.AddFirst(&node{
		Key:  2,
		Val:  2,
		Next: nil,
		Prev: nil,
	})
	l.AddFirst(&node{
		Key:  3,
		Val:  3,
		Next: nil,
		Prev: nil,
	})

	assert.Equal(t, l.Size(), 3)

	last := l.RemoveLast()
	assert.Equal(t, last.Key, 1)
	last = l.RemoveLast()
	assert.Equal(t, last.Key, 2)
	last = l.RemoveLast()
	assert.Equal(t, last.Key, 3)

	l.AddFirst(&node{
		Key:  1,
		Val:  1,
		Next: nil,
		Prev: nil,
	})
	l.AddFirst(&node{
		Key:  2,
		Val:  2,
		Next: nil,
		Prev: nil,
	})
	l.AddFirst(&node{
		Key:  3,
		Val:  3,
		Next: nil,
		Prev: nil,
	})
	l.Remove(node{
		Key: 2,
		Val: 2,
	})
	l.RemoveLast()
	last = l.RemoveLast()
	assert.Equal(t, last.Key, 3)
}
