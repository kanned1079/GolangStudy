package main

import (
	"fmt"
	"os"
)

func showArgs() {
	fmt.Println("Args = ", os.Args)
	list := os.Args
	println("Args nums: ", len(list))
}
