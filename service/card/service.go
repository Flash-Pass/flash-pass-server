package card

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/gin-gonic/gin"
)

type Service struct {
	cardRepo Repository
}

type IService interface {
	CreateCard(ctx *gin.Context, card *model.Card) error
	GetCard(ctx *gin.Context, id uint64) (*model.Card, error)
	UpdateCard(ctx *gin.Context, card *model.Card) (*model.Card, error)
	DeleteCard(ctx *gin.Context, cardId, userId uint64) error
	GetCardList(ctx *gin.Context, search string, userId uint64) ([]*model.Card, error)
}

type Repository interface {
	Create(ctx *gin.Context, card *model.Card) error
	GetById(ctx *gin.Context, cardId uint64) (*model.Card, error)
	Update(ctx *gin.Context, cardId uint64, question, answer string) (*model.Card, error)
	Delete(ctx *gin.Context, cardId, userId uint64) error
	GetList(ctx *gin.Context, search string, userId uint64) ([]*model.Card, error)
}

func NewService(repo Repository) *Service {
	return &Service{
		cardRepo: repo,
	}
}

func (s *Service) CreateCard(ctx *gin.Context, card *model.Card) error {
	return s.cardRepo.Create(ctx, card)
}

func (s *Service) GetCard(ctx *gin.Context, id uint64) (*model.Card, error) {
	return s.cardRepo.GetById(ctx, id)
}

func (s *Service) UpdateCard(ctx *gin.Context, card *model.Card) (*model.Card, error) {
	return s.cardRepo.Update(ctx, card.Id, card.Question, card.Answer)
}

func (s *Service) DeleteCard(ctx *gin.Context, cardId, userId uint64) error {
	return s.cardRepo.Delete(ctx, cardId, userId)
}

func (s *Service) GetCardList(ctx *gin.Context, search string, userId uint64) ([]*model.Card, error) {
	return s.cardRepo.GetList(ctx, search, userId)
}

var _ IService = (*Service)(nil)
