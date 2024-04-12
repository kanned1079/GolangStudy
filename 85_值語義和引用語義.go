package main

import "fmt"

type PersonEF struct {
	name string
	sex  byte
	age  int
}

// 接收者為普通變量 非指針 值語義 一份拷貝
func (tmp PersonEF) setInfoValue(n string, s byte, a int) {
	// 值傳遞改不了
	tmp.name = n
	tmp.sex = s
	tmp.age = a
}

// 接收者為指針變量 引用傳遞
func (tmp *PersonEF) setInfoValueViaPointer(n string, s byte, a int) {
	// 引用傳遞是可以改變值的
	tmp.name = n
	tmp.sex = s
	tmp.age = a
}

func main() {
	var per1 PersonEF
	per1.setInfoValue("kinggyo", 'F', 14)
	fmt.Println("per1 = ", per1)
	per1.setInfoValueViaPointer("kinggyo", 'F', 14)
	fmt.Println("per1 = ", per1)

}
