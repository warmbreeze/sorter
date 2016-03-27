package quick

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 0}
	QuickSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 ||
		values[4] != 4 || values[5] != 5 {
		t.Error("QuickSort() failed, Got", values, "Excepted 0 1 2 3 4 5")
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{0, 8, 12, 5}
	QuickSort(values)
	if values[0] != 0 || values[1] != 5 || values[2] != 8 ||
		values[3] != 12 {
		t.Error("QuickSort() failed, Got", values, "Excepted 0 5 8 12")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	}
}
