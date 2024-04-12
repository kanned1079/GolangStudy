package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "mike"

	// 這樣寫的比較多
	// 匿名函数 没有函数名字
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1() // 進行調用

	// 這樣寫的人比較好
	// 給一個函數類型起別名
	type FuncType3 func() // 函數沒有參數 沒有返回值
	var f2 FuncType3
	f2 = f1
	f2()

	// 定義匿名函數 同時調用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() // 這裏的()表示調用此匿名函數

	// 帶參數的匿名函數
	f3 := func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(1, 2) //調用

	func(i, j int) {
		fmt.Printf("i2 = %d, j2 = %d\n", i, j)
	}(10, 20) // 上述的也可以這樣調用

	// 匿名函數 有參數有返回值
	max, min := func(i, j int) (max, min int) {
		max, min = i, j
		if j > i {
			max, min = min, max
		}
		return
	}(100, 202)
	fmt.Printf("x = %d, y = %d\n", max, min)

}
