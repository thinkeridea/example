package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10, 10)
	fmt.Println(len(s), cap(s))
	s2 := make([]int, 40)

	s = append(s, s2...)
	fmt.Println(len(s), cap(s))

}
