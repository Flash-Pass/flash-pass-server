package api

import (
	"github.com/Flash-Pass/flash-pass-server/api/card"
	userhandler "github.com/Flash-Pass/flash-pass-server/api/user"
	"github.com/Flash-Pass/flash-pass-server/config"
	"github.com/Flash-Pass/flash-pass-server/db"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	cardrepo "github.com/Flash-Pass/flash-pass-server/db/repositories/card"
	cardservice "github.com/Flash-Pass/flash-pass-server/service/card"
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
	DB, err := db.InitMySQL(cfg.MySQL)
	if err != nil {
		panic(err)
	}
	query.SetDefault(DB)

	// TODO: load all repositories
	cardRepository := cardrepo.NewRepository(DB)

	// TODO: load all services
	cardService := cardservice.NewService(cardRepository)
	userService := userservice.NewService()

	// TODO: load all handlers
	cardHandler := card.NewHandler(cardService, 1)
	userHandler := userhandler.NewHandler(userService)

	h.AddRoutes(cardHandler, userHandler)
}
