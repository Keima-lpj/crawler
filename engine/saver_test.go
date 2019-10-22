package engine

import (
	"crawler/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	expect := Item{
		Id:   "1",
		Url:  "https://album.zhenai.com/u/1942899917",
		Type: "crawler",
		Payload: model.Person{
			Name:   "lpj",
			Gender: "男",
			Age:    "28",
			Height: "180",
			Weight: "65",
			Income: "*",
		},
	}

	result, err := Save(expect, "immoc")
	if err != nil {
		t.Errorf("error! result:%s, err:%s", result, err)
	}
	fmt.Println(result)

	//再从es中将这个数据取出来，判断两个数据是否相等
	getResult, err := Gets("1")
	if err != nil {
		t.Errorf("error! get data faild, error:%v", err)
	}

	var actual = Item{}
	err = json.Unmarshal(*getResult.Source, &actual)
	if err != nil {
		panic(err)
	}
	fmt.Println(actual)

	if actual != expect {
		t.Errorf("Expect:%v, Got:%v", expect, actual)
	}

}
