package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 9, 10)

	// 引用底层的数组地址
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))

	fmt.Println(len(s), cap(s))
	s = append(s, 1)

	// 引用底层的数组地址
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))

	fmt.Println(len(s), cap(s))
	s = append(s, 1)

	// 引用底层的数组地址
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))

	fmt.Println(len(s), cap(s))
}
