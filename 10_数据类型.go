package main

import (
	"fmt"
)

func main() {

	// 沒有賦初始值
	var a0 bool
	fmt.Println("a0 = ", a0)

	// 聲明變量
	var a bool
	a = true
	fmt.Println("a = ", a)

	// 自動推導
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)

	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	f2 := 3.1415
	fmt.Printf("f2 = %f, type is %T\n", f2, f2)
	// float64存儲小數的準確度比32高

	// 沒有存儲一個字符的char
	// 內存中存儲字符 ASCII碼 a => 97  A => 65
	var ch1 byte
	ch1 = 97
	//fmt.Println("ch1 = ", ch1)
	// 這個不行
	fmt.Printf("ch1 = %d %c\n", ch1, ch1)

	ch1 = 'A' // 字符使用單引號
	fmt.Printf("ch1 = %d %c\n", ch1, ch1)

	// 大小寫互相轉換
	ch1 = 'A'
	fmt.Printf("%c -> %c\n", ch1, ch1+32)
	ch1 = 'a'
	fmt.Printf("%c -> %c\n", ch1, ch1-32)

	// 轉義字符
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello2\n")

	// 字符串類型
	var str1 string
	str1 = "abcde"
	fmt.Println("str1 = ", str1)

	//自動推導
	str2 := "ijk"
	fmt.Printf("str2 = %s, type is %T\n", str2, str2)
	//獲取字符串長度
	fmt.Println("字符串str2的長度：", len(str2))

	var a1 string = "a"
	// 字符串最後一位隱藏掉了 \0 結束字符
	fmt.Printf(a1)

	// 只想操作字符串的某一個字符
	str3 := "Hello World"
	fmt.Printf("str3[0] = %c\n", str3[0])

	// 複數類型
	var t complex128
	t = 2.1 + 3.14i
	//  實屬   虛數
	fmt.Println("t = ", t)

	t2 := 3.3 + 4.4i
	fmt.Printf("t1 type is %T\n", t2)

	// 通過內建函數，取實部和虛部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))

}

/*
	* bool true或false 不可以用數字表示
	* byte 字節型
	  rune 可放置中文字符 => uint32
	* int, uint 整形 32位或64位 u代表無符號
	  int8, uint8 整形 8位
	  int16, uint16
	  int32, uint32
	  int64, uint64
	* float32 浮點型 小數精確到7位
	* float64 浮點型 消暑精確到15位
	  complex64 複數類型 8
	  complex128 複數類型 16
	  uintptr 整形 4/8 足以存儲指針的uint32或uint64整數
	* string 字符串 utf-8

	聲明變量的初始值要麼是0或false
*/
