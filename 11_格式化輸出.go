package main

import "fmt"

func main() {

	a := 10
	b := 20
	c := 'a'
	d := 3.14
	fmt.Printf("%T %T %T %T\n", a, b, c, d)

	/*
		%d 整形
		%s 字符串
		%c 字符
		%f 浮點
	*/

	fmt.Printf("a = %d, b = %d, c = %c, d = %f\n", a, b, c, d)
	// 也可以使用 %v 自動匹配格式
	// 但是字符類型不行
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)

}
