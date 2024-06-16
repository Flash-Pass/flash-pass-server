package card

import (
	"errors"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	card query.ICardDo
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CardRepositoryMocks
type IRepository interface {
	Create(ctx *gin.Context, card *model.Card) error
	GetById(ctx *gin.Context, cardId uint64) (*model.Card, error)
	Update(ctx *gin.Context, cardId uint64, question, answer string) (*model.Card, error)
	Delete(ctx *gin.Context, cardId, userId uint64) error
	GetList(ctx *gin.Context, search string, userId uint64) ([]*model.Card, error)
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CardRepositoryMocks
type ICard interface {
	query.ICardDo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		card: query.Card.WithContext(db.Statement.Context),
	}
}

func (r *Repository) Create(ctx *gin.Context, card *model.Card) error {
	logger := ctxlog.GetLogger(ctx)

	if err := r.card.Create(card); err != nil {
		logger.Error("create card defeat", zap.Error(err), zap.Any("card", card))
		return err
	}

	return nil
}

func (r *Repository) GetById(ctx *gin.Context, cardId uint64) (*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	card, err := r.card.WithContext(ctx).Where(query.Card.Id.Eq(cardId)).First()
	if err != nil {
		logger.Error("card id not found", zap.Error(err), zap.Any("card id", cardId))
		return nil, err
	}

	return card, nil
}

func (r *Repository) Update(ctx *gin.Context, cardId uint64, question, answer string) (*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	_, err := r.card.WithContext(ctx).Where(query.Card.Id.Eq(cardId)).Updates(model.Card{
		Question: question,
		Answer:   answer,
	})

	if err != nil {
		logger.Error("update card defeat", zap.Error(err), zap.Uint64("card id", cardId))
		return nil, err
	}

	return r.GetById(ctx, cardId)
}

func (r *Repository) Delete(ctx *gin.Context, cardId, userId uint64) error {
	logger := ctxlog.GetLogger(ctx)

	updateInfo, err := r.card.WithContext(ctx).Where(
		query.Card.Id.Eq(cardId),
		query.Card.CreatedBy.Eq(userId),
		query.Card.IsDeleted.Is(false)).Update(query.Card.IsDeleted, true)
	if err != nil {
		logger.Error("delete card defeat", zap.Error(err), zap.Uint64("card id", cardId))
		return err
	}

	if updateInfo.RowsAffected == 0 {
		logger.Error("delete card defeat", zap.Error(err), zap.Uint64("card id", cardId))
		return errors.New("card has been deleted or card not exist")
	}

	return nil
}

// GetList use search and userId to get a card list.
// If searched string in question or answer in a card, it will be added to the result list.
// If a card matched by the given user id, it also will be added to the result list.
func (r *Repository) GetList(ctx *gin.Context, search string, userId uint64) ([]*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	list, err := r.card.GetBySearchAndUserId(search, userId)
	if err != nil {
		logger.Error("search card list defeat", zap.Error(err), zap.String("search string", search), zap.Uint64("user id", userId))
		return nil, err
	}

	return list, nil
}

// GetListByIds TODO 待实现
func (r *Repository) GetListByIds(ctx *gin.Context, cardIds []int64) ([]*model.Card, error) {
	return nil, nil
}

var _ IRepository = (*Repository)(nil)
