package main

import "fmt"

func main() {
	// 支持比較 只支持 == 和 != 比較的是每個元素都一樣
	//兩個數組進行比較 數組類型都要相同
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	fmt.Println("a == b : ", a == b)
	fmt.Println("b == c : ", b == c)

	// 同類型的數組可以賦值
	var d [5]int
	d = b
	fmt.Println("d = ", d)

}
