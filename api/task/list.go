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

type ListRequest struct {
	IsActive string `json:"is_active"`
}

func (h *Handler) List(ctx *gin.Context) {
	logger := ctxlog.GetLogger(ctx)

	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		logger.Error("parse token error")
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	var req ListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if req.IsActive == "" {
		list, err := h.service.GetTaskList(ctx, userId.(int64))
		if err != nil {
			logger.Error("get task list failed", zap.Error(err))
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
		res.RespondSuccess(ctx, list)
		return
	} else {
		active := false
		if req.IsActive == "1" {
			active = true
		}

		list, err := h.service.GetTaskListByIsActive(ctx, userId.(int64), active)
		if err != nil {
			logger.Error("get task list failed", zap.Error(err))
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}

		res.RespondSuccess(ctx, list)
		return
	}
}
