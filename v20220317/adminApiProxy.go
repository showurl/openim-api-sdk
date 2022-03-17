package v20220317

import (
	"encoding/json"
	"errors"
	"github.com/showurl/openim-api-sdk/log"
	"github.com/showurl/openim-api-sdk/v20220317/utils"
)

func adminApiProxy(
	router string,
	req interface{},
	response interface{},
	options ...Option,
) (err error) {
	var opt *Options
	opt, err = getOptions(options)
	if err != nil {
		return
	}
	buf, err := utils.HttpClient.Post2Api(router, req, opt.token)
	if err != nil {
		if errors.Is(utils.StatusCodeErr, err) {
			// 删掉token
			log.Infof("可能token过期，重试一次")
			delToken(opt.userId)
			return adminApiProxyDontRetry(router, req, response, options...)
		}
		return err
	}
	err = json.Unmarshal(buf, response)
	return err
}

func adminApiProxyDontRetry(
	router string,
	req interface{},
	response interface{},
	options ...Option,
) (err error) {
	var opt *Options
	opt, err = getOptions(options)
	if err != nil {
		return
	}
	buf, err := utils.HttpClient.Post2Api(router, req, opt.token)
	if err != nil {
		log.Errorf("请求失败，重试过一次")
		return err
	}
	err = json.Unmarshal(buf, response)
	return err
}
