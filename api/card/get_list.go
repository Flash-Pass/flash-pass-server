package card

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCardListRequest struct {
	search string `json:"search"`
	userId string `json:"id"`
}

func (h *Handler) GetCardListController(ctx *gin.Context) {
	params := &GetCardListRequest{}
	if err := ctx.ShouldBind(&params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	cards, err := h.service.GetCardList(ctx, params.search, params.userId)
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, cards)
}
