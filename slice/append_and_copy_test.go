package slice

import "testing"

func BenchmarkExperiment3Append1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < 20; j++ {
			s = append(s, []int{j, j + 1, j + 2, j + 3, j + 4}...)
		}
	}
}

func BenchmarkExperiment3Append2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 100)
		for j := 0; j < 20; j++ {
			s = append(s, []int{j, j + 1, j + 2, j + 3, j + 4}...)
		}
	}
}

func BenchmarkExperiment3Copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 100)
		n := 0
		for j := 0; j < 20; j++ {
			n += copy(s[n:], []int{j, j + 1, j + 2, j + 3, j + 4})
		}
	}
}
