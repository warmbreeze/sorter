package merge

import "testing"

func TestMergeSort1(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	MergeSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 ||
		values[4] != 4 {
		t.Error("MergeSort() failed, Got", values, "Excepted 0 1 2 3 4")
	}
}

func TestMergeSort2(t *testing.T) {
	values := []int{0, 0}
	MergeSort(values)
	if values[0] != 0 || values[1] != 0 {
		t.Error("MergeSort() failed, Got", values, "Excepted 0 0")
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{9, 4, 3, 5, 2, 3432, 55})
	}
}
