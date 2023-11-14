package card

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetCardListRequest struct {
	Search string `json:"search"`
	UserId int64  `json:"id"`
}

func (h *Handler) GetCardListController(ctx *gin.Context) {
	params := &GetCardListRequest{}
	if err := ctx.Bind(params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	cards, err := h.service.GetCardList(ctx, params.Search, params.UserId)
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, cards)
}
