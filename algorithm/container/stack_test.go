package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, s.Size(), 3)
	assert.Equal(t, s.Pop(), 3)
	assert.Equal(t, s.Pop(), 2)
	assert.Equal(t, s.Pop(), 1)
	assert.Equal(t, s.Pop(), -1)
	s.Push(4)
	assert.Equal(t, s.Size(), 1)
}
