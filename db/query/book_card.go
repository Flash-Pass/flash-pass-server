package query

import (
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"gorm.io/gorm"
)

type BookCardQuery struct {
	db *gorm.DB
}

func NewBookCardQuery(db *gorm.DB) *BookCardQuery {
	return &BookCardQuery{db: db}
}

func (q *BookCardQuery) CheckExistByBookIdAndCardId(bookId, cardId uint64) (bool, error) {
	bookCard := model.BookCard{}
	if err := q.db.Where("book_id = ? AND card_id = ?", bookId, cardId).Find(&bookCard).Error; err != nil {
		return false, err
	}
	if bookCard == (model.BookCard{}) {
		return false, nil
	}
	return true, nil
}

func (q *BookCardQuery) Create(bookCard *model.BookCard) error {
	ok, err := q.CheckExistByBookIdAndCardId(bookCard.BookId, bookCard.CardId)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("book %d already has a card %d", bookCard.BookId, bookCard.CardId)
	}
	if err = q.db.Create(bookCard).Error; err != nil {
		return err
	}
	return nil
}

func (q *BookCardQuery) Delete(bookId, cardId, createdBy uint64) error {
	if err := q.db.Model(&model.BookCard{}).Where("book_id = ? AND card_id = ? AND created_by = ?", bookId, cardId, createdBy).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (q *BookCardQuery) GetBookCardList(bookId uint64) ([]*model.BookCard, error) {
	bookCardList := make([]*model.BookCard, 0)
	if err := q.db.Where("book_id = ?", bookId).Find(&bookCardList).Error; err != nil {
		return nil, err
	}
	return bookCardList, nil
}
