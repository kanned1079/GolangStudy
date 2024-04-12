package main

import "fmt"

func main() {
	var flag bool
	flag = true
	//fmt.Printf("flag = %d\n", flag)
	// 需要使用 %t
	fmt.Printf("flag = %t\n", flag)

	// 转换为int
	// bool不能转换为int
	//fmt.Printf("flag = %d\n", int(flag))

	//flag = 1
	//flag = bool(1)
	// 这俩都不能转
	// 不能转换的类型，叫做不兼容类型

	var ch byte
	ch = 'a' // 字符類型本質上就是int
	var t int
	//t = ch //這麼不行
	t = int(ch) // 這樣才行
	fmt.Println("t = %d\n", t)

}
