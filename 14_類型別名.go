package main

import "fmt"

// 類型別名 就是給一個類型起一個名稱而已
type DataType int64

func main() {

	type bigint int64 // 給int64起了個別名叫bigint
	var a bigint      // 等價於 var a int64
	fmt.Printf("type is %T\n", a)

	// 也可以這麼寫
	type (
		long int64
		char byte
	)
	var b long
	var c char
	b = 100
	c = 'C'
	fmt.Printf("type is %T\n", b)
	fmt.Printf("type is %T\n", c)

}
