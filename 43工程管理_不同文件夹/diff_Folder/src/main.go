package main

//import (
//	"diff_Folder/cal"
//	"fmt"
//)

import (
	"diff_Folder/cal"
	_ "diff_Folder/cal" // _的使用是该包 而不直接使用包裡的函數 而是調用裡包裡的init函數
	"fmt"
)

// 调用不同包里的函数
// 格式为 包名.函数名()
//调用别的包的函数，这个包的函数名如果是小写，那么无法被调用 必须首字母大写

func main() {
	fmt.Println("main")
	//showArgs()
	fmt.Println("add = ", cal.Add(10, 20))
	//cal.Add(10, 20)
}
