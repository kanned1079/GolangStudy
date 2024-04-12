package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// [low:high:max] 取下標從low開始的元素 len = high - low,  cap = max - low

	s1 := array[:] // 等價於 [0:len(array):len(array)]
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	// 操作某個元素 和操作數組相同
	data := array[0]
	fmt.Println("data = ", data)

	s2 := array[3:6:7] // a[3] a[4] a[5]  len: 6-3=3,  cap: 7-3=4
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	s3 := array[:6] // 從0開始 取6個元素 容量為6
	fmt.Println("s3 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[3:] // 從3開始 取到末尾
	fmt.Println("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
}
