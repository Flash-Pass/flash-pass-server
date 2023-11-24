package main

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/api"
	"github.com/Flash-Pass/flash-pass-server/config"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/fplog"
	_ "github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	middlewares "github.com/Flash-Pass/flash-pass-server/middlerwares"
	"github.com/gin-gonic/gin"
)

var (
	handler *api.Handler
)

func init() {
	ctx := context.Background()

	variables, err := config.LoadEnv(ctx)
	if err != nil {
		panic(err)
	}

	handler = api.NewHandler(
		api.WithRoot(gin.Default()),
		api.WithMiddleware(
			middlewares.PreRequest(),
			middlewares.CORS(variables.BASE.InDev),
			middlewares.GinLogger(),
			middlewares.GinRecovery(true),
			middlewares.Authorize(),
		),
	)
	handler.Load(variables)

	ctxlog.SetLogger(fplog.InitLogger(variables))
}

func main() {
	if err := handler.Root.Run(); err != nil {
		panic(err)
	}
}
