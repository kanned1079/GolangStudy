package main

// 傳統寫法
//import "fmt"
//import "os"

// 不建议这么写
//import . "fmt"

// 建議
//import (
//	"fmt"
//	"os"
//)

// 给包名起别名

import (
	io "fmt"
	"os"
	_ "sort" // 這樣可以忽略此包
)

func main() {
	io.Println("Args = ", os.Args)
}
