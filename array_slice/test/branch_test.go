package test

import "testing"

func BenchmarkSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 100)
		for i, v := range s {
			s[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkArray100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [100]int{}
		for i, v := range a {
			a[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 1000)
		for i, v := range s {
			s[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkArray1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [1000]int{}
		for i, v := range a {
			a[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 10000)
		for i, v := range s {
			s[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkArray10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [10000]int{}
		for i, v := range a {
			a[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 100000)
		for i, v := range s {
			s[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkArray100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [100000]int{}
		for i, v := range a {
			a[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkSlice1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 1000000)
		for i, v := range s {
			s[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkArray1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [1000000]int{}
		for i, v := range a {
			a[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkSlice10000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 10000000)
		for i, v := range s {
			s[i] = 1 + i
			_ = v
		}
	}
}

func BenchmarkArray10000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [10000000]int{}
		for i, v := range a {
			a[i] = 1 + i
			_ = v
		}
	}
}
