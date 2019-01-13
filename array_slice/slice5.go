package main

import "fmt"

func main() {
	var s = make([]int, 10)
	fmt.Println(s)

	var s1 = make([]*int, 10)
	fmt.Println(s1)

	var s2 = make([]string, 10)
	fmt.Println(s2)
}
