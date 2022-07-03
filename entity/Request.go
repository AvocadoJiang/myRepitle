package entity

//请求
type Request struct {
	Url string
	//内容解析函数（解析器）
	ParseFunc func([]byte) ParseResult
}
