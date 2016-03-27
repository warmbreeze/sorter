package insertion

import "testing"

func TestInsertionSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 0}
	InsertionSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 ||
		values[4] != 4 || values[5] != 5 {
		t.Error("InsertionSort() failed, Got", values, "Excepted 0 1 2 3 4 5")
	}
}

func TestInsertionSort2(t *testing.T) {
	values := []int{1, 6}
	InsertionSort(values)
	if values[0] != 1 || values[1] != 6 {
		t.Error("InsertionSort() failed, Got", values, "Excepted 1 6")
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort([]int{9, 8, 6, 4, 3, 2})
	}
}
