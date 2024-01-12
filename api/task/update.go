package task

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	TaskId   int64  `json:"task_id"`
	Name     string `json:"name"`
	IsActive string `json:"is_active"`
}

func (h *Handler) Update(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	active := false
	if req.IsActive == "true" || req.IsActive == "True" || req.IsActive == "TRUE" {
		active = true
	}

	task, err := h.service.Update(ctx, req.TaskId, req.Name, active)
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, task)
}
