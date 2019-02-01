package iouitl_readall

import (
	"bytes"
	"testing"
)

var data = bytes.Repeat([]byte("ABCD"), 1000)

func BenchmarkIouitlReadAll(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := IouitlReadAll(bytes.NewReader(data))
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}

func BenchmarkIoCopy(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := IoCopy(bytes.NewReader(data))
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}

func BenchmarkIouitlReadAllAndJson(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := IouitlReadAllAndJson(bytes.NewReader(data))
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}

func BenchmarkIoCopyAndJson(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := IoCopyAndJson(bytes.NewReader(data))
			if err != nil {
				b.Error(err.Error())
			}
		}
	})
}
