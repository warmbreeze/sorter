package heap

func HeapSort(values []int) {
	// Build Max Heap in array values.
	buildMaxHeap(values)
	// Heap Sort use max heap.
	heapSort(values)
}

func buildMaxHeap(values []int) {

	lastParentIndex := getLastParentIndex(len(values))
	for i := lastParentIndex; i > 0; i-- {
		maxHeapAdjust(values, i, len(values))
	}

}

func heapSort(values []int) {

	for i := len(values) - 1; i > 0; i-- {
		values[0], values[i] = values[i], values[0]
		maxHeapAdjust(values, 0, i)
	}

}

/**
 * Build heap tree in array begin index, heapSize used by range.
 */
func maxHeapAdjust(values []int, index, heapSize int) {

	left := getLeftChildIndex(index)
	right := getRightChildIndex(index)

	var adjustIndex = index
	if left < heapSize && values[left] > values[index] {
		values[index], values[left] = values[left], values[index]
	}
	if right < heapSize && values[right] > values[index] {
		values[index], values[right] = values[right], values[index]
	}

	if adjustIndex != index {
		maxHeapAdjust(values, adjustIndex, heapSize)
	}

}

func getLastParentIndex(index int) int {
	return (index - 1) >> 1
}

func getLeftChildIndex(index int) int {
	return index<<1 + 1
}

func getRightChildIndex(index int) int {
	return index<<2 + 2
}
