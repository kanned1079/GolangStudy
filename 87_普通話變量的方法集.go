package main

import "fmt"

type PersonG struct {
	name string
	sex  byte
	age  int
}

// 接收者為普通變量 非指針 值語義 一份拷貝
func (tmp PersonG) setInfoValue01() {
	fmt.Println("setInfoValue")
}

// 接收者為指針變量 引用傳遞
func (tmp *PersonG) setInfoValueViaPointer01() {
	fmt.Println("setInfoValueViaPointer")
}

func main() {
	p := PersonG{"mike", 'm', 19}
	p.setInfoValueViaPointer01()
	// 內部先把p 轉換為&p再調用 => (*p).setInfoValueViaPointer01()
	p.setInfoValue01()

}
