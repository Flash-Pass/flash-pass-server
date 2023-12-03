package model

type BookCard struct {
	Base
	BookId    int64 `json:"bookId"`
	Book      Book  `gorm:"ForeignKey:BookId;AssociationForeignKey:Id" json:"book"`
	CardId    int64 `json:"cardId"`
	Card      Card  `gorm:"ForeignKey:CardId;AssociationForeignKey:Id" json:"card"`
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
