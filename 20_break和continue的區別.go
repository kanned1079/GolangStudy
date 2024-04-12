package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for { // for後面不寫任何參數，那麼這就是個死循環
		i++
		time.Sleep(time.Second)

		if i == 5 {
			break
			//continue
		}
		fmt.Println("i = ", i)
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

}
