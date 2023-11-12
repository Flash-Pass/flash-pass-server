package model

type BookCard struct {
	Base
	BookId    int64 `json:"bookId"`
	CardId    int64 `json:"cardId"`
	CreatedBy int64 `json:"createdBy"`
}

func NewBookCard(id, bookId, cardId, createdBy int64) *BookCard {
	return &BookCard{
		Base: Base{
			Id: id,
		},
		BookId:    bookId,
		CardId:    cardId,
		CreatedBy: createdBy,
	}
}