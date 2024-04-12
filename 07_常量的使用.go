package main

import "fmt"

func main() {
	// 常量，程序運行期間，不可以改變的量，聲明需要使用 const
	const a int = 20
	//a = 30
	// 常量不允許被修改
	fmt.Printf("a = %d\n", a)

	const b = 10
	// 不能使用 :=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}
