package main

import "fmt"

// 實現兩個數相加

// 這是傳統面向過程的寫法
func Add01(a, b int) int {
	return a + b
}

// 面向對象 方法：給某個類型綁定一個函數
type long int

// 有一個參數變成了主語
// tmp叫接收者，接受者就是傳遞的一個參數
func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add01(10, 67) // 普通函數調用方式
	fmt.Println("result = ", result)

	var a long = 2
	// 調用方法的格式是： 變量名.函數(所需參數)
	result2 := a.Add02(100)
	fmt.Println("result2 = ", result2)

	// 面向對象只是換了一種表達方式

}
