package main

import "fmt"

func main() {
	// defer延遲調用 main函數結束前
	defer fmt.Println("bbbbbbbbbbb")

	fmt.Println("aaaaaaaaaaa")

}
