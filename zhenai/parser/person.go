package parser

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
)

//这里提前编译，因为每次单独编译的话会耗费性能
var (
	nameRegexp   = regexp.MustCompile(`<h1 class="nickName"[^>]+>([^<]+)</h1>`)
	ageRegexp    = regexp.MustCompile(`<div [^>]+>([0-9]+岁)</div>`)
	heightRegexp = regexp.MustCompile(`<div [^>]+>([0-9]+cm)</div>`)
	weightRegexp = regexp.MustCompile(`<div [^>]+>([0-9]+kg)</div>`)
	incomeRegexp = regexp.MustCompile(`<div [^>]+>月收入:([^<]+)</div>`)
	IdRegexp     = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

//这个是城市列表的解析器，通过解析城市列表页面的文本，返回下一级页面的request数组和对应的item
func ParsePerson(contents []byte, gender, url string) engine.ParserResult {

	//使用正则表达式来匹配到对应的城市和链接
	person := model.Person{}

	result := engine.ParserResult{}

	name := nameRegexp.FindSubmatch(contents)
	if len(name) >= 2 {
		person.Name = string(name[1])
	}

	person.Gender = gender

	age := ageRegexp.FindSubmatch(contents)
	if len(age) >= 2 {
		person.Age = string(age[1])
	}

	height := heightRegexp.FindSubmatch(contents)
	if len(height) >= 2 {
		person.Height = string(height[1])
	}

	weight := weightRegexp.FindSubmatch(contents)
	if len(weight) >= 2 {
		person.Weight = string(weight[1])
	}

	income := incomeRegexp.FindSubmatch(contents)
	if len(income) >= 2 {
		person.Height = string(income[1])
	}

	var id string
	idMatch := IdRegexp.FindSubmatch([]byte(url))
	if len(idMatch) == 2 {
		id = string(idMatch[1])
	} else {
		return result
	}

	//根据id查询ES中此人是否已经存在，如果已经存在，则不再次录入
	_, err := engine.Gets(id)
	if err == nil {
		log.Printf("jump save Id: %v. this Id has been saved", id)
		return result
	}

	result = engine.ParserResult{
		Item: []engine.Item{
			{
				Id:      id,
				Url:     url,
				Type:    "zhenai",
				Payload: person,
			},
		},
	}

	return result
}
