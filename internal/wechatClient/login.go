package wechatClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"net/http"
)

type WeChatLoginByCodeResponse struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"union_id"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type IWeChatLoginByCodeResponse interface {
	GetSessionKey() string
	GetUnionId() string
	GetOpenId() string
}

func (r *WeChatLoginByCodeResponse) GetSessionKey() string {
	return r.SessionKey
}

func (r *WeChatLoginByCodeResponse) GetUnionId() string {
	return r.UnionId
}

func (r *WeChatLoginByCodeResponse) GetOpenId() string {
	return r.OpenId
}

func (c *Client) LoginByCode(code string) (IWeChatLoginByCodeResponse, error) {
	url := fmt.Sprintf(constants.WeChatLoginByCodeUrl, c.WeChatAppid, c.WeChatAppSecret, code)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	wxLoginRes := WeChatLoginByCodeResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&wxLoginRes); err != nil {
		return nil, err
	}

	if wxLoginRes.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("wechat login by code defeat, error code is %d, error message is %s", wxLoginRes.ErrCode, wxLoginRes.ErrMsg))
	}

	return nil, nil
}

var _ IWeChatLoginByCodeResponse = (*WeChatLoginByCodeResponse)(nil)
