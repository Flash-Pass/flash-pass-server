package plan

import (
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getListRequest struct {
	userId uint64 `json:"id"`
}

func (h *Handler) GetList(ctx *gin.Context) {
	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.ParseTokenError, nil)
		return
	}

	param := &getListRequest{}
	if err := ctx.ShouldBind(param); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if param.userId != 0 {
		planList, err := h.service.GetList(ctx, param.userId)
		if err != nil {
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}

		res.RespondSuccess(ctx, planList)
	} else {
		planList, err := h.service.GetList(ctx, userId.(uint64))
		if err != nil {
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}

		res.RespondSuccess(ctx, planList)
	}
}
