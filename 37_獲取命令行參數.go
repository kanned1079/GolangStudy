package main

import (
	"fmt"
	"os"
)

func main() {
	// 接收用戶傳遞的參數 都是以字符串傳遞的
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n)

	// 執行 ./exec aaa bb c
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	for i2, data := range list {
		fmt.Printf("list[%d] = %s\n", i2, data)

	}
}
