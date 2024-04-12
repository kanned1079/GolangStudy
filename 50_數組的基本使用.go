package main

import "fmt"

func main() {
	// 定義一個數組 [10]int 和 [5]int 是兩個不同類型
	// [數字] 這個數字是作為數組的元素個數
	var a [10]int
	var b [5]int

	fmt.Printf("a len = %d, b len = %d\n", len(a), len(b))

	// 操作數組元素 從0開始 到len()-1 不對稱元素 這個操作時候的數組叫做下標
	// 這個下標可以是常量或者變量
	a[0] = 1
	i := 1
	a[i] = 2

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	// 第一個返回下標 第二個返回元素
	for i, data := range a {
		fmt.Printf("i = %d, data = %d\n", i, data)
	}
}
