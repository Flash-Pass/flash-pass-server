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

type FeedRequest struct {
	TaskId int64 `json:"task_id"`
}

func (h *Handler) Feed(ctx *gin.Context) {
	logger := ctxlog.GetLogger(ctx)

	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		logger.Error("parse token error")
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	var req FeedRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	list, err := h.service.Feed(ctx, userId.(int64), req.TaskId)
	if err != nil {
		logger.Error("feed task failed", zap.Error(err))
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, list)
}
