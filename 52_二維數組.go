package main

import "fmt"

func main() {

	// 有多少個[]就是多少緯
	// 有多少個[]就用多少個循環
	var a [3][4]int
	k := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d\t", i, j, k)
		}
		fmt.Println()
	}

	fmt.Println("a = ", a)

	// 部分初始化 没有初始化的默认为0
	b := [3][4]int{{1, 2, 3}, {5, 7, 8}, {9, 11, 12}}
	fmt.Println("b = ", b)

	// 只要前兩行
	c := [3][4]int{{1, 2, 3}, {5, 7, 8, 9}}
	fmt.Println("c = ", c)

	// 只要第兩行
	d := [3][4]int{1: {5, 7, 8, 9}}
	fmt.Println("d = ", d)

}
