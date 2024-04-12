package main

import "fmt"

func main() {
	a := 10
	b := 20

	/*
		defer func(a, b int) { // main函數結束前那一刻才執行
			fmt.Printf("a = %d, b = %d\n", a, b)
		}(a, b) // 把參數傳遞過去 已經先傳遞了參數 只是沒有調用
	*/
	// 上面的等價於

	defer func(a, b int) { // main函數結束前那一刻才執行
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(10, 20) // 把參數傳遞過去 已經先傳遞了參數 只是沒有調用

	a = 111 // 先會運行這裏
	b = 222
	fmt.Printf("外部： a = %d, b = %d\n", a, b)
}

func main001() {
	a := 10
	b := 20

	defer func() { // main函數結束前那一刻才執行
		fmt.Printf("a = %d, b = %d\n", a, b)
	}()

	a = 111
	b = 222
	fmt.Printf("外部： a = %d, b = %d\n", a, b)

}
