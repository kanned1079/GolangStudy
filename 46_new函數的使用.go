package main

import "fmt"

func main() {
	//a := 10 // 整形變量a

	var p *int
	//p = &a // 指向了一個合法內存

	// p是*int 指向int類型
	p = new(int)
	*p = 777
	fmt.Printf("*p = %v\n", *p)

	// Go語言只管申請空間 不用管釋放和垃圾回收

	// 也可以這麼個寫
	q := new(int)
	*q = 888
	fmt.Printf("*q = %v\n", *q)

}
