package plan

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type deletePlanRequest struct {
	planId int64 `json:"id" binding:"required"`
}

func (h *Handler) Delete(ctx *gin.Context) {
	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.ParseTokenError, nil)
		return
	}

	param := &deletePlanRequest{}
	if err := ctx.ShouldBind(param); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	ok, err := h.service.IsPlanBelongToUser(ctx, param.planId, userId.(int64))
	if !ok {
		res.RespondWithError(ctx, http.StatusForbidden, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if err := h.service.Delete(ctx, param.planId); err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}
