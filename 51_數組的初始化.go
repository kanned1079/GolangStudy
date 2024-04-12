package main

import "fmt"

func main() {
	// 聲明定義同時賦值 叫做初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5} // 這種叫做全部初始化
	fmt.Println("a = ", a)

	// 部分初始化 沒有初始化的元素 自動賦值為0
	b := [5]int{1, 2, 3}
	fmt.Println("b = ", b)

	// 指定某個元素初始化 沒賦值的就是0
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d = ", d)
}
