package main

import "fmt"

/*
	append 函数会智能地底层数组的容量增长
	一旦超过原底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的教据
*/

func main() {
	s := make([]int, 0, 1)
	oldCap := cap(s)
	for i := 0; i < 8; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap: %d --> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
