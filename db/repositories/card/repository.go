package card

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
)

type Repository struct {
	query    *query.Query
	card     query.ICardDo
	bookCard query.IBookCardDo
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CardRepositoryMocks
type IRepository interface {
	Create(ctx *gin.Context, card *model.Card) error
	CreateCardAndAddToBook(ctx *gin.Context, card *model.Card, bookId int64) error
	GetById(ctx *gin.Context, cardId int64) (*model.Card, error)
	Update(ctx *gin.Context, cardId int64, question, answer string) (*model.Card, error)
	Delete(ctx *gin.Context, cardId, userId int64) error
	GetList(ctx *gin.Context, search string, userId int64) ([]*model.Card, error)
	GetListByIds(ctx *gin.Context, cardIds []int64) ([]*model.Card, error)
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CardRepositoryMocks
type ICard interface {
	query.ICardDo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		query: query.Use(db),
		card:  query.Card.WithContext(db.Statement.Context),
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

func (r *Repository) CreateCardAndAddToBook(ctx *gin.Context, card *model.Card, bookId int64) error {
	logger := ctxlog.GetLogger(ctx)

	tx := r.query.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Card.WithContext(ctx).Create(card); err != nil {
		logger.Error("create card defeat", zap.Error(err), zap.Any("card", card))
		return err
	}

	if err := tx.BookCard.WithContext(ctx).Create(&model.BookCard{
		BookId:    bookId,
		CardId:    card.Id,
		CreatedBy: card.CreatedBy,
	}); err != nil {
		logger.Error("add card to book defeat", zap.Error(err), zap.Int64("card id", card.Id), zap.Int64("book id", bookId))
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetById(ctx *gin.Context, cardId int64) (*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	card, err := r.card.WithContext(ctx).Where(query.Card.Id.Eq(cardId)).First()
	if err != nil {
		logger.Error("card id not found", zap.Error(err), zap.Any("card id", cardId))
		return nil, err
	}

	return card, nil
}

func (r *Repository) Update(ctx *gin.Context, cardId int64, question, answer string) (*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	_, err := r.card.WithContext(ctx).Where(query.Card.Id.Eq(cardId)).Updates(model.Card{
		Question: question,
		Answer:   answer,
	})

	if err != nil {
		logger.Error("update card defeat", zap.Error(err), zap.Int64("card id", cardId))
		return nil, err
	}

	return r.GetById(ctx, cardId)
}

func (r *Repository) Delete(ctx *gin.Context, cardId, userId int64) error {
	logger := ctxlog.GetLogger(ctx)

	updateInfo, err := r.card.WithContext(ctx).Where(
		query.Card.Id.Eq(cardId),
		query.Card.CreatedBy.Eq(userId),
		query.Card.IsDeleted.Is(false)).Update(query.Card.IsDeleted, true)
	if err != nil {
		logger.Error("delete card defeat", zap.Error(err), zap.Int64("card id", cardId))
		return err
	}

	if updateInfo.RowsAffected == 0 {
		logger.Error("delete card defeat", zap.Error(err), zap.Int64("card id", cardId))
		return errors.New("card has been deleted or card not exist")
	}

	return nil
}

func (r *Repository) GetList(ctx *gin.Context, search string, userId int64) ([]*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	list, err := r.card.Where(
		query.Card.CreatedBy.Eq(userId),
		query.Card.IsDeleted.Is(false),
	).Where(
		query.Card.Question.Like("%"+search+"%"),
		query.Card.Answer.Like("%"+search+"%"),
	).Find()
	if err != nil {
		logger.Error("search card list defeat", zap.Error(err), zap.String("search string", search), zap.Int64("user id", userId))
		return nil, err
	}

	return list, nil
}

func (r *Repository) GetListByIds(ctx *gin.Context, cardIds []int64) ([]*model.Card, error) {
	logger := ctxlog.GetLogger(ctx)

	list, err := r.card.WithContext(ctx).Where(query.Card.Id.In(cardIds...)).Find()
	if err != nil {
		logger.Error("search card list by ids defeat", zap.Error(err), zap.Any("card_ids", cardIds))
		return nil, err
	}
	return list, nil
}

var _ IRepository = (*Repository)(nil)
