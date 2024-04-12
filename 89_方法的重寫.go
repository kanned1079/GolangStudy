package main

import "fmt"

type PersonI struct {
	name string
	sex  byte
	age  int
}

// Person類型 實現了一個方法
func (tmp *PersonI) PrintInfo() {
	fmt.Printf("Person-> name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

// 有個學生繼承Person字段 成員和方法都會被繼承
type StuStudent01 struct {
	PersonI // 匿名字段
	id      int
	addr    string
}

// Student也實現了一個方法 這個方法和Person同名 這個叫方法的重寫
func (tmp *StuStudent01) PrintInfo() {
	fmt.Printf("Student-> name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

func main() {
	s := StuStudent01{PersonI{"mike", 'm', 18}, 1, "bj"}
	//s.PrintInfo() // 繼承了Person的PrintInfo()方法
	s.PrintInfo() // 這裏調用的是Person的還是Student的方法
	// 調用的規則是 就近原則 先找本作用域的方法 找不到再去找父類的方法
	// 所以這裏調用的是Student的方法

	// 當然也可以用 顯示調用 父類Person的方法
	s.PersonI.PrintInfo()

}
