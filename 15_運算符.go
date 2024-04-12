package main

import "fmt"

func main() {
	fmt.Println("4 > 3 結果: ", 4 > 3)
	fmt.Println("4 != 3 結果: ", 4 != 3)

	// 非
	fmt.Println("!(4 > 3) 結果: ", !(4 > 3))

	// &&並且 左右都為真，那麼結果才為真
	fmt.Println("true && true 結果：", true && true)
	fmt.Println("true && false 結果：", true && false)

	// ||或者 左右都為假，結果才為假
	fmt.Println("true || false 結果：", true || false)
	fmt.Println("true || false 結果：", false && false)

	a := 8
	fmt.Printf("a = %d\n", a)
	//fmt.Println("0 <= a <= 10 結果為：", 0 <= a <= 10)
	// 不能這麼個寫
	// GO語言的bool類型和int不兼容
	fmt.Println("0 <= a <= 10 結果為：", 0 <= a && a <= 10)
	// 得這麼寫

	// & 取地址
	// * 取值

	/* 優先級別
	6  * / % << >> & &^
	5  _ - | ^
	4  == != < <= >= >
	3 <-
	2 &&
	1 ||
	*/

}
