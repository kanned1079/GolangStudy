package main

import "fmt"

func main() {
	// 不同類型變量的聲明（定義）

	// 傳統寫法
	//var a int
	//var b float64

	var (
		a int
		b float64
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//const i int = 10
	//const j float64 = 3.14

	//const (
	//	i int = 10
	//	j float64 = 3.14
	//)
	// 或者
	const (
		// 自動推導類型
		i = 10
		j = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
