package main

import "fmt"

func main() {
	Myfunc1(666, 777)

	Myfunc22()
	println("-----------------")
	Myfunc22(1)
	println("-----------------")
	Myfunc22(1, 2)
	println("-----------------")
	Myfunc22(1, 2, 3)
	println("-----------------")

	test1(1, 2, 3, 4)
}

// 這叫固定參數
func Myfunc1(a int, b int) {

}

// 這叫不定參數
// 不定參數 只能放在形式參數的最後一個參數
func Myfunc22(args ...int) {
	//fmt.Println("元素個數為 len(args) :", len(args))
	// 有什麼用
	//for i := 0; i < len(args); i++ {
	//	fmt.Printf("args[%d] = %d\n", i, args[i])
	//}

	// 使用迭代
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)

	}
}

// 不定參數 只能放在形式參數的最後一個參數
// 固定參數一定要傳遞參數 不定參數根絕需求傳參
func MyFunc3(a int, args ...int) {

}

// 不能這麼寫
//func MyFunc3(args ...int, a int)  {
//
//}

// 想將這個參數傳遞給另外一個函數使用
func test1(args ...int) {
	// 全部參數傳遞給test2()
	test2(args...)
	println("-----------------")
	// 只傳遞最後兩個參數
	test3(args[:2]...) // args[0] ～ args[2] 但是不包括數字args[2]，傳遞過去

	test3(args[2:]...) // 從args[2]開始（包括本身），把後面所有元素傳遞過去
}

func test2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("tmp = ", data)
	}
}

func test3(arg ...int) {
	for _, data := range arg {
		fmt.Println("tmp = ", data)
	}
}
