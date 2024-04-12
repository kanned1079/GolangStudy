package main

import "fmt"

type StudentsA3 struct {
	name string
	id   int
}

// Humaner03 定義接口類型
type Humaner03 interface { // 子集
	// 方法 方法只有聲明 沒有實現 由自定義類型實現
	sayHi03()
}

type Person03 interface { // 超集
	Humaner03 // 繼承字段 繼承sayHi()
	sing(lrc string)
}

// Student類型實現了方法
func (tmp *StudentsA3) sayHi03() {
	fmt.Printf("Student[%d, %s] say Hi!\n", tmp.id, tmp.name)
}

func (tmp *StudentsA3) sing(lrc string) {
	fmt.Printf("Student is sing %s\n", lrc)
}

func main() {
	// 定義一個接口類型的變量
	var i1 Humaner03
	var i2 Person03

	stu := StudentsA3{"kanna", 3}
	i1 = &stu
	i1.sayHi03() // 這個是繼承過來的方法

	i2 = &stu
	i2.sing("songA3") //

}
