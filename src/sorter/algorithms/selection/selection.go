package selection

func SelectionSort(values []int) {
	for i := 0; i < len(values); i++ {
		var min int = i
		for j := i + 1; j < len(values); j++ {
			if values[min] > values[j] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
}
