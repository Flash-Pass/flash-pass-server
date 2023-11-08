package main

import (
	"context"

	"github.com/Flash-Pass/flash-pass-server/api"
	"github.com/Flash-Pass/flash-pass-server/config"
	middlewares "github.com/Flash-Pass/flash-pass-server/middlerwares"
	"github.com/gin-gonic/gin"
)

var (
	//variables *config.EnvVariable
	handler *api.Handler
)

func init() {
	ctx := context.Background()

	variables, err := config.LoadEnv(ctx)
	if err != nil {
		panic(err)
	}

	handler = api.NewHandler(api.WithRoot(gin.Default()))
	handler.Load(variables)
}

func main() {
	r := handler.Root

	r.Use(middlewares.PreRequest())
	r.Use(middlewares.GinLogger())
	r.Use(middlewares.GinRecovery(true))

	if err := r.Run(); err != nil {
		panic(err)
	}
}
