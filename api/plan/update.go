package plan

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type updatePlanRequest struct {
	PlanId         int64  `json:"id" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	CycleSize      int    `json:"cycle_size" binding:"required"`
	LearnPerCycle  int    `json:"learn_per_cycle" binding:"required"`
	ReviewPerCycle int    `json:"review_per_cycle" binding:"required"`
	ReviewPerLearn int    `json:"review_per_learn" binding:"required"`
	GroupSize      int    `json:"group_size" binding:"required"`
	ReviewCycles   int    `json:"review_cycles" binding:"required"`
	LearnStrategy  string `json:"learn_strategy" binding:"required"`
	ReviewStrategy string `json:"review_strategy" binding:"required"`
}

func (h *Handler) Update(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.ParseTokenError, nil)
		return
	}

	param := &updatePlanRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	ok, err := h.service.IsPlanBelongToUser(ctx, param.PlanId, userId.(int64))
	if !ok {
		res.RespondWithError(c, http.StatusForbidden, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	plan, err := h.service.Update(context.Background(), &model.Plan{
		Base: model.Base{
			Id: param.PlanId,
		},
		Title:          param.Title,
		Description:    param.Description,
		CycleSize:      param.CycleSize,
		LearnPerCycle:  param.LearnPerCycle,
		ReviewPerCycle: param.ReviewPerCycle,
		ReviewPerLearn: param.ReviewPerLearn,
		GroupSize:      param.GroupSize,
		ReviewCycles:   param.ReviewCycles,
		LearnStrategy:  param.LearnStrategy,
		ReviewStrategy: param.ReviewStrategy,
	})
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, plan)
}
