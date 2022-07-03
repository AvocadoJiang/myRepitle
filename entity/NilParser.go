package entity

//添加一个空的解析器
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
