package main

import (
	"fmt"
)

// 回調函數 函數有一個參數是函數類型 那麼這個函數就是回調函數

func main() {
	aa := Calc(1, 3, Add)
	fmt.Println("1+3 = ", aa)

	aa = Calc(2, 4, Minus)
	fmt.Println("2-4 = ", aa)

	aa = Calc(2, 4, Mul)
	fmt.Println("2x4 = ", aa)

}

type FuncType2 func(int, int) int

//可以實現多態 調用同一個接口

// 這裏有何個參數是函數類型
// 計算器可以進行四則運算
func Calc(a, b int, fTest FuncType2) (result int) {
	fmt.Println("Calc")
	// 先有想法 后面再实现
	//result = Add(a, b) // 不要這麼寫 這樣直接寫死了只能實現加發
	result = fTest(a, b) // 這個才是執行者 但是這個函數還沒有實現
	return
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}
