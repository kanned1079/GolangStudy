package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike2", 2: "c++", 3: "go"}
	fmt.Println("m1 = ", m1)

	delete(m1, 1) // 刪除key為1的內容
	fmt.Println("m1 = ", m1)

}
