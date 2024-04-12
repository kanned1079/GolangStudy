package lei

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 這樣可以減少重複性 只有類型 沒有名字
	// 繼承了Person的成員
	id   int
	addr string
}
