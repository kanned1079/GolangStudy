package main

import "fmt"

// 定義一個結構體類型
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 順序初始化 每個成員都必須初始化
	var s1 Student = Student{1, "張三", 'M', 18, "上海宛平南路"}
	fmt.Println("s1 = ", s1)

	// 指定成員初始化 沒有初始化的成員自動賦值
	s2 := Student{name: "mike", addr: "北京"}
	fmt.Println("s2 = ", s2)
}
