package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"miao_sticker_server/media/constdef"
)

type MyRequest struct {
	client *http.Client
}

var myRequest *MyRequest = &MyRequest{}

func init() {
	myRequest.client = &http.Client{}
}

// http请求
func (r MyRequest) httpRequestGet(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// 设置请求投
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	// Do sends an HTTP request and returns an HTTP response
	// 发起一个HTTP请求，返回一个HTTP响应
	return r.client.Do(request)
}

// http请求
func (r MyRequest) httpRequestGetWithBody(url string, body interface{}) (*http.Response, error) {
	bodyBytes, _ := json.Marshal(body)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	// 设置请求投
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	// Do sends an HTTP request and returns an HTTP response
	// 发起一个HTTP请求，返回一个HTTP响应
	return r.client.Do(request)
}

// 编码识别
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Printf("fetch error : %v\n", err)
		// 如果没有识别到，返回一个UTF-8(默认)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}

// 根据URL提取
func Fetch(url string, body interface{}) ([]byte, error) {
	var resp *http.Response
	var err error
	if body != nil {
		resp, err = myRequest.httpRequestGetWithBody(url, body)
	} else {
		resp, err = myRequest.httpRequestGet(url)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d of %s", resp.StatusCode, url)
	}
	reader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, encoding.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func GetCode2Session(code string) ([]byte, error) {
	body := &map[string]string{
		"appid":      constdef.APPID,
		"secret":     constdef.APPSecret,
		"js_code":    code,
		"grant_type": constdef.GRANT_TYPE,
	}
	return Fetch(constdef.Code2Session_URL, body)
}
