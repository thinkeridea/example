package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	s := make([]int, 10, 20)

	size := unsafe.Sizeof(0)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size*2)))

	slice(s)

	s1 := s
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s1)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size*2)))

	fmt.Println(strings.Repeat("-", 50))

	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size)) = 20

	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size*2)))

	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s1)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size*2)))

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(strings.Repeat("-", 50))

	s2 := s
	s2 = append(s2, 1)

	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)

}

func slice(s []int) {
	size := unsafe.Sizeof(0)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size*2)))
}
