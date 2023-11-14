package model

type BookCard struct {
	Base
	BookId    uint64 `json:"bookId"`
	CardId    uint64 `json:"cardId"`
	CreatedBy uint64 `json:"createdBy"`
}

func NewBookCard(id, bookId, cardId, createdBy uint64) *BookCard {
	return &BookCard{
		Base: Base{
			Id: id,
		},
		BookId:    bookId,
		CardId:    cardId,
		CreatedBy: createdBy,
	}
}
