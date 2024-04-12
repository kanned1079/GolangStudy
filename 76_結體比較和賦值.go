package main

import "fmt"

// 定義一個結構體類型
type Student5 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var s1 Student5 = Student5{1, "張三", 'M', 18, "上海宛平南路"}
	var s2 Student5 = Student5{1, "張三", 'M', 18, "上海宛平南路"}
	fmt.Println("s1 == s2 : ", s1 == s2) // 結構體比較只有 == 和 !=

	var s3 Student5
	// 同類型的結果體可以相互賦值
	s3 = s2
	fmt.Println("s3 = ", s3)

}
