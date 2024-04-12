package main

import "fmt"

func main() {
	a, b := 10, 20
	swap001(a, b)
	fmt.Println("outer: a = ", a, ", b = ", b)

}

func swap001(a, b int) {
	a, b = b, a
	fmt.Println("inner: a = ", a, ", b = ", b)
}
