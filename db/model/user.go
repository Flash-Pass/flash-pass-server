package model

type User struct {
	Base
	OpenId   string `json:"open_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar"`
}
