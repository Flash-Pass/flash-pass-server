package api

import (
	"fmt"
	"github.com/Flash-Pass/flash-pass-server/config"

	userhandler "github.com/Flash-Pass/flash-pass-server/api/user"
	userservice "github.com/Flash-Pass/flash-pass-server/service/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Root *gin.Engine
}

type HandlerOption func(*Handler)

func NewHandler(options ...HandlerOption) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithRoot(root *gin.Engine) HandlerOption {
	return func(handler *Handler) {
		handler.Root = root
	}
}

type Router interface {
	AddRoutes(*gin.Engine)
}

func (h *Handler) AddRoutes(routers ...Router) {
	for _, router := range routers {
		router.AddRoutes(h.Root)
	}
}

func (h *Handler) GetRoutes() *gin.Engine {
	return h.Root
}

func (h *Handler) Load(cfg *config.EnvVariable) {
	// TODO: load all repositories
	fmt.Println(cfg)

	// TODO: load all services
	userService := userservice.NewService()

	// TODO: load all handlers
	userHandler := userhandler.NewHandler(userService)

	h.AddRoutes(userHandler)
}
