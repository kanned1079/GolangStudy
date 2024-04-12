package main

import "fmt"

/*
定義在函數外部的變量
*/

// a := 10 // 全局變量不支持自動推導
var a int = 10

func main() {
	fmt.Println("a = ", a)
	test005()
}

func test005() {
	fmt.Println("a2 = ", a)
}
