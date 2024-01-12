package model

type User struct {
	Base
	OpenId   string `gorm:"index" json:"open_id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Salt     string `json:"-"`
	Nickname string `json:"nickname"`
	Mobile   string `gorm:"type:char(11)" json:"mobile"`
	Avatar   string `gorm:"type:text" json:"avatar"`
}
