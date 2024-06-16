package card

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetCardRequest struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
}

func (h *Handler) GetCardController(ctx *gin.Context) {
	var params GetCardRequest
	if err := ctx.Bind(&params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	card, err := h.service.GetCard(ctx, params.Id)
	if err != nil {
		res.RespondWithError(ctx, http.StatusNotFound, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, card)
}
