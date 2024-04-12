package main

import "fmt"

func main() {
	// 做函數參數是引用傳遞
	m1 := map[int]string{1: "mike2", 2: "c++", 3: "go"}
	fmt.Println("m1 = ", m1)

	test007(m1) // 在函數中刪除某個key
	fmt.Println("m1 = ", m1)

}

func test007(m map[int]string) {
	delete(m, 1)
}
