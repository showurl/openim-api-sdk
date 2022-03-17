# openim-api-sdk

openim-api GO版本sdk

## 目录结构

|路径|描述|
|---|---|
|apiStruct/|api请求和响应的结构体|
|conf/|配置文件|
|tests/|单元测试文件|
|utils/|工具方法|
|||
|adminApiProxy.go|后台管理api请求代理|
|apiOption.go|请求选项|
|newOperationID.go|OperationID相关|
|tokenMgr.go|token管理相关|
|||
|CreateGroup.go|创建群组|
|InviteUserToGroup.go|邀请用户直接进群|
|KickGroup.go|从群里踢出用户|
|SendMsg.go|发送消息|
|UpdateUserInfo.go|更新用户的资料|
|UserRegister.go|注册新用户|
|UserToken.go|用户登录获取token|

## 环境变量

|key|描述|默认值|
|---|---|---|
|UrlPrefix|openim api服务的host|http://127.0.0.1:10000|
|DefaultUserID|openim 默认admin账号 当option不传时 则使用该账号|openIM123456|
|DefaultUserSecret|config.yaml中的secret|tuoyun|
|LogLevel|从debug,info,error,none中选一个|debug|