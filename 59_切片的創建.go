package main

import "fmt"

func main() {
	// 傳統的自動推導 同時初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	// 借助make函數 格式 make(切片類型, 長度, 容量)
	s2 := make([]int, 5, 20)
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))

	s3 := make([]int, 5) // 不指定長度 容量和長度相同
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))

}

func main0001() {
	// 切片和數組的區別
	// 數組[]中的長度是一個常量不能修改 固定長度
	a := [5]int{}
	fmt.Printf("a: len=%d, cap=%d\n", len(a), cap(a))

	// 切片
	// []裡為空或者... 切片的長度和容量不固定
	s := []int{}
	fmt.Printf("s: len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 11) // 給切片魔紋追加一個成員
	fmt.Printf("appended: len=%d, cap=%d\n", len(s), cap(s))

}
