package container

import (
	"fmt"
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

	if pq.Max() != 9 {
		t.Error("error max")
	}
	max := pq.DelMax()
	if max != 9 {
		t.Error("error max")
	}
	max = pq.DelMax()
	if max != 8 {
		t.Error("error max")
	}
	max = pq.DelMax()
	if max != 7 {
		t.Error("error max")
	}
	if pq.Max() != 6 {
		t.Error("error max")
	}
	pq.Insert(10)
	if pq.Max() != 10 {
		t.Error("error max")
	}
	pq.Insert(11)
	if pq.Max() != 11 {
		t.Error("error max")
	}
	pq.Insert(12)
	if pq.Max() != 12 {
		t.Error("error max")
	}
	pq.Insert(13)
	if pq.Max() != 13 {
		t.Error("error max")
	}
	fmt.Println(pq.heap)
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

	if pq.Min() != 0 {
		t.Error("error min")
	}
	max := pq.DelMin()
	if max != 0 {
		t.Error("error min")
	}
	max = pq.DelMin()
	if max != 1 {
		t.Error("error min")
	}
	max = pq.DelMin()
	if max != 2 {
		t.Error("error min")
	}
	if pq.Min() != 3 {
		t.Error("error min")
	}
	pq.Insert(-2)
	if pq.Min() != -2 {
		t.Error("error max")
	}
	pq.Insert(-3)
	if pq.Min() != -3 {
		t.Error("error max")
	}
	pq.Insert(-4)
	if pq.Min() != -4 {
		t.Error("error max")
	}
	pq.Insert(-5)
	if pq.Min() != -5 {
		t.Error("error max")
	}
	fmt.Println(pq.heap)
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
