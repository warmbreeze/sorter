package shell

func ShellSort(values []int) {
	var gap int = 1
	var i, j int = 0, 0
	for gap < len(values) {
		gap = gap*3 + 1 // gap will be: 1, 4, 13, 40, 121, ..., reversed
	}

	for ; gap > 0; gap /= 3 {
		for i = gap; i < len(values); i += gap {
			var temp int = values[i]
			for j = i - gap; j >= 0 && values[j] > temp; j -= gap {
				values[j+gap] = values[j]
			}
			values[j+gap] = temp
		}
	}
}
