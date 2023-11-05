package card

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCardRequest struct {
	id string `json:"id"`
}

func (h *Handler) GetCardController(ctx *gin.Context) {
	params := &GetCardRequest{}
	if err := ctx.ShouldBind(&params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	card, err := h.service.GetCard(ctx, params.id)
	if err != nil {
		res.RespondWithError(ctx, http.StatusNotFound, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, card)
}
