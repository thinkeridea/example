package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	var s []int
	s1 := make([]int, 0)
	s2 := make([]int, 0, 0)
	s3 := make([]int, 0, 100)

	arr := [10]int{}
	s4 := arr[:0]

	fmt.Println(strings.Repeat("--s--", 10))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s)))
	fmt.Println(s)
	fmt.Println(s == nil)

	fmt.Println(strings.Repeat("--s1--", 10))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	fmt.Println(strings.Repeat("--s2--", 10))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Println(s2)
	fmt.Println(s2 == nil)

	fmt.Println(strings.Repeat("--s3--", 10))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
	fmt.Println(s3)
	fmt.Println(s3 == nil)

	fmt.Println(strings.Repeat("--s4--", 10))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Println(s4)
	fmt.Println(s4 == nil)
}
