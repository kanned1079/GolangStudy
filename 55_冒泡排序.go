package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nums [30]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("nums = ", nums)

	// 冒泡排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]

			}
		}
	}
	fmt.Println("nums = ", nums)

}
