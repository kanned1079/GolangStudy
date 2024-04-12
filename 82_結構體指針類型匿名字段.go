package main

import "fmt"

type PersonD struct {
	name string
	sex  byte
	age  int
}

type StudentsD struct {
	*PersonD // 指針類型
	// 繼承了Person的成員
	id   int
	addr string
}

func main() {
	//stu1 := StudentsA{PersonsA{"kanna", 'M', 18}, 7, "dw"}
	// 以前這麼寫的

	stu1 := StudentsD{&PersonD{"kanna", 'M', 18}, 7, "dw"}
	fmt.Println(stu1.name, stu1.sex, stu1.age, stu1.id, stu1.addr)

	//s1 := StudentsD{ , 18, "bj"}

	// 換一種方法 先定義變量
	var s2 StudentsD
	s2.PersonD = new(PersonD) //分配空間
	s2.name = "yoyo"
	s2.sex = 'F'
	s2.age = 20
	s2.id = 01
	s2.addr = "nn"
	fmt.Println("s2 = ", s2)
}
