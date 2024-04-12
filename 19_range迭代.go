package main

import "fmt"

func main() {
	str1 := "abvc"
	// 通過for循環打印每一個字符

	// 傳統寫法
	//for i := 0; i < len(str1); i++ {
	//	fmt.Printf("str[%d] = %c\n", i, str1[i])
	//}

	// 使用range迭代打印每一個元素，默認返回2個值
	// 一個是元素的位置，一個是元素本身
	for i, data := range str1 {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	fmt.Println()

	for i := range str1 { // 第二個返回值默認丟棄，返回元素的位置（下標）
		fmt.Printf("str[%d] = %c\n", i, str1[i])
	}

	fmt.Println()

	// 上一個等價於
	for i, _ := range str1 { // 第二個返回值默認丟棄，返回元素的位置（下標）
		fmt.Printf("str[%d] = %c\n", i, str1[i])
	}

}
