package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//这里通过获取的url来返回对应解析成功的utf8格式的html文本
func Fetch(url string) ([]byte, error) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	//request.Header.Add("Cookie", "xxxxxx")
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	request.Header.Add("upgrade-insecure-requests", "1")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//判断响应码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Http code error, %d\n", resp.StatusCode)
	}

	bufReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bufReader)

	//转换编码
	reader := transform.NewReader(bufReader, encode.NewDecoder())
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("reader Html Error:%s\n", err)
	}

	return contents, nil
}

// 获取当前页面的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		// 如果没有获取到编码格式，则返回默认UTF-8编码格式
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
