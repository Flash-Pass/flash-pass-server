package wechatClient

type Client struct {
	WeChatAppid     string
	WeChatAppSecret string
}

type IClient interface {
	LoginByCode(code string) (IWeChatLoginByCodeResponse, error)
}

func New(appId, secret string) *Client {
	return &Client{
		WeChatAppid:     appId,
		WeChatAppSecret: secret,
	}
}

var _ IClient = (*Client)(nil)
