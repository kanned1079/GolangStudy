package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 0}
	s := a[0:3:5] // [開始:終點（不包含）:容量]
	/*
		[low:high:max]
		len = high - low
		cap = max - low
	*/
	fmt.Println("s = ", s)
	fmt.Println("s len = ", len(s))
	fmt.Println("s cap = ", cap(s))

	s = a[1:4:5] // [開始:終點（不包含）:容量]
	fmt.Println("s = ", s)
	fmt.Println("s len = ", len(s))
	fmt.Println("s cap = ", cap(s))

}
