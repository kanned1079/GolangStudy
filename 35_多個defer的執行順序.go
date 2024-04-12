package main

import "fmt"

/*
	先進後出
	如果一个函数中有多个 defer语句
	它们会以LIFO(后进先出)的顺序执行。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
*/

func main() {
	defer fmt.Println("aaaaa")
	defer fmt.Println("bbbbb")

	// 調用一個函數 導致內存出問題
	defer test003(0)
	// 到這裡程序崩潰了 所以下面就沒有被執行
	//fmt.Println("ccccc")
	defer fmt.Println("ccccc")

}

func test003(x int) {
	res := 100 / x
	fmt.Println("res = ", res)
}
