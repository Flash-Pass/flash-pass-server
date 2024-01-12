package task

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteRequest struct {
	TaskId int64 `json:"task_id"`
}

func (h *Handler) Delete(c *gin.Context) {
	ctx, logger := ctxlog.Export(c)

	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if err := h.service.DeleteTask(ctx, req.TaskId); err != nil {
		logger.Error("delete task failed", zap.Error(err))
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, nil)
}
