package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randNum int
	// 產生一個四位的隨機數
	CreateNum(&randNum)
	//fmt.Println("randNum = ", randNum)

	randSlice := make([]int, 4)
	// 保存四位數的每一位
	// 1234/1000 = 1

	GetNums(randSlice, randNum) // 求得每個位數
	fmt.Println("randSlice = ", randSlice)

	OnGame(randSlice) // 開始猜數
}

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	//num := rand.Intn(10000)
	// 使用循環
	//for {
	//	num = rand.Intn(10000)
	//	if num >= 1000 {
	//		break
	//	}
	//}
	// 這樣也行
	num = rand.Intn(9000) + 1000
	//fmt.Println("num = ", num)
	*p = num
	return
}

func GetNums(s []int, num int) {
	//千位
	s[0] = num / 1000
	s[1] = (num % 1000) / 100
	s[2] = (num % 100) / 10
	s[3] = num % 10
}

func OnGame(randSlice []int) {
	var num int
	for {
		for {
			fmt.Printf("输入数字：")
			fmt.Scan(&num)
			if num > 1000 && num < 10000 {
				break
			}
			fmt.Println("不符合要求")
		}
		//fmt.Println("randSlice = ", randSlice)
		var keySlice []int
		keySlice = make([]int, 4)
		GetNums(keySlice, num)
		//fmt.Println("keySlice = ", keySlice)
		var ok int = 0

		for n := 0; n < 4; n++ {
			if keySlice[n] < randSlice[n] {
				fmt.Printf("第%d位小了\n", n+1)
			} else if keySlice[n] > randSlice[n] {
				fmt.Printf("第%d位大了\n", n+1)
			} else {
				fmt.Printf("第%d猜对了\n", n+1)
				ok++
			}
		}
		if ok == 4 {
			fmt.Println("猜对啦猜对啦")
			break
		}
	}

}
