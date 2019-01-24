package slice

import "testing"

func TestSliceInt2String1(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := SliceInt2String1(s)
	if ss != "1,2,3,4,5,6,7,8,9,10" {
		t.Errorf("s:%v ss:%s", s, ss)
	}
}

func TestSliceInt2String2(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := SliceInt2String2(s)
	if ss != "1,2,3,4,5,6,7,8,9,10" {
		t.Errorf("s:%v ss:%s", s, ss)
	}
}

func TestSliceInt2String3(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := SliceInt2String3(s)
	if ss != "1,2,3,4,5,6,7,8,9,10" {
		t.Errorf("s:%v ss:%s", s, ss)
	}
}

func TestSliceInt2String4(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := [][]int{s, s, s, s, s, s, s, s, s}
	res := SliceInt2String4(ss)

	if len(ss) != len(res) {
		t.Errorf("len(ss):%d != len(res):%d", len(ss), len(res))
	}

	for _, v := range res {
		if v != "1,2,3,4,5,6,7,8,9,10" {
			t.Errorf("v:%s", v)
		}
	}
}

func TestSliceInt2String5(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := [][]int{s, s, s, s, s, s, s, s, s}
	res := SliceInt2String5(ss)

	if len(ss) != len(res) {
		t.Errorf("len(ss):%d != len(res):%d", len(ss), len(res))
	}

	for _, v := range res {
		if v != "1,2,3,4,5,6,7,8,9,10" {
			t.Errorf("v:%s", v)
		}
	}
}

func BenchmarkExperiment1SliceInt2String1(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ss := SliceInt2String1(s)
		_ = ss
	}
}

func BenchmarkExperiment1SliceInt2String2(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ss := SliceInt2String2(s)
		_ = ss
	}
}

func BenchmarkExperiment1SliceInt2String3(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ss := SliceInt2String3(s)
		_ = ss
	}
}

func BenchmarkExperiment2SliceInt2String4(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := [][]int{s, s, s, s, s, s, s, s, s}
	for i := 0; i < b.N; i++ {
		res := SliceInt2String4(ss)
		_ = res
	}
}

func BenchmarkExperiment2SliceInt2String5(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := [][]int{s, s, s, s, s, s, s, s, s}
	for i := 0; i < b.N; i++ {
		res := SliceInt2String5(ss)
		_ = res
	}
}
