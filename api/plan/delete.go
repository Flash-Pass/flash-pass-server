package plan

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type deletePlanRequest struct {
	planId int64 `json:"id" binding:"required"`
}

func (h *Handler) Delete(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.ParseTokenError, nil)
		return
	}

	param := &deletePlanRequest{}
	if err := c.ShouldBind(param); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	ok, err := h.service.IsPlanBelongToUser(ctx, param.planId, userId.(int64))
	if !ok {
		res.RespondWithError(c, http.StatusForbidden, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if err := h.service.Delete(context.Background(), param.planId); err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, nil)
}
