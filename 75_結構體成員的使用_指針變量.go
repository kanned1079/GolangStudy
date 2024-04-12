package main

import "fmt"

// 定義一個結構體類型
type Student4 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 一定要注意
	// 指針有合法指向後 才能操作成員
	// 限定一個一個普通結構體變量
	var s Student4
	// 再定義一個指針變量 保存s的地址
	var p1 *Student4
	p1 = &s

	// 通過指針操作成員 p1.id和 (*p1).id完全等價 只能使用.運算符
	p1.id = 4
	(*p1).name = "kanna" // 不管是不是指針都能這麼寫

	fmt.Println("p1 = ", p1)

	// 2. 通過new方法申請一個結構體
	p2 := new(Student4)
	p2.id = 144
	fmt.Println("p2 = ", p2)

}
