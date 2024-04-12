package main

import "fmt"

func main() {
	/*
		for 初始條件； 判斷條件； 條件變化 {
			語句
		}
	*/

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

}
