package task

import (
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type CreateRequest struct {
	PlanId int64  `json:"plan_id"`
	BookId int64  `json:"book_id"`
	Name   string `json:"name"`
}

func (h *Handler) Create(ctx *gin.Context) {
	logger := ctxlog.GetLogger(ctx)

	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		logger.Error("parse token error")
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	var req CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	task, err := h.service.CreateTask(ctx, req.PlanId, req.BookId, userId.(int64), req.Name)
	if err != nil {
		logger.Error("create task failed", zap.Error(err))
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, task)
}
