package main

import "fmt"

func main() {
	var a int = 10
	/*
		Go中指針的默認值是nil 不是NULL
		Go不支持指針運算 不支持-> 只支持.
		每一个变量会有两层含义 一个是内存 一個是地址
		內存 => 教室
		&a => 教室外邊的門牌號
		內存用來放東西	a = 10
		教室裡面也是放桌子	教室 = 桌子

		a ==> 內存的內容
		&a ==> 外面的標號 叫地址
	*/
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %p\n", &a)
	fmt.Printf("&a = %v\n", &a)

	// 保存某個變量的地址 需要指針類型 *int 保存 int的地址
	// **int 保存 *int 的地址
	var p *int
	p = &a // 指針變量指向誰 就把誰的地址賦給指針變量
	fmt.Printf("p = %v, &a = %v\n", p, &a)

	*p = 666 // 這個操作的不是p 而是p所指向的內存（就是a）
	fmt.Printf("p = %v, &a = %v\n", p, &a)
	fmt.Printf("p = %v, &a = %v\n", *p, &a)

}
