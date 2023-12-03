package card

import (
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/gin-gonic/gin"
)

type Service struct {
	cardRepo Repository
	bookRepo BookRepository
}

type IService interface {
	CreateCard(ctx *gin.Context, card *model.Card) error
	CreateCardAndAddToBook(ctx *gin.Context, card *model.Card, bookId int64) error
	GetCard(ctx *gin.Context, id int64) (*model.Card, error)
	UpdateCard(ctx *gin.Context, card *model.Card) (*model.Card, error)
	DeleteCard(ctx *gin.Context, cardId, userId int64) error
	GetCardList(ctx *gin.Context, search string, userId int64) ([]*model.Card, error)
}

type Repository interface {
	Create(ctx *gin.Context, card *model.Card) error
	CreateCardAndAddToBook(ctx *gin.Context, card *model.Card, bookId int64) error
	GetById(ctx *gin.Context, cardId int64) (*model.Card, error)
	Update(ctx *gin.Context, cardId int64, question, answer string) (*model.Card, error)
	Delete(ctx *gin.Context, cardId, userId int64) error
	GetList(ctx *gin.Context, search string, userId int64) ([]*model.Card, error)
}

type BookRepository interface {
	IsBookIdExist(ctx *gin.Context, bookId int64) (bool, error)
}

func NewService(repo Repository, bookRepo BookRepository) *Service {
	return &Service{
		cardRepo: repo,
		bookRepo: bookRepo,
	}
}

func (s *Service) CreateCard(ctx *gin.Context, card *model.Card) error {
	return s.cardRepo.Create(ctx, card)
}

func (s *Service) CreateCardAndAddToBook(ctx *gin.Context, card *model.Card, bookId int64) error {
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

func (s *Service) GetCard(ctx *gin.Context, id int64) (*model.Card, error) {
	return s.cardRepo.GetById(ctx, id)
}

func (s *Service) UpdateCard(ctx *gin.Context, card *model.Card) (*model.Card, error) {
	return s.cardRepo.Update(ctx, card.Id, card.Question, card.Answer)
}

func (s *Service) DeleteCard(ctx *gin.Context, cardId, userId int64) error {
	return s.cardRepo.Delete(ctx, cardId, userId)
}

func (s *Service) GetCardList(ctx *gin.Context, search string, userId int64) ([]*model.Card, error) {
	return s.cardRepo.GetList(ctx, search, userId)
}

var _ IService = (*Service)(nil)
