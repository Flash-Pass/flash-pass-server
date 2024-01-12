package book

import (
	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
	"github.com/Flash-Pass/flash-pass-server/service/book"
)

type Handler struct {
	service         book.IService
	snowflakeHandle snowflake.IHandle
}

type IHandler interface {
	CreateBookController(c *gin.Context)
	GetBookController(c *gin.Context)
	UpdateBookController(c *gin.Context)
	DeleteBookController(c *gin.Context)
	GetBookListController(c *gin.Context)
	AddCardToBookController(c *gin.Context)
	RemoveCardFromBookController(c *gin.Context)
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
				list := card.Group("/list")
				{
					list.GET("/", h.GetBookCardListController)
				}
			}
		}
	}
}
