package task

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddLearnStatusRequest struct {
	TaskCardRecordId int64  `json:"task_card_record_id"`
	Status           string `json:"status"`
}

func (h *Handler) AddLearnStatus(c *gin.Context) {
	ctx, logger := ctxlog.Export(c)

	var req AddLearnStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if err := h.service.AddLearnStatus(ctx, req.TaskCardRecordId, req.Status); err != nil {
		logger.Error("add learn status failed", zap.Error(err))
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, nil)
}
