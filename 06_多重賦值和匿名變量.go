package main

import "fmt"

// go函數可以返回多個返回值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	a, b := 10, 20

	//交換兩個變量的值
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a = %d, b = %d\n", a, b)

	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	//導致的問題
	i = 10
	j = 20
	// _ 匿名變量， 丟棄數據不處理，匿名變量配合函數返回值使用才有優勢
	temp, _ = i, j
	fmt.Println(temp)

	var c, d, e int
	c, d, e = test()
	fmt.Println(c, d, e)

	_, d, _ = test()
	fmt.Println(d)

	_, d, e = test()
	fmt.Println(d, e)
}
