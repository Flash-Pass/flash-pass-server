package entity

import (
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/db/model"
)

type CardVO struct {
	Id        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedBy string `json:"created_by"`
}

func ConvertToCardVO(card *model.Card) *CardVO {
	return &CardVO{
		Id:        fmt.Sprint(card.Id),
		Question:  card.Question,
		Answer:    card.Answer,
		CreatedBy: fmt.Sprint(card.CreatedBy),
	}
}
