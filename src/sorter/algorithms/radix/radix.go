package radix

func RadixSort(values []int) {
	radixSort(values, len(values))
}

func radixSort(values []int, size int) {

	bit := maxBit(values, size)
	tmp := make([]int, size) // temp value array
	count := make([]int, 10) // count by radix

	radix := 1
	for i := 1; i <= bit; i++ {
		for j := 0; j < 10; j++ {
			count[j] = 0
		}
		for j := 0; j < size; j++ {
			k := (values[j] / radix) % 10
			count[k]++
		}
		for j := 1; j < 10; j++ {
			// allocate position in tmp replace two dimension array
			count[j] = count[j-1] + count[j]
		}
		// put element into bucket
		for i := size - 1; i >= 0; i-- {
			k := (values[i] / radix) % 10
			tmp[count[k]-1] = values[i]
			count[k]--
		}
		for i := 0; i < size; i++ {
			values[i] = tmp[i]
		}

		radix *= 10
	}
}

func maxBit(values []int, size int) int {
	maxValue := values[0]
	for i := 1; i < size; i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	d := 1
	for p := 10; maxValue > p; d++ {
		p *= 10
	}

	return d
}
