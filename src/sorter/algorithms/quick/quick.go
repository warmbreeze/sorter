package quick

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, left, right int) {
	pivot := values[left]
	p := left
	i, j := left, right

	for i <= j {
		// look smaller from right
		for j >= p && values[j] >= pivot {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		// look bigger from left
		for i <= p && values[i] <= pivot {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = pivot

	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}
