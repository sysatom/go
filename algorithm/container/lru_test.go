package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	cache.put(1, 1)
	cache.put(2, 2)
	v := cache.get(1)
	assert.Equal(t, v, 1)

	cache.put(3, 3)
	v = cache.get(2)
	assert.Equal(t, v, -1)

	cache.put(1, 4)
	v = cache.get(1)
	assert.Equal(t, v, 4)
}
