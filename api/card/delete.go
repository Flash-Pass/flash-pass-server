package card

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type DeleteCardRequest struct {
	Id int64 `json:"id" binding:"required"`
}

func (h *Handler) DeleteCardController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
	}

	params := &DeleteCardRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	if err := h.service.DeleteCard(ctx, params.Id, userId.(int64)); err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}
