package main

import "fmt"

func main() {
	//a, b, c := myFunc03()
	a, b, c := myFunc04()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

// 多個返回值
func myFunc03() (int, int, int) {
	return 1, 2, 3
}

// GO推薦寫法
func myFunc04() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
