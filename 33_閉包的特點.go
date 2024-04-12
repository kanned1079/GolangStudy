package main

import "fmt"

func main() {
	// 返回值是一個匿名函數 所以返回是一個函數 通過f3來調用返回的匿名函數 f3來調用閉包函數
	f3 := test002()
	// 他不關心這些捕獲的變量和常量是否超出作用域
	// 只要閉包還在用 那麼變量就會存在
	fmt.Println("f3 = ", f3())
	fmt.Println("f3 = ", f3())
	fmt.Println("f3 = ", f3())

}

func main01() {
	fmt.Println("test001 = ", test001())
	fmt.Println("test001 = ", test001())
	fmt.Println("test001 = ", test001())
}

func test001() int {
	// 函數被調用的時候才給x分配空間 才會初始化為0
	var x int
	x++
	return x * x // 函數調用完後銷毀x
}

// 函數的返回值是一個匿名函數 就返回一個函數類型
func test002() func() int {
	var x int
	return func() int {
		x++
		return x * x // 函數調用完後銷毀x
	}
}

/*
// 函數的返回值是一個匿名函數 就返回一個函數類型
func test002() func() int {
	var x int
 ----------------------------------------
 ｜ return func() int {                 ｜
 ｜		x++                             ｜
 ｜		return x * x // 函數調用完後銷毀x  ｜
 ｜	}                                   ｜
 ----------------------------------------
}
*/
