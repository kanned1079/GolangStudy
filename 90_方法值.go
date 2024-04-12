package main

import "fmt"

type PersonK struct {
	name string
	sex  byte
	age  int
}

func (tmp PersonK) setInfoValue02() {
	// 這裏是值傳遞 所以獲取的地址會不相同 由於是值 後面需要加上&取地址
	fmt.Printf("setInfoValue02 : %p, %v\n", &tmp, tmp)
}

func (tmp *PersonK) setInfoValueViaPointer02() {
	// 這裏是引用傳遞 由於本身就是指針函數的tmp 則後邊的不需要加上取地址&符號
	fmt.Printf("setInfoValueViaPointer02 : %p, %v\n", tmp, tmp)
}

func main() {
	p1 := PersonK{"mike", 'f', 18}
	fmt.Printf("main's p1 = %p, val = %v\n", &p1, p1)

	// 以前是這樣的調用的
	p1.setInfoValueViaPointer02() // 引用傳遞
	p1.setInfoValue02()           // 值傳遞

	// 方法值的格式 f = p1.Function 隱藏了接收者
	// 想要保存這個函數入口地址
	pFunc := p1.setInfoValueViaPointer02 // 這個就是方法值 調用函數時 無須再傳遞接收者 隱藏了接收者
	pFunc()                              // 等價於 p1.setInfoValueViaPointer02()
	// 以後要或者這個函數的值 只需要調用pFunc()即可

	vFunc := p1.setInfoValue02 // 注意這裏不要 ()
	vFunc()                    // 也可以保存另一個函數的值裡了

}
