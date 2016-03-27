package bubble

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 0}
	BubbleSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 ||
		values[4] != 4 || values[5] != 5 {
		t.Error("BubbleSort() failed, Got", values, "Excepted 0 1 2 4 5")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 2, 4, 3}
	BubbleSort(values)
	if values[0] != 2 || values[1] != 3 || values[2] != 4 || values[3] != 5 ||
		values[4] != 5 {
		t.Error("BubbleSort() failed, Got", values, "Excepted 2 3 4 5 5")
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failed, Got", values, "Excepted 5")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}
