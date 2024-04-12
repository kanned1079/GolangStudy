package main

import "fmt"

type PersonsB struct {
	name string
	sex  byte
	age  int
}

type StudentsB struct {
	PersonsB // 這樣可以減少重複性 只有類型 沒有名字
	// 繼承了Person的成員
	id   int
	addr string
	name string // 這裏和Persons同名了
}

func main() {
	var stu1 StudentsB
	// 	默認規則：如果能在本作用域找到該成員 那麼就操作它
	//			如果找不到， 那就找父類的
	stu1.name = "mike" // 這裏操作的是Students裡的還是 Persons裡的
	stu1.age = 18
	stu1.addr = "bg"
	stu1.sex = 'M'

	fmt.Printf("stu1 = %+v\n", stu1)

	// 顯示調用
	stu1.PersonsB.name = "lily"
	fmt.Printf("stu1 = %+v\n", stu1)

}
