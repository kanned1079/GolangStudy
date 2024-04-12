package main

import "fmt"

func main() {
	//var num int64
	//fmt.Print("Input: ")
	//fmt.Scan(&num)
	//switch num { // 寫的是變量本身
	//case 1:
	//	fmt.Println("num = 1")
	//	//break // GO語言保留了break關鍵字，不寫就已經是包含
	//	fallthrough // 不跳出switch 後面無條件執行
	//case 2:
	//	fmt.Println("num = 2")
	//	fallthrough
	//	//break
	//case 3:
	//	fmt.Println("num = 3")
	//	//break
	//case 4:
	//	fmt.Println("num = 4")
	//	//break
	//default:
	//	fmt.Println("ERROR")
	//	//break
	//}

	//num2 := 1
	// 也支持一個初始化語句，初始化語句和變量本身，以分號分割
	switch num2 := 1; num2 { // 寫的是變量本身
	case 1:
		fmt.Println("num = 1")
		break // GO語言保留了break關鍵字，不寫就已經是包含
		//fallthrough // 不跳出switch 後面無條件執行
	case 2:
		fmt.Println("num = 2")
		//fallthrough
		//break
	case 3:
		fmt.Println("num = 3")
		//break
	case 4, 5, 6: // 這裏可以寫多個
		fmt.Println("num = 4")
		//break
	default:
		fmt.Println("ERROR")
		//break
	}

	score := 85
	switch { // 可以沒有條件
	case score > 90: // case後面可以放條件
		fmt.Println("優秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其他")

	}

}
