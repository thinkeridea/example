package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = [3]int{1, 2, 3}

	fmt.Println(unsafe.Sizeof(arr))
	size := unsafe.Sizeof(arr[0])

	// 获取数组指定索引元素的值
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 1*size)))

	// 设置数组指定索引元素的值
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 1*size)) = 10

	fmt.Println(arr[1])
}
