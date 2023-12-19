package main

import "fmt"

func main() {
	arr := make([]int, 10)

	for i := 1; i <= len(arr)/2; i++ {
		arr[i-1] = i
	}

	// src & dest overlapped, then how to copy deal with
	src := arr[0:5]
	dest := arr[1:6]

	fmt.Println(arr)
	n := copy(dest, src)
	fmt.Println(arr)
	fmt.Println(n)

}
