package main

import (
	"fmt"
	_ "fmt"
	_ "math/rand"
	"os"
)

func main() {
	createFile("temp.txt")
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	data := []byte("AAAAA")
	file.Write(data)
	file.Close()
	fmt.Println("创建了一个名为" + file.Name() + "的文件")
}

func getX(value string) int32 {
	fmt.Println("functio 1")
	return 0
}

//func findInKey(str string, km map[string]struct{}) bool {
//	for p := range km {
//		if wildcard.Match(p, str) {
//			return true
//		}
//	}
//
//	return false
//}
