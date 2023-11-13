package plan

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
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

func (h *Handler) Create(ctx *gin.Context) {
	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	params := &createRequest{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	plan, err := h.service.Create(ctx, &model.Plan{
		Title:          params.Title,
		Description:    params.Description,
		CycleSize:      params.CycleSize,
		LearnPerCycle:  params.LearnPerCycle,
		ReviewPerCycle: params.ReviewPerCycle,
		ReviewPerLearn: params.ReviewPerLearn,
		GroupSize:      params.GroupSize,
		ReviewCycles:   params.ReviewCycles,
		LearnStrategy:  params.LearnStrategy,
		ReviewStrategy: params.ReviewStrategy,
		CreatedBy:      userId.(uint64),
	})
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, plan)
}