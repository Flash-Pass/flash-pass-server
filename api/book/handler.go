package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
	"github.com/Flash-Pass/flash-pass-server/service/book"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service         book.IService
	snowflakeHandle snowflake.IHandle
}

type IHandler interface {
	CreateBookController(ctx *gin.Context)
	GetBookController(ctx *gin.Context)
	UpdateBookController(ctx *gin.Context)
	DeleteBookController(ctx *gin.Context)
	GetBookListController(ctx *gin.Context)
	AddCardToBookController(ctx *gin.Context)
	RemoveCardFromBookController(ctx *gin.Context)
}

func NewHandler(service book.IService, snowflake *snowflake.Handle) *Handler {
	return &Handler{
		service:         service,
		snowflakeHandle: snowflake,
	}
}

func (h *Handler) AddRoutes(r *gin.Engine) {
	root := r.Group("/")
	{
		book := root.Group("/book")
		{
			book.POST("/", h.CreateBookController)
			book.PUT("/", h.UpdateBookController)
			book.DELETE("/", h.DeleteBookController)
			book.GET("/", h.GetBookController)
			book.GET("/list", h.GetBookListController)
			card := book.Group("/card")
			{
				card.POST("/", h.AddCardToBookController)
				card.DELETE("/", h.RemoveCardFromBookController)
			}
		}
	}
}
