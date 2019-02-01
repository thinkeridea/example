package iouitl_readall

import (
	"strconv"
	"strings"
	"testing"
)

var request map[string]string

func init() {
	request = make(map[string]string, 100)
	for i := 0; i < 100; i++ {
		request["X"+strconv.Itoa(i)] = strings.Repeat("A", i/2)
	}
}
func BenchmarkJson(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := Json(request)
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}

func BenchmarkJsonIter(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := JsonIter(request)
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}

func BenchmarkJsonPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := JsonPool(request)
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}

func BenchmarkJsonIterPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := JsonIterPool(request)
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}
