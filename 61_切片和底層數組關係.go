package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("a = ", a)
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	s1 := a[2:5] // 從a[2]開始取三個元素
	s1[1] = 666
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	fmt.Println("a = ", a)
}
