package model

import "gorm.io/gen"

type Card struct {
	Base
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedBy string `json:"created_by" gorm:"index:created_by_user_id"`
}

type CardQueries interface {
	// SELECT * FROM @@table WHERE question LIKE '%@search%' OR answer LIKE '%@search%' OR created_by = @userId
	GetBySearchAndUserId(search, userId string) ([]*gen.T, error)
}

func NewCard(id, question, answer, createdBy string) *Card {
	return &Card{
		Base: Base{
			Id: id,
		},
		Question:  question,
		Answer:    answer,
		CreatedBy: createdBy,
	}
}
