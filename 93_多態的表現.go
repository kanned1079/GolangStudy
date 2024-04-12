package main

import "fmt"

// Humaner02 定義接口類型
type Humaner02 interface {
	// 方法 方法只有聲明 沒有實現 由自定義類型實現
	sayHi()
}

type StudentsA2 struct {
	name string
	id   int
}

// Student類型實現了方法
func (tmp *StudentsA2) sayHi() {
	fmt.Printf("Student[%d, %s] say Hi!\n", tmp.id, tmp.name)
}

type TeacherA2 struct {
	addr  string
	group string
}

// Teacher類型實現了方法
func (tmp *TeacherA2) sayHi() {
	fmt.Printf(
		"Teacher[%s, %s] say Hi!\n", tmp.addr, tmp.group)
}

type myStr2 string

// myStr類型實現了方法
func (tmp *myStr2) sayHi() {
	fmt.Printf("myStr[%s] say Hi!\n", *tmp)
}

// 定義一個普通函數 函數參數是接口類型
// 只有一個函數 但是可以有多個表現 這就是多態
func whoSayHi(i Humaner02) {
	i.sayHi()
}

func main() {
	// 調用了同一個函數 不同表現
	s := &StudentsA2{"kanna", 1}
	t := &TeacherA2{"bj", "group1"}
	var str myStr2 = "swd"

	whoSayHi(s)
	whoSayHi(t)
	whoSayHi(&str)
	fmt.Println()

	// 創建一個切片
	x := make([]Humaner02, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		i.sayHi()
	}
}

func main000001() {
	// 定義接口類型的變量
	var i Humaner02

	// 只要實現了此接口方法的類型 那麼這個類型的變量（接收者類型）就可以給i賦值
	s := &StudentsA2{"kanna", 1}
	i = s
	i.sayHi()

	t := &TeacherA2{"bj", "group1"}
	i = t
	i.sayHi()

	var str myStr2 = "swd"
	i = &str
	i.sayHi()

	// 都調用sayHi() 但是是不同的表現
}
