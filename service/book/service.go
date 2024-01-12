package book

import (
	"context"

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
	CreateBook(ctx context.Context, card *model.Book) error
	GetBook(ctx context.Context, id int64) (*entity.BookVO, error)
	UpdateBook(ctx context.Context, book *model.Book) error
	DeleteBook(ctx context.Context, bookId, userId int64) error
	GetBookList(ctx context.Context, search string, userId int64) ([]*entity.BookVO, error)
	AddCardToBook(ctx context.Context, bookCard *model.BookCard) error
	DeleteCardFromBook(ctx context.Context, bookId, cardId, userId int64) error
	GetBookCardList(ctx context.Context, bookId int64) ([]*model.Card, error)
}

type cardRepository interface {
	GetListByIds(ctx context.Context, ids []int64) ([]*model.Card, error)
}

func NewService(bookRepo book.IRepository, cardRepo cardRepository) *Service {
	return &Service{
		bookRepo, cardRepo,
	}
}

func (s *Service) CreateBook(ctx context.Context, book *model.Book) error {
	return s.bookRepo.Create(ctx, book)
}

func (s *Service) GetBook(ctx context.Context, id int64) (*entity.BookVO, error) {
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

func (s *Service) UpdateBook(ctx context.Context, book *model.Book) error {
	return s.bookRepo.Update(ctx, book)
}

func (s *Service) DeleteBook(ctx context.Context, bookId, userId int64) error {
	return s.bookRepo.Delete(ctx, bookId, userId)
}

func (s *Service) GetBookList(ctx context.Context, search string, userId int64) ([]*entity.BookVO, error) {
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

func (s *Service) AddCardToBook(ctx context.Context, bookCard *model.BookCard) error {
	return s.bookRepo.CreateBookCard(ctx, bookCard)
}

func (s *Service) DeleteCardFromBook(ctx context.Context, bookId, cardId, userId int64) error {
	return s.bookRepo.DeleteBookCard(ctx, bookId, cardId, userId)
}

func (s *Service) GetBookCardList(ctx context.Context, bookId int64) ([]*model.Card, error) {
	bookCards, err := s.bookRepo.GetBookCardList(ctx, bookId)
	if err != nil {
		return nil, err
	}
	cards := make([]*model.Card, 0, len(bookCards))
	for _, bookCard := range bookCards {
		cards = append(cards, &bookCard.Card)
	}
	return cards, nil
}
