package main

import "fmt"

/* 函數的格式
func [函數名]( [實參] ) [返回值] {
	[語句]
}
*/

// 程序從main方法開始執行
func main() {
	// 函數需要被調用
	MyFunc()
	MyFunc1(5)
	Myfunc2(666, 777)
	Myfunc3(222, "wwwwww", 1.54)
}

// 無參數 無返回值 的函數定義
// 函數放前後都可
func MyFunc() {
	a := 66
	fmt.Println("a = ", a)
}

// 有參數 無返回值 的函數定義
func MyFunc1(a int) {
	fmt.Println("a = ", a)
}

// 多個參數
func Myfunc2(a, b int) {
	fmt.Println("a = ", a, ", b = ", b)
}

// 多個參數，但是類型不同
func Myfunc3(a int, b string, c float64) {
	fmt.Printf("a = %d, b = %s, c = %f\n", a, b, c)
}
