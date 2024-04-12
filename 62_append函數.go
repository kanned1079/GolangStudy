package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	for i := 1; i <= 10; i++ {
		// 在末尾插入元素
		s1 = append(s1, i)
	}
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	fmt.Println("--------------------------------")

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	s2 = append(s2, 7)
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

}
