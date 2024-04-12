package main

import "fmt"

func main() {
	a, b := 10, 20
	swap002(&a, &b)
	fmt.Println("outer: a = ", a, ", b = ", b)

}

func swap002(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Println("inner: a = ", *p1, ", b = ", *p2)
}
