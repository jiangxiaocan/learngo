package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items []interface{} //任何类型都可
}

func NilFunction(byte []byte) ParseResult{
	return ParseResult{

	}
}
