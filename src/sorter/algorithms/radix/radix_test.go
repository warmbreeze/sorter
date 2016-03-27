package radix

import "testing"

func TestRadixSort1(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	RadixSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 || values[4] != 4 {
		t.Error("RadixSort() failed, Got", values, "Excepted 0 1 2 3 4")
	}
}

func TestRadixSort2(t *testing.T) {
	values := []int{0, 0}
	RadixSort(values)
	if values[0] != 0 || values[1] != 0 {
		t.Error("RadixSort() failed, Got", values, "Excepted 0 0")
	}
}

func BenchmarkRadixSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RadixSort([]int{9, 8, 3, 5, 2, 1, 0})
	}
}
