package main

import "fmt"

// 定義一個結構體類型
type Student6 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var s1 Student6 = Student6{1, "張三", 'M', 18, "上海宛平南路"}

	test009(&s1) // 值傳遞 形叁沒法改變實叁
	fmt.Println("s = ", s1)
}

func main00001() {
	var s1 Student6 = Student6{1, "張三", 'M', 18, "上海宛平南路"}

	test008(s1) // 值傳遞 形叁沒法改變實叁
	fmt.Println("s = ", s1)

}

func test008(s Student6) {
	s.id = 666
	fmt.Println("s = ", s)
}

// 地址傳遞（引用傳遞）
func test009(s *Student6) {
	s.id = 777 // 這樣是操作的指針所指向的地址
}
