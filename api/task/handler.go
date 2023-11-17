package task

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service IService
}

type IHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
	Feed(ctx *gin.Context)
	AddLearnStatus(ctx *gin.Context)
}

type IService interface {
	CreateTask(ctx *gin.Context, planId, bookId, userId int64, name string) (*model.Task, error)
	Active(ctx *gin.Context, taskId int64, isActive bool) (*model.Task, error)
	UpdateTaskName(ctx *gin.Context, taskId int64, name string) (*model.Task, error)
	DeleteTask(ctx *gin.Context, taskId int64) error
	GetTaskList(ctx *gin.Context, userId int64) ([]*model.Task, error)
	GetTaskListByIsActive(ctx *gin.Context, userId int64, isActive bool) ([]*model.Task, error)
	Feed(ctx *gin.Context, userId int64, taskId int64) ([]*model.TaskCardRecord, error)
	AddLearnStatus(ctx *gin.Context, taskCardRecordId int64, status string) error
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
