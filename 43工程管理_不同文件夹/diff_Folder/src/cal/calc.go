package cal

import "fmt"

// 导入包的时候会自动调用init函数
func init() {
	fmt.Println("cal的init函数")
}

func Add(a, b int) int {
	return a + b
}
