package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Students struct {
	Person // 這樣可以減少重複性 只有類型 沒有名字
	// 繼承了Person的成員
	id   int
	addr string
}

func main() {
	// 順序初始化
	var stu1 Students = Students{Person{"mike", 'M', 18}, 1, "bj"}
	fmt.Println("stu1 = ", stu1)

	// 自動推導類型
	stu2 := Students{Person{"lily", 'M', 18}, 1, "bj"}
	// %+v 可以讓打印更詳細
	fmt.Printf("stu2 = %+v\n", stu2)

	// 指定成員初始化 沒有初始化的默認值
	stu3 := Students{id: 1}
	fmt.Printf("stu3 = %+v\n", stu3)

	stu4 := Students{Person: Person{name: "kanna"}, id: 10}
	fmt.Printf("stu4 = %+v\n", stu4)
}
