package plan

import (
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetPlanRequest struct {
	PlanId uint64 `json:"id" form:"id" binding:"required"`
}

func (h *Handler) Get(ctx *gin.Context) {
	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	param := &GetPlanRequest{}
	if err := ctx.BindQuery(param); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	ok, err := h.service.IsPlanBelongToUser(ctx, param.PlanId, userId.(uint64))
	if !ok {
		res.RespondWithError(ctx, http.StatusForbidden, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	plan, err := h.service.Get(ctx, param.PlanId)
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, plan)
}
