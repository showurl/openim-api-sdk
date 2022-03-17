package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/showurl/openim-api-sdk/log"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
	"io/ioutil"
	"net/http"
	"time"
)

var HttpClient *httpClient
var StatusCodeErr = fmt.Errorf("resp.StatusCode != 200")

type httpClient struct {
	UrlPrefix string
}

func init() {
	HttpClient = &httpClient{
		UrlPrefix: conf.UrlPrefix,
	}
}

//application/json; charset=utf-8
func (c *httpClient) Post2Api(path string, data interface{}, token string) (content []byte, err error) {
	ct, err := c.postLogic(path, data, token)
	return ct, err
}

func (c *httpClient) postLogic(path string, data interface{}, token string) (content []byte, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	log.Debugf("请求体：%s", string(jsonStr))
	req, err := http.NewRequest("POST", c.UrlPrefix+path, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Add("content-type", "application/json")
	req.Header.Add("token", token)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Errorf("请求失败：resp.StatusCode = %d", resp.StatusCode)
		return nil, StatusCodeErr
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Debugf("响应体：%s", string(result))
	return result, nil
}
