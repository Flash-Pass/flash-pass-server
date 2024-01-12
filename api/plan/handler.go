package plan

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

type IHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetList(c *gin.Context)
}

type Service interface {
	Create(ctx context.Context, plan *model.Plan) (*model.Plan, error)
	IsPlanBelongToUser(ctx context.Context, planId, userId int64) (bool, error)
	Get(ctx context.Context, planId int64) (*model.Plan, error)
	Update(ctx context.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx context.Context, planId int64) error
	GetList(ctx context.Context, userId int64) ([]*model.Plan, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(r *gin.Engine) {
	router := r.Group("/plan")
	router.POST("/", h.Create)
	router.GET("/", h.Get)
	router.PUT("/", h.Update)
	router.DELETE("/", h.Delete)
	router.GET("/list", h.GetList)
}

var _ IHandler = (*Handler)(nil)
