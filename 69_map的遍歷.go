package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike2", 2: "c++", 3: "go"}
	for key, val := range m1 {
		// 第一個返回值是key 四二個返回是value
		// 遍歷的結果是無序的
		fmt.Printf("map[%d]%s\n", key, val)
	}

	// 怎麼判斷一個key是否存在
	// 第一個返回值為key所對應的value 第二個返回值是key是否存在的條件
	// 存在ok為true
	value, ok := m1[1]
	if ok == true {
		fmt.Println("存在m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}

}
