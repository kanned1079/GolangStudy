package main

import "fmt"

// 定義一個結構體類型
type Student2 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 別忘了取地址
	var p1 *Student2 = &Student2{1, "張三", 'M', 18, "上海宛平南路"}
	fmt.Println("*p1 = ", *p1)

	p2 := &Student2{name: "mike", addr: "北京"}
	fmt.Printf("p2's type is %T\n", p2)
	fmt.Println("p2 = ", p2) // 這麼也行

}
