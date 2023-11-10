package constants

const (
	CtxUserIdKey         = "userId"
	CtxOpenIdKey         = "openId"
	CtxSessionKeyKey     = "sessionKey"
	CtxUnionIdKey        = "unionId"
	WeChatLoginByCodeUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)
