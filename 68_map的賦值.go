package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "c++"}
	fmt.Println("m1 = ", m1)

	// 如果已存在的值 修改內容
	m1[1] = "c+++++"
	fmt.Println("m1 = ", m1)

	// 追加內存 map底層自動擴容 和append類似
	m1[3] = "golang"
	fmt.Println("m1 = ", m1)
	fmt.Printf("m1's type is %T\n", m1)
}
