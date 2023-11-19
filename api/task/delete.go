package task

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type DeleteRequest struct {
	TaskId int64 `json:"task_id"`
}

func (h *Handler) Delete(ctx *gin.Context) {
	logger := ctxlog.GetLogger(ctx)

	var req DeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if err := h.service.DeleteTask(ctx, req.TaskId); err != nil {
		logger.Error("delete task failed", zap.Error(err))
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}
