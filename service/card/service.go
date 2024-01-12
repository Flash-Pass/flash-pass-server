package card

import (
	"context"
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/db/model"
)

type Service struct {
	cardRepo Repository
	bookRepo BookRepository
}

type IService interface {
	CreateCard(ctx context.Context, card *model.Card) error
	GetCard(ctx context.Context, id int64) (*model.Card, error)
	UpdateCard(ctx context.Context, card *model.Card) (*model.Card, error)
	DeleteCard(ctx context.Context, cardId, userId int64) error
	GetCardList(ctx context.Context, search string, userId int64) ([]*model.Card, error)
}

type Repository interface {
	Create(ctx context.Context, card *model.Card) error
	GetById(ctx context.Context, cardId int64) (*model.Card, error)
	Update(ctx context.Context, cardId int64, question, answer string) (*model.Card, error)
	Delete(ctx context.Context, cardId, userId int64) error
	GetList(ctx context.Context, search string, userId int64) ([]*model.Card, error)
	CreateCardAndAddToBook(ctx context.Context, card *model.Card, bookId int64) error
}

type BookRepository interface {
	IsBookIdExist(ctx context.Context, bookId int64) (bool, error)
}

func NewService(repo Repository, bookRepo BookRepository) *Service {
	return &Service{
		cardRepo: repo,
		bookRepo: bookRepo,
	}
}

func (s *Service) CreateCard(ctx context.Context, card *model.Card) error {
	return s.cardRepo.Create(ctx, card)
}

func (s *Service) CreateCardAndAddToBook(ctx context.Context, card *model.Card, bookId int64) error {
	ok, err := s.bookRepo.IsBookIdExist(ctx, bookId)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("book %d not exist", bookId)
	}
	if err = s.cardRepo.CreateCardAndAddToBook(ctx, card, bookId); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetCard(ctx context.Context, id int64) (*model.Card, error) {
	return s.cardRepo.GetById(ctx, id)
}

func (s *Service) UpdateCard(ctx context.Context, card *model.Card) (*model.Card, error) {
	return s.cardRepo.Update(ctx, card.Id, card.Question, card.Answer)
}

func (s *Service) DeleteCard(ctx context.Context, cardId, userId int64) error {
	return s.cardRepo.Delete(ctx, cardId, userId)
}

func (s *Service) GetCardList(ctx context.Context, search string, userId int64) ([]*model.Card, error) {
	return s.cardRepo.GetList(ctx, search, userId)
}

var _ IService = (*Service)(nil)
