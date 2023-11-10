package model

import (
	"time"
)

type Base struct {
	Id        uint64     `gorm:"primary_key" json:"id,string"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	IsDeleted bool       `gorm:"index" json:"is_deleted"`
}
