package plan

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

type IHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetList(ctx *gin.Context)
}

type Service interface {
	Create(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	IsPlanBelongToUser(ctx *gin.Context, planId, userId uint64) (bool, error)
	Get(ctx *gin.Context, planId uint64) (*model.Plan, error)
	Update(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx *gin.Context, planId uint64) error
	GetList(ctx *gin.Context, userId uint64) ([]*model.Plan, error)
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
