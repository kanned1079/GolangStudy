package main

import "fmt"

func main() {
	a, b := maxNum(48, 19)
	fmt.Printf("max = %d, min = %d\n", a, b)
	//fmt.Println(max(1, 9))
}

// 手動求兩個數的最大值
func maxNum(x, y int) (max, min int) {
	max, min = x, y
	if y > x {
		max, min = min, max
	}
	return
}
