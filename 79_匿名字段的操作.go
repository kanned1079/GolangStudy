package main

import "fmt"

//import "lei"

type PersonsA struct {
	name string
	sex  byte
	age  int
}

type StudentsA struct {
	PersonsA // 這樣可以減少重複性 只有類型 沒有名字
	// 繼承了Person的成員
	id   int
	addr string
}

func main() {
	stu1 := StudentsA{PersonsA{"kanna", 'M', 18}, 7, "dw"}
	stu1.name = "kinggyo"
	stu1.sex = 'F'
	stu1.age = 16
	stu1.addr = "CZ"
	stu1.id = 22
	fmt.Printf("stu1.name = %s\n", stu1.name)
	fmt.Printf("stu1 = %+v\n", stu1)

	// 以一個整體來賦值
	stu1.PersonsA = PersonsA{"kanna", 'M', 18}
	fmt.Printf("stu1 = %+v\n", stu1)

}
