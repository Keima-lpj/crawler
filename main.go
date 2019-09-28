package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

const (
	URL  = "https://www.zhenai.com/zhenghun"
	URL1 = "https://album.zhenai.com/u/1930805153"
)

//珍爱网的单任务版爬虫
func main() {
	engine.Run(engine.Request{
		Url:        URL,
		ParserFunc: parser.ParseCityList,
	})

}
