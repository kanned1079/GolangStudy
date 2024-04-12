package main

import "fmt"

type PersonE struct {
	name string
	sex  byte
	age  int
}

/*
	方法总是绑定对象实例，并隐式将实例作为第一实参(receiver)，方法的语法如下:
	func (receiver ReceiverType) funcName (parameters)(results)

	1. 参数 receiver 可任意命名。如方法中未曾使用，可省略参数名。
	2. 參数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
	3. 不支持重载方法，也就是说，不能定义名字相同但是不同参数的方法。

*/

// 帶有接收者的函數叫方法
func (tmp PersonE) PrintInfo01() {
	fmt.Println("tmp = ", tmp)
}

// 通過一個函數給成員賦值
func (p *PersonE) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

//type long int

// 這裏的接收者是long
func (tmp long) test0001() {

}

type char byte

// 這裏的接收者是char
// 只要接收者的類型不同 即使這兩個方法同名 那也是不同方法 不會出現重複定義函數的錯誤
func (tmp char) test0001() {

}

// 這種可以寫
func (tmp *long) tesst0002() {

}

type pointer *int

// 這麼不行
// pointer為接收者類型 它本身不能是指針
//func (tmp pointer) test0003() {
//
//}

func main() {
	per1 := PersonE{"kanna", 'M', 16}
	per1.PrintInfo01()

	var per2 PersonE
	(&per2).SetInfo("kinggyo", 'F', 14)
	per2.PrintInfo01()
}
