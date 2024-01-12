package card

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
)

type Handler struct {
	service         Service
	snowflakeHandle snowflake.IHandle
}

//go:generate mockgen -source=handler.go -destination=./mocks/handler_mock.go -package CardHandlerMocks
type IHandler interface {
	CreateCardController(c *gin.Context)
	GetCardController(c *gin.Context)
	UpdateCardController(c *gin.Context)
	DeleteCardController(c *gin.Context)
	GetCardListController(c *gin.Context)
}

//go:generate mockgen -source=handler.go -destination=./mocks/handler_mock.go -package CardHandlerMocks
type Service interface {
	CreateCard(ctx context.Context, card *model.Card) error
	GetCard(ctx context.Context, id int64) (*model.Card, error)
	UpdateCard(ctx context.Context, card *model.Card) (*model.Card, error)
	DeleteCard(ctx context.Context, id, userId int64) error
	GetCardList(ctx context.Context, search string, userId int64) ([]*model.Card, error)
	CreateCardAndAddToBook(ctx context.Context, card *model.Card, bookId int64) error
}

func NewHandler(service Service, snowflakeNode int64) *Handler {
	return &Handler{
		service:         service,
		snowflakeHandle: snowflake.NewHandle(snowflakeNode),
	}
}

func (h *Handler) AddRoutes(r *gin.Engine) {
	router := r.Group("/card")
	router.POST("/", h.CreateCardController)
	router.GET("/", h.GetCardController)
	router.PUT("/", h.UpdateCardController)
	router.DELETE("/", h.DeleteCardController)
	router.GET("/list", h.GetCardListController)
}
