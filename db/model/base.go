package model

import "time"

type Base struct {
	Id        int64      `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	IsDeleted int        `json:"is_deleted"`
}
