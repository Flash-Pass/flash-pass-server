package task

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateRequest struct {
	PlanId int64  `json:"plan_id"`
	BookId int64  `json:"book_id"`
	Name   string `json:"name"`
}

func (h *Handler) Create(c *gin.Context) {
	ctx, logger := ctxlog.Export(c)

	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		logger.Error("parse token error")
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("parse params error", zap.Error(err))
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	task, err := h.service.CreateTask(ctx, req.PlanId, req.BookId, userId.(int64), req.Name)
	if err != nil {
		logger.Error("create task failed", zap.Error(err))
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, task)
}
