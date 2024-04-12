package main

import "fmt"

func main() {
	var a int
	fmt.Print("请输入：")
	// 格式1
	//fmt.Scanf("%d", &a)
	// 格式2 简洁一点
	fmt.Scan(&a)
	fmt.Println("a = ", a)

}
