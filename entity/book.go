package entity

import (
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/db/model"
)

type BookVO struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CardList    []*CardVO `json:"card_list"`
}

func ConvertToBookVO(book *model.Book) *BookVO {
	return &BookVO{
		Id:          fmt.Sprint(book.Id),
		Title:       book.Title,
		Description: book.Description,
		CreatedBy:   fmt.Sprint(book.CreatedBy),
	}
}
