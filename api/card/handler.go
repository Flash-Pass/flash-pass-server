package card

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service         Service
	snowflakeHandle snowflake.IHandle
}

//go:generate mockgen -source=handler.go -destination=./mocks/handler_mock.go -package CardHandlerMocks
type IHandler interface {
	CreateCardController(ctx *gin.Context)
	GetCardController(ctx *gin.Context)
	UpdateCardController(ctx *gin.Context)
	DeleteCardController(ctx *gin.Context)
	GetCardListController(ctx *gin.Context)
}

//go:generate mockgen -source=handler.go -destination=./mocks/handler_mock.go -package CardHandlerMocks
type Service interface {
	CreateCard(ctx *gin.Context, card *model.Card) error
	GetCard(ctx *gin.Context, id int64) (*model.Card, error)
	UpdateCard(ctx *gin.Context, card *model.Card) (*model.Card, error)
	DeleteCard(ctx *gin.Context, id, userId int64) error
	GetCardList(ctx *gin.Context, search string, userId int64) ([]*model.Card, error)
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
