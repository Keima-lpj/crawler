package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

const (
	URL = "https://www.zhenai.com/zhenghun"
)

//珍爱网的单任务版爬虫
func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:        URL,
		ParserFunc: parser.ParseCityList,
	})*/

	engine.ConcurrentEngine{
		Schedule:         &engine.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         engine.ItemSaver("immoc4"),
		RequestProcessor: engine.Work,
	}.Run(engine.Request{
		Url:    URL,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

}
