package user

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service         Service
	snowflakeHandle snowflake.IHandle
}

type IHandler interface {
	AddRoutes(r *gin.Context)
	loginViaWeChat(c *gin.Context)
	login(c *gin.Context)
	update(c *gin.Context)
	registerViaMobile(c *gin.Context)
	getInfo(c *gin.Context)
}

type Service interface {
	Login(ctx context.Context, mobile, password string) (token string, err error)
	LoginViaWeChat(ctx context.Context, code string) (token string, err error)
	Register(ctx context.Context, mobile, password string) (token string, err error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, openId, mobile string, userId int64) (*model.User, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(r *gin.Engine) {
	router := r.Group("/user")
	router.GET("/", h.getInfo)
	router.PUT("/", h.update)
	router.POST("/login/wx", h.loginViaWeChat)
	router.POST("/login", h.login)
	router.POST("/register/mobile", h.registerViaMobile)
}
