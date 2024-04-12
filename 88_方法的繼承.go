package main

import "fmt"

type PersonH struct {
	name string
	sex  byte
	age  int
}

// Person類型 實現了一個方法
func (tmp *PersonH) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

// 有個學生繼承Person字段 成員和方法都會被繼承
type StuStudent struct {
	PersonH // 匿名字段
	id      int
	addr    string
}

func main() {
	s := StuStudent{PersonH{"mike", 'm', 18}, 1, "bj"}
	s.PrintInfo() // 繼承了Person的PrintInfo()方法
}
