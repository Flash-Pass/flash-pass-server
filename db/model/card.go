package model

import "gorm.io/gen"

type Card struct {
	Base      `jsonUp:"true"`
	Question  string `gorm:"type:varchar(255)" json:"question"`
	Answer    string `gorm:"type:text" json:"answer"`
	CreatedBy uint64 `gorm:"index" json:"created_by,string"`
}

type CardQueries interface {
	// SELECT * FROM @@table WHERE question LIKE concat("%", @search,"%") OR answer LIKE concat("%", @search,"%") OR created_by = @userId
	GetBySearchAndUserId(search string, userId uint64) ([]*gen.T, error)
}

func NewCard(id uint64, question, answer string, createdBy uint64) *Card {
	return &Card{
		Base: Base{
			Id: id,
		},
		Question:  question,
		Answer:    answer,
		CreatedBy: createdBy,
	}
}
