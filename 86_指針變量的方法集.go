package main

import "fmt"

type PersonF struct {
	name string
	sex  byte
	age  int
}

// 接收者為普通變量 非指針 值語義 一份拷貝
func (tmp PersonF) setInfoValue() {
	fmt.Println("setInfoValue")
}

// 接收者為指針變量 引用傳遞
func (tmp *PersonF) setInfoValueViaPointer() {
	fmt.Println("setInfoValueViaPointer")
}

func main() {
	/*
		結構體變量是一個指針變量 它能夠調用哪些方法 這些方法就是一個方法集
	*/
	p := &PersonF{"aaa", 'a', 10}
	p.setInfoValueViaPointer()    // 為什麼這個是指針也能調用
	(*p).setInfoValueViaPointer() // 這樣也能調 把 (*p)轉換成了p
	/*
		用实例 value 和 pointer 调用方法(含匿名字段)不受方法集约束
		编译器编总是查找全部方法，并自动转换 receiver 实参。
	*/

	// 內部做了轉換
	// 把指針 p 轉換成 *p  再進行調用了
	(*p).setInfoValue()
	//p.setInfoValue()
}
