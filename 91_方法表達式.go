package main

import "fmt"

type PersonA1 struct {
	name string
	sex  byte
	age  int
}

func (tmp PersonA1) setInfoValue02() {
	// 這裏是值傳遞 所以獲取的地址會不相同 由於是值 後面需要加上&取地址
	fmt.Printf("setInfoValue02 : %p, %v\n", &tmp, tmp)
}

func (tmp *PersonA1) setInfoValueViaPointer02() {
	// 這裏是引用傳遞 由於本身就是指針函數的tmp 則後邊的不需要加上取地址&符號
	fmt.Printf("setInfoValueViaPointer02 : %p, %v\n", tmp, tmp)
}

func main() {
	p1 := PersonA1{"mike", 'f', 18}
	fmt.Printf("main's p1 = %p, val = %v\n", &p1, p1)

	/*
		和方法值的比較
			方法值的格式 f = p1.Function 隱藏了接收者
			方法表達式
	*/
	f := (*PersonA1).setInfoValueViaPointer02
	f(&p1) // 要顯式把接收者傳遞過去 相當於p1.setInfoValueViaPointer02()

	f2 := (PersonA1).setInfoValue02
	f2(p1)

	// 以前是這樣的調用的
	//p1.setInfoValueViaPointer02() // 引用傳遞
	//p1.setInfoValue02()           // 值傳遞

	// 想要保存這個函數入口地址
	//pFunc := p1.setInfoValueViaPointer02 // 這個就是方法值 調用函數時 無須再傳遞接收者 隱藏了接收者
	//pFunc()                              // 等價於 p1.setInfoValueViaPointer02()
	// 以後要或者這個函數的值 只需要調用pFunc()即可

	//vFunc := p1.setInfoValue02 // 注意這裏不要 ()
	//vFunc()                    // 也可以保存另一個函數的值裡了
}
