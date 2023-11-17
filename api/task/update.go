package task

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateRequest struct {
	TaskId   int64  `json:"task_id"`
	Name     string `json:"name"`
	IsActive string `json:"is_active"`
}

func (h *Handler) Update(ctx *gin.Context) {
	var req UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	var task *model.Task
	var err error
	if req.Name != "" {
		task, err = h.service.UpdateTaskName(ctx, req.TaskId, req.Name)
		if err != nil {
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
	}

	if req.IsActive != "" {
		active := false
		if req.IsActive == "true" || req.IsActive == "True" || req.IsActive == "TRUE" {
			active = true
		}
		task, err = h.service.Active(ctx, req.TaskId, active)
		if err != nil {
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
	}

	res.RespondSuccess(ctx, task)
}
