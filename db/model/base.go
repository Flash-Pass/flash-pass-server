package model

import "time"

type Base struct {
	Id        string     `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	IsDeleted bool       `json:"is_deleted"`
}
