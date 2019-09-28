package engine

import (
	"crawler/fetcher"
	"fmt"
)

//这里是运行主程序
func Run(seeds ...Request) {
	//声明一个运行队列，将seeds放入队列中，使用广度优先算法来跑
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		//获取单个的request
		r := requests[0]
		requests = requests[1:]
		fmt.Println("Fetching Url:", r.Url)

		//通过fetch获取到了utf8格式的html源代码
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			fmt.Printf("Got url body error : %s", err)
			continue
		}

		//再通过解析器，来解析源代码，得到下一级的[]requests
		parseResult := r.ParserFunc(body)

		//这里将parseResult的request数组放入队列中，并且打印出返回的item
		if len(parseResult.Requests) > 0 {
			requests = append(requests, parseResult.Requests...)
		}

		for _, v := range parseResult.Item {
			fmt.Printf("Got item :%v\n", v)
		}
	}
}
