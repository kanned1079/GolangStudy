package main

import "fmt"

type Stude struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 0                // intger
	i[1] = "hellow"         // string
	i[2] = Stude{"mike", 6} // struct

	// 那麼怎麼反推類型呢
	// 需要 類型查詢 => 類型斷言

	// 	1. 可以通過if來實現
	// index返回下標 data返回數據 data分別是 i[0] i[1] i[2]
	for index, data := range i {
		if val, status := data.(int); status == true {
			fmt.Printf("x[%d] 類型為int, 內容為 %d\n", index, val)
		} else if val, status := data.(string); status == true {
			fmt.Printf("x[%d] 類型為string, 內容為 %s\n", index, val)
		} else if val, status := data.(Stude); status == true {
			fmt.Printf("x[%d] 類型為string, 內容為 name=%s, id=%d\n", index, val.name, val.id)

		}
	}
	fmt.Println("-----------------------------")

	// 	2. 可以通過switch 來實現
	for index01, data01 := range i {
		switch value := data01.(type) {
		case int:
			fmt.Printf("x[%d] 類型為int, 內容為 %d\n", index01, value)
		case string:
			fmt.Printf("x[%d] 類型為string, 內容為 %s\n", index01, value)
		case Stude:
			fmt.Printf("x[%d] 類型為string, 內容為 name=%s, id=%d\n", index01, value.name, value.id)

		}

	}

}
