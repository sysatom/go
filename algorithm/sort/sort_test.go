package sort

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	r := SelectionSort(a)
	assert.Equal(t, r, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func BenchmarkSelectionSort(b *testing.B) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		SelectionSort(a)
	}
}

func TestInsertionSort(t *testing.T) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	r := InsertionSort(a)
	assert.Equal(t, r, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func BenchmarkInsertionSort(b *testing.B) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		InsertionSort(a)
	}
}

func TestBubbleSort(t *testing.T) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	r := BubbleSort(a)
	assert.Equal(t, r, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func BenchmarkBubbleSort(b *testing.B) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func TestCocktailSort(t *testing.T) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	r := CocktailSort(a)
	assert.Equal(t, r, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func BenchmarkCocktailSort(b *testing.B) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		CocktailSort(a)
	}
}

func TestQuickSort(t *testing.T) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	r := QuickSort(a)
	assert.Equal(t, r, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func BenchmarkQuickSort(b *testing.B) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	a := []int{4, 8, 2, 5, 6, 7, 1, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		sort.Ints(a)
	}
}
