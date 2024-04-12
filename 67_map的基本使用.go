package main

import "fmt"

/*
	map(映射、字典)是一种内置的数据结构
	它是一个无序的 key-value 对的集合
	比如以身份证号作为唯一键来标识一个人的信息

	map的格式為: map[keyType]valueType
	定義: 	info := map[int]string{
				110: "mike",
				111: "yoyo"	}

	在一个 map 里所有的键都是唯一的 而且必须是支持==和!=操作符的类型
	切片、函数以及包含切片的结构类型这些类型由于具有引用语义，不能作为映射的键，使用这些类型会造成编译错误

*/

func main() {
	// 定義一個變量 類型為map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	fmt.Println("m1's len = ", len(m1)) //對於map 只有len，沒有cpa

	//可以通過make創建
	m2 := make(map[int]string, 10) // 雖然會自動擴充 但是在一開始分配可以提高性能
	fmt.Println("m3 = ", m2)
	fmt.Println("m3's len = ", len(m2))
	m2[1] = "mike"
	m2[2] = "yoyo"
	m2[3] = "go" // 會自動擴充的

	fmt.Println("m3's len = ", len(m2))
	fmt.Println("m3 = ", m2)

	// 初始化
	// 鍵值是唯一的
	m3 := map[int]string{1: "mike2", 2: "c++", 3: "go"}
	//m3 := map[int]string{1: "mike2", 2: "c++", 3: "go", 3: "xxx"} // 這樣不行 因為鍵值唯一
	fmt.Println("m3 = ", m3)

}
