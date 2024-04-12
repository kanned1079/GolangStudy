package main

import "fmt"

// 定義接口類型
type Humaner interface {
	// 方法 方法只有聲明 沒有實現 由自定義類型實現
	sayHi()
}

type StudentsA0 struct {
	name string
	id   int
}

// Student類型實現了方法
func (tmp *StudentsA0) sayHi() {
	fmt.Printf("Student[%d, %s] say Hi!\n", tmp.id, tmp.name)
}

type TeacherA0 struct {
	addr  string
	group string
}

// Teacher類型實現了方法
func (tmp *TeacherA0) sayHi() {
	fmt.Printf("Teacher[%s, %s] say Hi!\n", tmp.addr, tmp.group)
}

type myStr string

// myStr類型實現了方法
func (tmp *myStr) sayHi() {
	fmt.Printf("myStr[%s] say Hi!\n", *tmp)
}

func main() {
	// 定義接口類型的變量
	var i Humaner

	// 只要實現了此接口方法的類型 那麼這個類型的變量（接收者類型）就可以給i賦值
	s := &StudentsA0{"kanna", 1}
	i = s
	i.sayHi()

	t := &TeacherA0{"bj", "group1"}
	i = t
	i.sayHi()

	var str myStr = "swd"
	i = &str
	i.sayHi()

	// 都調用sayHi() 但是是不同的表現
}
