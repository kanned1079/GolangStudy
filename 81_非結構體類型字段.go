package main

import "fmt"

type myster string // 自定義類型 給一個變量改名

type PersonsC struct {
	name string
	sex  byte
	age  int
}

type StudentsC struct {
	PersonsC // 這樣可以減少重複性 只有類型 沒有名字
	// 繼承了Person的成員
	int    // 基礎類型的匿名字段
	myster // 可這麼寫
}

func main() {
	stu1 := StudentsC{PersonsC{"mike", 'm', 18}, 666, "he"}
	fmt.Printf("stu1 = %+v\n", stu1)

	// 那麼該怎麼操作
	fmt.Println(stu1.name, stu1.age, stu1.int, stu1.myster) // 直接寫int就行

}
