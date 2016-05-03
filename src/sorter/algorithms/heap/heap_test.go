package heap

import "testing"

func TestHeapSort1(t *testing.T) {
	values := []int{6, 5, 4, 3, 2, 1, 0}
	HeapSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 ||
		values[4] != 4 || values[5] != 5 || values[6] != 6 {
		t.Error("HeapSort() failed, Got", values, "Excepted 0 1 2 3 4 5")
	}
}

func TestHeapSort2(t *testing.T) {
	values := []int{2, 0, 1}
	HeapSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 {
		t.Error("HeapSort() failed, Got", values, "Excepted 1 2 2")
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort([]int{8, 4, 2, 3, 3, 4, 6, 9})
	}
}
