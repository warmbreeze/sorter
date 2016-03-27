package insertion

func InsertionSort(values []int) {

	for i := 1; i < len(values); i++ {
		var sorting int = values[i] // the element will be inserted
		var j int = 0
		for j = i - 1; j >= 0 && values[j] > sorting; j-- {
			values[j+1] = values[j]
		}
		values[j+1] = sorting
	}

}
