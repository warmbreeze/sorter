package shell

import "testing"

func TestShellSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 0}
	ShellSort(values)
	if values[0] != 0 || values[1] != 1 || values[2] != 2 || values[3] != 3 || values[4] != 4 ||
		values[5] != 5 {
		t.Error("ShellSort() failed, Got", values, "Excepted 0 1 2 3 4 5")
	}
}

func TestShellSort2(t *testing.T) {
	values := []int{6}
	ShellSort(values)
	if values[0] != 6 {
		t.Error("ShellSort() failed, Got", values, "Excepted 6")
	}
}

func TestShellSort3(t *testing.T) {
	values := []int{10, 14, 73, 25, 23}
	ShellSort(values)
	if values[0] != 10 || values[1] != 14 || values[2] != 23 || values[3] != 25 ||
		values[4] != 73 {
		t.Error("ShellSort() failed, Got", values, "Excepted 10 14 23 25 73")
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort([]int{9, 4, 3, 2, 5, 0})
	}
}
