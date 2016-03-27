package selection

import "testing"

func TestSelectionSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 0}
	SelectionSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 ||
		values[4] != 4 || values[5] != 5 {
		t.Error("SelectibSort() failed, Got", values, "Excepted 0 1 2 3 4 5")
	}
}

func TestSelectionSort2(t *testing.T) {
	values := []int{6, 1}
	SelectionSort(values)
	if values[0] != 1 || values[1] != 6 {
		t.Error("SelectionSort() failed, Got", values, "Excepted 1 6")
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{9, 8, 4, 5, 7, 3, 1})
	}
}
