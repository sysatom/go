package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, q.Size(), 3)
	assert.Equal(t, q.Dequeue(), 1)
	assert.Equal(t, q.Dequeue(), 2)
	assert.Equal(t, q.Dequeue(), 3)
	assert.Equal(t, q.Dequeue(), -1)
	q.Enqueue(4)
	assert.Equal(t, q.Size(), 1)
}
