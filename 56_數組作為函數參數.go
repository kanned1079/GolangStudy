package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify01(a)
	a[4] = 666
	fmt.Println("main: a = ", a)

}

// 數組做函數參數 是值傳遞
// 實參數組的每個元素給形叁數組拷貝一份
// 形叁數組是實叁數組的複製品
func modify01(a [5]int) {
	fmt.Println("modify: a = ", a)

}
