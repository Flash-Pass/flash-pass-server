package book

import (
	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/repositories/book"
	"github.com/Flash-Pass/flash-pass-server/entity"
)

var _ IService = (*Service)(nil)

type Service struct {
	bookRepo book.IRepository
	cardRepo cardRepository
}

type IService interface {
	CreateBook(ctx *gin.Context, card *model.Book) error
	GetBook(ctx *gin.Context, id int64) (*entity.BookVO, error)
	UpdateBook(ctx *gin.Context, book *model.Book) error
	DeleteBook(ctx *gin.Context, bookId, userId int64) error
	GetBookList(ctx *gin.Context, search string, userId int64) ([]*entity.BookVO, error)
	AddCardToBook(ctx *gin.Context, bookCard *model.BookCard) error
	DeleteCardFromBook(ctx *gin.Context, bookId, cardId, userId int64) error
}

type cardRepository interface {
	GetListByIds(ctx *gin.Context, ids []int64) ([]*model.Card, error)
}

func NewService(bookRepo book.IRepository, cardRepo cardRepository) *Service {
	return &Service{
		bookRepo, cardRepo,
	}
}

func (s *Service) CreateBook(ctx *gin.Context, book *model.Book) error {
	return s.bookRepo.Create(ctx, book)
}

func (s *Service) GetBook(ctx *gin.Context, id int64) (*entity.BookVO, error) {
	book, err := s.bookRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	bookVO := entity.ConvertToBookVO(book)
	bookCards, err := s.bookRepo.GetBookCardList(ctx, id)
	if err != nil {
		return nil, err
	}

	cardIds := make([]int64, 0, len(bookCards))
	for _, bookCard := range bookCards {
		cardIds = append(cardIds, bookCard.CardId)
	}

	cards, err := s.cardRepo.GetListByIds(ctx, cardIds)
	if err != nil {
		return nil, err
	}

	bookVO.CardList = make([]*entity.CardVO, 0, len(bookCards))
	for _, card := range cards {
		bookVO.CardList = append(bookVO.CardList, entity.ConvertToCardVO(card))
	}
	return bookVO, nil
}

func (s *Service) UpdateBook(ctx *gin.Context, book *model.Book) error {
	return s.bookRepo.Update(ctx, book)
}

func (s *Service) DeleteBook(ctx *gin.Context, bookId, userId int64) error {
	return s.bookRepo.Delete(ctx, bookId, userId)
}

func (s *Service) GetBookList(ctx *gin.Context, search string, userId int64) ([]*entity.BookVO, error) {
	bookList, err := s.bookRepo.GetBookList(ctx, search, userId)
	if err != nil {
		return nil, err
	}
	bookVOList := make([]*entity.BookVO, 0, len(bookList))
	for _, book := range bookList {
		bookVOList = append(bookVOList, entity.ConvertToBookVO(book))
	}
	return bookVOList, nil
}

func (s *Service) AddCardToBook(ctx *gin.Context, bookCard *model.BookCard) error {
	return s.bookRepo.CreateBookCard(ctx, bookCard)
}

func (s *Service) DeleteCardFromBook(ctx *gin.Context, bookId, cardId, userId int64) error {
	return s.bookRepo.DeleteBookCard(ctx, bookId, cardId, userId)
}
