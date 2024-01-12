package plan

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetPlanRequest struct {
	PlanId int64 `json:"id" form:"id" binding:"required"`
}

func (h *Handler) Get(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}

	param := &GetPlanRequest{}
	if err := c.BindQuery(param); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	ok, err := h.service.IsPlanBelongToUser(ctx, param.PlanId, userId.(int64))
	if !ok {
		res.RespondWithError(c, http.StatusForbidden, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	plan, err := h.service.Get(ctx, param.PlanId)
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, plan)
}
