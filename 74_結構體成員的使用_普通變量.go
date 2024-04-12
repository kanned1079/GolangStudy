package main

// 定義一個結構體類型
type Student3 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 定義一個結構體普通變量
	var s Student3

	// 操作成員 需要使用.操作符
	s.id = 1
	s.name = "李四"
	s.age = 19
	s.addr = "重慶"
	s.sex = 'M'
}
