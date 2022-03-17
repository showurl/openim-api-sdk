package v20220317

import "github.com/showurl/openim-api-sdk/v20220317/conf"

type Options struct {
	token  string // token
	userId string // 管理员的id
	secret string // 管理员的密码
}
type Option func(*Options)

func WithToken(token string) Option {
	return func(options *Options) {
		options.token = token
	}
}
func WithUser(userId, secret string) Option {
	return func(options *Options) {
		options.userId = userId
		options.secret = secret
	}
}

func getOptions(options []Option) (*Options, error) {
	o := &Options{}
	for _, option := range options {
		option(o)
	}
	if o.token == "" {
		if o.userId == "" {
			o.userId = conf.DefaultUserID
			o.secret = conf.DefaultUserSecret
		}
		token, err := GetToken(o.userId, o.secret)
		if err != nil {
			return nil, err
		}
		return &Options{
			token:  token,
			userId: conf.DefaultUserID,
			secret: conf.DefaultUserSecret,
		}, nil
	} else {
		return o, nil
	}
}
