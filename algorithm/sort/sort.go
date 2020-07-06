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

func InsertionSort(items []int) []int {
	length := len(items)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && items[j] < items[j-1]; j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
	return items
}

func BubbleSort(items []int) []int {
	length := len(items)
	for i := 0; i < length; i++ {
		for j := i+1; j < length; j++ {
			if items[i] > items[j] {
				items[j], items[i] = items[i], items[j]
			}
		}
	}
	return items
}
