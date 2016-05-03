package merge

func MergeSort(values []int) {
	length := len(values)
	merged := make([]int, length)
	mergeSort(values, merged, 0, length-1)
}

/**
 * Recursive devide array and merge sub-array.
 */
func mergeSort(values, merged []int, start, end int) {
	if start >= end {
		return
	}

	length := end - start
	var mid = (length >> 1) + start
	var leftStart, leftEnd = start, mid
	var rightStart, rightEnd = mid + 1, end
	mergeSort(values, merged, leftStart, leftEnd)
	mergeSort(values, merged, rightStart, rightEnd)

	index := start
	for ; leftStart <= leftEnd && rightStart <= rightEnd; index++ {
		if values[leftStart] < values[rightStart] {
			merged[index] = values[leftStart]
			leftStart++
		} else {
			merged[index] = values[rightStart]
			rightStart++
		}
	}
	for ; leftStart <= leftEnd; index++ {
		merged[index] = values[leftStart]
		leftStart++
	}
	for ; rightStart <= rightEnd; index++ {
		merged[index] = values[rightStart]
		rightStart++
	}

	for i := start; i <= end; i++ {
		values[i] = merged[i]
	}
}
