package main

import "fmt"

type StudentsA4 struct {
	name string
	id   int
}

// Humaner04 定義接口類型
type Humaner04 interface { // 子集
	// 方法 方法只有聲明 沒有實現 由自定義類型實現
	sayHi04()
}

// Human <- Person
type Person04 interface { // 超集
	Humaner04 // 繼承字段 繼承sayHi()
	sing(lrc string)
}

// Student類型實現了方法 這裏的接收者類型是自定義的結構體 也就是只有此結構體和其子類才能調用這個方法
func (tmp *StudentsA4) sayHi04() {
	fmt.Printf("Student[%d, %s] say Hi!\n", tmp.id, tmp.name)
}

func (tmp *StudentsA4) sing(lrc string) {
	fmt.Printf("Student is sing %s\n", lrc)
}

func main() {
	/*
		超集可以轉換為子集 但是反過來不可以
	*/

	// 定義一個超集接口類型
	var iPro Person04 // 超集
	iPro = &StudentsA4{"kinggyo", 04}
	var i Humaner04 // 子集

	//iPro = i // 這樣不行
	i = iPro // 這個可以
	i.sayHi04()

}
