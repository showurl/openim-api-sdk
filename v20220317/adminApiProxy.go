package v20220317

import (
	"encoding/json"
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
		return err
	}
	err = json.Unmarshal(buf, response)
	return err
}
