package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [10]int{1, 2, 3}
	arr[7] = 100
	arr[9] = 200

	fmt.Println(arr)

	s1 := arr[:]
	s2 := arr[2:8]

	size := unsafe.Sizeof(0)
	fmt.Println("----------s1---------")
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s1)))
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&arr[0])))

	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size*2)))

	fmt.Println(s1)
	fmt.Println(*(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s1)))))

	fmt.Println("----------s2---------")
	fmt.Printf("%x\n", *(*uintptr)(unsafe.Pointer(&s2)))
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&arr[0]))+size*2)

	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s2)) + size)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s2)) + size*2)))

	fmt.Println(s2)
	fmt.Println(*(*[8]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s2)))))
}
