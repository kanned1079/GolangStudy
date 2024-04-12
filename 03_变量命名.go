package main

import "fmt" //導入了包必須使用

func main() {
	//聲名變量格式
	var a int
	//聲名了變量必須要使用
	//聲明沒有初始化的變量，默認值是0
	//同一個 {} 中聲明的變量名是唯一的

	a = 10 //先聲明再賦值
	fmt.Println(a)

	//var b, c int
	//可以同時聲明多個變量
	var b int = 20 //聲明變量並初始化
	fmt.Println(b)

	//可以自動推倒類型，但是必須初始化
	c := 30
	fmt.Println(c)
	fmt.Println("c type is %T", c)

}

/*
	字母，數字，下劃線
	不能數字開頭
	不能與關鍵字相同
	區分大小寫
*/
