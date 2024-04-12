package main

import "fmt"

func main() {
	// 這樣不行 所以要用數組
	//id1 := 1
	//id2 := 1
	//id3 := 1

	// 這麼寫好
	var id [50]int // 這個元素個數必須是常量才行
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}
}
