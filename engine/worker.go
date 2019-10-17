package engine

import (
	"crawler/fetcher"
	"fmt"
)

func work(r Request) (ParserResult, error) {
	//通过fetch获取到了utf8格式的html源代码
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		fmt.Printf("Got url body error : %s", err)
		return ParserResult{}, err
	}

	//再通过解析器，来解析源代码，得到下一级的[]requests
	return r.ParserFunc(body), nil
}
