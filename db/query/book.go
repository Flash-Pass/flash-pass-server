package query

import (
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"gorm.io/gorm"
)

type BookQuery struct {
	db *gorm.DB
}

func NewBookQuery(db *gorm.DB) *BookQuery {
	return &BookQuery{db: db}
}

func (q *BookQuery) CheckExistByTitleAndUserId(title string, bookId uint64) (bool, error) {
	book := model.Book{}
	if err := q.db.Where("title = ? AND created_by = ?", title, bookId).Find(&book).Error; err != nil {
		return false, err
	}
	if book == (model.Book{}) {
		return false, nil
	}
	return true, nil
}

func (q *BookQuery) Create(book *model.Book) error {
	ok, err := q.CheckExistByTitleAndUserId(book.Title, book.CreatedBy)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("user %d already has a book named %s", book.CreatedBy, book.Title)
	}
	if err = q.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (q *BookQuery) Update(book *model.Book) error {
	if err := q.db.Model(book).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func (q *BookQuery) GetById(bookId uint64) (*model.Book, error) {
	book := &model.Book{}
	if err := q.db.Where("id = ?", bookId).First(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (q *BookQuery) Delete(bookId, userId uint64) error {
	if err := q.db.Model(&model.Book{}).Where("id = ? AND created_by = ?", bookId, userId).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (q *BookQuery) GetBookList(search string, userId uint64) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	if err := q.db.Where("created_by = ? OR title like ? OR description like ?", userId, "%"+search+"%", "%"+search+"%").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
