package main

import "fmt"

/*
定義在 {} 裡的變量就是局部變量
作用域 只在 {} 裡有效

執行到定義變量的那句話時候 才開始分配空間 離開了就自動釋放
*/

func main() {
	{
		i := 20
		fmt.Println("i = ", i)
	}
	//i = 111 // 這個不行

	// 經常的錯誤
	if flag := 3; flag == 3 {
		fmt.Println("flag = ", flag)
	}
	//flag = 4 // 這樣不行 離開了if語句
}

func test004() {
	a := 10
	fmt.Println("a = ", a)
}
