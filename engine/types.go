package engine

type Request struct {
	//请求url
	Url string
	//解析函数
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	//新的request
	Request []Request
	//得到的内容
	Items []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}