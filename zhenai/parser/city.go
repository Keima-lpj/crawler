package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

//正则表达式
//注：[^>]表示匹配非'>'符号的任意字符串
const cityRegexpString = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
const sexRegexpString = `<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`

//这个是城市列表的解析器，通过解析城市列表页面的文本，返回下一级页面的request数组和对应的item
func ParseCity(contents []byte) engine.ParserResult {
	//使用正则表达式来匹配到对应的用户和链接
	re := regexp.MustCompile(cityRegexpString)
	match := re.FindAllSubmatch(contents, -1)
	re = regexp.MustCompile(sexRegexpString)
	matchSex := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for k, v := range match {
		//匹配这个用户的性别
		sex := string(matchSex[k][1])
		result.Requests = append(result.Requests, engine.Request{
			Url: strings.Replace(string(v[1]), "http", "https", 1),
			ParserFunc: func(bytes []byte) engine.ParserResult {
				return ParsePerson(bytes, sex)
			}, //这里应该传下一级解析器，也就是用户解析器
		})
		result.Item = append(result.Item, "User "+string(v[2]))
	}

	return result
}
