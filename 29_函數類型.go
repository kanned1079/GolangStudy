package main

import "fmt"

/*
	在Go中 函數也是一種數據類型 可以使用type來定義/改名
	函數的類型就是所有擁有相同的參數，相同的返回值的一種類型
*/

func main() {
	var result int
	result = AddNums(1, 1) // 傳統的調用方法
	fmt.Println("result = ", result)

	var F_Test1 FuncType1    // 聲明一個函數類型變量 變量名叫F_Test1
	F_Test1 = AddNums        // 是變量就可以賦值
	result = F_Test1(10, 20) // 等價於 AddNums(10, 20)
	fmt.Println("result2 = ", result)

	F_Test1 = MyMinus
	result = F_Test1(5, 10) // 等價於Minus(5, 10)
	fmt.Println("result3 = ", result)
	// 多態

}

func AddNums(a, b int) int {
	return a + b
}

func MyMinus(a, b int) int {
	return a - b
}

// 函數也是一種數據類型 可以使用type來定義/改名
// FuncType1是一個函數類型 可以定義變量
type FuncType1 func(int, int) int
