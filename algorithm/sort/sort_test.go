package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{4, 1, 2, 5, 6, 7, 8, 0, 3, 9, }
	r := SelectionSort(a)
	fmt.Println(r)
}

func BenchmarkSelectionSort(b *testing.B)  {
	a := []int{4, 1, 2, 5, 6, 7, 8, 0, 3, 9, }
	for i := 0; i < b.N; i++ {
		SelectionSort(a)
	}
}
