package sort

func SelectionSort(items []int) []int {
	length := len(items)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
	return items
}
