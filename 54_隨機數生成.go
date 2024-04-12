package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. 設置種子 （只需設置一次）
	// 如果種子相同 那麼每次生成的隨機數都相同
	rand.Seed(time.Now().UnixNano())
	// 2. 產生隨機數
	for i := 0; i < 5; i++ {
		//fmt.Println("rand = ", rand.Int()) // 隨機一個很大的數
		fmt.Println("rand = ", rand.Intn(100)) // Intn限定在100以內

	}

}
