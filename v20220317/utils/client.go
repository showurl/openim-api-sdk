package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
	"io/ioutil"
	"net/http"
	"time"
)

var HttpClient *httpClient

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
		return nil, fmt.Errorf("resp.StatusCode = %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//	fmt.Println(path, "Marshal data: ", string(jsonStr), string(result))
	return result, nil
}
