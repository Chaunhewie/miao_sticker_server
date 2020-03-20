package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/tidwall/gjson"

	"miao_sticker_server/index/logger"
	"miao_sticker_server/media/constdef"
	"miao_sticker_server/media/models"
)

type MyRequest struct {
	client *http.Client
}

var myRequest *MyRequest = &MyRequest{}

func init() {
	myRequest.client = &http.Client{}
}

// http请求
func (r MyRequest) httpRequest(url string) (*http.Response, error) {
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
func Fetch(url string) ([]byte, error) {
	resp, err := myRequest.httpRequest(url)
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

func FetchRepoInfoLoop(filePath string) {
	for true {
		resp, err := Fetch(constdef.REPO_URL)
		if err != nil {
			logger.Error("Fetch Repo Info Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		var repoInfo models.RepoInfo
		repoInfo.StargazersCount = int32(gjson.Get(string(resp), "stargazers_count").Num)
		repoInfo.WatchersCount = int32(gjson.Get(string(resp), "watchers_count").Num)
		repoInfo.Forks = int32(gjson.Get(string(resp), "forks_count").Num)
		repoInfoBytes, err := json.Marshal(repoInfo)
		if err != nil {
			logger.Error("Marshal Repo Info Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		f, err := os.OpenFile(filePath+"__", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			logger.Error("Open Repo Info Tmp File Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		if _, err = f.Write(repoInfoBytes); err != nil {
			logger.Error("Write Repo Info Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		if err = f.Close(); err != nil{
			logger.Error("Close Repo Info Tmp File Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		if err = os.Rename(filePath+"__", filePath); err != nil {
			logger.Error("Rename Repo Info File Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		logger.Info("Fetch Repo Info Success")
		time.Sleep(constdef.DURATION_REPO_REQ_SUCCESS)
	}
}
