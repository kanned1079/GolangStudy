package main

import "fmt"

func main() {
	//賦值前必須先聲明變量
	var a int
	a = 10
	fmt.Println(a)

	b := 20
	// := 先聲明變量b，再給b賦值
	fmt.Println(b)
}
