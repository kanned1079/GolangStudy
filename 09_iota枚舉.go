package main

import "fmt"

func main() {

	// iota	是一個常量自動生成器，每行一個，自動累加1
	// iota給常量賦值使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// iota遇到一個const，會重置為0
	const d = iota
	fmt.Println(d)

	const (
		// 可以只些一個iota
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	// 如果是同一行，值都相同
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j1, j2, j3, k)
}
