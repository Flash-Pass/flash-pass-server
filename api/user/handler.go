package user

import "github.com/gin-gonic/gin"

type Handler struct {
	service UserService
}

type IHandler interface {
	AddRoutes(r *gin.Engine)
	loginViaWeChat(c *gin.Context)
	login(c *gin.Context)
	update(c *gin.Context)
}

type UserService interface {
}

func NewHandler(service UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(r *gin.Engine) {
	router := r.Group("/user")
	router.PUT("/", h.update)
	router.POST("/login/wx", h.loginViaWeChat)
	router.POST("/login", h.login)
}
