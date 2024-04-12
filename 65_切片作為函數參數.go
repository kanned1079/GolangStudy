package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 切片是引用傳遞
func main() {
	var n int = 10
	s := make([]int, n)
	InitData01(s) // 初始化數據 隨機數
	fmt.Println("排序前的s = ", s)
	BubbleSort(s) // 冒泡排序
	fmt.Println("排序後的s = ", s)

}

// 這是引用傳遞
func InitData01(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
