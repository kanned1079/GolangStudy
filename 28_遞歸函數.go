package main

import "fmt"

func main() {
	t1(3)
	fmt.Println("main")
	fmt.Println("sum = ", SumFunc1(100))
}

func t1(a int) {
	if a == 1 { // 函數終止調用的條件 重要
		fmt.Println("a = ", a)
		return
	}
	t1(a - 1)
	fmt.Println("a = ", a)

}

// 使用遞歸實現 1+2+3++++100
func SumFunc1(n int) (res int) {
	i, sum := 1, 0
	for i = 1; i <= n; i++ {
		sum += i
	}
	return sum
}
