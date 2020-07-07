package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	pq := NewMaxPriorityQueue()
	pq.Insert(4)
	pq.Insert(8)
	pq.Insert(2)
	pq.Insert(5)
	pq.Insert(6)
	pq.Insert(9)
	pq.Insert(7)
	pq.Insert(1)
	pq.Insert(0)
	pq.Insert(3)

	assert.Equal(t, pq.Max(), 9)

	max := pq.DelMax()
	assert.Equal(t, max, 9)

	max = pq.DelMax()
	assert.Equal(t, max, 8)

	max = pq.DelMax()
	assert.Equal(t, max, 7)

	pq.Insert(10)
	assert.Equal(t, pq.Max(), 10)

	pq.Insert(11)
	assert.Equal(t, pq.Max(), 11)

	pq.Insert(12)
	assert.Equal(t, pq.Max(), 12)

	pq.Insert(13)
	assert.Equal(t, pq.Max(), 13)
}

func TestMinPriorityQueue(t *testing.T) {
	pq := NewMinPriorityQueue()
	pq.Insert(4)
	pq.Insert(8)
	pq.Insert(2)
	pq.Insert(5)
	pq.Insert(6)
	pq.Insert(9)
	pq.Insert(7)
	pq.Insert(1)
	pq.Insert(0)
	pq.Insert(3)

	assert.Equal(t, pq.Min(), 0)

	min := pq.DelMin()
	assert.Equal(t, min, 0)

	min = pq.DelMin()
	assert.Equal(t, min, 1)

	min = pq.DelMin()
	assert.Equal(t, min, 2)
	assert.Equal(t, pq.Min(), 3)

	pq.Insert(-2)
	assert.Equal(t, pq.Min(), -2)

	pq.Insert(-3)
	assert.Equal(t, pq.Min(), -3)

	pq.Insert(-4)
	assert.Equal(t, pq.Min(), -4)

	pq.Insert(-5)
	assert.Equal(t, pq.Min(), -5)
}

func BenchmarkMaxPriorityQueue(b *testing.B) {
	pq := NewMaxPriorityQueue()
	pq.Insert(4)
	pq.Insert(8)
	pq.Insert(2)
	pq.Insert(5)
	pq.Insert(6)
	pq.Insert(9)
	pq.Insert(7)
	pq.Insert(1)
	pq.Insert(0)
	pq.Insert(3)
	for i := 0; i < b.N; i++ {
		pq.Insert(10)
		pq.DelMax()
	}
}

func BenchmarkMinPriorityQueue(b *testing.B) {
	pq := NewMinPriorityQueue()
	pq.Insert(4)
	pq.Insert(8)
	pq.Insert(2)
	pq.Insert(5)
	pq.Insert(6)
	pq.Insert(9)
	pq.Insert(7)
	pq.Insert(1)
	pq.Insert(0)
	pq.Insert(3)
	for i := 0; i < b.N; i++ {
		pq.Insert(-2)
		pq.DelMin()
	}
}
