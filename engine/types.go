package engine

//request的结构体，其中包含将要访问的url，以及这个url对应的解析器
type Request struct {
	Url        string
	ParserFunc func([]byte, string) ParserResult
}

//解析器解析了文本之后，返回下一级页面的若干个request(这个request中包含了对应子页面的url和子页面的解析器)
//以及这些request对应的item
type ParserResult struct {
	Requests []Request
	Item     []Item
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}
