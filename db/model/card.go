package model

import "gorm.io/gen"

type Card struct {
	Base
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedBy int64  `json:"created_by" gorm:"index:created_by_user_id"`
}

type CardQueries interface {
	// SELECT * FROM @@table WHERE question LIKE concat("%", @search,"%") OR answer LIKE concat("%", @search,"%") OR created_by = @userId
	GetBySearchAndUserId(search, userId string) ([]*gen.T, error)
}

func NewCard(id int64, question, answer string, createdBy int64) *Card {
	return &Card{
		Base: Base{
			Id: id,
		},
		Question:  question,
		Answer:    answer,
		CreatedBy: createdBy,
	}
}
