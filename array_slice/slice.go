package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 10, 20)

	s[2] = 100
	s[9] = 200

	size := unsafe.Sizeof(0)
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size*2)))

	fmt.Println(*(*[20]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
}
