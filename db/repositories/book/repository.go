package book

import (
	"context"

	"gorm.io/gorm"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
)

var _ IRepository = (*Repository)(nil)

type Repository struct {
	bookQuery     *query.BookQuery
	bookCardQuery *query.BookCardQuery
}

// IRepository 防腐层，定义数据操作接口，不指定具体实现，操作 db 时的相关校验包含在 query 层实现
type IRepository interface {
	Create(ctx context.Context, book *model.Book) error
	GetById(ctx context.Context, bookId int64) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, bookId, userId int64) error
	GetBookList(ctx context.Context, search string, userId int64) ([]*model.Book, error)
	GetBookCardList(ctx context.Context, bookId int64) ([]*model.BookCard, error)
	CreateBookCard(ctx context.Context, bookCard *model.BookCard) error
	DeleteBookCard(ctx context.Context, bookId, cardId, createdBy int64) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		query.NewBookQuery(db),
		query.NewBookCardQuery(db),
	}
}

func (r *Repository) Create(ctx context.Context, book *model.Book) error {
	if err := r.bookQuery.Create(book); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetById(ctx context.Context, bookId int64) (*model.Book, error) {
	book, err := r.bookQuery.GetById(bookId)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *Repository) Update(ctx context.Context, book *model.Book) error {
	if err := r.bookQuery.Update(book); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, bookId, userId int64) error {
	if err := r.bookQuery.Delete(bookId, userId); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetBookList(ctx context.Context, search string, userId int64) ([]*model.Book, error) {
	books, err := r.bookQuery.GetBookList(search, userId)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *Repository) GetBookCardList(ctx context.Context, bookId int64) ([]*model.BookCard, error) {
	bookCards, err := r.bookCardQuery.GetBookCardList(bookId)
	if err != nil {
		return nil, err
	}
	return bookCards, nil
}

func (r *Repository) CreateBookCard(ctx context.Context, bookCard *model.BookCard) error {
	if err := r.bookCardQuery.Create(bookCard); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteBookCard(ctx context.Context, bookId, cardId, createdBy int64) error {
	if err := r.bookCardQuery.Delete(bookId, cardId, createdBy); err != nil {
		return err
	}
	return nil
}

func (r *Repository) IsBookIdExist(ctx context.Context, bookId int64) (bool, error) {
	cnt, err := r.bookQuery.CountById(bookId)
	if err != nil {
		return false, err
	}
	if cnt == 0 {
		return false, nil
	}
	return true, nil
}
