package api

import (
	"github.com/Flash-Pass/flash-pass-server/internal/encryptor"
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
	"github.com/Flash-Pass/flash-pass-server/internal/wechatClient"
	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/api/card"
	userhandler "github.com/Flash-Pass/flash-pass-server/api/user"
	"github.com/Flash-Pass/flash-pass-server/config"
	"github.com/Flash-Pass/flash-pass-server/db"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	cardrepo "github.com/Flash-Pass/flash-pass-server/db/repositories/card"
	userrepo "github.com/Flash-Pass/flash-pass-server/db/repositories/user"
	"github.com/Flash-Pass/flash-pass-server/internal/generator"
	cardservice "github.com/Flash-Pass/flash-pass-server/service/card"
	userservice "github.com/Flash-Pass/flash-pass-server/service/user"
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

func WithMiddleware(middleware ...gin.HandlerFunc) HandlerOption {
	return func(handler *Handler) {
		handler.Root.Use(middleware...)
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
	DB = DB.Debug()
	if err != nil {
		panic(err)
	}
	query.SetDefault(DB)

	// TODO: load all utils
	g := generator.New()
	e := encryptor.New()
	snowflakeHandle := snowflake.NewHandle(1)
	wxClient := wechatClient.New(cfg.WeChat.AppId, cfg.WeChat.Secret)

	// TODO: load all repositories
	cardRepository := cardrepo.NewRepository(DB)
	userRepository := userrepo.NewRepository(
		DB, g, e, snowflakeHandle,
	)

	// TODO: load all services
	cardService := cardservice.NewService(cardRepository)
	userService := userservice.NewService(userRepository, wxClient)

	// TODO: load all handlers
	cardHandler := card.NewHandler(cardService, 1)
	userHandler := userhandler.NewHandler(userService)

	h.AddRoutes(cardHandler, userHandler)
}
