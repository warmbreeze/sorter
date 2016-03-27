package bubble

func BubbleSort(values []int) {

	for i := len(values); i > 0; i-- {
		var sorted bool = true // used for fast return if the input values was sorted.

		for j := 0; j < i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}
}
