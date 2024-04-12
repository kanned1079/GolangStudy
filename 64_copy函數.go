package main

import "fmt"

func main() {
	srcSlice := []int{1, 3}
	dstSlice := []int{6, 6, 6, 6, 6}

	// copy(目標, 源)
	copy(dstSlice, srcSlice) // [1 3 6 6 6]
	fmt.Println("dstSlice = ", dstSlice)

}
