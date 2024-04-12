package main

import "fmt"

func main() {

	a := 10
	str := "mike"

	// 在這裏改了後會同步到外邊
	func() {
		// 閉包以引用方式捕獲外部變量
		a = 666
		str = "go"
		fmt.Printf("內部： a = %d, str = %s\n", a, str)
	}() // 調用
	fmt.Printf("外部： a = %d, str = %s\n", a, str)

}
