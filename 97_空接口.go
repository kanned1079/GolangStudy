package main

import "fmt"

// 空接口其實就是個萬能類型

func xxx(args ...interface{}) {

}

func main() {
	var i interface{} = 16
	i = "abc"
	fmt.Println("interface i = ", i)
}
