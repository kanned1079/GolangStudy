package main

import "fmt"

func main() {
	var p *int
	p = nil
	fmt.Printf("p = %v\n", p)

	//*p = 666 // 這樣不行err 因為p沒有合法的指向
	var a int
	p = &a
	*p = 666 //這樣才行
	fmt.Printf("p = %v\n", p)
}
