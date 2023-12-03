package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBookCardListRequest struct {
	BookId int64 `json:"id,string" form:"id,string" required:"true"`
}

func (h *Handler) GetBookCardListController(ctx *gin.Context) {
	params := &GetBookCardListRequest{}
	if err := ctx.Bind(params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	cards, err := h.service.GetBookCardList(ctx, params.BookId)
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, cards)
}
