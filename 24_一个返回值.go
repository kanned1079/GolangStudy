package main

import "fmt"

func main() {
	a := myFunc01()
	fmt.Println("a = ", a)

	c := myFunc02()
	fmt.Println("c = ", c)
}

// 無參有返回值 只有一個返回值
// 有返回值的函數需要有return返回結果
func myFunc01() int {
	return 666
}

// 給返回值起了一個變量名
// GO的推薦寫法
func myFunc02() (result int) {
	//return 777 // 可以這麼寫
	result = 777
	return // 也可以這麼寫 用的最多
}
