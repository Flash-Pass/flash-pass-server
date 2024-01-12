package task

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service IService
}

type IHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
	Feed(c *gin.Context)
	AddLearnStatus(c *gin.Context)
}

type IService interface {
	CreateTask(ctx context.Context, planId, bookId, userId int64, name string) (*model.Task, error)
	Active(ctx context.Context, taskId int64, isActive bool) (*model.Task, error)
	DeleteTask(ctx context.Context, taskId int64) error
	GetTaskList(ctx context.Context, userId int64) ([]*model.Task, error)
	GetTaskListByIsActive(ctx context.Context, userId int64, isActive bool) ([]*model.Task, error)
	Feed(ctx context.Context, userId int64, taskId int64) ([]*model.TaskCardRecord, error)
	AddLearnStatus(ctx context.Context, taskCardRecordId int64, status string) error
	Update(ctx context.Context, taskId int64, taskName string, isActive bool) (*model.Task, error)
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(r *gin.Engine) {
	router := r.Group("/task")
	router.POST("/", h.Create)
	router.PUT("/", h.Update)
	router.DELETE("/", h.Delete)
	router.GET("/list", h.List)
	router.GET("/feed", h.Feed)
	router.POST("/card", h.AddLearnStatus)
}

var _ IHandler = (*Handler)(nil)
