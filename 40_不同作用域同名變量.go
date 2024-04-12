package main

import "fmt"

var a2 byte // 全局變量

func main() {
	var a2 int
	// 不同作用域允許定義同名變量
	// 使用變量的原則是 就近原則
	fmt.Println("a2 = ", a2)
	fmt.Printf("a2's type is %T\n", a2)

	{
		var a2 float64
		fmt.Printf("a2's type is %T\n", a2)

	}

	test006()

}

func test006() {
	fmt.Printf("a2's type is %T\n", a2)
}
