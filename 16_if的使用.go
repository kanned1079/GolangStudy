package main

import "fmt"

func main() {
	s := "王思聰2"
	// if 和 { 就是條件，條件通常都是關係運算符
	if s == "王思聰" { // 左括號和if在同一行
		fmt.Println("aaa")
	}

	// if支持一個初始化語句，初始化語句和條件判斷以分號分割
	if a := 10; a == 10 {
		fmt.Println("yes")
	}

	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	// 多分支結構
	a = 12
	if a == 10 {
		fmt.Printf("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("impossible")
	}

}
